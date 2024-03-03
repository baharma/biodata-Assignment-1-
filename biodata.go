package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getDataByAbsen(absen int) Teman {
	var teman Teman

	if absen == 1 {
		teman = Teman{
			Nama:      "Aditya",
			Alamat:    "Denpasar Barat",
			Pekerjaan: "Developer",
			Alasan:    "Suka dengan pemrograman Golang",
		}
	} else if absen == 2 {
		teman = Teman{
			Nama:      "rages",
			Alamat:    "Jalan menderita",
			Pekerjaan: "Designer",
			Alasan:    "Ingin mendalami menggunakan Golang",
		}
	} else {
		teman = Teman{
			Nama:      "Agus",
			Alamat:    "Jalan Denpasar",
			Pekerjaan: "QA",
			Alasan:    "Ingin Mengetes Adit dan rages",
		}
	}

	return teman
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	absen := os.Args[1]

	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen tidak valid")
		return
	}

	dataTeman := getDataByAbsen(absenInt)

	fmt.Println("Data Teman:")
	fmt.Println("Nama:", dataTeman.Nama)
	fmt.Println("Alamat:", dataTeman.Alamat)
	fmt.Println("Pekerjaan:", dataTeman.Pekerjaan)
	fmt.Println("Alasan Memilih Kelas Golang:", dataTeman.Alasan)
}
