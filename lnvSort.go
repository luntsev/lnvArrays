package lnvArrays

// Структура элемента массива, где priority - его приоритет, а object - любой объект
type ArrayItem struct {
	priority int
	object   interface{}
}

// Массив элементов
type Array []ArrayItem

// Метод пузырьковой сортировки по возрастанию
func (slice *Array) AscBubbleSort() {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].priority > (*slice)[j+1].priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}

// Метод пузырьковой сортировки по убыванию
func (slice *Array) DescBubbleSort() {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].priority > (*slice)[j+1].priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}

// Метод сортировки слиянием по возрастанию
func (slice *Array) AscMergeSort() {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		slice = mergeArray(mergeSort(&left), mergeSort(&right))
	}
}

func mergeAscArrays(left, right *Array) {

}
