// Package for topological sort of the DAG.
package toposort

type Ordered []string

func (ord *Ordered) traverse(edges [][]string, start string, visited map[string]bool) {
	visited[start] = true
	for _, a := range edges {
		if a[0] == start && !visited[a[1]] {
			ord.traverse(edges, a[1], visited)
		}
	}
	*ord = append(*ord, start)
}

// Sort does topological sort on a DAG supplied as slice of edges.
// Edge holds two named vertices. Edge (slice) has vertex as element 0
// pointing to vertex element 1. Look at test cases for concrete examples.
func Sort(edges [][]string) []string {
	var ord Ordered
	visited := make(map[string]bool)

	for _, a := range edges {
		start := a[0]
		if visited[start] {
			continue
		}
		ord.traverse(edges, start, visited)
	}

	for i, j := 0, len(ord)-1; i < len(ord)/2; i, j = i+1, j-1 {
		ord[i], ord[j] = ord[j], ord[i]
	}
	return ord
}
