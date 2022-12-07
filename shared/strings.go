package shared

func newRepeatingString(length int, value string) string {

	result := ""
	for i := 0; i <= length; i++ {
		result = result + value
	}
	return result
}
