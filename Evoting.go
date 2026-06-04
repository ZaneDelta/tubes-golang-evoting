package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Kandidat struct {
	NomorUrut int
	Nama      string
	VisiMisi  string
	Suara     int
}

var data []Kandidat

func tambahKandidat(reader *bufio.Reader) {
	var nomor int
	var nama, visi string

	fmt.Print("Nomor Urut : ")
	fmt.Scanln(&nomor)

	fmt.Print("Nama Kandidat : ")
	nama, _ = reader.ReadString('\n')

	fmt.Print("Visi Misi : ")
	visi, _ = reader.ReadString('\n')

	data = append(data, Kandidat{
		NomorUrut: nomor,
		Nama:      strings.TrimSpace(nama),
		VisiMisi:  strings.TrimSpace(visi),
		Suara:     0,
	})

	fmt.Println("Kandidat berhasil ditambahkan!")
}

func tampilkanData() {
	if len(data) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	fmt.Println("\n===== DATA KANDIDAT =====")

	for i := 0; i < len(data); i++ {
		fmt.Println("Nomor Urut :", data[i].NomorUrut)
		fmt.Println("Nama       :", data[i].Nama)
		fmt.Println("Visi Misi  :", data[i].VisiMisi)
		fmt.Println("Suara      :", data[i].Suara)
		fmt.Println("-------------------------")
	}
}

func tambahSuara() {
	var nomor int

	fmt.Print("Masukkan nomor urut kandidat : ")
	fmt.Scanln(&nomor)

	for i := 0; i < len(data); i++ {
		if data[i].NomorUrut == nomor {
			data[i].Suara++
			fmt.Println("Suara berhasil ditambahkan!")
			return
		}
	}

	fmt.Println("Kandidat tidak ditemukan.")
}

func sequentialSearch() {
	var nomor int

	fmt.Print("Cari nomor urut : ")
	fmt.Scanln(&nomor)

	for i := 0; i < len(data); i++ {
		if data[i].NomorUrut == nomor {
			fmt.Println("Data ditemukan")
			fmt.Println("Nomor Urut :", data[i].NomorUrut)
			fmt.Println("Nama       :", data[i].Nama)
			fmt.Println("Visi Misi  :", data[i].VisiMisi)
			fmt.Println("Suara      :", data[i].Suara)
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func selectionSortNama() {
	for i := 0; i < len(data)-1; i++ {
		min := i

		for j := i + 1; j < len(data); j++ {
			if strings.ToLower(data[j].Nama) < strings.ToLower(data[min].Nama) {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}
}

func selectionSortNomor() {
	for i := 0; i < len(data)-1; i++ {
		min := i

		for j := i + 1; j < len(data); j++ {
			if data[j].NomorUrut < data[min].NomorUrut {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}
}

func binarySearch() {
	selectionSortNomor()

	var nomor int

	fmt.Print("Cari nomor urut : ")
	fmt.Scanln(&nomor)

	kiri := 0
	kanan := len(data) - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if data[tengah].NomorUrut == nomor {
			fmt.Println("Data ditemukan")
			fmt.Println("Nomor Urut :", data[tengah].NomorUrut)
			fmt.Println("Nama       :", data[tengah].Nama)
			fmt.Println("Visi Misi  :", data[tengah].VisiMisi)
			fmt.Println("Suara      :", data[tengah].Suara)
			return
		}

		if nomor < data[tengah].NomorUrut {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func selectionSortSuara() {
	for i := 0; i < len(data)-1; i++ {
		max := i

		for j := i + 1; j < len(data); j++ {
			if data[j].Suara > data[max].Suara {
				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}
}

func cariData() {
	var metode int

	fmt.Println("\n=== CARI DATA ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih metode : ")
	fmt.Scanln(&metode)

	if metode == 1 {
		sequentialSearch()
	} else if metode == 2 {
		binarySearch()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func urutData() {
	var urut int

	fmt.Println("\n=== URUTKAN DATA ===")
	fmt.Println("1. Nama")
	fmt.Println("2. Nomor Urut")
	fmt.Print("Pilih : ")
	fmt.Scanln(&urut)

	if urut == 1 {
		selectionSortNama()
		fmt.Println("Data berhasil diurutkan berdasarkan nama.")
	} else if urut == 2 {
		selectionSortNomor()
		fmt.Println("Data berhasil diurutkan berdasarkan nomor urut.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func statistik() {
	total := 0

	for i := 0; i < len(data); i++ {
		total += data[i].Suara
	}

	fmt.Println("\n===== STATISTIK SUARA =====")

	for i := 0; i < len(data); i++ {
		persen := 0.0

		if total > 0 {
			persen = float64(data[i].Suara) * 100 / float64(total)
		}

		fmt.Printf("%s : %d suara (%.2f%%)\n",
			data[i].Nama,
			data[i].Suara,
			persen)
	}

	fmt.Println("Total Suara :", total)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== MENU E-VOTING =====")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Tambah Suara")
		fmt.Println("4. Cari Data")
		fmt.Println("5. Statistik Suara")
		fmt.Println("6. Urutkan Data")
		fmt.Println("7. Keluar")

		var pilih int

		fmt.Print("Pilih menu : ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahKandidat(reader)
		case 2:
			tampilkanData()
		case 3:
			tambahSuara()
		case 4:
			cariData()
		case 5:
			statistik()
		case 6:
			urutData()
		case 7:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
