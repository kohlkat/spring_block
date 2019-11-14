package main

import "math"

var priceRipple = 0.27


func dropToXrp(drop float64) (xrp float64){
	xrp = drop/math.Pow(10, 6)
	return
}


func dropToPriceInUSD (drop int) (usd float64){
	return dropToXrp(float64(drop))*priceRipple

}
