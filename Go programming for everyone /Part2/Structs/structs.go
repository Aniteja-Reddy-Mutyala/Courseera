package main

import "fmt"
type Student struct{
	name string
	gpa uint8
}
func main(){
	fmt.Println("Using Student Struct")
	var s Student
	pointerToS:=&s;
	s.name="Ani";
	pointerToS.gpa=9;
	fmt.Println("Student is ",s);
	fmt.Println("Dereferenced student is ",*(pointerToS));
	fmt.Println("Student name is ",s.name);
	fmt.Println("Student gpa is ",pointerToS.gpa);


}
