package main
import "fmt"
var myInt int= 5;
func main(){
	fmt.Println("Simple ideas");
	fmt.Println("Printing global variable : ",myInt);
	{
		var myInt int =6;
		fmt.Println("The inner variable : ",myInt);
	}
	fmt.Println("myInt value is : ",myInt);
}