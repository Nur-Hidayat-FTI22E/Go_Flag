package main

func main() {
	for i := 0; i < 12 ; i ++ {
		for j := 0; j < 50; j++ {
			if i == 1 || i == 10 {
				print("*")
			} else if i == 2 || i == 9 {
				if j == 24 {
					for k := j; k < 27; k++ {
						print("\033[31m*\033[0m")
					}
					j = 26
				} else {
					print("*")
				}
			} else if i == 3 || i == 8 {
				if j == 22 {
					for k := j; k < 29; k++ {
						print("\033[31m*\033[0m")
					}
					j = 28
				} else {
					print("*")
				}
			} else if i == 4 || i == 7 {
				if j == 20 {
					for k := j; k < 31; k++ {
						print("\033[31m*\033[0m")
					}
					j = 30
				} else {
					print("*")
				}
			} else if i == 5 || i == 6 {
				if j == 18 {
					for k := j; k < 33; k++ {
						print("\033[31m*\033[0m")
					}
					j = 32
				} else {
					print("*")
				}
			} else {
				print("*")
			}
		}
		println("*")
	}
}