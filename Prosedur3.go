package main

import "fmt"

// Prosedur untuk mencetak deret bilangan
func cetakDeret(n_2204 int) {
	for n_2204 != 1 {
		fmt.Print(n_2204, " ")
		if n_2204%2 == 0 {
			n_22 = n_2204 / 2
		} else {
			n_2204 = 3*n_2204 + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
