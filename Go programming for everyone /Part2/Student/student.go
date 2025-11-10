package main

import "fmt"

type Name struct{
	first string
	last string
}
type Person struct{
	who Name
	address string
	age int
}
type Student struct{
	per Person
	testScores []int
	gradeAverage int
}
func (s *Student) testAverage(){
	sum :=0;
	for _,v :=range s.testScores{
		sum +=v;
	}
	s.gradeAverage=sum/len(s.testScores)
}
func inputScores(s *Student,testCount int){
	fmt.Printf("Input %d scores for %s",testCount,s.per.who);
	var score int;
	for i:=0;i< testCount;i++{
		fmt.Scanf("%d",&score);
		s.testScores[i]=score;
	}
}
func (s Student) printStudent(){
   fmt.Println("Student and student score ",s);
}
func main(){
    var p Person
	p.who.last="Coder";
	p.who.first="Ani";
	var s Student;
	p.address="North pole";
	p.age=29;
	s.per=p;
	ptrToS:=&s;
    var testCount int;
	fmt.Println("How many tests?");
	fmt.Scanf("%d",&testCount);
	s.testScores=make([]int,testCount);
	inputScores(&s,testCount);
	ptrToS.testAverage();
	ptrToS.printStudent();
	fmt.Println(s);
}