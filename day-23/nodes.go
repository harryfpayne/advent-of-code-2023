package main

import (
	"fmt"
)

type Node struct {
	Position Position
	Edges    []Edge
}

type Edge struct {
	To       Position
	Distance int
}

func NodesFromGrid(grid Grid, start Position) map[Position]Node {
	nodes := make(map[Position]Node)
	nodes[start] = Node{
		Position: start,
	}

	type QueueItem struct {
		node    Position
		start   Position
		exclude []Position
	}

	queue := []QueueItem{
		{start, start.Move(2), []Position{start}},
	}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		m := make(map[Position]struct{})
		for _, position := range s.exclude {
			m[position] = struct{}{}
		}
		m[s.node] = struct{}{}

		node, next, distance := followPathToNode(grid, s.start, m)

		if _, ok := nodes[node]; !ok {
			nodes[node] = Node{
				Position: node,
				Edges: []Edge{
					{s.node, distance},
				},
			}

			for _, position := range next {
				if _, ok := nodes[position]; ok {
					continue
				}
				queue = append(queue, QueueItem{
					node:    node,
					start:   position,
					exclude: []Position{node, position},
				})
			}
		} else {
			nodes[node] = Node{
				Position: node,
				Edges:    append(nodes[node].Edges, Edge{s.node, distance}),
			}
		}
		nodes[s.node] = Node{
			Position: s.node,
			Edges:    append(nodes[s.node].Edges, Edge{node, distance}),
		}
	}

	for p, node := range nodes {
		newNode := node
		newNode.Edges = dedupe(newNode.Edges)
		nodes[p] = newNode
	}

	return nodes
}

func dedupe(e []Edge) []Edge {
	a := make(map[Edge]struct{})
	for _, edge := range e {
		a[edge] = struct{}{}
	}
	var o []Edge
	for edge, _ := range a {
		o = append(o, edge)
	}
	return o
}

func followPathToNode(grid Grid, p Position, history map[Position]struct{}) (Position, []Position, int) {
	point, ok := grid.Get(p)
	if !ok {
		fmt.Println(p)
		panic("invalid position")
	}
	directions := point.AvailableDirections(true)

	var validPositions []Position
	for _, direction := range directions {
		position := p.Move(direction)
		point2, ok := grid.Get(position)
		if !ok || point2 == Tree {
			continue
		}
		if _, ok := history[position]; ok {
			continue
		}
		validPositions = append(validPositions, position)
	}
	if len(validPositions) == 0 { // I've found the end
		return p, validPositions, len(history) - 1
	}

	if len(validPositions) > 1 { // I've found a node
		return p, validPositions, len(history) - 1
	}
	history[validPositions[0]] = struct{}{}
	return followPathToNode(grid, validPositions[0], history)
}
