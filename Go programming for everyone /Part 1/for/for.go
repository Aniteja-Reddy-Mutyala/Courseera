package main
import ("fmt"
"math/rand"
"time"
)
func main(){
	var move ,machineMove int
	const (
		rock = 0
		paper = 1
		scissors = 2

	)
	const (
		cRock= 'R'
		cPaper= 'P'
		cScissors='S'
	)
	var cMove rune
	var draws,wins,machineWins int
	var rounds int
	fmt.Println("How many rounds do you want to play ?.");
	fmt.Scanf("%d",&rounds);
	for i:= 0; i < rounds; i++ {
		fmt.Println("\n Round " ,i+1 ,": Choose either R P or S")
		fmt.Scanf("%c\n", &cMove);
		if cMove==cRock{
			move =rock
		} else if cMove==cPaper{
			move = paper
		} else if cMove==cScissors{
			move = scissors
		} else{
			fmt.Println("Illegal move");
			i--
			continue
		}
		rand.Seed(time.Now().UnixNano())
		machineMove= rand.Intn(3)
		if machineMove==rock{
			cMove =cRock
		}else if machineMove==paper{
			cMove=cPaper
		} else if machineMove==scissors{
			cMove=cScissors
		}
		fmt.Printf("Machine plays %c",cMove);
		if move ==machineMove{
			fmt.Println("=> Draw")
			draws++
		} else if (move==rock) && (machineMove==paper){
			machineWins++
			fmt.Println(" Machine wins")
			continue
		} else if (move == paper) &&(machineMove==scissors){
			machineWins++
			fmt.Println(" Machine wins")
			continue
		} else if (move == scissors) &&(machineMove == rock){
			machineWins++
			fmt.Println(" Machine wins")
			continue
		} else{
			wins++
			fmt.Println("=> You win")
		}

	}
	fmt.Println("After ",rounds ," rounds:\n","you win: ",wins," machine wins: ",machineWins," with ",draws," draws")
}
