package main

type QueueEntry struct {
	p                Position
	d                int
	heatLoss         int
	straightDistance int
	path             []Position
}

type VisitEntry struct {
	p Position
	d int
	s int
}

func (e *QueueEntry) IsHigherPriority(than any) bool {
	return e.heatLoss < than.(*QueueEntry).heatLoss
}
