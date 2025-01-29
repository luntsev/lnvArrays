package lnvArrays

// Структура элемента массива, где priority - его приоритет, а object - любой объект
type Item struct {
	Priority int
	Object   interface{}
}

// Массив элементов, без объявления этого типа к нему не получится привязать методы
type Array []Item

// Метод пузырьковой сортировки по возрастанию
// Временная сложность - O(n^2). пространственная сложность - O(1)
// Внешний цикл выполняет количество проходов строго равное длинне сортируемого массива
// На каждой итерации внешнего цикла перебираются все кроме последних
// i элементов массива, где i - текущий номер прохода внешнего массива.
// На каждом проходе внутреннего массива значение текущего элемента сравнивается.
// Если их значения не соответствуют порядку сортировки они меняются местами.
func (slice *Array) AscBubbleSort() {
	for i := range *slice {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].Priority > (*slice)[j+1].Priority {
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
			if (*slice)[j].Priority > (*slice)[j+1].Priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}

// Метод сортировки выбором по возрастанию выбором
// Временная сложность - O(n^2). пространственная сложность - O(1)
// Внешний цикл выполняет количество проходов строго равное длинне сортируемого массива
// На каждой итерации внешнего цикла перебираются все кроме i первых отсортированных элементов
// В ходе перебора определяется наименьший из неотсортированных элементов его позиция меняется с
// очередным неотсортированным элементом.
func (slice *Array) SelectionAscSort() {
	for i, val := range *slice {
		minIndex := i
		for j := i + 1; j < len(*slice); j++ {
			if val.Priority < (*slice)[minIndex].Priority {
				minIndex = j
			}
			(*slice)[i], (*slice)[minIndex] = (*slice)[minIndex], (*slice)[i]
		}
	}
}

// Метод сортировки выбором по убыванию выбором
// Все аналогично предыдущему методу
func (slice *Array) SelectionDescSort() {
	for i, val := range *slice {
		minIndex := i
		for j := i + 1; j < len(*slice); j++ {
			if val.Priority > (*slice)[minIndex].Priority {
				minIndex = j
			}
			(*slice)[i], (*slice)[minIndex] = (*slice)[minIndex], (*slice)[i]
		}
	}
}

// Метод сортировки вставкой по возрастанию
// Временная сложность - O(n^2) или O(n) - если уже отсортирован. пространственная сложность - O(1)
func (slice *Array) InsertAscSort() {
	for i, val := range *slice {
		j := i - 1
		for j >= 0 && (*slice)[j].Priority > val.Priority {
			(*slice)[j+1] = (*slice)[j]
			j--
		}
		(*slice)[j+1] = val
	}
}

// Метод сортировки вставкой по убыванию
// Все аналогично предыдущему методу
func (slice *Array) InsertDescSort() {
	for i, val := range *slice {
		j := i - 1
		for j >= 0 && (*slice)[j].Priority < val.Priority {
			(*slice)[j+1] = (*slice)[j]
			j--
		}
		(*slice)[j+1] = val
	}
}

// Функция сортировки слиянием по возрастанию
// Так как данный алгоритм предусматривает рекурсивные вызовы реализовать его в виде метода
// не представляется возможным, разве что написать отдельную рекурсивную функцию, вызываемую из метода.
// Срез делится на 2 части и вызывается функция слияния в которой реккурсивно вызывается функция
// сортировки но уже для каждой половинки исходного массива. Деление происходит до тех пор пока в разделенных
// массивах не останется по 1 элементу, так как массив с 1 элементом сам по себе априори отсортирован
func MergeAscSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		slice = mergeAscArrays(MergeAscSort(&left), MergeAscSort(&right))
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
		if j == len(*right) || (i < len(*left) && (*left)[i].Priority <= (*right)[j].Priority) {
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
func MergeDescSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		slice = mergeDescArrays(MergeDescSort(&left), MergeDescSort(&right))
	}
	return slice
}

// Внутренняя функция слияния 2- массивов в порядке убывания элементов
func mergeDescArrays(left, right *[]Item) *[]Item {
	var i, j int
	slice := make([]Item, len(*left)+len(*right))

	for k, _ := range slice {
		if j == len(*right) || (i < len(*left) && (*left)[i].Priority >= (*right)[j].Priority) {
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
// Временная сложность в среднем - O(n Log n ), в худшем случае - O(n^2), пространственная сложность O(log n)
func QuickAscSort(slice *[]Item) *[]Item {
	if len(*slice) < 2 {
		return slice
	}

	pivot := 0
	var less, greater []Item

	for _, val := range (*slice)[1:] {
		if val.Priority <= (*slice)[pivot].Priority {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	*slice = append(*QuickAscSort(&less), (*slice)[pivot])
	*slice = append(*slice, *QuickAscSort(&greater)...)
	return slice
}
