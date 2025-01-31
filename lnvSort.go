package lnvArrays

// Структура элемента массива, где priority - его приоритет, а object - любой объект
type Item struct {
	Priority int
	Object   interface{}
}

// Массив элементов, без объявления этого типа к нему не получится привязать методы
type Array []Item

// Метод пузырьковой сортировки по возрастанию
// Временная сложность - O(n^2) (В самом лучшем случае O(1) - если не отсортирован). пространственная сложность - O(1)
// Внешний цикл выполняет количество проходов строго равное длине сортируемого массива
// На каждой итерации внешнего цикла перебираются все кроме последних
// i элементов массива, где i - текущий номер прохода внешнего массива.
// На каждом проходе внутреннего массива значение текущего элемента сравнивается.
// Если их значения не соответствуют порядку сортировки они меняются местами.
func (slice *Array) AscBubbleSort() {
	for i := range *slice {
		swap := false
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j].Priority > (*slice)[j+1].Priority {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
				swap = true
			}
		}
		if !swap {
			break
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
// Внешний цикл выполняет количество проходов строго равное длине сортируемого массива
// На каждой итерации внешнего цикла перебираются все кроме i первых отсортированных элементов
// В ходе перебора определяется наименьший из неотсортированных элементов его позиция меняется с
// очередным неотсортированным элементом.
func (slice *Array) AscSelectionSort() {
	for i, _ := range *slice {
		minIndex := i
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[minIndex].Priority > (*slice)[j].Priority {
				minIndex = j
			}
		}
		(*slice)[i], (*slice)[minIndex] = (*slice)[minIndex], (*slice)[i]
	}
}

// Метод сортировки выбором по убыванию выбором
// Все аналогично предыдущему методу
func (slice *Array) DesсSelectionSort() {
	for i, _ := range *slice {
		maxIndex := i
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[maxIndex].Priority < (*slice)[j].Priority {
				maxIndex = j
			}
		}
		(*slice)[i], (*slice)[maxIndex] = (*slice)[maxIndex], (*slice)[i]
	}
}

// Метод сортировки вставкой по возрастанию
// Временная сложность - O(n^2) или O(n) - если уже отсортирован. пространственная сложность - O(1)
func (slice *Array) AscInsertSort() {
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
func (slice *Array) DescInsertSort() {
	for i, val := range *slice {
		j := i - 1
		for j >= 0 && (*slice)[j].Priority < val.Priority {
			(*slice)[j+1] = (*slice)[j]
			j--
		}
		(*slice)[j+1] = val
	}
}

// Метод сортировки двоичной кучей по возрастанию
// Временная сложность: O(n log n), пространственная сложность: O(1).
// В кучу с минимальным корнем поочередно добавляются все элементы массива
// После чего из кучи поочередно извлекаются все элементы в массив
func (slice *Array) AscHeapSort() {
	var sortedHeap MinHeap

	for _, val := range *slice {
		sortedHeap.Push(HeapItem(val))
	}

	for i, _ := range *slice {
		heapItem, _ := sortedHeap.Pop()
		(*slice)[i] = Item(heapItem)
	}
}

// Метод сортировки двоичной кучей по убыванию
// Все аналогично предыдущему методу
func (slice *Array) DescHeapSort() {
	var sortedHeap MaxHeap

	for _, val := range *slice {
		sortedHeap.Push(HeapItem(val))
	}

	for i, _ := range *slice {
		heapItem, _ := sortedHeap.Pop()
		(*slice)[i] = Item(heapItem)
	}
}

// Метод сортировки слиянием по возрастанию с использованием Горутин
// Несмотря на то, что данный алгоритм предусматривает рекурсивные вызовы он реализован его в виде метода.
// Для этого определена инициализирующая функция (которая и представляет метод), которая вызывает аналогичную
// себе функцию, которая в свою очередь уже является рекурсивной. Этот вызов происходит 2 раза каждый из которых
// в отдельной горутине.
func (slice *Array) GoAscMergeSort() {
	if len(*slice) < 2 {
		return
	}
	m := len(*slice) / 2
	left := (*slice)[:m]
	right := (*slice)[m:]
	inLeftChan := make(chan Array, 1)
	inRightChan := make(chan Array, 1)
	go goAscMergeSort(&left, inLeftChan)
	go goAscMergeSort(&right, inRightChan)
	var leftArray, rightArray Array
	leftArray = <-inLeftChan
	rightArray = <-inRightChan
	*slice = goMergeAscArrays(&leftArray, &rightArray)
}

// Внутренняя функция для метода сортировки слиянием по возрастанию, представляет из себя клон метода AscMergeSort
// с той разницей, что предусматривает рекурсивные вызови и поэтому поддерживает связь между ними через каналы
func goAscMergeSort(slice *Array, outChannel chan Array) {
	if len(*slice) < 2 {
		outChannel <- *slice
		close(outChannel)
		return
	}
	m := len(*slice) / 2
	left := (*slice)[:m]
	right := (*slice)[m:]
	inLeftChan := make(chan Array, 1)
	inRightChan := make(chan Array, 1)
	go goAscMergeSort(&left, inLeftChan)
	go goAscMergeSort(&right, inRightChan)
	var leftArray, rightArray Array
	leftArray = <-inLeftChan
	rightArray = <-inRightChan
	outChannel <- goMergeAscArrays(&leftArray, &rightArray)
	close(outChannel)
}

// Внутренняя функция слияния 2- массивов в порядке возрастания элементов, для соответствующего метода
func goMergeAscArrays(left, right *Array) Array {
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
	return slice
}

// Метод сортировки слиянием по убыванию с использованием Горутин
// Все аналогично вышеуказанному
func (slice *Array) GoDescMergeSort() {
	if len(*slice) < 2 {
		return
	}
	m := len(*slice) / 2
	left := (*slice)[:m]
	right := (*slice)[m:]
	inLeftChan := make(chan Array)
	inRightChan := make(chan Array)
	go goDescMergeSort(&left, inLeftChan)
	go goDescMergeSort(&right, inRightChan)
	var leftArray, rightArray Array
	leftArray = <-inLeftChan
	rightArray = <-inRightChan
	*slice = goMergeDescArrays(&leftArray, &rightArray)
}

func goDescMergeSort(slice *Array, outChannel chan Array) {
	if len(*slice) < 2 {
		outChannel <- *slice
		close(outChannel)
		return
	}
	m := len(*slice) / 2
	left := (*slice)[:m]
	right := (*slice)[m:]
	inLeftChan := make(chan Array)
	inRightChan := make(chan Array)
	go goDescMergeSort(&left, inLeftChan)
	go goDescMergeSort(&right, inRightChan)
	var leftArray, rightArray Array
	leftArray = <-inLeftChan
	rightArray = <-inRightChan
	outChannel <- goMergeDescArrays(&leftArray, &rightArray)
}

func goMergeDescArrays(left, right *Array) Array {
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
	return slice
}

// Функция сортировки слиянием по возрастанию
// Так как данный алгоритм предусматривает рекурсивные вызовы реализовать его в виде метода
// не представляется возможным, разве что написать отдельную рекурсивную функцию, вызываемую из метода.
// Срез делится на 2 части и вызывается функция слияния в которой рекурсивно вызывается функция
// сортировки но уже для каждой половинки исходного массива. Деление происходит до тех пор пока в разделенных
// массивах не останется по 1 элементу, так как массив с 1 элементом сам по себе априори отсортирован
func AscMergeSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[:m]
		right := (*slice)[m:]
		*slice = mergeAscArrays(AscMergeSort(&left), AscMergeSort(&right))
	}
	return slice
}

