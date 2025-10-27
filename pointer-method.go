package main

type Hewan struct {
	Nama string
}

func (hewan *Hewan) Kucing() {
	hewan.Nama = "Kucing"
}

func main() {
	h := Hewan{}
	h.Kucing()
	println(h.Nama) // Output: Kucing
}