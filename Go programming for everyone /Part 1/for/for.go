package main
import "fmt"
func main(){
	var n int
	fmt.Println("Countdown enter n:");
	fmt.Scanf("%d",&n);
	for i:=n; i >= 0; i--{
		if i==0{
			fmt.Println("Blast off");
		} else{
			fmt.Println("Countdown : ",i);
		}
	}
}
