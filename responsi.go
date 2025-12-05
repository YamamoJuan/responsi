// // Nama: Yamamo Juan Alterico Situmorang
// // NIM: 103032530028
// // Kelas: IT-49-03
// // Tanggal: 14 November 2025

// package main
// import (
// 	"fmt"
// 	"time"
// )

// func layakMemilih() {
// 	var umur int

// 	fmt.Print("Masukkan umur: ")
// 	fmt.Scan(&umur)

// 	if umur >= 17{
// 		fmt.Println("Kamu sudah layak memilih!")
// 	}else{
// 		fmt.Println("Kamu belum layak memilih!")
// 	}
// }

// func nilaiUjian() {
// 	var nilai1 int
// 	var nilai2 int

// 	fmt.Print("Masukkan nilai 1: ")
// 	fmt.Scan(&nilai1)
// 	fmt.Print("Masukkan nilai 2: ")
// 	fmt.Scan(&nilai2)

// 	if nilai1 & nilai2 >= 60{
// 		fmt.Println("Lulus!")
// 	}else{
// 		fmt.Println("Tidak lulus!")
// 	}
// }

// func diskonBelanja() {
// 	var umur int
// 	var pernikahan string

// 	fmt.Print("Masukkan umur: ")
// 	fmt.Scan(&umur)
// 	fmt.Print("Masukkan status: ")
// 	fmt.Scan(&pernikahan)

// 	if umur >= 60 && pernikahan == "menikah"{
// 		fmt.Println("Anda mendapatkan diskon 30%")
// 	}else if umur >= 60 || pernikahan == "menikah"{
// 		fmt.Println("Anda mendapatkan diskon 15%")
// 	}else{
// 		fmt.Println("Anda tidak mendapatkan diskon")
// 	}
// }

// func main() {
// 	fmt.Println("writed by K")
// 	time.Sleep(1 * time.Second)

// 	layakMemilih()
// 	fmt.Println()
// 	nilaiUjian()
// 	fmt.Println()
// 	diskonBelanja()
// 	fmt.Println()
// }