package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"yandex-team.ru/bstask/internal/api"
)

type OrdersRepository struct {
	*sqlx.DB
}

// GetOrders (GET /orders)
/*
	Возвращает информацию о всех заказах, а также их дополнительную информацию:
	- вес заказа,
	- район доставки,
	- промежутки времени, в которые удобно принять заказ.
*/
func (or *OrdersRepository) GetOrders(ctx echo.Context, params api.GetOrdersParams) error {
	const (
		OFFSET_DEFAULT_VALUE = 0
		LIMIT_DEFAULT_VALUE  = 1
	)
	ordersResponse := make([]api.OrderDto, 0)

	rows, err := or.DB.Queryx("SELECT order_id, weight, regions, array_to_json(delivery_hours), cost, completed_time FROM orders")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Cannot select * from orders in GetOrders")
	}

	var start, end int32
	if params.Offset == nil {
		start = OFFSET_DEFAULT_VALUE
	} else {
		start = *params.Offset
	}

	if params.Limit == nil {
		end = start + LIMIT_DEFAULT_VALUE
	} else {
		end = start + *params.Limit
	}

	var idx int32 = 0

	for rows.Next() {
		if idx < start {
			idx++
			continue
		}
		if idx >= end {
			break
		}
		idx++

		order := new(api.OrderDto)

		var deliveryHoursStr string
		var deliveryHours []string

		err := rows.Scan(
			&order.OrderId,
			&order.Weight,
			&order.Regions,
			&deliveryHoursStr,
			&order.Cost,
			&order.CompletedTime,
		)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot fill OrderDto in GetOrders")
		}

		json.Unmarshal([]byte(deliveryHoursStr), &deliveryHours)

		order.DeliveryHours = deliveryHours

		ordersResponse = append(ordersResponse, *order)
	}

	return ctx.JSON(http.StatusOK, ordersResponse)
}

// (POST /orders)
func (or *OrdersRepository) CreateOrder(ctx echo.Context) error {
	orders := new(api.CreateOrderRequest)
	ordersResponse := make([]api.OrderDto, 0)

	//couriersResponse := new(api.Order)
	if err := ctx.Bind(&orders); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid format for CreateOrder JSON")
	}
	for _, order := range orders.Orders {
		curResp := api.OrderDto{
			Cost:          order.Cost,
			DeliveryHours: order.DeliveryHours,
			Regions:       order.Regions,
			Weight:        order.Weight,
		}
		fmt.Println(curResp)

		row := or.DB.QueryRow(`INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`,
			order.Weight,
			order.Regions,
			order.DeliveryHours,
			order.Cost,
		)

		var id int64
		if err := row.Scan(&id); err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot insert into table orders")
		}

		curResp.OrderId = id
		ordersResponse = append(ordersResponse, curResp)
	}

	return ctx.JSON(http.StatusOK, ordersResponse)
}

// Распределение заказов по курьерам
// (POST /orders/assign)
func (or *OrdersRepository) OrdersAssign(ctx echo.Context, params api.OrdersAssignParams) error {
	return nil
}

// (POST /orders/complete)
func (or *OrdersRepository) CompleteOrder(ctx echo.Context) error {
	var err error
	//var sqlStatement string
	var id int64
	completeRequest := new(api.CompleteOrderRequestDto)
	if err = ctx.Bind(&completeRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid format for CompleteOrder JSON")
	}

	for _, order := range completeRequest.CompleteInfo {

		err = or.DB.QueryRowx(`SELECT order_id FROM orders WHERE order_id=$1`, order.OrderId).Scan(
			&id,
		)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot find order in CompleteOrder")
		}
		// --------------------
		var orderId, groupId, courierId int64
		err = or.DB.QueryRowx(`SELECT group_id, order_id FROM groups_orders WHERE order_id=$1`, id).Scan(
			&groupId,
			&orderId,
		)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot select from groups_orders in CompleteOrder")
		}

		err = or.DB.QueryRowx(`SELECT courier_id, group_id FROM couriers_groups WHERE group_id=$1`, groupId).Scan(
			&courierId,
			&groupId,
		)

		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot select from couriers_groups in CompleteOrder")
		}

		if courierId != order.CourierId {
			return ctx.JSON(http.StatusBadRequest, "order assigned to another courier")
		}
		// --------------------

		row := or.DB.QueryRow(`UPDATE orders SET completed_time=$1 WHERE order_id=$2 RETURNING order_id`, order.CompleteTime, order.OrderId)
		if err = row.Scan(&id); err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot update order in table orders")
		}
	}
	return ctx.JSON(http.StatusOK, id)
}

// (GET /orders/{order_id})
func (or *OrdersRepository) GetOrder(ctx echo.Context, orderId int64) error {
	order := new(api.OrderDto)

	var deliveryHoursStr string
	var deliveryHours []string

	err := or.DB.QueryRow(`SELECT order_id, weight, regions, array_to_json(delivery_hours), cost, completed_time FROM orders WHERE order_id=$1`, orderId).Scan(
		&order.OrderId,
		&order.Weight,
		&order.Regions,
		&deliveryHoursStr,
		&order.Cost,
		&order.CompletedTime,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot select from orders in GetOrder")
	}

	json.Unmarshal([]byte(deliveryHoursStr), &deliveryHours)

	order.DeliveryHours = deliveryHours

	return ctx.JSON(http.StatusOK, order)
}
