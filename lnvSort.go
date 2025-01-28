package lnvArrays

type ArrayItem struct {
	priority int
	object   interface{}
}

type Array []ArrayItem

func BubbleAscSort(slice *Array) *Array {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].priority > (*slice)[j+1].priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
	return slice
}

func BubbleDescSort(slice *Array) *Array {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].priority > (*slice)[j+1].priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
	return slice
}
