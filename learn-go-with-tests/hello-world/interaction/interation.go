package interaction

func Repeat(char string) string {
	var repeat string 

	for i:= 0; i < 5; i++ {
		repeat = repeat + char
	}

	return repeat
}