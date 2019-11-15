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
	c := make(chan int)
	liquidOptimizer := NewOptimizer(*addr, c)
	go liquidOptimizer.ConstructTxGraph()


	for {
			<-c
			all_offers, cycle := liquidOptimizer.Graph.GetProfitableOffers()
			if all_offers != nil {
				log.Println("Found profitable cycle:", cycle)
				for i, offers := range all_offers {
					for _, offer := range offers {
						log.Println(cycle[i], "->", cycle[(i+1)%len(cycle)], offer.Rate, offer.Hash)
					}
				}
				return
			}
	}







}
