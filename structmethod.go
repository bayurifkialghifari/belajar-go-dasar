package main

type Anggota struct {
	Nia string
	Nama string
	TanggalLahir string
	Alamat string
	Age int
}

func (anggota Anggota) getInfo() string {
	return "NIA: " + anggota.Nia + ", Nama: " + anggota.Nama + ", Tanggal Lahir: " + anggota.TanggalLahir + ", Alamat: " + anggota.Alamat + ", Age: " + string(anggota.Age)
}

func main() {
	anggota := Anggota{
		Nia: "201524002",
		Nama: "Budi",
		TanggalLahir: "01-01-2000",
		Alamat: "Jakarta",
		Age: 24,
	}
	println(anggota.getInfo())
}
