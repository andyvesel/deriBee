package trade

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type jsonData struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
}

type params struct {
	GrantType      string `json:"grant_type"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Currency       string `json:"currency"`
	Extended       bool   `json:"extended"`
	IndexName      string `json:"index_name"`
	InstrumentName string `json:"instrument_name"`
	Amount         int32  `json:"amount"`
	PostOnly       bool   `json:"post_only"`
	Price          int32  `json:"price"`
	OrderId        int    `json:"order_id"`
	Type           string `json:"type"`
}

type Message struct {
	jsonData
	Params params `json:"params"`
}

type result struct {
	AccessToken  string  `json:"access_token"`
	IndexPrice   float32 `json:"index_price"`
	State        string  `json:"state"`
	AveragePrice float32 `json:"average_price"`
	Order        order   `json:"order"`
}

type order struct {
	OrderId string `json:"order_id"`
}

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int32  `json: id"`
	Result  result `json:"result"`
}

var resp Response

func GetIndexPrice(conn *websocket.Conn) float32 {
	message := Message{
		jsonData{
			Jsonrpc: "2.0",
			Method:  "public/get_index_price",
		},
		params{
			IndexName: "btc_usd",
		},
	}
	conn.WriteJSON(message)

	err := conn.ReadJSON(&resp)
	if err != nil {
		log.Fatal(err)
	}

	float := resp.Result.IndexPrice
	return float
}

func EditOrder(conn *websocket.Conn) {

}

func BuyLimit(conn *websocket.Conn) {
	// adding 1000 to the actual price, in order to set buyLimit order
	// always near the actual price edge (PostOnly should be ALWAYS set to true)
	price := int32(GetIndexPrice(conn)) + 1000
	message := Message{
		jsonData{
			Jsonrpc: "2.0",
			Method:  "private/buy",
		},
		params{
			InstrumentName: "BTC-PERPETUAL",
			Amount:         10,
			Price:          price,
			PostOnly:       true,
			Type:           "limit",
		},
	}
	conn.WriteJSON(message)

	err := conn.ReadJSON(&resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Price %d", int32(price))
}
