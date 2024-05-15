package main

func main() {
	for i := 1; i < 12; i++ {
		for j := 0; j < 50; j++ {
			if i == 1 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[30m*")
				}
			}else if i == 2 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[30m*")
				}
			}else if i == 3 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[30m*")
				}
			}else if i == 4 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[30m*")
				}
			}else if i == 5 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[97m*")
				}
			}else if i == 6 {
				if j == 0 {
					for k := 0; k < (i+i); k++ {
						print("\033[31m*")
					}
					j = (i*2)-1
				} else {
					print("\033[97m*")
				}
			}else if i == 7 {
				if j == 0 {
					for k := 10; k > 0; k-- {
						print("\033[31m*")
					}
					j = 9
				} else {
					print("\033[97m*")
				}
			}else if i == 8 {
				if j == 0 {
					for k := 8; k > 0; k-- {
						print("\033[31m*")
					}
					j = 7
				} else {
					print("\033[32m*")
				}
			}else if i == 9 {
				if j == 0 {
					for k := 6; k > 0; k-- {
						print("\033[31m*")
					}
					j = 5
				} else {
					print("\033[32m*")
				}
			}else if i == 10 {
				if j == 0 {
					for k := 4; k > 0; k-- {
						print("\033[31m*")
					}
					j = 3
				} else {
					print("\033[32m*")
				}
			} else {
				if j == 0 {
					for k := 2; k > 0; k-- {
						print("\033[31m*")
					}
					j = 1
				} else {
					print("\033[32m*")
				}
			}
		}
		println("\033[0m")
	}
}
