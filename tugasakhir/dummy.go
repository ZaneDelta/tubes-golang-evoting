package main

func inisialisasiDataDummy() {
	dataKandidat[0] = Kandidat{nomorUrut: 1, nama: "Prabowo Subianto", visiMisi: "MBG No 1", jumlahSuara: 2}
	dataKandidat[1] = Kandidat{nomorUrut: 2, nama: "Rafi Maulana", visiMisi: "Aku Suka Rodok", jumlahSuara: 1}
	dataKandidat[2] = Kandidat{nomorUrut: 3, nama: "Ilman Baruna", visiMisi: "Kudus Negeri Kretek", jumlahSuara: 0}
	jumlahKandidat = 3

	dataPemilih[0] = Pemilih{idPemilih: 101, nama: "Siti Aminah", sudahMemilih: true}
	dataPemilih[1] = Pemilih{idPemilih: 102, nama: "Raka Saputra", sudahMemilih: true}
	dataPemilih[2] = Pemilih{idPemilih: 103, nama: "Dini Lestari", sudahMemilih: true}
	dataPemilih[3] = Pemilih{idPemilih: 104, nama: "Bima Arya", sudahMemilih: false}
	dataPemilih[4] = Pemilih{idPemilih: 105, nama: "Nadia Putri", sudahMemilih: false}
	jumlahPemilih = 5
}