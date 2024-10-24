package main

import "fmt"

// Menghitung faktorial
func factorial(n_204 int) int {
	if n_204 == 0 {
		return 1
	}
	return n_204 * factorial(n_204-1)
}

// Menghitung permutasi
func permutation(n_204, r_204 int) int {
	return factorial(n_204) / factorial(n_204-r_204)
}

// Menghitung kombinasi
func combination(n_204, r_204 int) int {
	return factorial(n_204) / (factorial(r_204) * factorial(n_204-r_204))
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Permutasi dan kombinasi untuk a terhadap c
	p_a_c := permutation(a, c)
	c_a_c := combination(a, c)
	fmt.Printf("P(%d, %d) = %d\nC(%d, %d) = %d\n", a, c, p_a_c, a, c, c_a_c)

	// Permutasi dan kombinasi untuk b terhadap d
	p_b_d := permutation(b, d)
	c_b_d := combination(b, d)
	fmt.Printf("P(%d, %d) = %d\nC(%d, %d) = %d\n", b, d, p_b_d, b, d, c_b_d)
}
