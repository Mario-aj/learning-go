package interaction


const repeatQuantity = 5;

func Repeat(char string) string {
	var repeat string 

	for i:= 0; i < repeatQuantity; i++ {
		repeat = repeat + char
	}

	return repeat
}