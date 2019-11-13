package main

import (
	"flag"

	xrpapi "github.com/gaspardpeduzzi/spring_block/data"
)

func main() {
	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	xrpapi.XrpGetLedgerSeq(addr)

}
