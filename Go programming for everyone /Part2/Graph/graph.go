package main

import "fmt"

type Edge struct{
	source,destination int
	weight float64
}
type Graph struct{
	vertices int
	edges []Edge
}
func NewGraph(vertices int) *Graph{
	return &Graph{vertices,make([]Edge, 0) }
}
func (g *Graph) AddEdge(src,dest int,weight float64){
	g.edges=append(g.edges, Edge{src,dest,weight})
	g.edges = append(g.edges, Edge{dest,src,weight})
}
func main(){
    g:=NewGraph(5);
	g.AddEdge(0,1,4);
	g.AddEdge(0,4,8);
	g.AddEdge(1,2,8);
	g.AddEdge(1,3,11);
	g.AddEdge(1,4,3);
	g.AddEdge(2,3,1);
	g.AddEdge(3,2,1);
	g.AddEdge(3,4,2);
	g.AddEdge(4,3,2);
	fmt.Println("Graph: ",g);

}
