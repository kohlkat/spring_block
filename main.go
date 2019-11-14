package main

import (
	"flag"
	"github.com/gaspardpeduzzi/spring_block/data"
	"log"
)

func main() {
	//go s.LaunchServer()
	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")

	log.Println("Checking last sequence number", data.GetLastLedgerSeq(addr))





	liquidOptimizer := NewOptimizer(*addr)
	liquidOptimizer.ConstructTxGraph()





}




