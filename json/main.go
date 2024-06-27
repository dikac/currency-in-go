package main

import (
	"encoding/json"
	"log"
)

type Data struct {
	Decimal string  `json:"decimal"`
	Float   float64 `json:"float"`
}

func main() {

	payload := []byte(`{"decimal":"999.99999999999999","float":999.99999999999999}`)

	result := Data{}

	_ = json.Unmarshal(payload, &result)

	log.Print("Decimal: ", result.Decimal)
	log.Print("Float: ", result.Float)

}
