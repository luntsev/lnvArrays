package lnvArrays

import (
	"errors"
)

type heapItem struct {
	priority int
	object   interface{}
}

type maxHeap []heapItem
type minHeap []heapItem

func (pq *maxHeap) init() {
	var nullItem heapItem
	*pq = append((*pq)[:0], nullItem)
}

func (pq *maxHeap) push(elem heapItem) {
	*pq = append(*pq, elem)
	if pq.len() > 1 {
		pq.embed(pq.len() - 1)
	}
}

func (pq *maxHeap) pop() (heapItem, error) {
	if pq.len() > 0 {
		rootElem := (*pq)[0]
		pq.swap(0, pq.len()-1)
		*pq = (*pq)[:pq.len()-1]
		if pq.len() > 1 {
			pq.rebuild(0)
		}
		return rootElem, nil
	} else {
		var nullElement heapItem
		err := errors.New("The heap is empty, there is nothing to return")
		return nullElement, err
	}
}

func (pq *maxHeap) get() heapItem {
	rootElem := (*pq)[0]
	return rootElem
}

func (pq *maxHeap) swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *maxHeap) len() int {
	return len(*pq)
}

func (pq *maxHeap) embed(itemId int) {
	if itemId > 0 {
		if (itemId)%2 != 0 {
			if (*pq)[itemId].priority > (*pq)[(itemId-1)/2].priority {
				pq.swap(itemId, (itemId-1)/2)
				pq.embed((itemId - 1) / 2)
			}
		} else {
			if (*pq)[itemId].priority > (*pq)[(itemId-2)/2].priority {
				pq.swap(itemId, (itemId-2)/2)
				pq.embed((itemId - 2) / 2)
			}
		}
	}
}

func (pq *maxHeap) rebuild(itemId int) {
	switch {
	case (itemId*2)+1 == pq.len()-1:
		leftId := (itemId * 2) + 1
		if (*pq)[leftId].priority > (*pq)[itemId].priority {
			pq.swap(itemId, leftId)
		}
	case (itemId*2)+2 == pq.len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].priority > (*pq)[rightId].priority {
			if (*pq)[leftId].priority > (*pq)[itemId].priority {
				pq.swap(itemId, leftId)
			}
		} else {
			if (*pq)[rightId].priority > (*pq)[itemId].priority {
				pq.swap(itemId, rightId)
			}
		}
	case (itemId*2)+2 < pq.len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].priority > (*pq)[rightId].priority {
			if (*pq)[leftId].priority > (*pq)[itemId].priority {
				pq.swap(itemId, leftId)
				pq.rebuild(leftId)
			}
		} else {
			if (*pq)[rightId].priority > (*pq)[itemId].priority {
				pq.swap(itemId, rightId)
				pq.rebuild(rightId)
			}
		}
	}
}
