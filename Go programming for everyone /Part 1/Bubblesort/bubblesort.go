package main

import "fmt"

func main(){
	elements :=[]int{10,5,3,9,12};
	for i:=0;i <len(elements)-1;i++{
		for j:=0;j<len(elements)-1-i;j++{
			if elements[j]>elements[j+1]{
				elements[j],elements[j+1]=elements[j+1],elements[j]
			}
		}
	}
	fmt.Println(elements);
}