package datastructures

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	UNDISCOVERED = iota
	DISCOVERED
	PROCESSED
)

var parent []int

type Edge struct {
	Y      int
	weight int
}

type Graph struct {
	maxNodes  int          //the maximum number of vertices the graph can contain
	edges     []LinkedList //an array of linked lists to represh all edges to vertexes
	degree    []int        //number of edges for a vertex
	nVertices int
	nEdges    int
	directed  bool
}

//size is the total number of vertices that can be allowed
func NewGraph(size int, directed bool) (Graph, error) {

	if size <= 0 {
		return Graph{}, errors.New("Graph size is supplied is 0 ")
	}

	graph := Graph{maxNodes: size, edges: make([]LinkedList, size), degree: make([]int, size), nVertices: size, nEdges: 0, directed: directed}
	for i, _ := range graph.degree {
		graph.degree[i] = 0
	}
	return graph, nil
}

func (g *Graph) insertEdge(vertex1 string, vertex2 string) {

	vertex1 = strings.Trim(vertex1, "\n")
	vertex2 = strings.Trim(vertex2, "\n")

	vertex1_64, error1 := strconv.ParseInt(vertex1, 0, 64)
	vertex2_64, error2 := strconv.ParseInt(vertex2, 0, 64)

	if error1 != nil {
		fmt.Println(error1)
	}

	if error2 != nil {
		fmt.Println(error2)
	}

	if error1 == nil && error2 == nil {
		vertex1_int := int(vertex1_64)
		vertex2_int := int(vertex2_64)

		if vertex1_int > g.nVertices || vertex2_int > g.nVertices {
			fmt.Println("error node does not exist")
		}

		g.edges[vertex1_int].Insert(&Edge{Y: vertex2_int, weight: 1})
		g.nEdges = g.nEdges + 1
	}
}

//the data should have a proper format
func (g *Graph) ReadGraph(data []byte) error {
	dataString := string(data[:])

	stringsData := strings.Split(dataString, ",")
	fmt.Println("splitting done, number of rows:", len(stringsData))

	graphMetaData := strings.Split(stringsData[0], " ")
	fmt.Println("meta data:", len(graphMetaData))

	if (len(graphMetaData)) != 2 {
		return errors.New("too less information")
	}

	numberOfVertices64, err := strconv.ParseInt(graphMetaData[0], 0, 64)
	if err != nil {
		return errors.New("error parsing vertices")
	}

	numberOfVertices := int(numberOfVertices64)
	fmt.Println("number of vertices:", numberOfVertices)
	g.nVertices = numberOfVertices

	numberOfEdges64, err := strconv.ParseInt(graphMetaData[1], 0, 64)
	if err != nil {
		return errors.New("error parsing vertices")
	}
	numberOfEdges := int(numberOfEdges64)
	fmt.Println("number of edges:", numberOfEdges)

	g.nEdges = numberOfEdges

	for i := 1; i < len(stringsData); i++ {
		edgesData := strings.Split(stringsData[i], " ")
		if len(edgesData) < 2 {
			return errors.New("not enough nodes to form graphs")
		}
		for j := 1; j < len(edgesData); j++ {
			g.insertEdge(edgesData[0], edgesData[j])
		}

	}
	return nil
}

func (g *Graph) DisplayMatrix() {
	adjecencyMatrix := make([][]string, g.nVertices+1)

	for i := 1; i <= g.nVertices; i++ {
		adjecencyMatrix[i] = make([]string, g.nVertices+1)
	}

	for i := 1; i <= g.nVertices; i++ {
		g.edges[i].Traverse(func(listItem *ListItem) {
			adjecencyMatrix[i][listItem.item.(*Edge).Y] = "1"
		})
	}

	for i := 1; i <= g.nVertices; i++ {
		for j := 1; j <= g.nVertices; j++ {

			if adjecencyMatrix[i][j] == "" {
				adjecencyMatrix[i][j] = "0"
			}

			fmt.Print(adjecencyMatrix[i][j], " ")
		}
		fmt.Println()
	}

}

func (g *Graph) TraverseBFS(processEdge func(edge *Edge)) {

	fmt.Println("beginning breath first traversal...")
	if processEdge == nil {
		return
	}

	//loop through edges to get the first vertex which is not nil
	var headVertex int
	for i, _ := range g.edges {
		if g.edges[i].Length() != 0 {
			headVertex = i
			break
		}
	}
	fmt.Println("Vertex selected as head:", headVertex)
	g.bfs(processEdge, headVertex)
}

func (g *Graph) bfs(processEdge func(edge *Edge), headVertex int) {
	nVertextState := make([]int, g.nVertices+1)
	parent = make([]int, g.nVertices+1)
	for i := 1; i <= g.nVertices; i++ {
		nVertextState[i] = UNDISCOVERED
	}

	nVertextState[headVertex] = DISCOVERED

	//0 is no parent
	parent[headVertex] = 0

	queue := NewQueue(g.nVertices * g.nVertices)
	queue.Enque(headVertex)

	for queue.Length() != 0 {

		vertex := queue.Deque().(int)

		g.edges[vertex].Traverse(func(listItem *ListItem) {

			fmt.Print("from vertex:", vertex, " to ")

			processEdge(listItem.item.(*Edge))

			adjecentVertex := listItem.item.(*Edge).Y
			if nVertextState[adjecentVertex] == UNDISCOVERED {
				nVertextState[adjecentVertex] = DISCOVERED
				parent[adjecentVertex] = vertex
				queue.Enque(adjecentVertex)
			}

		})
		nVertextState[vertex] = PROCESSED
	}

	for i := 1; i <= g.nVertices; i++ {
		fmt.Println("parent of vertex:", i, " is vertex:", parent[i])
	}
}

func (g *Graph) FindPath(vertex int) {
	fmt.Println("the path")
	if parent != nil && len(parent) != 0 {
		path(g, vertex)
	}
	fmt.Println()
}

func path(g *Graph, i int) {

	if parent[i] == 0 {
		fmt.Print("start at ", i, " ")
		return
	} else {
		path(g, parent[i])
		fmt.Print(i, " ")
	}

}
