package main

import (
	"flag"
	"strconv"

	"github.com/gaspardpeduzzi/spring_block/data"
	display "github.com/gaspardpeduzzi/spring_block/display_cli"
)

func main() {
	//go s.LaunchServer()
	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", true, "Display more information")
	flag.Parse()
	display.VERBOSE = *verb

	// Init Display
	display.Init()
	display.DisplayVerbose("Checking last sequence number ", strconv.Itoa(data.GetLastLedgerSeq(addr)))

	c := make(chan int)
	liquidOptimizer := NewOptimizer(*addr, c)
	go liquidOptimizer.ConstructTxGraph()

	for {
		<-c
		offers := liquidOptimizer.Graph.GetProfitableOffers()
		if offers != nil {
			display.DisplayVerbose("offers", offers)
			return
		}
	}

}
