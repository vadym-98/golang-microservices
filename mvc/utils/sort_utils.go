package utils

import "sort"

func BubbleSort(elements []int) {
	isNeedContinue := true
	for isNeedContinue {
		isNeedContinue = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				isNeedContinue = true
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)

		return
	}

	sort.Ints(elements)
}
