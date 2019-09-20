package main

func bubbleSort(values []int) []int {
	len := len(values)
	index := len - 1
	for i := 0; i < index; i++ {
		swap := false
		for j := 0; j < index-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swap = true
			}
		}

		if swap == false {
			break
		}
	}

	return values
}

func chooseSort(values []int) []int {
	len := len(values)
	index := len - 1
	for i := 0; i < index; i++ {
		for j := i + 1; j < len; j++ {
			if values[j] < values[i] {
				values[j], values[i] = values[i], values[j]
			}
		}
	}
	return values
}

func insertSort(valuse []int) []int {
	var l, tmp, i, j int
	l = len(valuse)
	for i = 1; i < l; i++ {
		tmp = valuse[i]
		for j = i - 1; j >= 0; j -= 1 {
			if valuse[j] > tmp {
				valuse[j+1] = valuse[j]
				valuse[j] = tmp
			} else {
				break
			}
		}
	}
	return valuse
}

func shellSort(s []int) []int {
	var gap, l, tmp, i, j int
	l = len(s)

	for gap = l / 2; gap > 0; gap /= 2 {
		for i = gap; i < l; i++ {
			tmp = s[i]
			for j = i - gap; j >= 0; j -= gap {
				if s[j] > tmp {
					s[j+gap] = s[j]
					s[j] = tmp
				} else {
					break
				}
			}
		}
	}
	return s
}

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}

	// 实现了“quickSort(left) + v + quickSort(right)”的操作
	return append(append(quickSort(left), v), quickSort(right)...)
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	//递[归]
	middle := len(data) / 2
	//不断地进行左右对半划分
	left := mergeSort(data[:middle])
	right := mergeSort(data[middle:])
	//合[并]
	return merge(left, right)
}

func bucketSort(s []int) []int {
	max := s[0]
	for _, value := range s {
		if value > max {
			max = value
		}
	}
	bucket := make([]int, max)
	for _, value := range s {
		bucket[value]++
	}

}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	// 注意：[左右]对比，是指左的第一个元素，与右边的第一个元素进行对比，哪个小，就先放到结果的第一位，然后左或右取出了元素的那边的索引进行++
	for l < len(left) && r < len(right) {
		//从小到大排序.
		if left[l] > right[r] {
			result = append(result, right[r])
			//因为处理了右边的第r个元素，所以r的指针要向前移动一个单位
			r++
		} else {
			result = append(result, left[l])
			//因为处理了左边的第r个元素，所以r的指针要向前移动一个单位
			l++
		}
	}
	// 比较完后，还要分别将左，右的剩余的元素，追加到结果列的后面(不然就漏咯）。
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
