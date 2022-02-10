package dsgraph

import (
	"fmt"
	"sidsin/srcmodule/graph"
	"sidsin/srcmodule/list"
)

type DSGraph struct {
	vertexset  []*map[int]graph.GraphNode
	adjlist    []*list.LinkedList
	edgeset    map[graph.Edge]bool
	isdirected bool
}

//implementing makeset with a map
//creates a map with x as key and a new graphnode as value
//graphnode.value is also set to x
func MakeSet(x int) *map[int]graph.GraphNode {
	m := make(map[int]graph.GraphNode)
	gnp := graph.MakeGraphNode(x, "WHITE", 0, nil, 0)
	m[x] = *gnp
	return &m
}

func PrintSet(p *map[int]graph.GraphNode) {

	if p != nil {
		m := *p
		fmt.Println(m)
	} else {
		fmt.Println("Given map is empty.")
	}
}

func MakeDSGraph(vertexcount int, isdirected bool) *DSGraph {

	//note: single slice holding all the pointers to maps
	//vertexset is a slice of map pointers
	vertexset := make([]*map[int]graph.GraphNode, vertexcount)

	//make a new map with key = i
	//and set ith slice value to point to the new map
	for i := range vertexset {

		mp := MakeSet(i)
		vertexset[i] = mp

	}

	adjlist := make([]*list.LinkedList, vertexcount)

	//initialize adj lists
	for i := 0; i < vertexcount; i++ {
		adjlist[i] = list.CreateEmptyLinkedList()
	}

	edgeset := make(map[graph.Edge]bool)

	dsgraph := &DSGraph{vertexset: vertexset, adjlist: adjlist, edgeset: edgeset, isdirected: isdirected}

	return dsgraph
}

//assumes that both x and y are not nil
//all the contents of y are added to x
//preserves the condition that keys are not repeated
//after union even if x and y have the same element
//after call to Union, set y to nil
func Union(x, y *map[int]graph.GraphNode) *map[int]graph.GraphNode {

	m1 := *x
	m2 := *y

	//add all elements of y to x
	for k, gn := range m2 {
		m1[k] = gn
		delete(m2, k)
	}

	return x
}

func FindSet(x int, gp *DSGraph) *map[int]graph.GraphNode {

	dsg := *gp

	var mapwithx *map[int]graph.GraphNode
	var keyfound bool

	for _, mp := range dsg.vertexset {
		if mp != nil {

			//get the map to which mp points to
			m := *mp
			for range m {
				_, ok := m[x]

				if ok {
					mapwithx = mp
					keyfound = true
					break
				}
			}

			if keyfound {
				break
			}
		}
	}

	return mapwithx
}
