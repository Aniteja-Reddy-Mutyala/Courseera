package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data, left: nil, right: nil}
	} else if data < root.data {
		root.left = insert(root.left, data)
	} else if data > root.data {
		root.right = insert(root.right, data)
	}
	return root
}
func search(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	if root.data > data {
		return search(root.left, data)
	}
	return search(root.right, data)

}
func main(){
	var root *Node
	keys:=[]int{8,3,10,1,6,14,4,7,13};
	for _,data:=range keys{
		root=insert(root,data);
	}
	fmt.Println("Enter the number you want to search in binary tree");
	var searchInt int
	fmt.Scanf("%d",&searchInt)
	fmt.Println("",search(root,searchInt))
}