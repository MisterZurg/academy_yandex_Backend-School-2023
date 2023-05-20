package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type TraceContext struct {
	TraceID string `json:"trace_id"` // по условию string
	Offer   Offer  `json:"offer,omitempty"`
}

// Offer описывает Товарное предложение в базе
type Offer struct {
	ID             *string         `json:"id"`
	Price          *int            `json:"price,omitempty"`
	StockCount     *int            `json:"stock_count,omitempty"`
	PartnerContent *PartnerContent `json:"partner_content,omitempty"`
}

type PartnerContent struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Subscriber struct {
	Triggers  []string
	Shipments []string
}

// частичное обновление товарных предложений в базе данных
func updateOffer(req Offer, db map[string]map[string]bool) {
	// offerID : [field : processed]
	// ANTI PATTERN
	v := reflect.ValueOf(req)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

// уведомление сервисов-подписчиков при обновлении данных
func alertSubs() {

}

func parseInput() ([]Subscriber, []TraceContext) {
	n, m := parseNM()
	subs := parseSubscribers(n)
	reqs := parseRequests(m)
	return subs, reqs
}

func parseNM() (int, int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nm := strings.Split(sc.Text(), " ")
	n, err := strconv.Atoi(nm[0])
	if err != nil {
		panic(err)
	}

	m, err := strconv.Atoi(nm[1])
	if err != nil {
		panic(err)
	}
	return n, m
}

func parseSubscribers(n int) []Subscriber {
	sc := bufio.NewScanner(os.Stdin)
	subs := make([]Subscriber, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		curr := strings.Split(sc.Text(), " ")
		trigLen, err := strconv.Atoi(curr[0])
		if err != nil {
			panic(err)
		}
		shipLen, err := strconv.Atoi(curr[1])
		if err != nil {
			panic(err)
		}
		var subscriber Subscriber
		idx := 2
		if trigLen > 0 {
			subscriber.Triggers = curr[idx : trigLen+2]
			idx += trigLen + 2
		}

		if shipLen > 0 {
			subscriber.Shipments = curr[idx:]
		}
		subs[i] = subscriber
	}
	return subs
}

func parseRequests(m int) []TraceContext {
	sc := bufio.NewScanner(os.Stdin)
	traces := make([]TraceContext, m)

	for i := 0; i < m; i++ {
		sc.Scan()
		var trace TraceContext
		err := json.Unmarshal(sc.Bytes(), &trace)
		if err != nil {
			panic(err)
		}
		traces[i] = trace
	}
	return traces
}

func main() {
	subs, reqs := parseInput()

	// offerID : [field : processed]
	dataBase := make(map[string]map[string]bool)
	for _, req := range reqs {
		trID := req.TraceID
		updateOffer(req.Offer, dataBase)
	}
	// fmt.Println(subs)
}

func printStructs() {
	//for _, sub := range subs {
	//	fmt.Printf("%v\n", sub)
	//}
	//
	//fmt.Println()
	//for _, r := range reqs {
	//	fmt.Printf("%s\n", r)
	//}
}

//2 5
//2 0 price stock_count
//1 0 partner_content
//{"trace_id": "1", "offer": {"id": "1", "price": 9990}}
//{"trace_id": "2", "offer": {"id": "1", "stock_count": 100}}
//{"trace_id": "3", "offer": {"id": "2", "partner_content": {"title": "Backpack"}}}
//{"trace_id": "4", "offer": {"id": "1", "stock_count": 100}}
//{"trace_id": "5", "offer": {"id": "2", "partner_content": {"description": "Best backpack ever"}}}
func TestCase() {
	//2 5
	//2 0 price stock_count
	//1 0 partner_content
	//{"trace_id": "1", "offer": {"id": "1", "price": 9990}}
	//{"trace_id": "2", "offer": {"id": "1", "stock_count": 100}}
	//{"trace_id": "3", "offer": {"id": "2", "partner_content": {"title": "Backpack"}}}
	//{"trace_id": "4", "offer": {"id": "1", "stock_count": 100}}
	//{"trace_id": "5", "offer": {"id": "2", "partner_content": {"description": "Best backpack ever"}}}

	//{"trace_id":"1","offer":{"id":"1","price":9990}}
	//{"trace_id":"2","offer":{"id":"1","price":9990,"stock_count":100}}
	//{"trace_id":"3","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
	//{"trace_id":"5","offer":{"id":"2","partner_content":{"description":"Best backpack ever","title":"Backpack"}}}
}
