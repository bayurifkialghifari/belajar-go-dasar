package main

type Anggota struct {
	Nama  string
	Umur  int
	Alamat string
}

func main() {
	anggota1 := new(Anggota)
	anggota2 := anggota1

	anggota2.Nama = "Budi"
	anggota2.Umur = 25
	anggota2.Alamat = "Jakarta"

	println("Anggota 1 Nama:", anggota1.Nama)
	println("Anggota 1 Umur:", anggota1.Umur)
	println("Anggota 1 Alamat:", anggota1.Alamat)

	println("Anggota 2 Nama:", anggota2.Nama)
	println("Anggota 2 Umur:", anggota2.Umur)
	println("Anggota 2 Alamat:", anggota2.Alamat)
}