package main

import "fmt"

type Anggota struct {
	Nia string
	Nama string
	TanggalLahir string
	Alamat string
	Age int
}

func main() {
	var anggota Anggota
	anggota.Nia = "201524002"
	anggota.Nama = "Budi"
	anggota.TanggalLahir = "01-01-2000"
	anggota.Alamat = "Jakarta"
	anggota.Age = 24

	println("NIA:", anggota.Nia)
	println("Nama:", anggota.Nama)
	println("Tanggal Lahir:", anggota.TanggalLahir)
	println("Alamat:", anggota.Alamat)
	println("Age:", anggota.Age)
	fmt.Println(anggota)

	// Struct literal
	anggota2 := Anggota{
		Nia: "201524003",
		Nama: "Siti",
		TanggalLahir: "02-02-2001",
		Alamat: "Bandung",
		Age: 23,
	}

	fmt.Println(anggota2)
}