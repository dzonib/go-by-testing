package iteration

func main() {

}

func Repeat(s string, repeatCount int) string {

	result := ""

	for i := 0; i < repeatCount; i++ {
		result += s
	}
	return result
}
