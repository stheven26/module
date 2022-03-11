package module

import "fmt"

type Verifikasi func(string, string) bool

func VerifikasiAkses(name string, password string, verifikasi Verifikasi) (string, string) {
	if verifikasi(name, password) {
		fmt.Print("Pesan: ")
	} else {
		fmt.Print("Pesan: ")
		fmt.Printf("Selamat datang %s\n", name)
	}
	nama := name
	Password := password
	return nama, Password
}

func FeetToMeter(num float64) float64 {
	meter := num * 0.3048006096
	return meter
}

func MileToKM(num float64) float64 {
	km := num * 1.60934
	return km
}

func CCToLiter(num float64) float64 {
	liter := num / 1000
	return liter
}
