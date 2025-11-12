package main

import "fmt"
type Edge struct{
	src,dest int
	weight float64
}
type Graph struct{
	vertices int
	edges []Edge
}
const Large=100000000
func NewGraph(vertices int) *Graph{
  return &Graph{vertices,make([]Edge, 0)}
}
func (g *Graph) AddEdge(src,dest int,weight float64){
   g.edges=append(g.edges, Edge{src,dest,weight})
   g.edges=append(g.edges, Edge{dest,src,weight})

}
func (g *Graph) Dijkstra(source int) []int{
	dist:=make([]int,g.vertices)
	visited:=make([]bool,g.vertices)
	for i:=range dist{
		dist[i]=Large;
	}
	dist[source]=0;
	for i:=0;i<g.vertices-1;i++{
	 u:=g.minDistance(dist,visited)
	 visited[u]=true;
	 for _,e :=range g.edges{
		if !visited[e.dest] &&e.src==u{
			newDist:=dist[u]+int(e.weight)
			if newDist<dist[e.dest]{
				dist[e.dest]=newDist
			}
		}
	 }
	}
	return dist
}
func (g *Graph)minDistance(dist []int,visited []bool)int{
	minDist:=Large;
	minIndex:=-1;
	for v:=0;v< g.vertices;v++{
		if !visited[v] &&dist[v]<minDist{
			minDist=dist[v];
			minIndex=v;
		}
	}
	return minIndex
}
func main(){
	g:=NewGraph(10);
	g.AddEdge(0,1,3);
	g.AddEdge(0,2,6);
	g.AddEdge(0,3,7);
	g.AddEdge(1,2,1);
	g.AddEdge(1,4,4);
	g.AddEdge(2,5,2);
	g.AddEdge(3,5,3);
	g.AddEdge(3,6,4);
	g.AddEdge(4,7,1);
	g.AddEdge(5,7,1);
	g.AddEdge(5,8,2);
	g.AddEdge(6,8,5);
	g.AddEdge(7,8,2);
	dist:=g.Dijkstra(0);
	fmt.Println("Shortest distance from vertex 0: ");
	for v,d:=range dist{
		fmt.Printf("Vertex %d:%d \n",v,d);
	}

}