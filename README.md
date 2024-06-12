# Aplikasi-Pemilihan-Umum

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
