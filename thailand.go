package main



func main() {
	for i := 0; i < 11; i++ {
		for j := 0; j < 50; j++ {
			if i == 0 || i == 1 || i == 9 ||i == 10 {
				print("\033[31m*")
			} else if i == 4|| i == 5 || i == 6 {
				print("\033[34m*")
			}else{
			print("\033[0m*")
			}
		}
		println()
	}
	print("\033[0m")
}