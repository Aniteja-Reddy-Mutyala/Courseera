package main

import "fmt"
func convertToKm(miles *int ,yards *int) float32{
      return 1.6 * (float32(*miles) + float32(*yards)/1760.0)
}
func main(){
	var miles =[]int{10,26,45}
	var yards=[] int{0,385,44}
	for i,_:=range miles{
		fmt.Println("convert Miles and yards to kilometers:\n",miles[i]," mi",yards[i]," yd");
		fmt.Printf("Answer is  %f kilometers.\n\n",convertToKm(&miles[i],&yards[i]));
	}
}
