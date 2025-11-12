package main

import "fmt"

type Node struct{
	data string
	next *Node
}
type Stack struct{
	top *Node
}
func (s *Stack) Push(str string){
	newNode:=&Node{data:str,next:s.top};
	s.top=newNode;
}
func (s *Stack) Pop() string{
	if s.top==nil{
		return "Empty";
	}
	
		value:=s.top.data;
		s.top=s.top.next;
		return value;

	

}
func (s *Stack) Top() string{
	if s.top==nil{
		return "Empty";
	}else{
		return  s.top.data;
	}
}
func (s *Stack) isEmpty() bool{
	return s.top==nil;
}
func main(){
	stack:=Stack{};
	var howMany int;
	fmt.Println("How many elements?");
	fmt.Scanf("%d",&howMany);
	var word string
	for i:=0; i<howMany;i++{
		fmt.Scanf("%s",&word);
		stack.Push(word);
	}
	fmt.Println("Top of the stack ",stack.Top());
	fmt.Println("Stack ",stack);
	fmt.Println("poping elements from stack ",stack.Pop());
    for !stack.isEmpty(){
		println(" ",stack.Pop())
	}
}