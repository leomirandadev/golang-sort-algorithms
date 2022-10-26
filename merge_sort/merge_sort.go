package merge_sort

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Sortable interface {
	Number | string
}

func Sort[s Sortable](values []s, order int) {
	sizeArray := len(values)

	if sizeArray < 2 {
		return
	}

	sizeFirstArray := int(sizeArray / 2)

	firstArray := make([]s, sizeFirstArray)
	for i := 0; i < sizeFirstArray; i++ {
		firstArray[i] = values[i]
	}

	sizeSecondArray := sizeArray - sizeFirstArray
	secondArray := make([]s, sizeSecondArray)
	for i := 0; i < sizeSecondArray; i++ {
		secondArray[i] = values[i+sizeFirstArray]
	}

	Sort(firstArray, order)
	Sort(secondArray, order)

	result := merge(firstArray, secondArray, order)

	copy(values, result)

	return
}

func merge[s Sortable](arrayA, arrayB []s, order int) []s {
	result := make([]s, len(arrayA)+len(arrayB))

	indexFirst := 0
	indexSecond := 0

	for i := range result {
		if indexFirst == len(arrayA) {
			result[i] = arrayB[indexSecond]
			indexSecond++
			continue
		}

		if indexSecond == len(arrayB) {
			result[i] = arrayA[indexFirst]
			indexFirst++
			continue
		}

		if order == 1 && arrayA[indexFirst] < arrayB[indexSecond] || order == -1 && arrayA[indexFirst] > arrayB[indexSecond] {
			result[i] = arrayA[indexFirst]
			indexFirst++
		} else {
			result[i] = arrayB[indexSecond]
			indexSecond++
		}
	}

	return result
}
