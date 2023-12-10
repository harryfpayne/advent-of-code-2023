package main

type Direction int

const (
	Right Direction = 0
	Left  Direction = 1
)

type Node struct {
	Name  string
	Right string
	Left  string
}

func (n Node) IsStart() bool {
	return n.Name[2] == 'A'
}

func (n Node) IsEnd() bool {
	return n.Name[2] == 'Z'
}

func (n Node) Get(d Direction) string {
	if d == Left {
		return n.Left
	}
	if d == Right {
		return n.Right
	}
	panic("invalid direction")
}
