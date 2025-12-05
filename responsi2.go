package main

import (
	"fmt"
	"time"
)

func dataTinggiBadan() {
	// array untuk menampung data tinggi badan
	var tinggi [5]int

	fmt.Println("Masukkan 5 data tinggi badan: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("Data ke-%d: ", i+1)
		fmt.Scan(&tinggi[i])
	}

	// bubble sort
	n := len(tinggi)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if tinggi[j] > tinggi[j+1] {
				temp := tinggi[j]
				tinggi[j] = tinggi[j+1]
				tinggi[j+1] = temp
			}
		}
	}

	// output
	fmt.Println("\nTinggi setelah diurutkan:", tinggi)
	fmt.Println("Tinggi yang paling pendek:", tinggi[0], "cm")
	fmt.Println("Tinggi yang paling tinggi:", tinggi[n-1], "cm")
	fmt.Println()
}

func peminjamanBuku() {
	var peminjaman [7]int

	fmt.Println("Jumlah peminjaman buku selama 7 hari")
	for i := 0; i < 7; i++ {
		fmt.Printf("Hari ke-%d: ", i+1)
		fmt.Scan(&peminjaman[i])
	}

	// bubble sort
	n := len(peminjaman)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if peminjaman[j] > peminjaman[j+1] {
				temp := peminjaman[j]
				peminjaman[j] = peminjaman[j+1]
				peminjaman[j+1] = temp
			}
		}
	}

	fmt.Println("\nData setelah diurutkan:", peminjaman)
	fmt.Println("Dua jumlah peminjaman tertinggi:", peminjaman[n-1], "dan", peminjaman[n-2])

	// hitung total & rata-rata
	total := 0
	for i := 0; i < n; i++ {
		total += peminjaman[i]
	}
	rata := float64(total) / float64(n)

	fmt.Println("Total buku dipinjam:", total)
	fmt.Println("Rata-rata per hari:", rata)
	fmt.Println()
}

func main() {
	fmt.Println("ditulis dengan sepenuh hati oleh YamamoJuan (on github) \n")
	time.Sleep(1 * time.Second)
	
	fmt.Println("Responsi opsi 1 (soal easy dan medium) \n")
	time.Sleep(1 * time.Second)

	fmt.Println("Soal Data Tinggi Badan")
	time.Sleep(700 * time.Millisecond)
	dataTinggiBadan()

	fmt.Println("Soal Peminjaman Buku")
	time.Sleep(700 * time.Millisecond)
	peminjamanBuku()
}
