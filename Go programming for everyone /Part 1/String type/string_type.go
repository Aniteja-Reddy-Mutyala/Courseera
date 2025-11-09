package main

import "fmt"

func main(){
	var c byte ='a';
	var strOne string
	strOne="My Idea?";
	fmt.Println("String program");
	fmt.Println("First one is :",strOne);
	fmt.Println("c is :",c);
	sum :=len(strOne);
	for i:=0; i< sum ;i++{
		fmt.Printf("char %d is %c\n",i,strOne[i]);
	}
}