

A. Tujuan dan Alasan Program Ini Dibuat

Program ini dibuat untuk mengelola data akademik secara sederhana, meliputi Mahasiswa, Dosen, dan Kelas. Tujuan utama dari program ini adalah:

•	Memudahkan pengguna (pengguna / ADMISI) dalam melakukan pengolahan data akademik;

•	Menyediakan fitur input, validasi, pencarian, pengurutan, dan tampilan data dalam satu aplikasi CLI (Command Line Interface);

•	Menjadi latihan atau implementasi konsep struktur data (array, sorting, searching) dan pemrograman prosedural dalam bahasa Go;

B. Fitur-Fitur Program

•	Input Data : Menginput data Mahasiswa, Dosen, dan Kelas dengan validasi nama dan ID;

•	Tampilkan Data : Menampilkan semua data yang telah diinput dalam format tabel;

•	Sorting Data : Mengurutkan data berdasarkan berbagai kriteria (nama, NIM, NIDN, ruangan) menggunakan selection/insertion sort;

•	Searching Data : Mencari data berdasarkan nama (Mahasiswa/Dosen) atau nama kelas, menggunakan linear/binary search;

•	Validasi Input : Memastikan input berupa angka/huruf sesuai kebutuhan;

•	Menu Interaktif : Menu utama dan sub-menu yang memandu pengguna menggunakan fitur yang ada;

C.	Alur Kerja Setiap Fitur;

1.	Input Data;

•	Ada beberapa fungsi dalam input data yaitu :

	a)	inputDataMahasiswa();

	b)	inputDataDosen();

	c)	inputDataKelas();

•	Alur : 

a)	User memilih menu input data;

b)	Program meminta jumlah data yang ingin diinput;

c)	Untuk setiap data:

i.	Meminta nama → divalidasi dengan validasiNama();

ii.	Meminta NIM/NIDN/Kode kelas/ruangan → divalidasi dengan validasiAngka();

d)	Data disimpan ke array mahasiswa, dosen, atau kelas;

2.	Tampilkan Data;

•	Ada beberapa Fungsi untuk tampilkan data yaitu :

	a)	tampilkanDataMahasiswa();

	b)	tampilkanDataDosen();

	c)	tampilkanDataKelas();

•	Alur :

a)	User memilih menu tampilkan data;

b)	Program menampilkan tabel data sesuai pilihan (Mahasiswa/Dosen/Kelas);

3.	Sorting (Pengurutan);

•	Ada 6 Fungsi dalam sorting yaitu :

	a)	selectionSortMahasiswaByNIM();

	b)	insertionSortMahasiswaByNama();

	c)	selectionSortDosenByNIDN();
	
	d)	insertionSortDosenByNama();

	e)	selectionSortKelasByRuangan();

	f)	insertionSortKelasByNamaKelas();

•	Alur :

a)	User memilih entitas dan metode sorting;

b)	Program mengurutkan array berdasarkan kriteria:

i.	Mahasiswa: Nama (A–Z) / NIM (kecil–besar);

ii.	Dosen: Nama / NIDN;

iii.	Kelas: Nama / Ruangan;

c)	Hasil sorting ditampilkan langsung;

4.	Searching (Pencarian)

•	Ada 6 Fungsi dalam searching yaitu :

	a)	linearSearchMahasiswaByNama();

	b)	binarySearchMahasiswaByNama();

	c)	linearSearchDosenByNama();

	d)	binarySearchDosenByNama();

	e)	linearSearchKelasByNamaKelas();

	f)	binarySearchKelasByNamaKelas();

•	Alur:

a)	User memilih entitas dan metode searching;

b)	Input kata kunci nama yang dicari;

c)	Jika binary search: Array diurutkan terlebih dahulu (ascending);

d)	Program mencari nama sesuai input:

i.	Jika ditemukan → tampilkan data;

ii.	Jika tidak → tampilkan pesan tidak ditemukan;

5.	Validasi;

•	Ada 2 Fungsi validasi :

	a)	validasiAngka() : "hanya menerima digit 0–9";

	b)	validasiNama() : "hanya menerima huruf A–Z, a–z, dan tanpa menerima spasi";

•	Alur:

a)	Dipanggil setiap kali user menginput nama atau ID;

b)	Jika tidak valid → user diminta input ulang / kembali ke menu awal; 

6.	Menu interaktif;

•	Alur:

a)	Ditangani dalam fungsi main();

b)	Menggunakan switch dan for untuk menavigasi pilihan pengguna;

c)	Program akan terus berjalan sampai user memilih keluar (0);






Anggota Kelompok :

Ridho Bintang Adwitya - 103072400015;

Restu Fadilah Al Fatah - 103072400081;

Albertrio Suranta Ginting - 103072400128;

