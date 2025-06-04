package main

import (
	"fmt"
	"sort"
)

type DataMahasiswa struct {
	Nama string
	NIM  string
}

type DataDosen struct {
	Nama string
	NIDN string
}

type DataKelas struct {
	NamaKelas string
	Ruangan   string
}

type Mahasiswa [100]DataMahasiswa
type Dosen [100]DataDosen
type Kelas [100]DataKelas

func selectionSortingMahasiswa(data *[100]DataMahasiswa, n int, ascending bool) {
    for i := 0; i < n-1; i++ {
        minmaxIndex := i
        for j := i + 1; j < n; j++ {
            if ascending {
                if data[j].NIM < data[minmaxIndex].NIM {
                    minmaxIndex = j
                }
            } else {
                if data[j].NIM > data[minmaxIndex].NIM {
                    minmaxIndex = j
                }
            }
        }
        data[i], data[minmaxIndex] = data[minmaxIndex], data[i]
    }
}

func insertionSortingMahasiswa(data *[100]DataMahasiswa, n int, ascending bool) {
    for i := 1; i < n; i++ {
        key := data[i]
        j := i - 1
        for j >= 0 && ((ascending && data[j].Nama > key.Nama) || (!ascending && data[j].Nama < key.Nama)) {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = key
    }
}

func linearSearchMahasiswa(data *[100]DataMahasiswa, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].Nama) >= len(prefix) && data[i].Nama[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}

func binarySearchPrefixMahasiswa(data *[100]DataMahasiswa, n int, prefix string) []int {
    temp := make([]DataMahasiswa, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
    bawah, atas := 0, n-1
    var hasil []int

    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].Nama) >= len(prefix) && temp[tengah].Nama[:len(prefix)] == prefix {
            i := tengah
            for i >= 0 && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--

			}
            i = tengah + 1

            for i < n && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i++
            }
            break

        } else if temp[tengah].Nama < prefix {
            bawah = tengah + 1
        } else {
            atas = tengah - 1
        }
    }
    return hasil
}

func selectionSortingDosen(data *[100]DataDosen, n int, ascending bool) {
    for i := 0; i < n-1; i++ {
        minmaxIndex := i
        for j := i + 1; j < n; j++ {
            if ascending {
                if data[j].NIDN < data[minmaxIndex].NIDN {
                    minmaxIndex = j
                }
            } else {
                if data[j].NIDN > data[minmaxIndex].NIDN {
                    minmaxIndex = j
                }
            }
        }
        data[i], data[minmaxIndex] = data[minmaxIndex], data[i]
    }
}

func insertionSortingDosen(data *[100]DataDosen, n int, ascending bool) {
    for i := 1; i < n; i++ {
        key := data[i]
        j := i - 1
        for j >= 0 && ((ascending && data[j].Nama > key.Nama) || (!ascending && data[j].Nama < key.Nama)) {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = key
    }
}

func linearSearchDosen(data *[100]DataDosen, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].Nama) >= len(prefix) && data[i].Nama[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}

func binarySearchDosen(data *[100]DataDosen, n int, prefix string) []int {
    temp := make([]DataDosen, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
    bawah, atas := 0, n-1
    var hasil []int

    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].Nama) >= len(prefix) && temp[tengah].Nama[:len(prefix)] == prefix {
            i := tengah
            for i >= 0 && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--
            }
            i = tengah + 1

            for i < n && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i++
            }
            break

        } else if temp[tengah].Nama < prefix {
            bawah = tengah + 1
        } else {
            atas = tengah - 1
        }
    }
    return hasil
}

func selectionSortingKelas(data *[100]DataKelas, n int, ascending bool) {
    for i := 0; i < n-1; i++ {
        minmaxIndex := i
        for j := i + 1; j < n; j++ {
            if ascending {
                if data[j].Ruangan < data[minmaxIndex].Ruangan {
                    minmaxIndex = j
                }
            } else {
                if data[j].Ruangan > data[minmaxIndex].Ruangan {
                    minmaxIndex = j
                }
            }
        }
        data[i], data[minmaxIndex] = data[minmaxIndex], data[i]
    }
}

