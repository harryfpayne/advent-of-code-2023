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

//type PriorityQueue []*QueueEntry
//
//var _ heap.Interface = new(PriorityQueue)
//
//func (p PriorityQueue) Len() int {
//	return len(p)
//}
//
//func (p PriorityQueue) Less(i, j int) bool {
//	return p[i].heatLoss < p[j].heatLoss
//}
//
//func (p PriorityQueue) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//
//func (p *PriorityQueue) Push(x any) {
//	item := x.(*QueueEntry)
//	*p = append(*p, item)
//}
//
//func (p *PriorityQueue) Pop() any {
//	old := *p
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil
//	*p = old[0 : n-1]
//	return item
//}
