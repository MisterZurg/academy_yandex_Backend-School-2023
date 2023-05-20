package handler

import (
	tb "github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	tbe "github.com/didip/tollbooth_echo"
	"yandex-team.ru/bstask/internal/api"
)

func RegisterHandlersWithBaseURLUsingRateLimiter(router api.EchoRouter, si api.ServerInterface, baseURL string) {
	wrapper := api.ServerInterfaceWrapper{
		Handler: si,
	}

	const MAX_RPS_FOR_ENDPOINT float64 = 10
	var opts *limiter.ExpirableOptions

	router.GET(baseURL+"/couriers", wrapper.GetCouriers, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.POST(baseURL+"/couriers", wrapper.CreateCourier, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.GET(baseURL+"/couriers/assignments", wrapper.CouriersAssignments, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.GET(baseURL+"/couriers/meta-info/:courier_id", wrapper.GetCourierMetaInfo, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.GET(baseURL+"/couriers/:courier_id", wrapper.GetCourierById, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.GET(baseURL+"/orders", wrapper.GetOrders, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.POST(baseURL+"/orders", wrapper.CreateOrder, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.POST(baseURL+"/orders/assign", wrapper.OrdersAssign, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.POST(baseURL+"/orders/complete", wrapper.CompleteOrder, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))
	router.GET(baseURL+"/orders/:order_id", wrapper.GetOrder, tbe.LimitHandler(tb.NewLimiter(MAX_RPS_FOR_ENDPOINT, opts)))

}
