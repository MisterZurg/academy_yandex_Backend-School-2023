// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for CourierDtoCourierType.
const (
	CourierDtoCourierTypeAUTO CourierDtoCourierType = "AUTO"
	CourierDtoCourierTypeBIKE CourierDtoCourierType = "BIKE"
	CourierDtoCourierTypeFOOT CourierDtoCourierType = "FOOT"
)

// Defines values for CreateCourierDtoCourierType.
const (
	CreateCourierDtoCourierTypeAUTO CreateCourierDtoCourierType = "AUTO"
	CreateCourierDtoCourierTypeBIKE CreateCourierDtoCourierType = "BIKE"
	CreateCourierDtoCourierTypeFOOT CreateCourierDtoCourierType = "FOOT"
)

// Defines values for GetCourierMetaInfoResponseCourierType.
const (
	AUTO GetCourierMetaInfoResponseCourierType = "AUTO"
	BIKE GetCourierMetaInfoResponseCourierType = "BIKE"
	FOOT GetCourierMetaInfoResponseCourierType = "FOOT"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse = map[string]interface{}

// CompleteOrder defines model for CompleteOrder.
type CompleteOrder struct {
	CompleteTime time.Time `json:"complete_time"`
	CourierId    int64     `json:"courier_id"`
	OrderId      int64     `json:"order_id"`
}

// CompleteOrderRequestDto defines model for CompleteOrderRequestDto.
type CompleteOrderRequestDto struct {
	CompleteInfo []CompleteOrder `json:"complete_info"`
}

// CourierDto defines model for CourierDto.
type CourierDto struct {
	CourierId    int64                 `json:"courier_id"`
	CourierType  CourierDtoCourierType `json:"courier_type"`
	Regions      []int32               `json:"regions"`
	WorkingHours []string              `json:"working_hours"`
}

// CourierDtoCourierType defines model for CourierDto.CourierType.
type CourierDtoCourierType string

// CouriersGroupOrders defines model for CouriersGroupOrders.
type CouriersGroupOrders struct {
	CourierId int64         `json:"courier_id"`
	Orders    []GroupOrders `json:"orders"`
}

// CreateCourierDto defines model for CreateCourierDto.
type CreateCourierDto struct {
	CourierType  CreateCourierDtoCourierType `json:"courier_type"`
	Regions      []int32                     `json:"regions"`
	WorkingHours []string                    `json:"working_hours"`
}

// CreateCourierDtoCourierType defines model for CreateCourierDto.CourierType.
type CreateCourierDtoCourierType string

// CreateCourierRequest defines model for CreateCourierRequest.
type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers"`
}

// CreateCouriersResponse defines model for CreateCouriersResponse.
type CreateCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
}

// CreateOrderDto defines model for CreateOrderDto.
type CreateOrderDto struct {
	Cost          int32    `json:"cost"`
	DeliveryHours []string `json:"delivery_hours"`
	Regions       int32    `json:"regions"`
	Weight        float32  `json:"weight"`
}

// CreateOrderRequest defines model for CreateOrderRequest.
type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders"`
}

// GetCourierMetaInfoResponse defines model for GetCourierMetaInfoResponse.
type GetCourierMetaInfoResponse struct {
	CourierId    int64                                 `json:"courier_id"`
	CourierType  GetCourierMetaInfoResponseCourierType `json:"courier_type"`
	Earnings     *int32                                `json:"earnings,omitempty"`
	Rating       *int32                                `json:"rating,omitempty"`
	Regions      []int32                               `json:"regions"`
	WorkingHours []string                              `json:"working_hours"`
}

// GetCourierMetaInfoResponseCourierType defines model for GetCourierMetaInfoResponse.CourierType.
type GetCourierMetaInfoResponseCourierType string

// GetCouriersResponse defines model for GetCouriersResponse.
type GetCouriersResponse struct {
	Couriers []CourierDto `json:"couriers"`
	Limit    int32        `json:"limit"`
	Offset   int32        `json:"offset"`
}

// GroupOrders defines model for GroupOrders.
type GroupOrders struct {
	GroupOrderId int64      `json:"group_order_id"`
	Orders       []OrderDto `json:"orders"`
}

// NotFoundResponse defines model for NotFoundResponse.
type NotFoundResponse = map[string]interface{}

// OrderAssignResponse defines model for OrderAssignResponse.
type OrderAssignResponse struct {
	Couriers []CouriersGroupOrders `json:"couriers"`
	Date     openapi_types.Date    `json:"date"`
}

// OrderDto defines model for OrderDto.
type OrderDto struct {
	CompletedTime *time.Time `json:"completed_time,omitempty"`

	// Cost Стоимость доставки заказа
	Cost          int32    `json:"cost"`
	DeliveryHours []string `json:"delivery_hours"`
	OrderId       int64    `json:"order_id"`
	Regions       int32    `json:"regions"`
	Weight        float32  `json:"weight"`
}

// GetCouriersParams defines parameters for GetCouriers.
type GetCouriersParams struct {
	// Limit Максимальное количество курьеров в выдаче. Если параметр не передан, то значение по умолчанию равно 1.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Количество курьеров, которое нужно пропустить для отображения текущей страницы. Если параметр не передан, то значение по умолчанию равно 0.
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// CouriersAssignmentsParams defines parameters for CouriersAssignments.
type CouriersAssignmentsParams struct {
	// Date Дата распределения заказов. Если не указана, то используется текущий день
	Date *openapi_types.Date `form:"date,omitempty" json:"date,omitempty"`

	// CourierId Идентификатор курьера для получения списка распредленных заказов. Если не указан, возвращаются данные по всем курьерам.
	CourierId *int `form:"courier_id,omitempty" json:"courier_id,omitempty"`
}

// GetCourierMetaInfoParams defines parameters for GetCourierMetaInfo.
type GetCourierMetaInfoParams struct {
	// StartDate Rating calculation start date
	StartDate openapi_types.Date `form:"startDate" json:"startDate"`

	// EndDate Rating calculation end date
	EndDate openapi_types.Date `form:"endDate" json:"endDate"`
}

// GetOrdersParams defines parameters for GetOrders.
type GetOrdersParams struct {
	// Limit Максимальное количество заказов в выдаче. Если параметр не передан, то значение по умолчанию равно 1.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Количество заказов, которое нужно пропустить для отображения текущей страницы. Если параметр не передан, то значение по умолчанию равно 0.
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// OrdersAssignParams defines parameters for OrdersAssign.
type OrdersAssignParams struct {
	// Date Дата распределения заказов. Если не указана, то используется текущий день
	Date *openapi_types.Date `form:"date,omitempty" json:"date,omitempty"`
}

// CreateCourierJSONRequestBody defines body for CreateCourier for application/json ContentType.
type CreateCourierJSONRequestBody = CreateCourierRequest

// CreateOrderJSONRequestBody defines body for CreateOrder for application/json ContentType.
type CreateOrderJSONRequestBody = CreateOrderRequest

// CompleteOrderJSONRequestBody defines body for CompleteOrder for application/json ContentType.
type CompleteOrderJSONRequestBody = CompleteOrderRequestDto
