package main

import "fmt"

func main1() {
	const (
		kuis     float64 = 0.15
		tugas    float64 = 0.15
		Pratikum float64 = 0.20
		UTS      float64 = 0.25
		UAS      float64 = 0.25
	)
	type mahasiswa struct {
		NIM           string
		Nama          string
		NilaiTugas    float64
		NilaiKuis     float64
		NilaiPratikum float64
		NilaiUTS      float64
		NilaiUAS      float64
		NilaiAkhir    float64
	}

	var DataMahasiswa mahasiswa
	var nilai_Akhir float64

	fmt.Scan(&DataMahasiswa.Nama, &DataMahasiswa.NIM, &DataMahasiswa.NilaiTugas, &DataMahasiswa.NilaiKuis, &DataMahasiswa.NilaiPratikum, &DataMahasiswa.NilaiUTS, &DataMahasiswa.NilaiUAS)

	nilai_Akhir = DataMahasiswa.NilaiTugas*tugas + DataMahasiswa.NilaiKuis*kuis + DataMahasiswa.NilaiPratikum*Pratikum + DataMahasiswa.NilaiUTS*UTS + DataMahasiswa.NilaiUAS*UAS

	fmt.Println("Nama: ", DataMahasiswa.Nama)
	fmt.Println("NIM: ", DataMahasiswa.NIM)
	fmt.Println("Nilai Tugas: ", DataMahasiswa.NilaiTugas)
	fmt.Println("Nilai Kuis: ", DataMahasiswa.NilaiKuis)
	fmt.Println("Nilai Pratikum: ", DataMahasiswa.NilaiPratikum)
	fmt.Println("Nilai UTS: ", DataMahasiswa.NilaiUTS)
	fmt.Println("Nilai UAS: ", DataMahasiswa.NilaiUAS)
	fmt.Println("Nilai Akhir: ", nilai_Akhir)
}