func insertionSortingKelas(data *[100]DataKelas, n int, ascending bool) {
    for i := 1; i < n; i++ {
        key := data[i]
        j := i - 1
        for j >= 0 && ((ascending && data[j].NamaKelas > key.NamaKelas) || (!ascending && data[j].NamaKelas < key.NamaKelas)) {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = key
    }
}

func linearSearchKelas(data *[100]DataKelas, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].NamaKelas) >= len(prefix) && data[i].NamaKelas[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}

func binarySearchKelas(data *[100]DataKelas, n int, prefix string) []int {
    temp := make([]DataKelas, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].NamaKelas < temp[j].NamaKelas })
    bawah, atas := 0, n-1
    var hasil []int
	
    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].NamaKelas) >= len(prefix) && temp[tengah].NamaKelas[:len(prefix)] == prefix {
            i := tengah
            for i >= 0 && len(temp[i].NamaKelas) >= len(prefix) && temp[i].NamaKelas[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--
            }
            i = tengah + 1

            for i < n && len(temp[i].NamaKelas) >= len(prefix) && temp[i].NamaKelas[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i++
            }
            break

        } else if temp[tengah].NamaKelas < prefix {
            bawah = tengah + 1
        } else {
            atas = tengah - 1
        }
    }
    return hasil
}


