package main

import (
	"fmt"
)

type A struct {
	b string
}

func main() {
	a := []*A{}
	b := new(A)
	b.b = "sdas"
	a = append(a, b)
	c(a)
	drivers := make(map[string]string)
	drivers["driverName"] = "sads"
	driveri := drivers["driverName"]
	fmt.Println("dr", driveri)
	// fmt.Println("ok", ok)
}

func c(a []*A) {
	fmt.Println(a)
	fmt.Println(a[0].b)
}
