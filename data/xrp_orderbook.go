package data

// Basic exemple usage of https://xrpl.org/websocket-api-tool.html#book_offers

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

// websocket address

func XrpGetOrderbook(addr *string) {
	flag.Parse()
	log.SetFlags(0)

	var m message

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

	// Exemple values
	m.Command = "book_offers"
	//m.Taker = "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59"
	m.TakerGets.Currency = "XRP"
	m.TakerPays.Currency = "USD"
	m.TakerPays.Issuer = "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
	m.Limit = 10

	// struct to JSON marshalling
	msg, _ := json.Marshal(m)
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
