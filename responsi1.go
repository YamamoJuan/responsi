// // Nama: Yamamo Juan Alterico Situmorang
// // NIM: 103032530028
// // Kelas: IT-49-03
// // Tanggal: 21 November 2025

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func bilanganGenap() {
// 	var n int
// 	fmt.Print("Masukkan bilangan: ")
// 	fmt.Scan(&n)

// 	for i := 1; i <= n; i++ {
// 		if i%2 == 0 {
// 			fmt.Print(i, " \n")
// 		}
// 	}
// 	fmt.Println()
// }

// func angkaPositif() {
// 	var n int
// 	var positif int
// 	for {
// 		fmt.Print("Masukkan bilangan: ")
// 		fmt.Scan(&n)

// 		if n == 0 {
// 			fmt.Printf("Jumlah bilangan positif ada %d \n", positif)
// 			break
// 		} else if n > 0 {
// 			positif++
// 		}
// 	}
// }

// func passwordLogin() {
// 	var password string
// 	var verifikasi int
// 	for {
// 		fmt.Print("Masukkan password: ")
// 		fmt.Scan(&password)

// 		if password == "golang123" { // bagian 1
// 			fmt.Println("Password benar")
// 			break
// 		} else {
// 			verifikasi++
// 			if verifikasi == 3 {
// 				fmt.Println("Anda salah memasukkan password sebanyak 3 kali")
// 				break
// 			} else {
// 				fmt.Println("Password salah")
// 			}
// 		}
// 	}
// 	var n int
// 	fmt.Print("Masukkan angka N: ")
// 	fmt.Scan(&n)

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	fmt.Println("ditulis dengan sepenuh hati oleh YamamoJuan (github) \n")
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("soal 1")
// 	time.Sleep(1 * time.Millisecond)
// 	bilanganGenap()

// 	fmt.Println("soal 2")
// 	time.Sleep(1 * time.Millisecond)
// 	angkaPositif()

// 	fmt.Println("soal 3")
// 	time.Sleep(1 * time.Millisecond)
// 	passwordLogin()
// }
