package utils

func CompareSlice(slice1, slice2 []string) bool {
	i := 0
	for i < len(slice1) {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
