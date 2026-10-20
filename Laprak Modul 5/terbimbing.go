package main

import "fmt"

func main3() {
	type buku struct {
		judul  string
		genre  string
		tahun  int
		rating float64
	}
	var buku1, buku2 buku
	fmt.Scan(&buku1.judul, &buku2.judul)
	fmt.Scan(&buku1.genre, &buku2.genre)
	fmt.Scan(&buku1.tahun, &buku2.tahun)
	fmt.Scan(&buku1.rating, &buku2.rating)

	fmt.Println("--- Book 1 ---")
	fmt.Println("Book title: ", buku1.judul)
	fmt.Println("Genre: ", buku1.genre)
	fmt.Println("tahun: ", buku1.tahun)
	fmt.Println("Rating: ", buku1.rating)

	fmt.Println("--- Book 2 ---")
	fmt.Println("Book title: ", buku2.judul)
	fmt.Println("Genre: ", buku2.genre)
	fmt.Println("tahun: ", buku2.tahun)
	fmt.Println("Rating: ", buku2.rating)
}

//program databuku

//kamus
//type buku : struct <
//    judul : string
//    genre : string
//    tahunrilis : integer
//    rating : real
//>

//buku1, buku2 : buku

//algoritma
//input(buku1.judul, buku2.judul)
//input(buku1.genre, buku2.genre)
//input(buku1.tahunrilis, buku2.tahunrilis)
//input(buku1.rating, buku2.rating)

//output("--- book 1 ---")
//output(buku1.judul)
//output(buku1.genre)
//output(buku1.tahunrilis)
//output(buku1.rating)

//output("--- book 2 ---")
//output(buku2.judul)
//output(buku2.genre)
//output(buku2.tahunrilis)
//output(buku2.rating)
//endprogram
