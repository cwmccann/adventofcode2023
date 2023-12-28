package main

import (
	"adventofcode2023/utils"
	"fmt"
	"github.com/gammazero/deque"
	"math/rand"
	"strings"
	"sort"
)

func main() {
	text := utils.InputToString()

	solution := SolvePart1(text)
	fmt.Printf("Solution Part 1: %d\n", solution)

	solution = SolvePart2(text)
	fmt.Printf("Solution Part 2: %d\n", solution)
}

func SolvePart1(input string) int {
	adjList := make(map[string][]string)
	nodes := make([]string, 0)
	//Make the initial Adjacency List
	for _, line := range utils.TrimAndRemoveEmptyLines(strings.Split(input, "\n")) {
		parts := strings.Split(line, ": ")
		id := parts[0]
		nodes = append(nodes, id)
		neighbors := strings.Split(parts[1], " ")
		for _, neighbor := range neighbors {
			addItemToAdjacencyList(adjList, id, neighbor)
		}
	}

	numberOfRuns := 1000
	i := 0
	tracking := make(map[string]int)

	for i < numberOfRuns {
		i++
		//find the path between two random nodes
		start, end := choose2Nodes(nodes)
		path := bfs(adjList, start, end)
		for i := 0; i < len(path)-1; i++ {
			a := path[i]
			b := path[i+1]

			if a < b {
				tracking[a+","+b]++
			} else {
				tracking[b+","+a]++
			}
		}
	}

	// Convert the map to a slice of its values
	keys := make([]string, 0, len(tracking))
	for key := range tracking {
		keys = append(keys, key)
	}

	// Sort the slice in descending order
	sort.Slice(keys, func(i, j int) bool {
		return tracking[keys[i]] > tracking[keys[j]]
	})

	// Take the first 3 elements of the sorted slice
	largestValues := keys[:3]

	//Cut the edges
	for _, value := range largestValues {
		a := strings.Split(value, ",")[0]
		b := strings.Split(value, ",")[1]

		adjList[a] = utils.RemoveStringFromSlice(adjList[a], b)
		adjList[b] = utils.RemoveStringFromSlice(adjList[b], a)
	}

	//Count the connected components in each side
	connectedComponents := countConnectedComponents(adjList)
	if len(connectedComponents) != 2 {
		panic(fmt.Sprintf("Expected 2 groups of connected components.  There was %d.", len(connectedComponents)))
	}
	ans := 1
	for _, count := range connectedComponents {
		ans *= count
	}
	return ans
}

func bfs(adjList map[string][]string, start string, end string) []string {
	q := deque.New[string]()
	q.PushBack(start)
	visited := make(map[string]bool)
	parent := make(map[string]string)

	for q.Len() > 0 {
		currentNode := q.PopFront()
		visited[currentNode] = true
		if currentNode == end {
			//we found the path
			break
		}
		for _, neighbor := range adjList[currentNode] {
			if !visited[neighbor] {
				q.PushBack(neighbor)
				parent[neighbor] = currentNode
			}
		}
	}

	// Build the path from end to start
	path := []string{}
	for node := end; node != ""; node = parent[node] {
		path = append(path, node)
	}

	// Reverse the path to get it from start to end
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func countConnectedComponents(adjList map[string][]string) []int {
	visited := make(map[string]bool)
	components := make([]int, 0)

	var dfs func(node string, count int) int
	dfs = func(node string, count int) int {
        visited[node] = true

        for _, neighbor := range adjList[node] {
            if !visited[neighbor] {
                count = dfs(neighbor, count + 1)
            }
        }
		return count
    }

	for u := range adjList {
		if !visited[u] {
			components = append(components, dfs(u, 1))
		}
	}

	return components
}

func choose2Nodes(nodes []string) (string, string) {
	// Pick two random indices
	index1 := rand.Intn(len(nodes))
	index2 := rand.Intn(len(nodes))

	// Make sure the indices are different
	for index1 == index2 {
		index2 = rand.Intn(len(nodes))
	}

	return nodes[index1], nodes[index2]
}

func addItemToAdjacencyList(adjacencyList map[string][]string, a string, b string) {
	if _, ok := adjacencyList[a]; !ok {
		adjacencyList[a] = make([]string, 0)
	}
	adjacencyList[a] = utils.AddUniqueToList(adjacencyList[a], b)

	if _, ok := adjacencyList[b]; !ok {
		adjacencyList[b] = make([]string, 0)
	}
	adjacencyList[b] = utils.AddUniqueToList(adjacencyList[b], a)
}

func SolvePart2(input string) int {
	return -1
}

/*
type Graph struct {
    adjList map[string][]string
    edges   []Edge
}
type Edge struct {
    u, v string
	originalU, originalV string
}
func NewEdge(u, v string) Edge {
	return Edge{u, v, u, v}
}

func NewGraph(adjList map[string][]string) *Graph {
    g := &Graph{adjList: adjList, edges: []Edge{}}
    for u, neighbors := range adjList {
        for _, v := range neighbors {
            if u < v { // Ensure each edge is added only once
                g.edges = append(g.edges, NewEdge(u, v))
            }
        }
    }
    return g
}
func (g *Graph) ContractEdge(u, v string) {
    // Update edges: replace all occurrences of v with u
    for i, edge := range g.edges {
        if edge.u == v {
            g.edges[i].u = u
        }
        if edge.v == v {
            g.edges[i].v = u
        }
    }

	// Update adjacency list
    // Merge v into u, and remove v from the graph
    for _, neighbor := range g.adjList[v] {
        // Avoid adding self-loops
        if neighbor != u {
            g.adjList[u] = append(g.adjList[u], neighbor)
            // Replace v with u in the neighbor's adjacency list
            for i, n := range g.adjList[neighbor] {
                if n == v {
                    g.adjList[neighbor][i] = u
                }
            }
        }
    }
    delete(g.adjList, v)
}
func (g *Graph) RandomEdge() (string, string) {
    var u string
    var v string
    for u = range g.adjList {
        break
    }
    v = g.adjList[u][rand.Intn(len(g.adjList[u]))]
    return u, v
}
func (g *Graph) KargersMinCutEdges() []Edge {

	//Contract edges until only 2 nodes remain
    for len(g.adjList) > 2 {
        u, v := g.RandomEdge()
        g.ContractEdge(u, v)
    }

    // Filter out self-loops to identify the cut edges
    cutEdges := []Edge{}
    for _, edge := range g.edges {
        if edge.u != edge.v && !edgeExists(cutEdges, edge) {
            cutEdges = append(cutEdges, edge)
        }
    }
    return cutEdges
}

func (g *Graph) CountConnectedComponents() int {
	visited := make(map[string]bool)
	count := 0

	var dfs func(node string)

	dfs = func(node string) {
        visited[node] = true

        for _, neighbor := range g.adjList[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
    }

	for u := range g.adjList {
		if !visited[u] {
			count++
			dfs(u)
		}
	}

	return count
}

func edgeExists(edges []Edge, edge Edge) bool {
    for _, e := range edges {
        if (e == edge) {
            return true
        }
    }
    return false
}

*/
