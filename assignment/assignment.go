package main



import "fmt"

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Absen     int
}

func main() {
	absen1 := Biodata{
		Nama:      "Riwanto",
		Alamat:    "Jakarta Timur",
		Pekerjaan: "Freelancer",
		Absen:     1,
	}

	absen2 := Biodata{
		Nama:      "Lestaria",
		Alamat:    "Bengkulu",
		Pekerjaan: "Pelajar",
		Absen:     2,
	}
	fmt.Println(absen1, absen2)

	absen1.entryData(1)
	absen2.entryData(2)
}

func (data Biodata) entryData(absen int) {
	fmt.Println("Nomor urut absen ke : ", absen, "adalah", data.Nama)
	fmt.Println("Nomor urut absen ke : ", absen, "adalah", data.Absen)
}
