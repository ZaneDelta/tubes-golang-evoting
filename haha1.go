package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ================= DATA MODEL ================= (gabriel edbert)

const MaksData = 100

type Kandidat struct {
	nomorUrut   int
	nama        string
	visiMisi    string
	jumlahSuara int
}

type Pemilih struct {
	idPemilih    int
	nama         string
	sudahMemilih bool
}

var dataKandidat [MaksData]Kandidat
var jumlahKandidat int

var dataPemilih [MaksData]Pemilih
var jumlahPemilih int

// eel (gabriel edbert)
// @jebb_24 (gabriel edbert)
var eel = "CLI Kampus"
var jebb24 = 24

// ================= MAIN PROGRAM ================= (gabriel edbert)

func main() {
	reader := bufio.NewReader(os.Stdin)

	inisialisasiDataDummy()
	menuUtama(reader)
}

func inisialisasiDataDummy() {
	dataKandidat[0] = Kandidat{nomorUrut: 1, nama: "Aditya Pratama", visiMisi: "Transparansi pemilihan", jumlahSuara: 2}
	dataKandidat[1] = Kandidat{nomorUrut: 2, nama: "Bianca Safira", visiMisi: "Digitalisasi organisasi", jumlahSuara: 1}
	dataKandidat[2] = Kandidat{nomorUrut: 3, nama: "Chandra Wijaya", visiMisi: "Pelayanan aktif", jumlahSuara: 0}
	jumlahKandidat = 3

	dataPemilih[0] = Pemilih{idPemilih: 101, nama: "Siti Aminah", sudahMemilih: true}
	dataPemilih[1] = Pemilih{idPemilih: 102, nama: "Raka Saputra", sudahMemilih: true}
	dataPemilih[2] = Pemilih{idPemilih: 103, nama: "Dini Lestari", sudahMemilih: true}
	dataPemilih[3] = Pemilih{idPemilih: 104, nama: "Bima Arya", sudahMemilih: false}
	dataPemilih[4] = Pemilih{idPemilih: 105, nama: "Nadia Putri", sudahMemilih: false}
	jumlahPemilih = 5
}

// ================= MENU ================= (gabriel edbert)

