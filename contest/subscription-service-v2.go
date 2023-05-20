package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const Schema string = `{
   "$id": "message.schema.json",
   "type": "object",
   "properties": {
     "trace_id": {
       "type": "string"
     },
     "offer": {
       "$ref": "offer.schema.json"
     }
   },
   "required": [
     "trace_id",
     "offer"
   ]
}`

type PartnerContent struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type Offer struct {
	Id             *string         `json:"id,omitempty"`
	Price          *int            `json:"price,omitempty"`
	StockCount     *int            `json:"stock_count,omitempty"`
	PartnerContent *PartnerContent `json:"partner_content,omitempty"`
}

type Message struct {
	TraceId string `json:"trace_id,omitempty"`
	Offer   Offer  `json:"offer,omitempty"`
}

type Subscriber struct {
	Trig       int
	Ship       int
	TrigFields []string
	ShipFields []string
}

func deepFields(contA *PartnerContent, contB PartnerContent) []string {
	ans := make([]string, 0)

	if contB.Title != nil {
		ans = append(ans, "Title")
		contA.Title = contB.Title
	}
	if contB.Description != nil {
		ans = append(ans, "Description")
		contA.Description = contB.Description
	}
	return ans
}

func updateFields(offerA *Offer, offerB Offer) []string {
	ans := make([]string, 0)

	aVal := reflect.ValueOf(offerA).Elem()
	aTyp := aVal.Type()
	bVal := reflect.ValueOf(offerB)

	for i := 0; i < aVal.NumField(); i++ {
		if bVal.Field(i).IsNil() {
			continue
		}
		if aTyp.Field(i).Name == "PartnerContent" {
			deepAns := deepFields(offerA.PartnerContent, *offerB.PartnerContent)
			ans = append(ans, deepAns...)
			continue
		}
		aVal.Field(i).Set(bVal.Field(i))
		ans = append(ans, aTyp.Field(i).Name)
	}
	return ans
}

func read() ([2]int, []Subscriber, []Message) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nm := strings.Split(sc.Text(), " ")
	n, e := strconv.Atoi(nm[0])
	if e != nil {
		log.Fatal(e)
	}
	m, e := strconv.Atoi(nm[1])
	if e != nil {
		log.Fatal(e)
	}

	ans1 := [2]int{n, m}

	ans2 := make([]Subscriber, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		cur := strings.Split(sc.Text(), " ")
		a, e := strconv.Atoi(cur[0])
		if e != nil {
			log.Fatal(e)
		}
		b, e := strconv.Atoi(cur[1])
		if e != nil {
			log.Fatal(e)
		}
		af := cur[2 : 2+a]
		bf := cur[2+a:]

		for i := range af {
			conv := strings.Split(af[i], "_")
			for j := range conv {
				conv[j] = strings.Title(conv[j])
			}
			af[i] = strings.Join(conv, "")
		}

		for i := range bf {
			conv := strings.Split(bf[i], "_")
			for j := range conv {
				conv[j] = strings.Title(conv[j])
			}
			bf[i] = strings.Join(conv, "")
		}

		ans2[i] = Subscriber{a, b, af, bf}
	}

	ans3 := make([]Message, 0)
	for i := 0; i < m; i++ {
		var cur Message
		sc.Scan()
		e = json.Unmarshal(sc.Bytes(), &cur)
		if e != nil {
			log.Fatal(e)
		}
		if cur.Offer.PartnerContent == nil {
			var curDeep PartnerContent
			cur.Offer.PartnerContent = &curDeep
		}
		ans3 = append(ans3, cur)
	}

	return ans1, ans2, ans3
}

func printTriggers(subs []Subscriber, diff []string, id, trid string, bd map[string]Offer) {
	for _, sub := range subs {
		for _, field := range diff {
			for _, el := range sub.TrigFields {
				if el == field {
					alert(sub, id, trid, bd)
					return
				}
			}
		}
	}

}

func alert(s Subscriber, id, trid string, bd map[string]Offer) {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	allFields := append(s.ShipFields, s.TrigFields...)
	var mes Message
	mes.TraceId = trid
	var off Offer
	mes.Offer = off
	var deep PartnerContent
	mes.Offer.PartnerContent = &deep

	curOffer := bd[id]
	oVal := reflect.ValueOf(mes.Offer)
	oTyp := oVal.Type()
	for i := 0; i < oTyp.NumField(); i++ {
		for _, f := range allFields {
			if f == oTyp.Field(i).Name {
				// I FUCKING DIED
				bVal := reflect.ValueOf(curOffer)
				bTyp := bVal.Type()
				for j := 0; j < bTyp.NumField(); j++ {
					if f == bTyp.Field(i).Name {
						oVal.Field(i).Set(bVal.Field(i))
					}
				}
			}
		}
	}

	ans, e := json.Marshal(mes)
	if e != nil {
		log.Fatal(e)
	}

	// TBD: get all fields of bd[id]
	// which are kept in s.TrigFields and s.ShipFields
	// fmt.Print(mes.Offer.Id == nil)

	wr.Write(ans)
}

func main() {
	nm, subs, msgs := read()
	fmt.Println(nm)

	for i := range subs {
		fmt.Println(subs[i])
	}
	// fmt.Println(msgs)

	for i := range msgs {
		fmt.Printf("%s\n", msgs[i])
	}
	//bd := make(map[string]Offer)
	//
	//for i := 0; i < nm[1]; i++ {
	//	curMsg := msgs[i]
	//
	//	if _, ok := bd[*curMsg.Offer.Id]; !ok {
	//		bd[*curMsg.Offer.Id] = curMsg.Offer
	//		continue
	//	}
	//
	//	copy := bd[*curMsg.Offer.Id]
	//	diffFields := updateFields(&copy, curMsg.Offer)
	//	printTriggers(subs, diffFields, *curMsg.Offer.Id, curMsg.TraceId, bd)
	//	bd[*curMsg.Offer.Id] = copy
	//}

	// fmt.Println("id", *bd["1"].Id)
	// fmt.Println("price", *bd["1"].Price)
	// fmt.Println("stock", *bd["1"].StockCount)

	// fmt.Println("id", *bd["2"].Id)
	// fmt.Println("desc", *bd["2"].PartnerContent.Description)
	// fmt.Println("title", *bd["2"].PartnerContent.Title)
}
