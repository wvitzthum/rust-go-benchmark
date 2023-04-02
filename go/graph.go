package main

type Graph struct {
    adjList [][]int // Adjacency list representation of the graph
}

func NewGraph(numVertices int) *Graph {
    adjList := make([][]int, numVertices)
    return &Graph{adjList}
}

func (g *Graph) AddEdge(u, v int) {
	// Add an undirected edge between nodes u and v
    g.adjList[u] = append(g.adjList[u], v)
    g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) BFS(start int) []int {
    visited := make([]bool, len(g.adjList)) // Keep track of visited nodes
    queue := make([]int, 0) // FIFO queue for BFS traversal
    traversal := make([]int, 0) // Store the traversal order

    visited[start] = true
    queue = append(queue, start)

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        traversal = append(traversal, node)

		// Visit each neighbor of the current node that has not been visited yet
        for _, neighbor := range g.adjList[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }

    return traversal
}

func exampleGraph() *Graph {
	graph := NewGraph(6) // Create a new graph with 6 nodes
    graph.AddEdge(0, 1) // Add edges between the nodes
    graph.AddEdge(0, 2)
    graph.AddEdge(1, 3)
    graph.AddEdge(2, 3)
    graph.AddEdge(2, 4)
    graph.AddEdge(3, 4)
    graph.AddEdge(3, 5)

	return graph
}