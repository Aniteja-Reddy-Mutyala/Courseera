package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var pair,howMany int
	source :=rand.NewSource(time.Now().UnixNano())
	rng:=rand.New(source)
	for i:=0; i < 10000000; i++ {
		pair=rng.Intn(6) +rng.Intn(6) + 2
		if pair ==7{
			howMany ++
		}
	}
	fmt.Println(howMany);
	fmt.Printf("the probability of 7 is %g\n",float32(howMany)/10000000.0)

}