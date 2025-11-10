package main

import "fmt"
func main(){
	i:=55;
	pointerToI:=&i;
	fmt.Println("I is :" ,i);
	fmt.Println("address of i is ",pointerToI);
	*pointerToI++;
	fmt.Println("Now i is :",i);
	j:=45;
	pointerToI=&j;
	fmt.Println("address of j is ",pointerToI);
	pointerToI=nil;
	fmt.Println("value of pointerToI is ",pointerToI);

}