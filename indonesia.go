package main


func main() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 50; j++ {
			print("\033[31m*")
		}
		println("*")
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 50; j++ {
			print("\033[0m*")
		}
		println("*")
	}
}