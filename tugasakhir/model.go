package main

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

var bannerDipilih string