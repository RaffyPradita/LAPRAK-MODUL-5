package main

import "fmt"

func main() {

	var hargaAwal, hargaAkhir float64

	fmt.Scan(&hargaAwal)

	hargaAkhir = hargaAwal * 0.90 * 0.80

	fmt.Printf("%.2f\n", hargaAkhir)
}

//program Diskonbertahap

//Kamus
//hargaAwal : real
//hargaAkhir : real

//Algoritma
//input hargaAwal
//hargaAkhir <- hargaAwal * 0.90 * 0.80

//output (hargaAkhir)
//endprogram
