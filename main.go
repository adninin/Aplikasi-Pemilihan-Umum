/*
Tugas Besar UAS Algoritma Pemrogaman Telkom University
Program ini dibuat untuk menyelesaikan tugas besar mata kuliah algrotima pemrograman semester 2

APLIKASI PEMILIHAN UMUM

Deskripsi: Aplikasi digunakan untuk melakukan pemilihan umum calon legislatif dan partai tertentu. Pengguna aplikasi ini 
           adalah pemilih dan juga petugas kpu (admin).

Spesifikasi:
         a. Pengguna (admin) bisa menambahkan, mengubah (edit), dan juga menghapus data calon dan pemilih.
         b. Pemilih bisa melakukan pemilihan pada durasi waktu yang ditentukan saja, di luar itu, hanya bisa melihat daftar calon saja.
         c. Pengguna bisa menampilkan data terurut berdasarkan hasil perolehan suara. Berdasarkan partai, berdasarkan nama calon dan partai.
         d. Tentukan nilai threshold atau ambang batas suatu calon untuk bisa terpilih.
         e. Pengguna (admin) bisa melakukan pencarian data calon yang berasal dari partai tertentu, pencarian berdasarkan nama calon, 
            dan juga nama pemilih dari calon tertentu.

Authors:
         a. Adnin Atmadewati Ashrini    (103062300029)
		 b. Daniel Nadeak               (103062330016)
		 c. Virani Yulfina              (103062300028)

Cara Penggunaan Aplikasi Sebagai Admin:
         a. Login sebagai admin dengan nama pengguna "admin" lalu input data calon legislatif dengan format (nama partai)
		 b. Input data pemilih dengan format (nama)
		 c. Atur rentang waktu pada program dengan format time.Date(2024, time.June, 11, 8, 0, 0, 0, time.Local)   
		    Contoh: 11 Juni 2024 pukul 08:00                       tahun    bulan    tgl  jam

Cara Penggunaan Aplikasi Sebagai Pemilih:
         a. Login sebagai pemilih dengan memasukkan nama pengguna yang diberi oleh admin
		 b. Lakukan pemilihan selama rentang waktu yang diberikan

Notes: 
         a. opsi no 9 baru akan bisa bekerja efektif apabila pemilihan sudah dilakukan
		 b. opsi no 12 => opsi 3 baru akan bisa bekerja efektif apabila pemilihan sudah dilakukan
		 c. file "kanidat.txt" akan otomatis terbuat apabila admin sudah melakukan input data calon legislatif
		 d. file "pemilih.txt" akan otomatis terbuat apabila admin sudah melakukan input data pemilih
		 e. file "hasil_voting.txt" akan otomatis terbuat apabila pemilih sudah melakukan voting
		 f. pilih opsi 14 untuk mereset data voting 

Aspek:
         a. Program ini menggunakan procedure (Fungsi dan Metode)
         b. Array => var kandidats []Kandidat dsb
         c. Tipe bentukan (struct)
         d. Searching pada func TampilkanKandidatTerpilih
         e. Sorting pada func TampilkanDataTerurut


PS: susunan procedure masih berantakan akan tetapi semua function dan procedure berjalan sesuai kegunaannya
    kode yang dikomen bisa dihapus, hanya draft function awal saja jadi tidak akan ngaruh ke fungsi
	jika memiliki pertanyaan boleh DM ke instagram @adninin_

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Struct untuk kandidat
type Kandidat struct {
	ID     int
	Nama   string
	Partai string
	Suara  int
}

// Struct untuk pengguna
type Pengguna struct {
	Nama     string
	Peran    string // admin atau pemilih
	Kandidat []int  
}

// Struct untuk menyimpan pemilih
type Pemilih struct {
	Nama string
}

// Struct untuk merepresentasikan pemilih dan kandidat yang dipilihnya
type Suara struct {
	Pemilih  string
	Kandidat string
}

// Database kandidat
var kandidats []Kandidat

// Database pengguna
var pengguna = map[string]Pengguna{
	"admin": {"admin", "admin", nil},
}

// Threshold
var ambangSuara = 2

// Database pemilih
var pemilih []Pemilih

// Variabel global untuk melacak status pemilih
var telahMemilih = false

// Variabel hasilSuara untuk menyimpan data hasil suara
var hasilSuara []Suara

// Definisikan batas waktu pemilihan
var waktuMulaiPemilihan = time.Date(2024, time.June, 12, 8, 0, 0, 0, time.Local)    // 11 Juni 2024 pukul 08:00
var waktuSelesaiPemilihan = time.Date(2024, time.June, 12, 17, 0, 0, 0, time.Local) // 11 Juni 2024 pukul 17:00

// Fungsi untuk memeriksa apakah waktu saat ini berada dalam rentang waktu pemilihan
func WaktuPemilihanValid() bool {
	waktuSekarang := time.Now()
	return waktuSekarang.After(waktuMulaiPemilihan) && waktuSekarang.Before(waktuSelesaiPemilihan)
}

// Fungsi untuk memeriksa apakah saat ini adalah waktu pemilihan yang valid
// func WaktuPemilihanValid() bool {
// now := time.Now()
// startTime := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())  // Mulai pukul 08:00
// endTime := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())   // Berakhir pukul 17:00

// return now.After(startTime) && now.Before(endTime)
// }

// Fungsi untuk menambahkan kandidat
func TambahKandidat(nama, partai string) {
	// Periksa apakah nama kandidat sudah ada
	for _, k := range kandidats {
		if k.Nama == nama {
			fmt.Println("Kandidat dengan nama tersebut sudah tersedia!")
			return
		}
	}

	id := len(kandidats) + 1
	kandidat := Kandidat{ID: id, Nama: nama, Partai: partai}
	kandidats = append(kandidats, kandidat)
	fmt.Println("Kandidat baru berhasil ditambahkan.")
	SimpanDataKandidat() // Simpan data kandidat ke file
}

// Fungsi untuk menyimpan data kandidat ke dalam file
func SimpanDataKandidat() {
	file, err := os.Create("kandidat.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan data kandidat!", err)
		return
	}
	defer file.Close()

	for _, kandidat := range kandidats {
		_, err := fmt.Fprintf(file, "%d|%s|%s|%d\n", kandidat.ID, kandidat.Nama, kandidat.Partai, kandidat.Suara)
		if err != nil {
			fmt.Println("Gagal menyimpan data kandidat!", err)
			return
		}
	}
}

// Fungsi untuk membaca data Calon Legislatif dari file
func BacaDataKandidat() {
	file, err := os.Open("kandidat.txt")
	if err != nil {
		fmt.Println("File kandidat tidak ditemukan!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 4 {
			fmt.Println("Format data kandidat tidak valid!", line)
			continue
		}
		id, _ := strconv.Atoi(fields[0])
		suara, _ := strconv.Atoi(fields[3])
		kandidat := Kandidat{ID: id, Nama: fields[1], Partai: fields[2], Suara: suara}
		kandidats = append(kandidats, kandidat)
	}
}

// Fungsi untuk menampilkan Calon Legislatif terurut berdasarkan suara
func TampilkanKandidatTerurut(opsi string) {
	switch opsi {
	case "suara":
		sort.Slice(kandidats, func(i, j int) bool {
			return kandidats[i].Suara > kandidats[j].Suara
		})
	case "nama":
		sort.Slice(kandidats, func(i, j int) bool {
			return kandidats[i].Nama < kandidats[j].Nama
		})
	case "partai":
		sort.Slice(kandidats, func(i, j int) bool {
			if kandidats[i].Partai != kandidats[j].Partai {
				return kandidats[i].Partai < kandidats[j].Partai
			}
			return kandidats[i].Suara > kandidats[j].Suara
		})
	default:
		fmt.Println("Opsi tidak valid!")
		return
	}

	fmt.Println("\nPENCARIAN DATA CALON ANGGOTA LEGISLATIF TERURUT")
	switch opsi {
	case "suara":
		fmt.Println("\nBerdasarkan Jumlah Suara")
	case "nama":
		fmt.Println("\nBerdasarkan Nama Calon")
	case "partai":
		fmt.Println("\nBerdasarkan Partai")
	}
	fmt.Println("\nID    Nama                 Partai               Suara")
	for _, kandidat := range kandidats {
		fmt.Printf("%-5d %-20s %-20s %-5d\n", kandidat.ID, kandidat.Nama, kandidat.Partai, kandidat.Suara)
	}
}

// Fungsi untuk menampilkan daftar Calon Legislatif
func TampilkanKandidat() {
	fmt.Println("\nDAFTAR CALON LEGISLATIF YANG TERSEDIA")
	fmt.Println("\nID    Nama                 Partai")
	for _, kandidat := range kandidats {
		fmt.Printf("%-5d %-20s %-20s\n", kandidat.ID, kandidat.Nama, kandidat.Partai)
	}
}

// Fungsi untuk menampilkan menu berdasarkan peran pengguna
func TampilkanMenu(nama string) {
	if pengguna[nama].Peran == "admin" {
		fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("\nHALO ADMIN! APA YANG INGIN ANDA LAKUKAN?")
		fmt.Println("\n1. Menambahkan Calon Legislatif")
		fmt.Println("2. Mengedit Calon Legislatif")
		fmt.Println("3. Menghapus Calon Legislatif")
		fmt.Println("4. Menampilkan List Calon Legislatif")
		fmt.Println("5. Menambahkan Pemilih Baru")
		fmt.Println("6. Mengedit Pemilih")
		fmt.Println("7. Menghapus Pemilih")
		fmt.Println("8. Menampilkan Pemilih")
		fmt.Println("9. Menampilkan Calon Legislatif Terurut Berdasarkan Perolehan Suara")
		fmt.Println("10. Menampilkan Calon Legislatif Terurut Berdasarkan Nama")
		fmt.Println("11. Menampilkan Calon Legislatif Terurut Berdasarkan Partai")
		fmt.Println("12. Pencarian Data Calon Legislatif")
		fmt.Println("13. Menampilkan Pemenang Yang Memenuhi Ketentuan")
		fmt.Println("14. Mereset Data Voting")
		fmt.Println("15. Keluar")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	} else if pengguna[nama].Peran == "pemilih" {
		// Menampilkan menu untuk pemilih
		if !telahMemilih {
			// Jika pemilih belum memilih, tampilkan opsi pemilihan
			fmt.Print("\nSELAMAT DATANG DI APLIKASI PEMILIHAN UMUM.")
			fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("\nApa yang ingin anda lakukan?")
			fmt.Println("\n1. Memilih Calon Legislatif")
			fmt.Println("2. Keluar")
		} else {
			// Jika pemilih sudah memilih, hilangkan opsi pemilihan
			fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("\nTerimakasih Telah Berpartisipasi Dalam Kegiatan Pemilihan!")
			fmt.Println("2. Keluar")
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}

// Fungsi untuk admin menambahkan Calon Legislatif
func TambahKandidatAdmin() {
	fmt.Print("\nMasukkan data calon legislatif : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	data := strings.Split(strings.TrimSpace(input), " ")
	if len(data) != 2 {
		fmt.Println("Format input tidak valid!")
		return
	}
	TambahKandidat(data[0], data[1])
}

// Fungsi untuk admin mengubah data kandidat
func UbahKandidatAdmin() {
	fmt.Print("\nMasukkan ID calon legislatif yang ingin diubah : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID tidak valid!")
		return
	}

	// Meminta data baru untuk calon legislatif
	fmt.Print("\nMasukkan data baru untuk calon legislatif : ")
	input, _ = reader.ReadString('\n')
	data := strings.Split(strings.TrimSpace(input), " ")
	if len(data) != 2 {
		fmt.Println("Format input tidak valid!")
		return
	}

	// Periksa apakah nama baru sudah digunakan oleh kandidat lain
	for _, k := range kandidats {
		if k.Nama == data[0] && k.ID != id {
			fmt.Println("Kandidat dengan nama tersebut sudah tersedia!")
			return
		}
	}

	// Lakukan perubahan jika nama baru tidak digunakan oleh kandidat lain
	UbahKandidat(id, data[0], data[1])
}

// Fungsi untuk admin menghapus data kandidat
func HapusKandidatAdmin() {
	fmt.Print("\nMasukkan ID kandidat yang ingin dihapus : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID tidak valid!")
		return
	}
	HapusKandidat(id)
}

// Fungsi untuk melakukan voting
func PilihKandidat(nama string, id int) {
	penggunaStruct, ok := pengguna[nama]
	if !ok {
		fmt.Println("Pengguna tidak ditemukan!")
		return
	}

	penggunaStruct.Kandidat = append(penggunaStruct.Kandidat, id)

	// Meningkatkan jumlah suara kandidat
	for i := range kandidats {
		if kandidats[i].ID == id {
			kandidats[i].Suara++
			break
		}
	}

	pengguna[nama] = penggunaStruct

	// Simpan hasil suara
	hasilSuara = append(hasilSuara, Suara{Pemilih: nama, Kandidat: strconv.Itoa(id)})

	// Simpan data kandidat ke file setiap kali ada perubahan suara
	SimpanDataKandidat()
	// Simpan hasil suara setiap kali pemilih memilih
	SimpanHasilVoting()
}

// LoadVotingResults memuat hasil voting dari file ke dalam map
func LoadVotingResults(fileName string) (map[string]string, error) {
    // Buat map untuk menyimpan hasil voting
    votingResults := make(map[string]string)

    // Buka file hasil_voting.txt
    file, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Buat scanner untuk membaca file baris per baris
    scanner := bufio.NewScanner(file)

    // Loop melalui setiap baris file
    for scanner.Scan() {
        // Split baris menjadi dua bagian, nama pemilih dan kandidat yang dipilih
        parts := strings.Split(scanner.Text(), "\t")
        if len(parts) != 2 {
            return nil, fmt.Errorf("Format file tidak valid: %s", scanner.Text())
        }

        // Trim spasi dari nama pemilih dan kandidat
        pemilih := strings.TrimSpace(parts[0])
        kandidat := strings.TrimSpace(parts[1])

        // Tambahkan hasil voting ke dalam map
        votingResults[pemilih] = kandidat
    }

    // Periksa apakah ada kesalahan dalam proses pemindaian file
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return votingResults, nil
}

func TampilkanKandidatPemilih(namaPemilih string) {
    // Buka file hasil_voting.txt
    file, err := os.Open("hasil_voting.txt")
    if err != nil {
        fmt.Println("Gagal membuka file:", err)
        return
    }
    defer file.Close()

    // Buat scanner untuk membaca file baris per baris
    scanner := bufio.NewScanner(file)

    // Loop melalui setiap baris file
    found := false // Menandai apakah pemilih ditemukan atau tidak
    for scanner.Scan() {
        // Split baris menjadi dua bagian, nama pemilih dan kandidat yang dipilih
        parts := strings.Split(scanner.Text(), "\t")
        if len(parts) != 2 {
            fmt.Println("Format file tidak valid:", scanner.Text())
            continue
        }

        // Trim spasi dari nama pemilih dan kandidat
        pemilih := strings.TrimSpace(parts[0])
        kandidat := strings.TrimSpace(parts[1])

        // Cek apakah nama pemilih cocok dengan yang dicari
        if pemilih == namaPemilih {
            fmt.Printf("Pemilih %s memilih kandidat %s\n", pemilih, kandidat)
            found = true
            break // Keluar dari loop setelah menemukan pemilih
        }
    }

    // Jika nama pemilih tidak ditemukan
    if !found {
        fmt.Printf("Pemilih %s tidak ditemukan atau belum memilih kandidat\n", namaPemilih)
    }
}

// Fungsi untuk menampilkan Kandidat yang memenuhi ambang batas suara
func TampilkanKandidatTerpilih(ambangSuara int) {
	fmt.Printf("\nKANDIDAT YANG MEMENUHI KETENTUAN SUARA : %d SUARA\n", ambangSuara)
	fmt.Println("ID    Nama                 Partai               Suara")
	for _, kandidat := range kandidats {
		if kandidat.Suara >= ambangSuara {
			fmt.Printf("%-5d %-20s %-20s %-5d\n", kandidat.ID, kandidat.Nama, kandidat.Partai, kandidat.Suara)
		}
	}
}

// Fungsi untuk mengubah data kandidat berdasarkan ID
func UbahKandidat(id int, nama, partai string) {
	for i, kandidat := range kandidats {
		if kandidat.ID == id {
			kandidats[i].Nama = nama
			kandidats[i].Partai = partai
			fmt.Println("Data berhasil diubah!")
			SimpanDataKandidat()
			return
		}
	}
	fmt.Println("Data tidak ditemukan!")
}

// Fungsi untuk menghapus data kandidat berdasarkan ID
func HapusKandidat(id int) {
	for i, kandidat := range kandidats {
		if kandidat.ID == id {
			kandidats = append(kandidats[:i], kandidats[i+1:]...)
			fmt.Println("Data kandidat berhasil dihapus!")
			SimpanDataKandidat()
			return
		}
	}
	fmt.Println("Data tidak ditemukan!")
}

// Fungsi untuk menambahkan pemilih baru
func TambahPemilih(nama, peran string) {
	_, ok := pengguna[nama]
	if ok {
		fmt.Println("Pemilih dengan nama pengguna tersebut sudah tersedia!")
		return
	}
	pengguna[nama] = Pengguna{Nama: nama, Peran: peran, Kandidat: nil}
	fmt.Println("Pemilih baru berhasil ditambahkan!")
	SimpanDataPemilih() // Simpan data pemilih ke file
}

// Fungsi untuk mengubah data pemilih
func UbahPemilih(namaBaru, namaLama string) {
	pemilih, ok := pengguna[namaLama]
	if !ok {
		fmt.Println("Pemilih tidak ditemukan!")
		return
	}
	pemilih.Nama = namaBaru
	delete(pengguna, namaLama)
	pengguna[namaBaru] = pemilih
	fmt.Println("Data pemilih berhasil diubah!")
	SimpanDataPemilih() // Simpan data pemilih ke file
}

// Fungsi untuk menghapus data pemilih
func HapusPemilih(nama string) {
	_, ok := pengguna[nama]
	if !ok {
		fmt.Println("Pemilih tidak ditemukan!")
		return
	}
	delete(pengguna, nama)
	fmt.Println("Data pemilih berhasil dihapus!")
	HapusDataPemilih(nama) // Hapus data pemilih dari file
}

// Fungsi untuk menyimpan data pemilih ke dalam file
func SimpanDataPemilih() {
	file, err := os.Create("pemilih.txt")
	if err != nil {
		fmt.Println("Gagal menyimpan data pemilih!", err)
		return
	}
	defer file.Close()

	for _, p := range pengguna {
		if p.Peran == "pemilih" {
			_, err := fmt.Fprintf(file, "%s|%s\n", p.Nama, p.Peran)
			if err != nil {
				fmt.Println("Gagal menyimpan data pemilih!", err)
				return
			}
		}
	}
}

// Fungsi untuk membaca data pemilih dari file
func BacaDataPemilih() {
	file, err := os.Open("pemilih.txt")
	if err != nil {
		fmt.Println("File pemilih tidak ditemukan!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 2 {
			fmt.Println("Format data pemilih tidak valid!", line)
			continue
		}
		nama := fields[0]
		peran := fields[1]
		pengguna[nama] = Pengguna{Nama: nama, Peran: peran, Kandidat: nil}
	}
}

// Fungsi untuk menghapus data pemilih dari file
func HapusDataPemilih(nama string) {
	file, err := os.ReadFile("pemilih.txt")
	if err != nil {
		fmt.Println("Gagal membaca data pemilih!", err)
		return
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		if strings.Contains(line, nama) {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}
	output := strings.Join(lines, "\n")
	if err := os.WriteFile("pemilih.txt", []byte(output), 0644); err != nil {
		fmt.Println("Gagal menghapus data pemilih!", err)
	}
}

// Fungsi untuk menyimpan hasil voting
func SimpanHasilVoting() {
    file, err := os.OpenFile("hasil_voting.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Gagal menyimpan hasil voting!", err)
        return
    }
    defer file.Close()

    for _, hasil := range hasilSuara {
        _, err := fmt.Fprintf(file, "%s\t%s\n", hasil.Pemilih, hasil.Kandidat)
        if err != nil {
            fmt.Println("Gagal menyimpan hasil suara!", err)
            return
        }
    }
    fmt.Println("Hasil voting berhasil disimpan.")
}


// Fungsi untuk mengatur ulang suara semua kandidat menjadi 0
func ResetSuaraKandidatAdmin() {
	for i := range kandidats {
		kandidats[i].Suara = 0
	}
	SimpanDataKandidat() // Simpan perubahan ke file
}

// Fungsi untuk menghapus file data voting dan membuat file baru
func HapusDataVotingAdmin() {
    // Hapus file hasil suara
    err := os.Remove("hasil_voting.txt")
    if err != nil {
        fmt.Println("Gagal menghapus file hasil suara!", err)
        return
    }
    // Buat file baru untuk hasil suara
    _, err = os.Create("hasil_voting.txt")
    if err != nil {
        fmt.Println("Gagal membuat file hasil suara baru!", err)
        return
    }

    fmt.Println("Data voting berhasil dihapus.")
}



// Fungsi untuk menampilkan daftar pemilih
func TampilkanPemilih() {
	fmt.Println("\nLIST PEMILIH YANG TERDAFTAR\n")
	for nama, p := range pengguna {
		if p.Peran == "pemilih" {
			fmt.Printf("Nama : %s, sebagai %s\n", nama, p.Peran)
		}
	}
}

// Fungsi untuk menampilkan hasil suara
func TampilkanHasilSuara() {
	fmt.Println("Daftar Nama Pemilih dan Kandidat yang Dipilih")
	fmt.Println("ID     Nama          Kandidat")
	for _, hasil := range hasilSuara {
		pemilih := hasil.Pemilih
		var kandidat string
		for _, k := range kandidats {
			if strconv.Itoa(k.ID) == hasil.Kandidat {
				kandidat = k.Nama
				break
			}
		}
		fmt.Printf("%-6s %-13s %s\n", hasil.Kandidat, pemilih, kandidat)
	}
}

// Fungsi untuk mengatur ulang status pemilih menjadi belum memilih
func ResetStatusPemilihAdmin() {
	for nama := range pengguna {
		if pengguna[nama].Peran == "pemilih" {
			pengguna[nama] = Pengguna{
				Nama:     pengguna[nama].Nama,
				Peran:    "pemilih",
				Kandidat: nil,
			}
		}
	}
	SimpanDataPemilih() // Simpan perubahan ke file
}

// Fungsi untuk menangani opsi yang dipilih oleh pengguna
func HandleOpsi(nama string, opsi string) {
	reader := bufio.NewReader(os.Stdin)
	switch opsi {
	case "1":
		// Opsi untuk menambahkan calon legislatif
		if pengguna[nama].Peran == "admin" {
			TambahKandidatAdmin()
		} else {
			// Opsi untuk pemilih
			if !WaktuPemilihanValid() {
				fmt.Println("Waktu pemilihan sudah berakhir! Silahkan konfirmasi ke admin.")
				TampilkanKandidat()
				break
			}
			if telahMemilih {
				fmt.Println("Terimakasih! Telah Berpartisipasi Dalam Kegiatan Pemilihan.")
			} else {
				TampilkanKandidat()
				fmt.Print("\nMasukkan ID calon legislatif yang ingin Anda pilih : ")
				input, _ := reader.ReadString('\n')
				id, err := strconv.Atoi(strings.TrimSpace(input))
				if err != nil {
					fmt.Println("ID kandidat tidak valid!")
					return
				}
				PilihKandidat(nama, id)

				// Setelah pemilih memilih, ubah status menjadi true
				telahMemilih = true
			}
		}

	case "2":
		// Opsi untuk mengedit calon legislatif
		if pengguna[nama].Peran == "admin" {
			UbahKandidatAdmin()
		}
		if pengguna[nama].Peran == "pemilih" {
			fmt.Println("\nTerima kasih sudah menggunakan aplikasi pemilihan umum. Sampai jumpa!")
			os.Exit(0)
		}

	case "3":
		// Opsi untuk menghapus calon legislatif
		if pengguna[nama].Peran == "admin" {
			HapusKandidatAdmin()
		}

	case "4":
		// Opsi untuk menampilkan daftar calon legislatif
		TampilkanKandidat()

	case "5":
		// Opsi untuk menambahkan pemilih baru (hanya admin)
		if pengguna[nama].Peran == "admin" {
			fmt.Print("\nMasukkan nama pemilih baru : ")
			input, _ := reader.ReadString('\n')
			nama := strings.TrimSpace(input)
			TambahPemilih(nama, "pemilih")
		}

	case "6":
		// Opsi untuk mengubah data pemilih (hanya admin)
		if pengguna[nama].Peran == "admin" {
			fmt.Print("\nMasukkan nama pemilih yang ingin diubah : ")
			input, _ := reader.ReadString('\n')
			namaLama := strings.TrimSpace(input)
			fmt.Print("\nMasukkan nama baru untuk pemilih : ")
			input, _ = reader.ReadString('\n')
			namaBaru := strings.TrimSpace(input)
			UbahPemilih(namaBaru, namaLama)
		}

	case "7":
		// Opsi untuk menghapus data pemilih (hanya admin)
		if pengguna[nama].Peran == "admin" {
			fmt.Print("\nMasukkan nama pemilih yang ingin dihapus : ")
			input, _ := reader.ReadString('\n')
			nama := strings.TrimSpace(input)
			HapusPemilih(nama)
		}

	case "8":
		// Opsi untuk menampilkan daftar pemilih (hanya admin)
		if pengguna[nama].Peran == "admin" {
			TampilkanPemilih()
		}

	case "9":
		// Opsi untuk menampilkan calon legislatif terurut berdasarkan perolehan suara
		TampilkanKandidatTerurut("suara")

	case "10":
		// Opsi untuk menampilkan calon legislatif terurut berdasarkan nama
		TampilkanKandidatTerurut("nama")

	case "11":
		// Opsi untuk menampilkan calon legislatif terurut berdasarkan partai
		TampilkanKandidatTerurut("partai")

	case "12":
		// Opsi untuk pencarian data calon
		if pengguna[nama].Peran == "admin" {
			fmt.Print("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("\nPENCARIAN DATA CALON LEGISLATIF : ")
			fmt.Println("\n1. Pencarian Data Berdasarkan Nama Calon")
			fmt.Println("2. Pencarian Data Berdasarkan Partai Calon")
			fmt.Println("3. Pencarian Data Berdasarkan Nama Pemilih")
			fmt.Println("4. Kembali Ke Menu Utama")
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Print("\nMasukkan opsi pencarian : ")
			opsiPencarian, _ := reader.ReadString('\n')
			opsiPencarian = strings.TrimSpace(opsiPencarian)
			HandlePencarian(nama, opsiPencarian)
		} else {
			fmt.Println("Opsi tidak valid!")
		}

	case "13":
		// Opsi untuk menampilkan kandidat yang memenuhi ambang batas suara
		TampilkanKandidatTerpilih(ambangSuara)

	case "14":
		// Opsi untuk mereset data voting (hanya admin)
		if pengguna[nama].Peran == "admin" {
			ResetSuaraKandidatAdmin()
			HapusDataVotingAdmin()
			ResetStatusPemilihAdmin()
		}

	case "15":
		// Opsi untuk keluar
		fmt.Println("\nTerima kasih sudah menggunakan aplikasi pemilihan umum. Sampai jumpa!")
		os.Exit(0)

	default:
		fmt.Println("Opsi tidak valid")
	}
}

// Fungsi untuk menangani opsi pencarian data calon
func HandlePencarian(nama string, opsi string) {
	reader := bufio.NewReader(os.Stdin)
	switch opsi {
	case "1":
		// Pencarian berdasarkan nama calon
		fmt.Print("\nMasukkan nama calon legislatif yang ingin dicari : ")
		namaCalon, _ := reader.ReadString('\n')
		namaCalon = strings.TrimSpace(namaCalon)
		CariCalonBerdasarkanNama(namaCalon)

	case "2":
		// Pencarian berdasarkan partai calon
		fmt.Print("\nMasukkan partai calon legislatif yang ingin dicari : ")
		partaiCalon, _ := reader.ReadString('\n')
		partaiCalon = strings.TrimSpace(partaiCalon)
		CariCalonBerdasarkanPartai(partaiCalon)

	case "3":
		// Pencarian berdasarkan nama pemilih dari calon tertentu
		fmt.Print("\nMasukkan nama pemilih : ")
		namaPemilih, _ := reader.ReadString('\n')
		namaPemilih = strings.TrimSpace(namaPemilih)
	
		// Panggil fungsi TampilkanKandidatPemilih untuk menampilkan kandidat yang dipilih oleh pemilih dengan nama tertentu
		TampilkanKandidatPemilih(namaPemilih)
	

	case "4":
		// Kembali ke menu utama
		fmt.Println("Kembali ke Menu Utama")

	default:
		fmt.Println("Opsi tidak valid")
	}

	// Simpan hasil voting setiap kali pemilih memilih
	if opsi == "1" && pengguna[nama].Peran != "admin" {
		SimpanHasilVoting()
	}
}

// Fungsi untuk mencari calon berdasarkan nama
func CariCalonBerdasarkanNama(nama string) {
	found := false
	for _, kandidat := range kandidats {
		if strings.Contains(strings.ToLower(kandidat.Nama), strings.ToLower(nama)) {
			fmt.Printf("ID: %d, Nama: %s, Partai: %s\n", kandidat.ID, kandidat.Nama, kandidat.Partai)
			found = true
		}
	}
	if !found {
		fmt.Println("Calon dengan nama tersebut tidak ditemukan.")
	}
}

// Fungsi untuk mencari calon berdasarkan partai
func CariCalonBerdasarkanPartai(partai string) {
	found := false
	for _, kandidat := range kandidats {
		if strings.Contains(strings.ToLower(kandidat.Partai), strings.ToLower(partai)) {
			fmt.Printf("ID: %d, Nama: %s, Partai: %s\n", kandidat.ID, kandidat.Nama, kandidat.Partai)
			found = true
		}
	}
	if !found {
		fmt.Println("Calon dari partai tersebut tidak ditemukan.")
	}
}

// func CariNamaPemilihDariCalon(idCalon int) []string {
// 	fmt.Printf("Hasil Pencarian Nama Pemilih dari Calon ID %d:\n", idCalon)
// 	var pemilih []string
// 	found := false
// 	for _, hasil := range hasilSuara {
// 		if hasil.Kandidat == strconv.Itoa(idCalon) {
// 			fmt.Printf("Nama Pemilih: %s\n", hasil.Pemilih)
// 			pemilih = append(pemilih, hasil.Pemilih)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Println("Pemilih dari calon tersebut tidak ditemukan.")
// 	}
// 	return pemilih
// }

func main() {
	BacaDataKandidat()
	BacaDataPemilih()

	// Masukkan nama pengguna di awal saja
	fmt.Print("Masukkan nama pengguna : ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	for {
		_, ok := pengguna[nama]
		if !ok {
			fmt.Println("Pengguna tidak ditemukan!")
			fmt.Print("\nMasukkan nama pengguna : ")
			nama, _ = reader.ReadString('\n')
			nama = strings.TrimSpace(nama)
			continue
		}

		for {
			// Tampilkan menu setelah nama pengguna ditemukan
			TampilkanMenu(nama)

			// Pilih opsi
			fmt.Print("\nMasukkan opsi : ")
			opsi, _ := reader.ReadString('\n')
			opsi = strings.TrimSpace(opsi)

			// Tangani opsi yang dipilih
			HandleOpsi(nama, opsi)

			// Kembali ke menu awal
			if opsi == "13" {
				break
			}
		}
	}
}
