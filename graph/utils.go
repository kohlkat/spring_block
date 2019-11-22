package graph

import (
	"fmt"
	"log"
	"math"
	"os/exec"
)

var priceRipple = 0.27


func DropToXrp(drop float64) (xrp float64){
	xrp = drop/math.Pow(10, 6)
	return
}


func DropToPriceInUSD(drop int) (usd float64){
	return DropToXrp(float64(drop))*priceRipple
}


func removeOffer(slice []*Offer, s int) []*Offer {
	return append(slice[:s], slice[s+1:]...)
}

func (offer *Offer) Submit_Transaction(seq_nb int) {
	out, err := exec.Command("./submit.sh", offer.Account, offer.CreatorWillPay, fmt.Sprintf("%f", offer.Quantity), offer.Issuer, fmt.Sprintf("%d", seq_nb)).Output()
	log.Println("out, err", string(out), err)
}



