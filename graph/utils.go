package graph

import (
	"math"
	"os/exec"
	"log"
	"fmt"
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

func Submit_Transaction(cycle []Offer) {
	// value1 := fmt.Sprintf("%f", offer.Quantity)
	// value2 := fmt.Sprintf("%f", offer.Quantity * offer.Rate)

	args := fmt.Sprintf("%s %f", cycle[0].CreatorWillPay, 1.0)
	for i, offer := range cycle[1:] {
		args = fmt.Sprintf("%s %s %s", args, offer.CreatorWillPay, cycle[i].Issuer)
	}

	out, err := exec.Command("./submit.sh", args).Output()
	log.Println("out, err", string(out), err)
}
