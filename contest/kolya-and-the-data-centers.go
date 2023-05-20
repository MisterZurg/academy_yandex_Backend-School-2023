package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	metric     int
	dataCenter int
}

func main() {
	// число дата-центров,
	// число серверов в каждом из дата-центров
	// и число событий соответственно.
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	db := make([]int, n) // битовая маска включенных машин
	da := make([]int, n)
	for i := range da {
		da[i] = m // число включённыых машин
	}
	dr := make([]int, n) // число перезапусков
	dw := make([]int, n) // метрика r*a

	// В кучу надо класть только ДЦ, с серверами которых проводились действия
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	for i := 0; i < n; i++ {
		maxHeap.Push(Pair{0, i})
		minHeap.Push(Pair{0, i})
	}

	// fmt.Println(maxHeap)

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		// fmt.Println(maxHeap)
		sc.Scan()
		words := strings.Split(sc.Text(), " ")

		switch words[0] {
		case "RESET":
			dcNum, _ := strconv.Atoi(words[1])
			j := dcNum - 1
			db[j] = 0
			da[j] = m
			dr[j]++
			dw[j] = dr[j] * da[j]

			heap.Push(minHeap, Pair{dw[j], j})
			heap.Push(maxHeap, Pair{-dw[j], j})
		// в i-м дата-центре был выключен j-й сервер
		case "DISABLE":
			dcNum, _ := strconv.Atoi(words[1])
			srvNum, _ := strconv.Atoi(words[2])

			dcNum--
			srvNum--
			t := 1 << srvNum
			if (db[dcNum] & t) == 1 {
				continue
			}
			db[dcNum] |= t
			da[dcNum] -= 1
			dw[dcNum] -= dr[dcNum]

			heap.Push(minHeap, Pair{dw[dcNum], dcNum})
			heap.Push(maxHeap, Pair{-dw[dcNum], dcNum})
		// был перезагружен i-й дата-центр
		case "GETMAX":
			for -(*maxHeap)[0].metric != dw[(*maxHeap)[0].dataCenter] {
				heap.Pop(maxHeap)
			}
			fmt.Println((*maxHeap)[0].dataCenter + 1)
		case "GETMIN":
			for (*minHeap)[0].metric != dw[(*minHeap)[0].dataCenter] {
				heap.Pop(minHeap)
			}
			fmt.Println((*minHeap)[0].dataCenter + 1)
		}
	}
}

// "container/heap"
type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].metric > h[j].metric && h[i].dataCenter < h[j].dataCenter
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].metric < h[j].metric // && h[i].dataCenter < h[j].dataCenter
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

//
//type DataCenters struct {
//	dataCenters []DataCenter
//	currMax     int
//	currMin    int
//}
//
//func (dcs *DataCenters) UpdateMax(v int) {
//	if
//}
//
//type DataCenter struct {
//	servers  map[int]bool
//	restarts int
//	active   int
//}
//
//func InitDSs(n, m int) *DataCenters {
//	dsc := make([]DataCenter, n)
//	for i := 0; i < n; i++ {
//		srv := make(map[int]bool)
//		dsc[i].servers = srv
//		dsc[i].restarts = 0
//		dsc[i].active = m
//	}
//	return &DataCenters{dataCenters: dsc, currMax: 1, currMin: 1}
//}
//
//func (dataCenter *DataCenter) RESET() int {
//	for i := range dataCenter.servers {
//		dataCenter.servers[i] = true
//	}
//	dataCenter.active = len(dataCenter.servers)
//	dataCenter.restarts++
//	return dataCenter.active * dataCenter.restarts
//}
//
//func (dataCenter *DataCenter) DISABLE(srv int) {
//	v := dataCenter.servers[srv]
//	if v != true {
//		dataCenter.servers[srv] = false
//		dataCenter.active--
//	}
//}
//
//func main() {
//	// число дата-центров,
//	// число серверов в каждом из дата-центров
//	// и число событий соответственно.
//	var n, m, q int
//	fmt.Scan(&n, &m, &q)
//	dcs := InitDSs(n, m)
//	sc := bufio.NewScanner(os.Stdin)
//
//	for i := 0; i < q; q++ {
//		sc.Scan()
//		words := strings.Split(sc.Text(), " ")
//
//		switch words[0] {
//		// в i-м дата-центре был выключен j-й сервер
//		case "DISABLE":
//			dcNum, _ := strconv.Atoi(words[1])
//			srvNum, _ := strconv.Atoi(words[2])
//			dcs.dataCenters[dcNum].DISABLE(srvNum)
//		// был перезагружен i-й дата-центр
//		case "RESET":
//			dcNum, _ := strconv.Atoi(words[1])
//			restsActives := dcs.dataCenters[dcNum].RESET()
//			dcs.UpdateMax(restsActives)
//		case "GETMAX":
//		case "GETMIN":
//		}
//	}
//}
