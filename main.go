package module

import (
	"fmt"
)

type Verifikasi func(string, string) bool

type Author struct {
	Nama, Password string
}

var belanjaChannel = make(chan string)
var membayarChannel = make(chan string)

// func main() {
// 	authorAccess := Author{
// 		Nama:     "stheven",
// 		Password: "123456",
// 	}

// 	blokirAkses := func(nama, password string) bool {
// 		return nama != "stheven"
// 	}
// 	_ = authorAccess.VerifikasiAkses(authorAccess.Nama, authorAccess.Password, blokirAkses)

// 	if authorAccess.Nama == "stheven" && authorAccess.Password == "123456" {
// 		fmt.Print("\nActivity:\n1.Belajar\n2.Melakukan Konversi\n3.Berbelanja\n4.Operasi Matematika\n")
// 		fmt.Print("Enter Option Activity: ")
// 		var OptionActivity int
// 		fmt.Scan(&OptionActivity)
// 		fmt.Println("")

// 		if OptionActivity == 1 {
// 			Belajar()
// 		} else if OptionActivity == 2 {
// 			fmt.Print("Convert:\n1.Convert Feet to Meter\n2.Convert Mile to KM\n3.Convert Liter to CC\n")
// 			fmt.Print("Enter Option Convert: ")
// 			var OptionConvert int
// 			fmt.Scan(&OptionConvert)
// 			fmt.Println("")

// 			if OptionConvert == 1 {
// 				fmt.Print("Enter Feet: ")
// 				var feet float64
// 				fmt.Scan(&feet)
// 				meter := FeetToMeter(feet)
// 				fmt.Printf("'Hasil Konversi %.2f Feet adalah %.2f Meter'\n", feet, meter)
// 			} else if OptionConvert == 2 {
// 				fmt.Print("Enter Mile: ")
// 				var Mile float64
// 				fmt.Scan(&Mile)
// 				km := MileToKM(Mile)
// 				fmt.Printf("'Hasil Konversi %.2f Mile adalah %.2f KM\n'", Mile, km)
// 			} else if OptionConvert == 3 {
// 				fmt.Print("Enter CC: ")
// 				var CC float64
// 				fmt.Scan(&CC)
// 				liter := CCToLiter(CC)
// 				fmt.Printf("'Hasil Konversi %.1f CC adalah %.2f Liter'\n", CC, liter)
// 			} else {
// 				fmt.Println("Masukkan pilihan konversi anda dengan benar!")
// 			}
// 		} else if OptionActivity == 3 {
// 			go Belanja()
// 			time.Sleep(1 * time.Millisecond)
// 			go Membayar()
// 			time.Sleep(3 * time.Millisecond)
// 			go SelesaiBelanja()
// 			time.Sleep(5 * time.Millisecond)
// 		} else if OptionActivity == 4 {
// 			Matematika()
// 		} else {
// 			fmt.Println("Masukkan pilihan aktivitas anda dengan benar!")
// 		}
// 	} else if authorAccess.Nama == "stheven" && authorAccess.Password != "123456" {
// 		fmt.Println("'Masukkan password dengan benar!'")
// 	} else if authorAccess.Nama != "stheven" && authorAccess.Password == "123456" {
// 		fmt.Println("Masukkan nama dengan benar!'")
// 	} else {
// 		fmt.Print("Anda tidak memiliki akses!'")
// 	}
// }

