package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//fmt.Println(time.Now().Unix())

	//ret := rand.Int()
	ret2 := rand.Intn(100) //上限为100
	//ret1 := ret % 100
	fmt.Println(ret2)

}
