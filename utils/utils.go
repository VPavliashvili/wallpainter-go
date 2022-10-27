package utils

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Flattern(arr2 [][]string) (flatterned []string) {
    for _, arr := range arr2 {
        flatterned = append(flatterned, arr...)
    }
    return flatterned
}
