package main

import (
	"bufio"
	"fmt"
)

func tambahPemilih(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("TAMBAH PEMILIH")

	if jumlahPemilih >= MaksData {
		fmt.Println("Data pemilih sudah penuh.")
		pause(reader)
		return
	}

	id := inputIntMin(reader, "Masukkan ID pemilih: ", 1)
	if sequentialSearchPemilih(id) != -1 {
		fmt.Println("ID pemilih sudah digunakan.")
		pause(reader)
		return
	}

	nama := inputString(reader, "Masukkan nama pemilih: ")

	dataPemilih[jumlahPemilih] = Pemilih{
		idPemilih:    id,
		nama:         nama,
		sudahMemilih: false,
	}
	jumlahPemilih++

	fmt.Println("Pemilih berhasil ditambahkan.")
	pause(reader)
}

func tampilkanPemilih() {
	if jumlahPemilih == 0 {
		fmt.Println("Data pemilih masih kosong.")
		return
	}

	fmt.Println("DAFTAR PEMILIH")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("| %-8s | %-25s | %-15s |\n", "ID", "Nama", "Status")
	fmt.Println("------------------------------------------------------------")

	for i := 0; i < jumlahPemilih; i++ {
		status := "Belum Memilih"
		if dataPemilih[i].sudahMemilih {
			status = "Sudah Memilih"
		}

		fmt.Printf("| %-8d | %-25s | %-15s |\n",
			dataPemilih[i].idPemilih,
			potongTeks(dataPemilih[i].nama, 25),
			status)
	}

	fmt.Println("------------------------------------------------------------")
}

func editPemilih(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("EDIT PEMILIH")

	if jumlahPemilih == 0 {
		fmt.Println("Data pemilih masih kosong.")
		pause(reader)
		return
	}

	tampilkanPemilih()
	id := inputIntMin(reader, "Masukkan ID pemilih yang diedit: ", 1)
	index := sequentialSearchPemilih(id)

	if index == -1 {
		fmt.Println("Pemilih tidak ditemukan.")
		pause(reader)
		return
	}

	fmt.Println("Kosongkan nama jika tidak ingin mengubah.")
	idBaru := inputIntMin(reader, "ID baru (0 jika tetap): ", 0)

	if idBaru != 0 {
		indexDuplikat := sequentialSearchPemilih(idBaru)
		if indexDuplikat != -1 && indexDuplikat != index {
			fmt.Println("ID baru sudah digunakan pemilih lain.")
			pause(reader)
			return
		}
		dataPemilih[index].idPemilih = idBaru
	}

	namaBaru := inputStringBolehKosong(reader, "Nama baru: ")
	if namaBaru != "" {
		dataPemilih[index].nama = namaBaru
	}

	fmt.Println("Data pemilih berhasil diedit.")
	pause(reader)
}

func hapusPemilih(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("HAPUS PEMILIH")

	if jumlahPemilih == 0 {
		fmt.Println("Data pemilih masih kosong.")
		pause(reader)
		return
	}

	tampilkanPemilih()
	id := inputIntMin(reader, "Masukkan ID pemilih yang dihapus: ", 1)
	index := sequentialSearchPemilih(id)

	if index == -1 {
		fmt.Println("Pemilih tidak ditemukan.")
		pause(reader)
		return
	}

	fmt.Printf("Pemilih yang akan dihapus: %s\n", dataPemilih[index].nama)

	if dataPemilih[index].sudahMemilih {
		fmt.Println("Catatan: pemilih sudah voting, suara yang masuk tetap tercatat di kandidat.")
	}

	if !inputKonfirmasi(reader, "Yakin hapus pemilih? (y/n): ") {
		fmt.Println("Hapus pemilih dibatalkan.")
		pause(reader)
		return
	}

	for i := index; i < jumlahPemilih-1; i++ {
		dataPemilih[i] = dataPemilih[i+1]
	}

	dataPemilih[jumlahPemilih-1] = Pemilih{}
	jumlahPemilih--

	fmt.Println("Pemilih berhasil dihapus.")
	pause(reader)
}