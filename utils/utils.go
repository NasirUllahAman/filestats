package utils

func CharCount(data []byte) int {
	return len(data)
}
func WordCount(data []byte) int {
	count := 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == " " || i == len(data)-1 {
			count++
		}
	}
	return count
}
