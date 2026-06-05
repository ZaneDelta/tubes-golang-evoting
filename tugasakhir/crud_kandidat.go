package main

import (
	"bufio"
	"fmt"
)

func tambahKandidat(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("TAMBAH KANDIDAT")

	if jumlahKandidat >= MaksData {
		fmt.Println("Data kandidat sudah penuh.")
		pause(reader)
		return
	}

	nomor := inputIntMin(reader, "Masukkan nomor urut: ", 1)
	if sequentialSearchKandidat(nomor) != -1 {
		fmt.Println("Nomor urut sudah digunakan.")
		pause(reader)
		return
	}

	nama := inputString(reader, "Masukkan nama kandidat: ")
	visiMisi := inputString(reader, "Masukkan visi misi: ")

	dataKandidat[jumlahKandidat] = Kandidat{
		nomorUrut:   nomor,
		nama:        nama,
		visiMisi:    visiMisi,
		jumlahSuara: 0,
	}
	jumlahKandidat++

	fmt.Println("Kandidat berhasil ditambahkan.")
	pause(reader)
}

func tampilkanKandidat() {
	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong.")
		return
	}

	fmt.Println("DAFTAR KANDIDAT")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("| %-5s | %-22s | %-30s | %-6s |\n", "No", "Nama", "Visi Misi", "Suara")
	fmt.Println("--------------------------------------------------------------------------------")

	for i := 0; i < jumlahKandidat; i++ {
		fmt.Printf("| %-5d | %-22s | %-30s | %-6d |\n",
			dataKandidat[i].nomorUrut,
			potongTeks(dataKandidat[i].nama, 22),
			potongTeks(dataKandidat[i].visiMisi, 30),
			dataKandidat[i].jumlahSuara)
	}

	fmt.Println("--------------------------------------------------------------------------------")
}

func editKandidat(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("EDIT KANDIDAT")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong.")
		pause(reader)
		return
	}

	tampilkanKandidat()
	nomor := inputIntMin(reader, "Masukkan nomor urut kandidat yang diedit: ", 1)
	index := sequentialSearchKandidat(nomor)

	if index == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		pause(reader)
		return
	}

	fmt.Println("Kosongkan nama atau visi misi jika tidak ingin mengubah.")
	nomorBaru := inputIntMin(reader, "Nomor urut baru (0 jika tetap): ", 0)

	if nomorBaru != 0 {
		indexDuplikat := sequentialSearchKandidat(nomorBaru)
		if indexDuplikat != -1 && indexDuplikat != index {
			fmt.Println("Nomor urut baru sudah digunakan kandidat lain.")
			pause(reader)
			return
		}
		dataKandidat[index].nomorUrut = nomorBaru
	}

	namaBaru := inputStringBolehKosong(reader, "Nama baru: ")
	if namaBaru != "" {
		dataKandidat[index].nama = namaBaru
	}

	visiMisiBaru := inputStringBolehKosong(reader, "Visi misi baru: ")
	if visiMisiBaru != "" {
		dataKandidat[index].visiMisi = visiMisiBaru
	}

	fmt.Println("Data kandidat berhasil diedit.")
	pause(reader)
}

func hapusKandidat(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("HAPUS KANDIDAT")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong.")
		pause(reader)
		return
	}

	tampilkanKandidat()
	nomor := inputIntMin(reader, "Masukkan nomor urut kandidat yang dihapus: ", 1)
	index := sequentialSearchKandidat(nomor)

	if index == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		pause(reader)
		return
	}

	fmt.Printf("Kandidat yang akan dihapus: %s\n", dataKandidat[index].nama)
	fmt.Println("Catatan: jumlah suara kandidat tersebut juga ikut terhapus.")

	if !inputKonfirmasi(reader, "Yakin hapus kandidat? (y/n): ") {
		fmt.Println("Hapus kandidat dibatalkan.")
		pause(reader)
		return
	}

	for i := index; i < jumlahKandidat-1; i++ {
		dataKandidat[i] = dataKandidat[i+1]
	}

	dataKandidat[jumlahKandidat-1] = Kandidat{}
	jumlahKandidat--

	fmt.Println("Kandidat berhasil dihapus.")
	pause(reader)
}

func tampilkanDetailKandidat(index int) {
	fmt.Println("Data kandidat ditemukan:")
	fmt.Printf("Nomor Urut   : %d\n", dataKandidat[index].nomorUrut)
	fmt.Printf("Nama         : %s\n", dataKandidat[index].nama)
	fmt.Printf("Visi Misi    : %s\n", dataKandidat[index].visiMisi)
	fmt.Printf("Jumlah Suara : %d\n", dataKandidat[index].jumlahSuara)
}