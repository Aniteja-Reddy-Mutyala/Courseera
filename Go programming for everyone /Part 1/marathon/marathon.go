package main
import "fmt"
func main(){
 var miles,yards int32=26,385;
 var kilometers float32;
 kilometers =1.6 * (float32(miles) + float32(yards)/1760.0)
 fmt.Printf("A marathon is %g kilometers",kilometers);
  
}
