package selection_sort

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Sortable interface {
	Number | string
}

func Sort[s Sortable](values []s, order int) {
	for i := range values {
		// if you choose order equal 1, this variable is store the smaller number
		// if you choose order equal -1, this variable is store the bigger number
		positionOfRightLetter := i

		for j := i + 1; j < len(values); j++ {
			if order == 1 && values[j] < values[positionOfRightLetter] || order == -1 && values[j] > values[positionOfRightLetter] {
				positionOfRightLetter = j
			}
		}

		values[i], values[positionOfRightLetter] = values[positionOfRightLetter], values[i]
	}
}
