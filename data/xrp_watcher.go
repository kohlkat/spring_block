package data

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

// websocket address

func GetLastLedgerSeq(addr *string) int {
	flag.Parse()
	log.SetFlags(0)

	var ls LedgerSeq

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

	ls.Id = 2
	ls.Command = "ledger_current"

	msg, _ := json.Marshal(ls)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(msg)))
	if err != nil {
		log.Println("write:", err)
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	log.Printf("recv: %s", message)

	var rc ResponseCurrent
	err = json.Unmarshal(message, &rc)
	if err != nil {
		log.Println("Issue unmarshalling sequence response", err.Error())
	}
	return rc.Result.LedgerCurrentIndex
}
