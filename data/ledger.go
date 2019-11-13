package data

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var lastLedger = 51366888

func GetLedgerData(addr *string) {
	flag.Parse()
	log.SetFlags(0)

	// check for interrupts and cleanly close the connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	// make the connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	// on exit close
	defer c.Close()
	done := make(chan struct{})

	var lr LedgerRequest

	// Example values
	lr.Id = lastLedger
	lr.Command = "ledger"
	lr.LedgerIndex = "validated"
	lr.Full = false
	lr.Accounts = false
	lr.Transactions = true
	lr.Expand = false
	lr.OwnerFunds = false

	// struct to JSON marshalling
	msg, _ := json.Marshal(lr)
	// write to the websocket
	err = c.WriteMessage(websocket.TextMessage, []byte(string(msg)))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// read from the websocket
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	// print the response from the XRP Ledger
	log.Printf("recv: %s", message)

	response := &LedgerResponse{}
	err = json.Unmarshal(message, response)
	if err != nil {
		panic("Error unmarshalling" + err.Error())
	}

	log.Println("if this is good this is nice", response.Result.Ledger.Transactions[0])

	// handle interrupt
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
