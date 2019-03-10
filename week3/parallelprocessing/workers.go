package main

import "container/heap"

type worker struct {
	id          int
	releaseTime int
	index       int
}

type priorityQueue []*worker

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].releaseTime == pq[j].releaseTime {
		return pq[i].id < pq[j].id
	}
	return pq[i].releaseTime < pq[j].releaseTime
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	worker := x.(*worker)
	worker.index = n
	*pq = append(*pq, worker)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	worker := old[n-1]
	worker.index = -1
	*pq = old[0 : n-1]
	return worker
}

func Solve(numWorkers int, jobs []int) (result [][2]int) {
	pq := initWorkers(numWorkers)
	for _, job := range jobs {
		w := heap.Pop(pq).(*worker)
		result = append(result, [2]int{w.id, w.releaseTime})
		w.releaseTime += job
		heap.Push(pq, w)
	}

	return result
}

func initWorkers(numWorkers int) *priorityQueue {
	var pq = make(priorityQueue, numWorkers)
	for i := 0; i < numWorkers; i++ {
		w := worker{id: i, releaseTime: 0, index: i}
		pq[i] = &w
	}
	heap.Init(&pq)
	return &pq
}
