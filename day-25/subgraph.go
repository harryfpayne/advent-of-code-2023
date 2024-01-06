package main

func Subgraph(components map[string]Component) int {
	subGraphs := make(map[string]int)

	score := 1
	i := 0
	curr := randItemExcluding(components, subGraphs)
	for curr != "" {
		connected := Floodfill(components, curr)

		for conn, _ := range connected {
			subGraphs[conn] = i
		}
		i++
		score *= len(connected)
		curr = randItemExcluding(components, subGraphs)
	}
	return score
}

func Floodfill(components map[string]Component, start string) map[string]struct{} {
	queue := []string{
		start,
	}

	visited := make(map[string]struct{})
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr]; ok {
			continue
		}
		visited[curr] = struct{}{}

		comp := components[curr]
		for s := range comp.Connections {
			queue = append(queue, s)
		}
	}

	return visited
}

func randItemExcluding[K comparable, V1, V2 any](m map[K]V1, exclude map[K]V2) K {
	for i := range m {
		if _, ok := exclude[i]; !ok {
			return i
		}
	}
	return *new(K)
}
