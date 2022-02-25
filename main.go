package main

import (
	"fmt"
)

var (
	BuildVersion string = "0.1"
)

func main() {
	fmt.Printf("wireguard whatchdog | Version %s", BuildVersion)
}
