package graph

import "math"

var priceRipple = 0.27


func DropToXrp(drop float64) (xrp float64){
	xrp = drop/math.Pow(10, 6)
	return
}


func DropToPriceInUSD(drop int) (usd float64){
	return DropToXrp(float64(drop))*priceRipple

}
