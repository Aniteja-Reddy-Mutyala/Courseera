package main

import "fmt"
func main(){
var miles=0 ;
var yards=0;
fmt.Println("Enter the value of miles and yards respectively.");
fmt.Scanf("%d%d\n",&miles,&yards);
fmt.Println("if miles or yards is negative ,then program closes");
if miles<0 || yards<0 {
	return
}


var result=convertToKm(&miles,&yards)
fmt.Println("The answer is ",result);
}
func convertToKm(miles *int ,yards *int) float32{
      return 1.6 * (float32(*miles) + float32(*yards)/1760.0)
}