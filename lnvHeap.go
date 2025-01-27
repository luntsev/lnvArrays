package lnvArrays

import (
	"errors"
)

type HeapItem struct {
	Priority int
	object   interface{}
}

type MaxHeap []HeapItem

func (pq *MaxHeap) Len() int {
	return len(*pq)
}

func (pq *MaxHeap) Push(elem HeapItem) {
	*pq = append(*pq, elem)
	if pq.Len() > 1 {
		pq.embed(pq.Len() - 1)
	}
}

func (pq *MaxHeap) Pop() (HeapItem, error) {
	if pq.Len() > 0 {
		rootElem := (*pq)[0]
		pq.swap(0, pq.Len()-1)
		*pq = (*pq)[:pq.Len()-1]
		if pq.Len() > 1 {
			pq.rebuild(0)
		}
		return rootElem, nil
	} else {
		var nullElement HeapItem
		err := errors.New("The heap is empty, there is nothing to return")
		return nullElement, err
	}
}

func (pq *MaxHeap) Get() HeapItem {
	rootElem := (*pq)[0]
	return rootElem
}

func (pq *MaxHeap) swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *MaxHeap) embed(itemId int) {
	if itemId > 0 {
		if (itemId)%2 != 0 {
			if (*pq)[itemId].Priority > (*pq)[(itemId-1)/2].Priority {
				pq.swap(itemId, (itemId-1)/2)
				pq.embed((itemId - 1) / 2)
			}
		} else {
			if (*pq)[itemId].Priority > (*pq)[(itemId-2)/2].Priority {
				pq.swap(itemId, (itemId-2)/2)
				pq.embed((itemId - 2) / 2)
			}
		}
	}
}

func (pq *MaxHeap) rebuild(itemId int) {
	switch {
	case (itemId*2)+1 == pq.Len()-1:
		leftId := (itemId * 2) + 1
		if (*pq)[leftId].Priority > (*pq)[itemId].Priority {
			pq.swap(itemId, leftId)
		}
	case (itemId*2)+2 == pq.Len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].Priority > (*pq)[rightId].Priority {
			if (*pq)[leftId].Priority > (*pq)[itemId].Priority {
				pq.swap(itemId, leftId)
			}
		} else {
			if (*pq)[rightId].Priority > (*pq)[itemId].Priority {
				pq.swap(itemId, rightId)
			}
		}
	case (itemId*2)+2 < pq.Len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].Priority > (*pq)[rightId].Priority {
			if (*pq)[leftId].Priority > (*pq)[itemId].Priority {
				pq.swap(itemId, leftId)
				pq.rebuild(leftId)
			}
		} else {
			if (*pq)[rightId].Priority > (*pq)[itemId].Priority {
				pq.swap(itemId, rightId)
				pq.rebuild(rightId)
			}
		}
	}
}

type MinHeap []HeapItem

func (pq *MinHeap) Len() int {
	return len(*pq)
}

func (pq *MinHeap) Push(elem HeapItem) {
	*pq = append(*pq, elem)
	if pq.Len() > 1 {
		pq.embed(pq.Len() - 1)
	}
}

func (pq *MinHeap) Pop() (HeapItem, error) {
	if pq.Len() > 0 {
		rootElem := (*pq)[0]
		pq.swap(0, pq.Len()-1)
		*pq = (*pq)[:pq.Len()-1]
		if pq.Len() > 1 {
			pq.rebuild(0)
		}
		return rootElem, nil
	} else {
		var nullElement HeapItem
		err := errors.New("The heap is empty, there is nothing to return")
		return nullElement, err
	}
}

func (pq *MinHeap) Get() HeapItem {
	rootElem := (*pq)[0]
	return rootElem
}

func (pq *MinHeap) swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *MinHeap) embed(itemId int) {
	if itemId > 0 {
		if (itemId)%2 != 0 {
			if (*pq)[itemId].Priority < (*pq)[(itemId-1)/2].Priority {
				pq.swap(itemId, (itemId-1)/2)
				pq.embed((itemId - 1) / 2)
			}
		} else {
			if (*pq)[itemId].Priority < (*pq)[(itemId-2)/2].Priority {
				pq.swap(itemId, (itemId-2)/2)
				pq.embed((itemId - 2) / 2)
			}
		}
	}
}

func (pq *MinHeap) rebuild(itemId int) {
	switch {
	case (itemId*2)+1 == pq.Len()-1:
		leftId := (itemId * 2) + 1
		if (*pq)[leftId].Priority < (*pq)[itemId].Priority {
			pq.swap(itemId, leftId)
		}
	case (itemId*2)+2 == pq.Len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].Priority < (*pq)[rightId].Priority {
			if (*pq)[leftId].Priority < (*pq)[itemId].Priority {
				pq.swap(itemId, leftId)
			}
		} else {
			if (*pq)[rightId].Priority < (*pq)[itemId].Priority {
				pq.swap(itemId, rightId)
			}
		}
	case (itemId*2)+2 < pq.Len()-1:
		leftId := (itemId * 2) + 1
		rightId := (itemId * 2) + 2
		if (*pq)[leftId].Priority < (*pq)[rightId].Priority {
			if (*pq)[leftId].Priority < (*pq)[itemId].Priority {
				pq.swap(itemId, leftId)
				pq.rebuild(leftId)
			}
		} else {
			if (*pq)[rightId].Priority < (*pq)[itemId].Priority {
				pq.swap(itemId, rightId)
				pq.rebuild(rightId)
			}
		}
	}
}
