package repository

import (
	"encoding/json"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"time"
	"yandex-team.ru/bstask/internal/api"
)

type CouriersRepository struct {
	*sqlx.DB
}

// (GET /couriers)
func (cr *CouriersRepository) GetCouriers(ctx echo.Context, params api.GetCouriersParams) error {
	const (
		OFFSET_DEFAULT_VALUE = 0
		LIMIT_DEFAULT_VALUE  = 1
	)
	getCouriersResponse := new(api.GetCouriersResponse)

	// Loop through rows using only one struct
	rows, err := cr.DB.Queryx("SELECT courier_id, courier_type, array_to_json(regions), array_to_json(working_hours) FROM couriers")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Cannot select * from couriers in GetCouriers")
	}

	var start, end int32
	if params.Offset == nil {
		getCouriersResponse.Offset = OFFSET_DEFAULT_VALUE
		start = OFFSET_DEFAULT_VALUE
	} else {
		start = *params.Offset
	}

	if params.Limit == nil {
		getCouriersResponse.Limit = LIMIT_DEFAULT_VALUE
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

		courier := new(api.CourierDto)

		var regsStr, wohoStr string
		var regs []int32
		var woho []string

		err := rows.Scan(
			&courier.CourierId,
			&courier.CourierType,
			&regsStr,
			&wohoStr,
		)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "Cannot fill CourierDto in GetCouriers")
		}

		json.Unmarshal([]byte(regsStr), &regs)
		json.Unmarshal([]byte(wohoStr), &woho)

		courier.Regions = regs
		courier.WorkingHours = woho

		getCouriersResponse.Couriers = append(getCouriersResponse.Couriers, *courier)
	}

	return ctx.JSON(http.StatusOK, getCouriersResponse)
}

// (POST /couriers)
func (cr *CouriersRepository) CreateCourier(ctx echo.Context) error {
	couriers := new(api.CreateCourierRequest)
	couriersResponse := new(api.CreateCouriersResponse)
	if err := ctx.Bind(&couriers); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid JSON format for CreateCourierRequest")
	}

	for _, courier := range couriers.Couriers {
		curResp := api.CourierDto{
			Regions:      courier.Regions,
			WorkingHours: courier.WorkingHours,
			CourierType:  api.CourierDtoCourierType(courier.CourierType),
		}

		row := cr.DB.QueryRow(`INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3) RETURNING courier_id`,
			courier.CourierType,
			courier.Regions,
			courier.WorkingHours,
		)

		var id int64
		if err := row.Scan(&id); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Cannot insert into Courier")
		}

		curResp.CourierId = id
		couriersResponse.Couriers = append(couriersResponse.Couriers, curResp)
	}

	return ctx.JSON(http.StatusOK, couriers)
}

// Список распределенных заказов
// (GET /couriers/assignments)
func (cr *CouriersRepository) CouriersAssignments(ctx echo.Context, params api.CouriersAssignmentsParams) error {
	return nil
}

// (GET /couriers/meta-info/{courier_id})
func (cr *CouriersRepository) GetCourierMetaInfo(ctx echo.Context, courierId int64, params api.GetCourierMetaInfoParams) error {
	var start, end time.Time
	var err error
	//var sqlStatement string
	var groupId, orderId int64
	var rows *sqlx.Rows
	var totalCost, multiSalary, multiRating int
	var orderNum float64 = 0
	var regsStr, wohoStr string

	ans := new(api.GetCourierMetaInfoResponse)

	groupIds := make([]int64, 0)
	orderIds := make([]int64, 0)
	start, err = time.Parse("2006-01-02", params.StartDate.String())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot parse start_date")
	}

	end, err = time.Parse("2006-01-02", params.EndDate.String())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot parse end_date")
	}

	err = cr.DB.QueryRowx(`SELECT courier_id, courier_type, array_to_json(regions), array_to_json(working_hours) FROM couriers WHERE courier_id=$1`, courierId).Scan(
		&ans.CourierId,
		&ans.CourierType,
		&regsStr,
		&wohoStr,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot find courier with given id")
	}
	json.Unmarshal([]byte(regsStr), &ans.Regions)
	json.Unmarshal([]byte(wohoStr), &ans.WorkingHours)

	switch ans.CourierType {
	case api.FOOT:
		multiSalary = 2
		multiRating = 3
	case api.BIKE:
		multiSalary = 3
		multiRating = 2
	case api.AUTO:
		multiSalary = 4
		multiRating = 1
	}

	rows, err = cr.DB.Queryx(`SELECT group_id FROM couriers_groups WHERE courier_id=$1`, courierId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot find order groups for courier")
	}

	for rows.Next() {
		err = rows.Scan(&groupId)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot scan for order group")
		}
		groupIds = append(groupIds, groupId)
	}
	for _, grid := range groupIds {
		rows, err = cr.DB.Queryx(`SELECT order_id FROM groups_orders WHERE group_id=$1`, grid)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot find order groups for courier")
		}
		for rows.Next() {
			err = rows.Scan(&orderId)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, "cannot scan for order")
			}
			orderIds = append(orderIds, orderId)
		}
	}

	oidMap := make(map[int64]bool)
	for _, oid := range orderIds {
		oidMap[oid] = true
	}
	for oid := range oidMap {
		rows, err = cr.DB.Queryx(`SELECT cost FROM orders WHERE order_id=$1 AND completed_time >= $2 AND completed_time < $3`, oid, start, end)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "cannot find orders")
		}
		for rows.Next() {
			curCost := 0
			err = rows.Scan(&curCost)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, "cannot scan for order")
			}
			totalCost += curCost
			orderNum++
		}
	}
	ansSalary := int32(multiSalary * totalCost)
	ansRating := int32(multiRating) * int32(math.RoundToEven(orderNum/end.Sub(start).Hours()))

	ans.Earnings = &ansSalary
	ans.Rating = &ansRating
	return ctx.JSON(http.StatusOK, ans)
}

// (GET /couriers/{courier_id})
func (cr *CouriersRepository) GetCourierById(ctx echo.Context, courierId int64) error {
	courier := new(api.CourierDto)

	var regsStr, wohoStr string
	var regs []int32
	var woho []string

	if err := cr.DB.QueryRow(`SELECT courier_id, courier_type, array_to_json(regions), array_to_json(working_hours) FROM couriers WHERE courier_id=$1`, courierId).Scan(
		&courier.CourierId,
		&courier.CourierType,
		&regsStr,
		&wohoStr,
	); err != nil {
		return ctx.JSON(http.StatusBadRequest, "cannot select from couriers in GetCourierById")
	}

	json.Unmarshal([]byte(regsStr), &regs)
	json.Unmarshal([]byte(wohoStr), &woho)

	courier.Regions = regs
	courier.WorkingHours = woho

	return ctx.JSON(http.StatusOK, courier)
}
