package main

type QueueItem struct {
	Position     Position
	Steps        int
	VisitHistory map[Position]int
}

func (q QueueItem) IsHigherPriority(than any) bool {
	return q.Steps > than.(QueueItem).Steps
}
