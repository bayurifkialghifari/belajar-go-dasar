package main

type Binatang interface {
	GetNama() string
	GetSpesies() string
	GetUmur() int
}

type Anjing struct {
	nama   string
	spesies string
	umur   int
}

// Implement contract methods for Anjing
func (a Anjing) GetNama() string {
	return a.nama
}

func (a Anjing) GetSpesies() string {
	return a.spesies
}

func (a Anjing) GetUmur() int {
	return a.umur
}

type Kucing struct {
	nama   string
	spesies string
	umur   int
}

// Implement contract methods for Kucing
func (k Kucing) GetNama() string {
	return k.nama
}

func (k Kucing) GetSpesies() string {
	return k.spesies
}

func (k Kucing) GetUmur() int {
	return k.umur
}

func printBinatangInfo(b Binatang) {
	println("Hewan Nama: " + b.GetNama())
	println("Spesies: " + b.GetSpesies())
	println("Umur: " + string(b.GetUmur()) + " tahun")
}

func main() {
	anjing := Anjing{
		nama:   "Buddy",
		spesies: "Golden Retriever",
		umur:   3,
	}

	kucing := Kucing{
		nama:   "Whiskers",
		spesies: "Siamese",
		umur:   2,
	}

	printBinatangInfo(anjing)
	printBinatangInfo(kucing)
}