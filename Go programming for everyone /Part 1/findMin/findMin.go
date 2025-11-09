package main

import (
	"fmt"
	"math/rand"
	"time"
)
 func minSlice(data []float64 ) float64{
   var min float64=1000.0
   for _,v := range data{
	 if v < min{
		min = v
	 }
   }
   return min
 }
 func maxSlice(data [] float64) float64{
	var max float64=0.0000;
	for _,v := range data{
		if v > max{
			max =v
		}
	}
	return max
 }

func main(){
	rand.Seed(time.Now().UnixNano());
	fmt.Println("The size of slice is ");
	size :=0
	fmt.Scanf("%d",&size);
	sizePtr:=&size;
	data:=make([]float64,*sizePtr)
	for i:=0;i<size;i++{
		data[i]=rand.Float64()*1000;
	}
	startTime:=time.Now();

   fmt.Println("min of ",size ," rand.Float64() is ",minSlice(data))
   fmt.Println("max of ",size , " rand.Float64() is ",maxSlice(data))
   duration :=time.Since(startTime);
   fmt.Println("running time of ",duration)


}