func (author Author) VerifikasiAkses(name string, password string, verifikasi Verifikasi) string {
	if verifikasi(name, password) {
		fmt.Print("'Pesan: ")
		return name
	} else {
		fmt.Print("'Pesan: ")
		fmt.Printf("Selamat datang %s!'\n", author.Nama)
		return name
	}
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

func Belajar() {
	fmt.Print("Lesson:\n1.Belajar Golang\n2.Belajar Python\n3.Belajar Javascript\n")
	fmt.Print("Enter Option Lesson: ")
	var Option int
	fmt.Scan(&Option)

	if Option == 1 {
		bahasaPemograman := struct {
			Nama string
			SKS  int
		}{
			Nama: "Golang",
			SKS:  3,
		}
		fmt.Printf("'Selamat belajar %s dengan Jumlah %d SKS dan estimasi waktu belajar %d menit'", bahasaPemograman.Nama, bahasaPemograman.SKS, bahasaPemograman.SKS*50)
	} else if Option == 2 {
		bahasaPemograman := struct {
			Nama string
			SKS  int
		}{
			Nama: "Python",
			SKS:  4,
		}
		fmt.Printf("'Selamat belajar %s dengan Jumlah %d SKS dan estimasi waktu belajar %d menit'", bahasaPemograman.Nama, bahasaPemograman.SKS, bahasaPemograman.SKS*50)
	} else if Option == 3 {
		bahasaPemograman := struct {
			Nama string
			SKS  int
		}{
			Nama: "Javascript",
			SKS:  2,
		}
		fmt.Printf("'Selamat belajar %s dengan Jumlah %d SKS dan estimasi waktu belajar %d menit'", bahasaPemograman.Nama, bahasaPemograman.SKS, bahasaPemograman.SKS*50)
	} else {
		fmt.Println("Masukkan pilihan belajar anda dengan benar!")
	}
}

func Belanja() {
	listBelanja := [...]string{"kopi", "sendal", "burger", "tas", "kopi", "sepatu", "burger", "sendal"}
	go func() {
		for _, item := range listBelanja {
			if item == "kopi" {
				fmt.Printf("memasukkan %s ke dalam keranjang\n", item)
				belanjaChannel <- item
			} else if item == "burger" {
				fmt.Printf("memasukkan %s ke dalam keranjang\n", item)
				belanjaChannel <- item
			} else if item == "sendal" {
				fmt.Printf("memasukkan %s ke dalam keranjang\n", item)
				belanjaChannel <- item
			} else if item == "sepatu" {
				fmt.Printf("memasukkan %s ke dalam keranjang\n", item)
				belanjaChannel <- item
			} else if item == "tas" {
				fmt.Printf("memasukkan %s ke dalam keranjang\n", item)
				belanjaChannel <- item
			}
		}
	}()
}

func Membayar() {
	for membayarItem := range belanjaChannel {
		fmt.Println("Berhasil membayar", membayarItem)
		membayarChannel <- membayarItem
	}
}

func SelesaiBelanja() {
	for finished := range membayarChannel {
		fmt.Println("Telah selesai belanja", finished)
	}
}

func Matematika() {
	channelKeliling := make(chan float64)
	channelLuas := make(chan float64)
	fmt.Print("\nMath:\n1.Lingkaran\n2.Balok\n3.Persegi\n")
	fmt.Print("Enter Math: ")
	var OptionMath int
	fmt.Scan(&OptionMath)
	fmt.Println("")

	if OptionMath == 1 {
		go func() (<-chan float64, <-chan float64) {
			fmt.Print("Masukkan nilai r: ")
			var r float64
			fmt.Scan(&r)
			luas := 3.14 * r * r
			keliling := 3.14 * 2 * r
			channelLuas <- luas
			channelKeliling <- keliling
			return channelLuas, channelKeliling
		}()
		luas := <-channelLuas
		keliling := <-channelKeliling
		fmt.Printf("luas lingkaran: %.2f\n", luas)
		fmt.Printf("keliling lingkaran: %.2f\n", keliling)
	} else if OptionMath == 2 {
		go func() (<-chan float64, <-chan float64) {
			fmt.Print("Masukkan nilai p: ")
			var p float64
			fmt.Scan(&p)
			fmt.Print("Masukkan nilai l: ")
			var l float64
			fmt.Scan(&l)
			fmt.Print("Masukkan nilai t: ")
			var t float64
			fmt.Scan(&t)
			luas := 2 * (p*l + p*t + l*t)
			keliling := 4 * (p + l + t)
			channelLuas <- luas
			channelKeliling <- keliling
			return channelLuas, channelKeliling
		}()
		luas := <-channelLuas
		keliling := <-channelKeliling
		fmt.Printf("luas balok: %.2f\n", luas)
		fmt.Printf("keliling balok: %.2f\n", keliling)
	} else if OptionMath == 3 {
		go func() (<-chan float64, <-chan float64) {
			fmt.Print("Masukkan nilai s: ")
			var s float64
			fmt.Scan(&s)
			luas := s * s
			keliling := 4 * s
			channelLuas <- luas
			channelKeliling <- keliling
			return channelLuas, channelKeliling
		}()
		luas := <-channelLuas
		keliling := <-channelKeliling
		fmt.Printf("luas persegi: %.2f\n", luas)
		fmt.Printf("keliling persegi: %.2f\n", keliling)

	} else {
		fmt.Println("Masukkan pilihan matematika anda dengan benar!")
	}
}
