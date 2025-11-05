package main
import "fmt"
func main(){
	fmt.Println("Results of if");
	x :=5;
	y :=6;
	i :=3;
	fmt.Println("x is ",x," y is ",y)
	if x<y{
		fmt.Println("x < y");
	}else{
       fmt.Println("x => y");
	}
	if i :=9; i < x{
		fmt.Println(" i < x", " i is ",i," x is ",x);
	} else{
		fmt.Println("i >=x","i is ",i,"x is ",x);
	}
	fmt.Println(" i is ",i);

}