// Внутренняя функция слияния 2- массивов в порядке возрастания элементов
// Сливает элементу двух массивов в один беря текущие наименьшие из каждого массива так как сливаемые
// массивы сами по себе отсортированы
func mergeAscArrays(left, right *[]Item) []Item {
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
	return slice
}

// Функция сортировки слиянием по убыванию
func DescMergeSort(slice *[]Item) *[]Item {
	if len(*slice) > 1 {
		m := len(*slice) / 2
		left := (*slice)[0:m]
		right := (*slice)[m:len(*slice)]
		*slice = mergeDescArrays(DescMergeSort(&left), DescMergeSort(&right))
	}
	return slice
}

// Внутренняя функция слияния 2- массивов в порядке убывания элементов
func mergeDescArrays(left, right *[]Item) []Item {
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
	return slice
}

// Функция быстрой сортировки по возрастанию
// Временная сложность в среднем - O(n Log n ), в худшем случае - O(n^2), пространственная сложность O(log n)
func AscQuickSort(slice *[]Item) *[]Item {
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
	*slice = append(*AscQuickSort(&less), (*slice)[pivot])
	*slice = append(*slice, *AscQuickSort(&greater)...)
	return slice
}

func (slice *Array) GoAscQuickSort() {
	if len(*slice) < 2 {
		return
	}

	pivot := 0
	var less, greater Array

	for _, val := range (*slice)[1:] {
		if val.Priority <= (*slice)[pivot].Priority {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	lessChan := make(chan Array, 2)
	greaterChan := make(chan Array, 2)
	go goAscQuickSort(&less, lessChan)
	go goAscQuickSort(&greater, greaterChan)
	*slice = append(<-lessChan, (*slice)[pivot])
	*slice = append(*slice, <-greaterChan...)
}

func goAscQuickSort(slice *Array, outChan chan Array) {
	if len(*slice) < 2 {
		return
		close(outChan)
	}

	pivot := 0
	var less, greater Array

	for _, val := range (*slice)[1:] {
		if val.Priority <= (*slice)[pivot].Priority {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	lessChan := make(chan Array, 2)
	greaterChan := make(chan Array, 2)
	go goAscQuickSort(&less, lessChan)
	go goAscQuickSort(&greater, greaterChan)
	returnedSlice := append(<-lessChan, (*slice)[pivot])
	close(lessChan)
	returnedSlice = append(*slice, <-greaterChan...)
	close(greaterChan)
	outChan <- returnedSlice
}

// Функция быстрой сортировки по убыванию
// Временная сложность в среднем - O(n Log n ), в худшем случае - O(n^2), пространственная сложность O(log n)
func DescQuickSort(slice *[]Item) *[]Item {
	if len(*slice) < 2 {
		return slice
	}

	pivot := 0
	var less, greater []Item

	for _, val := range (*slice)[1:] {
		if val.Priority >= (*slice)[pivot].Priority {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}
	*slice = append(*DescQuickSort(&less), (*slice)[pivot])
	*slice = append(*slice, *DescQuickSort(&greater)...)
	return slice
}
