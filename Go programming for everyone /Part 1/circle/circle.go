package main
import ("fmt" 
 "math"
)
func main(){
	var area ,radius float64;
	//const pi=3.1459;
	fmt.Println("input radius in meters");
	fmt.Scanf("%f",&radius);
	
	area=math.Pi*radius *radius;
	fmt.Printf("The area is %g ,radius is %g\n",area,radius);

}
