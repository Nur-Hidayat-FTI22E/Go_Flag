package main

func main() {
	for i := 0; i < 15; i++ {
		for j := 0; j < 61; j++ {
			if i == 3 {
				if j == 31 {
					print("\033[33m*")
				} else {
					print("\033[31m*")
				}
			} else if i == 4 {
				if j == 30 {
					for k := j; k < 33; k++ {
						print("\033[33m*")
					}
					j = 32
				} else {
					print("\033[31m*")
				}

			} else if i == 5 {
				if j == 29 {
					for k := j; k < 34; k++ {
						print("\033[33m*")
					}
					j = 33
				} else {
					print("\033[31m*")
				}

			} else if i == 6 {
				if j == 23 {
					for k := j; k < 40; k++ {
						print("\033[33m*")
					}
					j = 39
				} else {
					print("\033[31m*")
				}

			} else if i == 7 {
				if j == 25 {
					for k := j; k < 38; k++ {
						print("\033[33m*")
					}
					j = 37
				} else {
					print("\033[31m*")
				}

			} else if i == 8 {
				if j == 27 {
					for k := j; k < 36; k++ {
						print("\033[33m*")
					}
					j = 35
				} else {
					print("\033[31m*")
				}

			} else if i == 9 {
				if j == 26 {
					for k := j; k < 37; k++ {
						print("\033[33m*")
					}
					j = 36
				} else {
					print("\033[31m*")
				}

			} else if i == 10 {
				if j == 25  {
					for k := 0; k < 4; k++ {
						print("\033[33m*")
						j = 28
					}
				}else if j == 34 {
					for k := j; k < 38; k++ {
						print("\033[33m*")
					}
					j = 37
				} else {
					print("\033[31m*")
				}

			}else if i == 11 {
				if j == 24  {
					for k := 0; k < 2; k++ {
						print("\033[33m*")
						j = 26
					}
				}else if j == 38 {
					for k := j; k < 40; k++ {
						print("\033[33m*")
					}
					j = 38
				} else {
					print("\033[31m*")
				}

			} else {
				print("\033[31m*")
			}
		}
		println()
	}
	print("\033[0m")
}
