package pq

import (
	"container/heap"
	"sync"
)

type Prioritisable interface {
	IsHigherPriority(than any) bool
}

type priorityQueue[T Prioritisable] []T

func (p priorityQueue[T]) Len() int {
	return len(p)
}

func (p priorityQueue[T]) Less(i, j int) bool {
	return p[i].IsHigherPriority(p[j])
}

func (p priorityQueue[T]) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue[T]) Push(x any) {
	item := x.(T)
	*p = append(*p, item)
}

func (p *priorityQueue[T]) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

type PriorityQueue[T Prioritisable] struct {
	q  priorityQueue[T]
	mx sync.Mutex
}

func NewPriorityQueue[T Prioritisable](items ...T) PriorityQueue[T] {
	var pq priorityQueue[T]
	pq = append(pq, items...)
	heap.Init(&pq)
	return PriorityQueue[T]{
		q:  pq,
		mx: sync.Mutex{},
	}
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.q.Len()
}

func (pq *PriorityQueue[T]) Push(item T) {
	pq.mx.Lock()
	defer pq.mx.Unlock()
	heap.Push(&pq.q, item)
}

func (pq *PriorityQueue[T]) Pop() T {
	pq.mx.Lock()
	defer pq.mx.Unlock()
	val := heap.Pop(&pq.q)
	return val.(T)
}
