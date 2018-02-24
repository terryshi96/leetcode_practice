package main

func partition(a []int, l, h int) int {
	p := a[h]
	// 对数组循环
	for j := l; j < h; j++ {
		// 如果比p小就自己和比p大的数交换
		if a[j] < p {
			a[j], a[l] = a[l], a[j]
			l++
		}
	}
	// 循环结束在l之前的数都比p小，l和l之后的数都比p大
	a[l], a[h] = a[h], a[l]
	// l和h 交换，此时l之前的数都比p小，l之后的数都比p大
	return l
}

func quickSort(a []int, l, h int) {
	if l > h {
		return
	}
	p := partition(a, l, h)
	quickSort(a, l, p-1)
	quickSort(a, p+1, h)
}