func inputDataMahasiswa(data *[100]DataMahasiswa, n *int) {
    var tambah int
    fmt.Print("Masukkan jumlah mahasiswa yang ingin ditambahkan: ")

    valid := false
    for !valid {
        _, err := fmt.Scan(&tambah)
        if err != nil || tambah <= 0 || *n+tambah > 100 {
            fmt.Println("Input jumlah mahasiswa harus berupa angka dan total tidak lebih dari 100!")
            fmt.Print("Masukkan jumlah mahasiswa yang ingin ditambahkan: ")
            var dummy string
            fmt.Scanln(&dummy)
        } else {
            valid = true
        }
    }

    for i := *n; i < *n+tambah; i++ {
        validNama := false
        for !validNama {
            fmt.Printf("Nama mahasiswa ke-%d: ", i+1)
            _, err := fmt.Scan(&data[i].Nama)
            if err != nil || data[i].Nama == "" {
                fmt.Println("Nama tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                isAngka := true
                for _, c := range data[i].Nama {
                    if c < '0' || c > '9' {
                        isAngka = false
                        break
                    }
                }
                if isAngka {
                    fmt.Println("Nama tidak boleh berupa angka!")
                } else {
                    validNama = true
                }
            }
        }
        validNIM := false
        for !validNIM {
            fmt.Printf("NIM mahasiswa ke-%d: ", i+1)
            _, err := fmt.Scan(&data[i].NIM)
            if err != nil || data[i].NIM == "" {
                fmt.Println("NIM tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                isAngka := true
                for _, c := range data[i].NIM {
                    if c < '0' || c > '9' {
                        isAngka = false
                        break
                    }
                }
                if !isAngka {
                    fmt.Println("NIM harus berupa angka!")
                } else {
                    validNIM = true
                }
            }
        }
    }
    *n += tambah // tambah jumlah mahasiswa
}

func inputDataDosen(data *[100]DataDosen, n *int) {
    var tambah int
    fmt.Print("Masukkan jumlah dosen yang ingin ditambah: ")

    valid := false
    for !valid {
        _, err := fmt.Scan(&tambah)
        if err != nil || tambah <= 0 || *n+tambah > 100 {
            fmt.Println("Input jumlah dosen harus berupa angka dan total tidak boleh lebih dari 100!")
            fmt.Print("Masukkan jumlah dosen yang ingin ditambah: ")
            var dummy string
            fmt.Scanln(&dummy)
        } else {
            valid = true
        }
    }

    for i := 0; i < tambah; i++ {
        idx := *n + i
        validNama := false
        for !validNama {
            fmt.Printf("Nama dosen ke-%d: ", idx+1)
            _, err := fmt.Scan(&data[idx].Nama)
            if err != nil || data[idx].Nama == "" {
                fmt.Println("Nama tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                isAngka := true
                for _, c := range data[idx].Nama {
                    if c < '0' || c > '9' {
                        isAngka = false
                        break
                    }
                }
                if isAngka {
                    fmt.Println("Nama tidak boleh berupa angka!")
                } else {
                    validNama = true
                }
            }
        }
        validNIDN := false
        for !validNIDN {
            fmt.Printf("NIDN dosen ke-%d: ", idx+1)
            _, err := fmt.Scan(&data[idx].NIDN)
            if err != nil || data[idx].NIDN == "" {
                fmt.Println("NIDN tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                isAngka := true
                for _, c := range data[idx].NIDN {
                    if c < '0' || c > '9' {
                        isAngka = false
                        break
                    }
                }
                if !isAngka {
                    fmt.Println("NIDN harus berupa angka!")
                } else {
                    validNIDN = true
                }
            }
        }
    }
    *n += tambah
}

func inputDataKelas(data *[100]DataKelas, n *int) {
    var tambah int
    fmt.Print("Masukkan jumlah kelas yang ingin ditambah: ")

    valid := false
    for !valid {
        _, err := fmt.Scan(&tambah)
        if err != nil || tambah <= 0 || *n+tambah > 100 {
            fmt.Println("Input jumlah kelas harus berupa angka dan total tidak boleh lebih dari 100!")
            fmt.Print("Masukkan jumlah kelas yang ingin ditambah: ")
            var dummy string
            fmt.Scanln(&dummy)
        } else {
            valid = true
        }
    }

    for i := 0; i < tambah; i++ {
        idx := *n + i
        validNamaKelas := false
        for !validNamaKelas {
            fmt.Printf("Nama kelas ke-%d: ", idx+1)
            _, err := fmt.Scan(&data[idx].NamaKelas)
            if err != nil || data[idx].NamaKelas == "" {
                fmt.Println("Nama kelas tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                isAngka := true
                for _, c := range data[idx].NamaKelas {
                    if c < '0' || c > '9' {
                        isAngka = false
                        break
                    }
                }
                if isAngka {
                    fmt.Println("Nama kelas tidak boleh berupa angka!")
                } else {
                    validNamaKelas = true
                }
            }
        }
        validRuangan := false
        for !validRuangan {
            fmt.Printf("Ruangan kelas ke-%d: ", idx+1)
            _, err := fmt.Scan(&data[idx].Ruangan)
            if err != nil || data[idx].Ruangan == "" {
                fmt.Println("Ruangan tidak boleh kosong!")
                var dummy string
                fmt.Scanln(&dummy)
            } else {
                validRuangan = true
            }
        }
    }
    *n += tambah
}

func tampilkanDataMahasiswa(data *[100]DataMahasiswa, n int) {
	fmt.Println("Daftar Mahasiswa:")
	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s, NIM: %s\n", i+1, data[i].Nama, data[i].NIM)
	}
}

func tampilkanDataDosen(data *[100]DataDosen, n int) {
	fmt.Println("Daftar Dosen:")
	if n == 0 {
		fmt.Println("Belum ada data dosen.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s, NIDN: %s\n", i+1, data[i].Nama, data[i].NIDN)
	}
}

func tampilkanDataKelas(data *[100]DataKelas, n int) {
	fmt.Println("Daftar Kelas:")
	if n == 0 {
		fmt.Println("Belum ada data kelas.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama Kelas: %s, Ruangan: %s\n", i+1, data[i].NamaKelas, data[i].Ruangan)
	}
}

func menuSortMahasiswa(data *[100]DataMahasiswa, n int) {
    var pilih, urut int
    fmt.Print("1. Selection Sort Mahasiswa\n2. Insertion Sort Mahasiswa\nPilih: ")
    fmt.Scan(&pilih)
    fmt.Print("Urutan: 1. Ascending 2. Descending\nPilih: ")
    fmt.Scan(&urut)
    ascending := urut == 1

    if pilih == 1 {
        selectionSortingMahasiswa(data, n, ascending)
        fmt.Println("Setelah selection sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].Nama, data[i].NIM)
        }
    } else if pilih == 2 {
        insertionSortingMahasiswa(data, n, ascending)
        fmt.Println("Setelah insertion sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].Nama, data[i].NIM)
        }
    } else {
        fmt.Println("Pilihan tidak valid!")
        return
    }
}

func menuSortDosen(data *[100]DataDosen, n int) {
    var pilih, urut int
    fmt.Print("1. Selection Sort Dosen\n2. Insertion Sort Dosen\nPilih: ")
    fmt.Scan(&pilih)
    fmt.Print("Urutan: 1. Ascending 2. Descending\nPilih: ")
    fmt.Scan(&urut)
    ascending := urut == 1

    if pilih == 1 {
        selectionSortingDosen(data, n, ascending)
        fmt.Println("Setelah selection sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].Nama, data[i].NIDN)
        }
    } else if pilih == 2 {
        insertionSortingDosen(data, n, ascending)
        fmt.Println("Setelah insertion sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].Nama, data[i].NIDN)
        }
    } else {
        fmt.Println("Pilihan tidak valid!")
        return
    }
}



func menuSortKelas(data *[100]DataKelas, n int) {
    var pilih, urut int
    fmt.Print("1. Selection Sort Kelas\n2. Insertion Sort Kelas\nPilih: ")
    fmt.Scan(&pilih)
    fmt.Print("Urutan: 1. Ascending 2. Descending\nPilih: ")
    fmt.Scan(&urut)
    ascending := urut == 1

    if pilih == 1 {
        selectionSortingKelas(data, n, ascending)
        fmt.Println("Setelah selection sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].NamaKelas, data[i].Ruangan)
        }
    } else if pilih == 2 {
        insertionSortingKelas(data, n, ascending)
        fmt.Println("Setelah insertion sort:")
        for i := 0; i < n; i++ {
            fmt.Println(data[i].NamaKelas, data[i].Ruangan)
        }
    } else {
        fmt.Println("Pilihan tidak valid!")
        return
    }
}

func menuSearchMahasiswa(data *[100]DataMahasiswa, n int) {
    var pilih int
    var target string
    fmt.Print("1. Linear Search Mahasiswa\n2. Binary Search Mahasiswa\nPilih: ")
    fmt.Scan(&pilih)
    var hasil []int

    if pilih == 1 {
	hasil = linearSearchMahasiswa(data, n, target)
	} else if pilih == 2 {
		hasil = binarySearchPrefixMahasiswa(data, n, target)
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Print("Masukkan nama / awalan yang dicari: ")
	fmt.Scan(&target)

    if len(hasil) > 0 {
        fmt.Println("Ditemukan pada indeks:")
        for _, idx := range hasil {
            fmt.Printf("%d. %s %s\n", idx+1, data[idx].Nama, data[idx].NIM)
        }
    } else {
        fmt.Println("Tidak ditemukan")
    }
}
func menuSearchDosen(data *[100]DataDosen, n int) {
    var pilih int
    var target string
    fmt.Print("1. Linear Search Dosen \n2. Binary Search Dosen \nPilih: ")
    fmt.Scan(&pilih)
    var hasil []int

    if pilih == 1 {
        hasil = linearSearchDosen(data, n, target)
    } else if pilih == 2 {
        hasil = binarySearchDosen(data, n, target)
    } else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Print("Masukkan nama / awalan yang dicari: ")
    fmt.Scan(&target)

    if len(hasil) > 0 {
        fmt.Println("Ditemukan pada indeks:")
        for _, idx := range hasil {
            fmt.Printf("%d. %s %s\n", idx+1, data[idx].Nama, data[idx].NIDN)
        }
    } else {
        fmt.Println("Tidak ditemukan")
    }
}

func menuSearchKelas(data *[100]DataKelas, n int) {
    var pilih int
    var target string
    fmt.Print("1. Linear Search Kelas \n2. Binary Search Kelas \nPilih: ")
    fmt.Scan(&pilih) 
    var hasil []int

    if pilih == 1 {
        hasil = linearSearchKelas(data, n, target)
    } else if pilih == 2 {
        hasil = binarySearchKelas(data, n, target)
    } else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Print("Masukkan nama kelas / awalan yang dicari: ")
    fmt.Scan(&target)

    if len(hasil) > 0 {
        fmt.Println("Ditemukan pada indeks:")
        for _, idx := range hasil {
            fmt.Printf("%d. %s %s\n", idx+1, data[idx].NamaKelas, data[idx].Ruangan)
        }
    } else {
        fmt.Println("Tidak ditemukan")
    }
}

// --- Fungsi Edit Data ---
func editDataMahasiswa(data *[100]DataMahasiswa, n int) {
    if n == 0 {
        fmt.Println("Belum ada data mahasiswa.")
        return
    }
    var idx int
    tampilkanDataMahasiswa(data, n)
    fmt.Print("Masukkan nomor mahasiswa yang ingin diedit: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx-- // indeks array mulai dari 0
    fmt.Print("Nama baru: ")
    fmt.Scan(&data[idx].Nama)
    fmt.Print("NIM baru: ")
    fmt.Scan(&data[idx].NIM)
    fmt.Println("Data berhasil diubah.")
}

func editDataDosen(data *[100]DataDosen, n int) {
    if n == 0 {
        fmt.Println("Belum ada data dosen.")
        return
    }
    var idx int
    tampilkanDataDosen(data, n)
    fmt.Print("Masukkan nomor dosen yang ingin diedit: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx--
    fmt.Print("Nama baru: ")
    fmt.Scan(&data[idx].Nama)
    fmt.Print("NIDN baru: ")
    fmt.Scan(&data[idx].NIDN)
    fmt.Println("Data berhasil diubah.")
}

func editDataKelas(data *[100]DataKelas, n int) {
    if n == 0 {
        fmt.Println("Belum ada data kelas.")
        return
    }
    var idx int
    tampilkanDataKelas(data, n)
    fmt.Print("Masukkan nomor kelas yang ingin diedit: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx--
    fmt.Print("Nama Kelas baru: ")
    fmt.Scan(&data[idx].NamaKelas)
    fmt.Print("Ruangan baru: ")
    fmt.Scan(&data[idx].Ruangan)
    fmt.Println("Data berhasil diubah.")
}

// --- Fungsi Hapus Data ---
func hapusDataMahasiswa(data *[100]DataMahasiswa, n *int) {
    if *n == 0 {
        fmt.Println("Belum ada data mahasiswa.")
        return
    }
    var idx int
    tampilkanDataMahasiswa(data, *n)
    fmt.Print("Masukkan nomor mahasiswa yang ingin dihapus: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > *n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx--
    for i := idx; i < *n-1; i++ {
        data[i] = data[i+1]
    }
    *n--
    fmt.Println("Data berhasil dihapus.")
}

func hapusDataDosen(data *[100]DataDosen, n *int) {
    if *n == 0 {
        fmt.Println("Belum ada data dosen.")
        return
    }
    var idx int
    tampilkanDataDosen(data, *n)
    fmt.Print("Masukkan nomor dosen yang ingin dihapus: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > *n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx--
    for i := idx; i < *n-1; i++ {
        data[i] = data[i+1]
    }
    *n--
    fmt.Println("Data berhasil dihapus.")
}

func hapusDataKelas(data *[100]DataKelas, n *int) {
    if *n == 0 {
        fmt.Println("Belum ada data kelas.")
        return
    }
    var idx int
    tampilkanDataKelas(data, *n)
    fmt.Print("Masukkan nomor kelas yang ingin dihapus: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > *n {
        fmt.Println("Nomor tidak valid!")
        return
    }
    idx--
    for i := idx; i < *n-1; i++ {
        data[i] = data[i+1]
    }
    *n--
    fmt.Println("Data berhasil dihapus.")
}

// --- Main Program ---
func main() {
    var mahasiswa [100]DataMahasiswa
    var dosen [100]DataDosen
    var kelas [100]DataKelas
    var nMahasiswa, nDosen, nKelas int
    var pilih, subpilih int

    nMahasiswa = 3
    mahasiswa[0] = DataMahasiswa{Nama: "Andi", NIM: "12345"}
    mahasiswa[1] = DataMahasiswa{Nama: "Budi", NIM: "23456"}
    mahasiswa[2] = DataMahasiswa{Nama: "Cici", NIM: "34567"}

    nDosen = 2
    dosen[0] = DataDosen{Nama: "Pak Dedi", NIDN: "11111"}
    dosen[1] = DataDosen{Nama: "Bu Sari", NIDN: "22222"}

    nKelas = 2
    kelas[0] = DataKelas{NamaKelas: "IF-41-01", Ruangan: "Lab 1"}
    kelas[1] = DataKelas{NamaKelas: "IF-41-02", Ruangan: "Lab 2"}

    exit := false
    for !exit {
        fmt.Println("-----Kemahasiswaan-----")
        fmt.Println("1. Input data Mahasiswa")
        fmt.Println("2. Input data Dosen")
        fmt.Println("3. Input data Kelas")
        fmt.Println("4. Sort Data")
        fmt.Println("5. Search Data")
        fmt.Println("6. Tampilkan Data")
        fmt.Println("7. Edit Data")
        fmt.Println("8. Hapus Data")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        if _, err := fmt.Scan(&pilih); err != nil {
            fmt.Println("Input harus berupa angka!")
            var dummy string
            fmt.Scanln(&dummy)
            fmt.Println()
        } else if pilih == 0 {
            exit = true
        } else if pilih == 1 {
            inputDataMahasiswa(&mahasiswa, &nMahasiswa)
        } else if pilih == 2 {
            inputDataDosen(&dosen, &nDosen)
        } else if pilih == 3 {
            inputDataKelas(&kelas, &nKelas)
        } else if pilih == 4 {
            fmt.Println("Sort untuk: 1. Mahasiswa 2. Dosen 3. Kelas")
            fmt.Print("Pilih: ")
            if _, err := fmt.Scan(&subpilih); err != nil {
                fmt.Println("Input harus berupa angka!")
                var dummy string
                fmt.Scanln(&dummy)
                fmt.Println()
            } else if subpilih == 1 {
                menuSortMahasiswa(&mahasiswa, nMahasiswa)
            } else if subpilih == 2 {
                menuSortDosen(&dosen, nDosen)
            } else if subpilih == 3 {
                menuSortKelas(&kelas, nKelas)
            } else {
                fmt.Println("Pilihan tidak valid!")
            }
        } else if pilih == 5 {
            fmt.Println("Search untuk: 1. Mahasiswa 2. Dosen 3. Kelas")
            fmt.Print("Pilih: ")
            if _, err := fmt.Scan(&subpilih); err != nil {
                fmt.Println("Input harus berupa angka!")
                var dummy string
                fmt.Scanln(&dummy)
                fmt.Println()
            } else if subpilih == 1 {
                menuSearchMahasiswa(&mahasiswa, nMahasiswa)
            } else if subpilih == 2 {
                menuSearchDosen(&dosen, nDosen)
            } else if subpilih == 3 {
                menuSearchKelas(&kelas, nKelas)
            } else {
                fmt.Println("Pilihan tidak valid!")
            }
        } else if pilih == 6 {
            fmt.Println("Tampilkan data: 1. Mahasiswa 2. Dosen 3. Kelas")
            fmt.Print("Pilih: ")
            if _, err := fmt.Scan(&subpilih); err != nil {
                fmt.Println("Input harus berupa angka!")
                var dummy string
                fmt.Scanln(&dummy)
                fmt.Println()
            } else if subpilih == 1 {
                tampilkanDataMahasiswa(&mahasiswa, nMahasiswa)
            } else if subpilih == 2 {
                tampilkanDataDosen(&dosen, nDosen)
            } else if subpilih == 3 {
                tampilkanDataKelas(&kelas, nKelas)
            } else {
                fmt.Println("Pilihan tidak valid!")
            }
        } else if pilih == 7 {
            fmt.Println("Edit data: 1. Mahasiswa 2. Dosen 3. Kelas")
            fmt.Print("Pilih: ")
            if _, err := fmt.Scan(&subpilih); err != nil {
                fmt.Println("Input harus berupa angka!")
                var dummy string
                fmt.Scanln(&dummy)
                fmt.Println()
            } else if subpilih == 1 {
                editDataMahasiswa(&mahasiswa, nMahasiswa)
            } else if subpilih == 2 {
                editDataDosen(&dosen, nDosen)
            } else if subpilih == 3 {
                editDataKelas(&kelas, nKelas)
            } else {
                fmt.Println("Pilihan tidak valid!")
            }
        } else if pilih == 8 {
            fmt.Println("Hapus data: 1. Mahasiswa 2. Dosen 3. Kelas")
            fmt.Print("Pilih: ")
            if _, err := fmt.Scan(&subpilih); err != nil {
                fmt.Println("Input harus berupa angka!")
                var dummy string
                fmt.Scanln(&dummy)
                fmt.Println()
            } else if subpilih == 1 {
                hapusDataMahasiswa(&mahasiswa, &nMahasiswa)
            } else if subpilih == 2 {
                hapusDataDosen(&dosen, &nDosen)
            } else if subpilih == 3 {
                hapusDataKelas(&kelas, &nKelas)
            } else {
                fmt.Println("Pilihan tidak valid!")
            }
        } else {
            fmt.Println("Pilihan menu tidak valid!")
        }
        fmt.Println()
    }
}