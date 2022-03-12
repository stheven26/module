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

func (author Author) VerifikasiAkses(name string, password string, verifikasi Verifikasi) {
	if verifikasi(name, password) {
		fmt.Print("'Pesan: ")
	} else {
		fmt.Print("'Pesan: ")
		fmt.Printf("Selamat datang %s!'\n", author.Nama)
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

func Sembayar() {
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
