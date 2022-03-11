package module

import "fmt"

type Verifikasi func(string, string) bool

func main() {
	blokirAkses := func(nama, password string) bool {
		return nama != "stheven"
	}
	nama := "stheven"
	password := "123456"
	dataNama, dataPassword := verifikasiAkses(nama, password, blokirAkses)

	if dataNama == "stheven" && dataPassword == "123456" {
		fmt.Print("Option:\n1.Convert Feet to Meter\n2.Convert Mile to KM\n3.Convert Liter to CC\n")
		fmt.Print("Enter Option Function: ")
		var Option int
		fmt.Scan(&Option)

		if Option == 1 {
			fmt.Print("Enter Feet: ")
			var feet float64
			fmt.Scan(&feet)
			meter := FeetToMeter(feet)
			fmt.Printf("Hasil Konversi %.2f Feet adalah %.2f Meter\n", feet, meter)
		} else if Option == 2 {
			fmt.Print("Enter Mile: ")
			var Mile float64
			fmt.Scan(&Mile)
			km := MileToKM(Mile)
			fmt.Printf("Hasil Konversi %.2f Mile adalah %.2f KM\n", Mile, km)
		} else if Option == 3 {
			fmt.Print("Enter CC: ")
			var CC float64
			fmt.Scan(&CC)
			liter := CCToLiter(CC)
			fmt.Printf("Hasil Konversi %.1f CC adalah %.2f Liter\n", CC, liter)
		} else {
			fmt.Println("Masukkan pilihan anda dengan benar!")
		}
	} else if dataNama == "stheven" && dataPassword != "123456" {
		fmt.Println("'Masukkan password yang benar!'")
	} else if dataNama != "stheven" && dataPassword == "123456" {
		fmt.Println("'Anda bukan stheven, jadi anda tidak memiliki akses'")
	} else {
		fmt.Print("'Koreksi kembali data yang anda masukkan!'")
	}

}

func verifikasiAkses(name string, password string, verifikasi Verifikasi) (string, string) {
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
