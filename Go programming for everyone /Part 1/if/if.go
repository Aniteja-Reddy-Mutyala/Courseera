package main
import ("fmt"
   "math/rand"
   "time"
)
func main(){
	var move rune;
	fmt.Println("Enter either R P or S");
	fmt.Scanf("%c",&move);
	if move == 'R'{
		fmt.Printf("My move is %c\n ",move);

	} else if move == 'P'{
		fmt.Printf("My move is %c \n",move);
	} else if move == 'S'{
		fmt.Printf("My move is %c\n",move);
	} else{
		fmt.Printf("Illegal move is %c\n",move);
	}
     rand.Seed(time.Now().UnixNano())
	 machineMove := rand.Intn(3);
	 if machineMove==0 && move =='R'{
		fmt.Println("Machine wins");
	 } else if machineMove == 1 && move =='S'{
		fmt.Println("Machine wins");
	 } else if machineMove == 2 && move =='P'{
		fmt.Println("Machine wins");

	 } else {
		fmt.Println("Draw or you win");
	 }

	

}
