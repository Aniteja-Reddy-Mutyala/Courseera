package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var pair,howMany,value int
	source :=rand.NewSource(time.Now().UnixNano())
	rng:=rand.New(source)
	fmt.Printf("Enter which value's probablilty you want in a dice");
	fmt.Scanf("%d",&value)
	for i:=0; i < 10000000; i++ {
		pair=rng.Intn(6) +rng.Intn(6) + 2
		if pair ==value{
			howMany ++
		}
	}
	fmt.Println(howMany);
	fmt.Printf("the probability of 7 is %g\n",float32(howMany)/10000000.0)

}