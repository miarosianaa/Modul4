package main

import "fmt"

func hitungScore(peserta_204 []int) (int, int) {
	soal := 0
	skor := 0
	for _, waktu := range peserta_204 {
		if waktu <= 301 {
			soal++
			skor += waktu
		}
	}
	return soal, skor
}

func main() {
	var nama, pemenang string
	var peserta [][]int
	var maxSoal204, minSkor204 int

	for {
		fmt.Print("Masukkan nama peserta ('selesai' untuk berhenti): ")
		fmt.Scan(&nama)
		if nama == "selesai" {
			break
		}

		var waktu [8]int
		fmt.Println("Masukkan waktu pengerjaan soal (dalam menit):")
		for i := 0; i < 8; i++ {
			fmt.Scanf("%d", &waktu[i])
		}

		peserta = append(peserta, waktu[:])
		soal, skor := hitungScore(waktu[:])
		if soal > maxSoal204 || (soal == maxSoal204 && skor < minSkor204) {
			maxSoal204 = soal
			minSkor204 = skor
			pemenang = nama
		}
	}
	fmt.Printf("Pemenang: %s\nJumlah soal: %d\nTotal skor: %d\n", pemenang, maxSoal204, minSkor204)
}