func menuUtama(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("1. Kelola Kandidat")
		fmt.Println("2. Kelola Pemilih")
		fmt.Println("3. Voting")
		fmt.Println("4. Statistik Hasil Voting")
		fmt.Println("5. Sorting Data Kandidat")
		fmt.Println("6. Searching Kandidat")
		fmt.Println("0. Keluar")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			menuKelolaKandidat(reader)
		case 2:
			menuKelolaPemilih(reader)
		case 3:
			prosesVoting(reader)
		case 4:
			tampilkanStatistik(reader)
		case 5:
			menuSorting(reader)
		case 6:
			menuSearching(reader)
		case 0:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuKelolaKandidat(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("MENU KELOLA KANDIDAT")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Edit Kandidat")
		fmt.Println("4. Hapus Kandidat")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			tambahKandidat(reader)
		case 2:
			tampilkanKandidat()
			pause(reader)
		case 3:
			editKandidat(reader)
		case 4:
			hapusKandidat(reader)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuKelolaPemilih(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("MENU KELOLA PEMILIH")
		fmt.Println("1. Tambah Pemilih")
		fmt.Println("2. Tampilkan Pemilih")
		fmt.Println("3. Edit Pemilih")
		fmt.Println("4. Hapus Pemilih")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		switch pilihan {
		case 1:
			tambahPemilih(reader)
		case 2:
			tampilkanPemilih()
			pause(reader)
		case 3:
			editPemilih(reader)
		case 4:
			hapusPemilih(reader)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuSorting(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("SORTING DATA KANDIDAT")
		fmt.Println("1. Selection Sort berdasarkan jumlah suara")
		fmt.Println("2. Insertion Sort berdasarkan nomor urut")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		if pilihan == 0 {
			return
		}

		if jumlahKandidat == 0 {
			fmt.Println("Data kandidat masih kosong.")
			pause(reader)
			continue
		}

		switch pilihan {
		case 1:
			ascending := pilihUrutanAscending(reader)
			selectionSortJumlahSuara(ascending)
			fmt.Println("Data berhasil diurutkan dengan Selection Sort.")
			tampilkanKandidat()
			pause(reader)
		case 2:
			ascending := pilihUrutanAscending(reader)
			insertionSortNomorUrut(ascending)
			fmt.Println("Data berhasil diurutkan dengan Insertion Sort.")
			tampilkanKandidat()
			pause(reader)
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

func menuSearching(reader *bufio.Reader) {
	for {
		clearScreen()
		tampilkanHeader()
		fmt.Println("SEARCHING KANDIDAT")
		fmt.Println("1. Sequential Search berdasarkan nomor urut")
		fmt.Println("2. Binary Search berdasarkan nomor urut")
		fmt.Println("0. Kembali")
		fmt.Println("------------------------------------------------------------")

		pilihan := inputInt(reader, "Pilih menu: ")

		if pilihan == 0 {
			return
		}

		if jumlahKandidat == 0 {
			fmt.Println("Data kandidat masih kosong.")
			pause(reader)
			continue
		}

		nomor := inputIntMin(reader, "Masukkan nomor urut kandidat: ", 1)

		switch pilihan {
		case 1:
			index := sequentialSearchKandidat(nomor)
			if index == -1 {
				fmt.Println("Kandidat tidak ditemukan.")
			} else {
				tampilkanDetailKandidat(index)
			}
			pause(reader)
		case 2:
			insertionSortNomorUrut(true)
			fmt.Println("Data diurutkan ascending terlebih dahulu sebelum Binary Search.")
			index := binarySearchKandidat(nomor)
			if index == -1 {
				fmt.Println("Kandidat tidak ditemukan.")
			} else {
				tampilkanDetailKandidat(index)
			}
			pause(reader)
		default:
			fmt.Println("Pilihan tidak tersedia.")
			pause(reader)
		}
	}
}

// ================= CRUD KANDIDAT ================= (gabriel edbert)

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

// ================= CRUD PEMILIH ================= (gabriel edbert)

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

// ================= VOTING ================= (gabriel edbert)

func prosesVoting(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("VOTING")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong, voting belum bisa dilakukan.")
		pause(reader)
		return
	}

	if jumlahPemilih == 0 {
		fmt.Println("Data pemilih masih kosong, voting belum bisa dilakukan.")
		pause(reader)
		return
	}

	tampilkanPemilih()

	id := inputIntMin(reader, "Masukkan ID pemilih: ", 1)
	indexPemilih := sequentialSearchPemilih(id)

	if indexPemilih == -1 {
		fmt.Println("ID pemilih tidak ditemukan.")
		pause(reader)
		return
	}

	if dataPemilih[indexPemilih].sudahMemilih {
		fmt.Println("Pemilih ini sudah menggunakan hak suara.")
		pause(reader)
		return
	}

	fmt.Println()
	tampilkanKandidat()

	nomor := inputIntMin(reader, "Pilih nomor urut kandidat: ", 1)
	indexKandidat := sequentialSearchKandidat(nomor)

	if indexKandidat == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		pause(reader)
		return
	}

	dataKandidat[indexKandidat].jumlahSuara++
	dataPemilih[indexPemilih].sudahMemilih = true

	waktuVoting := time.Now().Format("02-01-2006 15:04:05")

	fmt.Println("Voting berhasil.")
	fmt.Printf("Pemilih      : %s\n", dataPemilih[indexPemilih].nama)
	fmt.Printf("Kandidat     : %s\n", dataKandidat[indexKandidat].nama)
	fmt.Printf("Waktu voting : %s\n", waktuVoting)

	pause(reader)
}

// ================= STATISTIK ================= (gabriel edbert)

func tampilkanStatistik(reader *bufio.Reader) {
	clearScreen()
	tampilkanHeader()
	fmt.Println("STATISTIK HASIL VOTING")

	if jumlahKandidat == 0 {
		fmt.Println("Data kandidat masih kosong.")
		pause(reader)
		return
	}

	totalSuara := hitungTotalSuara()

	fmt.Printf("Total suara masuk: %d\n", totalSuara)
	fmt.Println("------------------------------------------------------------")

	for i := 0; i < jumlahKandidat; i++ {
		persentase := 0.0

		if totalSuara > 0 {
			persentase = float64(dataKandidat[i].jumlahSuara) / float64(totalSuara) * 100
		}

		fmt.Printf("Kandidat %d - %-22s : %.2f%% (%d suara)\n",
			dataKandidat[i].nomorUrut,
			potongTeks(dataKandidat[i].nama, 22),
			persentase,
			dataKandidat[i].jumlahSuara)
	}

	fmt.Println("------------------------------------------------------------")

	if totalSuara == 0 {
		fmt.Println("Belum ada suara masuk.")
	} else {
		suaraTertinggi := cariSuaraTerbanyak()

		fmt.Println("Kandidat dengan suara terbanyak:")
		for i := 0; i < jumlahKandidat; i++ {
			if dataKandidat[i].jumlahSuara == suaraTertinggi {
				fmt.Printf("- Nomor %d, %s dengan %d suara\n",
					dataKandidat[i].nomorUrut,
					dataKandidat[i].nama,
					dataKandidat[i].jumlahSuara)
			}
		}
	}

	pause(reader)
}

func hitungTotalSuara() int {
	total := 0

	for i := 0; i < jumlahKandidat; i++ {
		total += dataKandidat[i].jumlahSuara
	}

	return total
}

func cariSuaraTerbanyak() int {
	if jumlahKandidat == 0 {
		return 0
	}

	maks := dataKandidat[0].jumlahSuara

	for i := 1; i < jumlahKandidat; i++ {
		if dataKandidat[i].jumlahSuara > maks {
			maks = dataKandidat[i].jumlahSuara
		}
	}

	return maks
}

// ================= SORTING ================= (gabriel edbert)

// proses selection sort manual berdasarkan jumlah suara (gabriel edbert)
func selectionSortJumlahSuara(ascending bool) {
	for i := 0; i < jumlahKandidat-1; i++ {
		posisi := i

		for j := i + 1; j < jumlahKandidat; j++ {
			if ascending {
				if dataKandidat[j].jumlahSuara < dataKandidat[posisi].jumlahSuara {
					posisi = j
				}
			} else {
				if dataKandidat[j].jumlahSuara > dataKandidat[posisi].jumlahSuara {
					posisi = j
				}
			}
		}

		if posisi != i {
			temp := dataKandidat[i]
			dataKandidat[i] = dataKandidat[posisi]
			dataKandidat[posisi] = temp
		}
	}
}

// proses insertion sort manual berdasarkan nomor urut (gabriel edbert)
func insertionSortNomorUrut(ascending bool) {
	for i := 1; i < jumlahKandidat; i++ {
		key := dataKandidat[i]
		j := i - 1

		if ascending {
			for j >= 0 && dataKandidat[j].nomorUrut > key.nomorUrut {
				dataKandidat[j+1] = dataKandidat[j]
				j--
			}
		} else {
			for j >= 0 && dataKandidat[j].nomorUrut < key.nomorUrut {
				dataKandidat[j+1] = dataKandidat[j]
				j--
			}
		}

		dataKandidat[j+1] = key
	}
}

// ================= SEARCHING ================= (gabriel edbert)

// proses mencari kandidat berdasarkan nomor urut secara sequential (gabriel edbert)
func sequentialSearchKandidat(nomor int) int {
	for i := 0; i < jumlahKandidat; i++ {
		if dataKandidat[i].nomorUrut == nomor {
			return i
		}
	}

	return -1
}

// proses binary search setelah data kandidat diurutkan ascending (gabriel edbert)
func binarySearchKandidat(nomor int) int {
	kiri := 0
	kanan := jumlahKandidat - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if dataKandidat[tengah].nomorUrut == nomor {
			return tengah
		} else if dataKandidat[tengah].nomorUrut < nomor {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

// proses mencari pemilih berdasarkan id secara sequential (gabriel edbert)
func sequentialSearchPemilih(id int) int {
	for i := 0; i < jumlahPemilih; i++ {
		if dataPemilih[i].idPemilih == id {
			return i
		}
	}

	return -1
}

// ================= HELPER INPUT DAN TAMPILAN ================= (gabriel edbert)

func pilihUrutanAscending(reader *bufio.Reader) bool {
	for {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")

		pilihan := inputInt(reader, "Pilih urutan: ")

		if pilihan == 1 {
			return true
		}

		if pilihan == 2 {
			return false
		}

		fmt.Println("Pilihan urutan tidak tersedia.")
	}
}

func tampilkanHeader() {
	fmt.Println("============================================================")
	fmt.Println("+++ APLIKASI SISTEM PEMUNGUTAN SUARA DIGITAL (E-VOTING) +++")
	fmt.Printf("Mode: %s | Kode Project: JEBB-%d\n", eel, jebb24)
	fmt.Println("============================================================")
}

func clearScreen() {
	for i := 0; i < 25; i++ {
		fmt.Println()
	}
}

func pause(reader *bufio.Reader) {
	fmt.Print("\nTekan ENTER untuk melanjutkan...")
	reader.ReadString('\n')
}

func inputString(reader *bufio.Reader, pesan string) string {
	for {
		fmt.Print(pesan)

		teks, _ := reader.ReadString('\n')
		teks = strings.TrimSpace(teks)

		if teks == "" {
			fmt.Println("Input tidak boleh kosong.")
		} else {
			return teks
		}
	}
}

func inputStringBolehKosong(reader *bufio.Reader, pesan string) string {
	fmt.Print(pesan)

	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}

func inputInt(reader *bufio.Reader, pesan string) int {
	for {
		fmt.Print(pesan)

		teks, _ := reader.ReadString('\n')
		teks = strings.TrimSpace(teks)

		angka, err := strconv.Atoi(teks)
		if err != nil {
			fmt.Println("Input harus berupa angka.")
		} else {
			return angka
		}
	}
}

func inputIntMin(reader *bufio.Reader, pesan string, minimal int) int {
	for {
		angka := inputInt(reader, pesan)

		if angka < minimal {
			fmt.Printf("Input minimal adalah %d.\n", minimal)
		} else {
			return angka
		}
	}
}

func inputKonfirmasi(reader *bufio.Reader, pesan string) bool {
	for {
		jawaban := strings.ToLower(inputString(reader, pesan))

		if jawaban == "y" || jawaban == "ya" {
			return true
		}

		if jawaban == "n" || jawaban == "tidak" {
			return false
		}

		fmt.Println("Masukkan y atau n.")
	}
}

func potongTeks(teks string, batas int) string {
	if len(teks) <= batas {
		return teks
	}

	if batas <= 3 {
		return teks[:batas]
	}

	return teks[:batas-3] + "..."
}