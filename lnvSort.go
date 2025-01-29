package lnvArrays

// Структура элемента массива, где priority - его приоритет, а object - любой объект
type Item struct {
	priority int
	object   interface{}
}

// Массив элементов, без объявления этого типа к нему не получится привязать методы
type Array []Item

// Метод пузырьковой сортировки по возрастанию
// Внешний цикл количество проходов строго равное длинне сортируемого массива
// На каждой итерации внешнего цикра перебирается все кроме последних
// i элементов массива, где i - текущий номер прохода внешнего массива.
// На каждом проходе внутреннего массива значение текущего элемента сравнивается.
// Если их значения не соответствуют порядку сортировки они меняются местами.
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
// Все аналогично предыдущему методу
func (slice *Array) DescBubbleSort() {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].priority > (*slice)[j+1].priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}

// Функция сортировки слиянием по возрастанию
// Так как данный алгоритм предусматривает рекурсивные вызовы реализовать его в виде метода
// не представляется возможным, разве что написать отдельную рекурсивную функцию, вызываемую из метода.
// Срез делится на 2 части и вызывается функция слияния в которой реккурсивно вызывается функция
// сортировки но уже для каждой половинки исходного массива. Деление происходит до тех пор пока в разделенных
// массивах не останется по 1 элементу, так как массив с 1 элементом сам по себе априори отсортирован
func AscMergeSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		slice = mergeAscArrays(AscMergeSort(&left), AscMergeSort(&right))
	}
	return slice
}

// Внутренняя функция слияния 2- массивов в порядке возрастания элементов
// Сливает элементу двух массивов в один беря текущие наименьшие из каждого массива так как сливаемые
// массивы сами по себе отсортированны
func mergeAscArrays(left, right *[]Item) *[]Item {
	var i, j int
	slice := make([]Item, len(*left)+len(*right))

	for k, _ := range slice {
		if j == len(*right) || (i < len(*left) && (*left)[i].priority <= (*right)[j].priority) {
			slice[k] = (*left)[i]
			i++
		} else {
			slice[k] = (*right)[j]
			j++
		}
	}
	return &slice
}

// Функция сортировки слиянием по убыванию
func DescMergeSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		slice = mergeDescArrays(DescMergeSort(&left), DescMergeSort(&right))
	}
	return slice
}

// Внутренняя функция слияния 2- массивов в порядке убывания элементов
func mergeDescArrays(left, right *[]Item) *[]Item {
	var i, j int
	slice := make([]Item, len(*left)+len(*right))

	for k, _ := range slice {
		if j == len(*right) || (i < len(*left) && (*left)[i].priority >= (*right)[j].priority) {
			slice[k] = (*left)[i]
			i++
		} else {
			slice[k] = (*right)[j]
			j++
		}
	}
	return &slice
}

// Функция быстрой сортировки
func quickSort(slice *[]Item) *[]Item {
	if len(*slice) < 2 {
		return slice
	}

	pivot := 0
	var less, greater []Item

	for _, val := range (*slice)[1:] {
		if val.priority <= (*slice)[pivot].priority {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	*slice = append(*quickSort(&less), (*slice)[pivot])
	*slice = append(*slice, *quickSort(&greater)...)
	return slice
}
