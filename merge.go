package main

func Merge(collection1, collection2, collection3 []int) []int {
	n := len(collection1) + len(collection2) + len(collection3)
	result := make([]int, 0, n)

	i := len(collection1) - 1
	j := 0
	k := 0

	for i >= 0 || j < len(collection2) || k < len(collection3) {
		switch {
		case i >= 0 && (j >= len(collection2) || collection1[i] <= collection2[j]) &&
			(k >= len(collection3) || collection1[i] <= collection3[k]):
			result = append(result, collection1[i])
			i--
		case j < len(collection2) && (k >= len(collection3) || collection2[j] <= collection3[k]):
			result = append(result, collection2[j])
			j++
		default:
			result = append(result, collection3[k])
			k++
		}
	}

	return result
}
