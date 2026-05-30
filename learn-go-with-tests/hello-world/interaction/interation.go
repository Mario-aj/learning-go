package interaction

func Repeat(char string, quantity int) string {
	var repeat string 

	for i:= 0; i < quantity; i++ {
		repeat = repeat + char
	}

	return repeat
}