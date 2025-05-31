package main

// Import
import (
	"fmt"
	"sort"
)

// Struct
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

// Array
type Mahasiswa [100]DataMahasiswa
type Dosen [100]DataDosen
type Kelas [100]DataKelas

// Mahasiswa
func selectionSortingMahasiswa(data *Mahasiswa, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data[j].NIM < data[minIndex].NIM {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
func insertionSortingMahasiswa(data *Mahasiswa, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Nama > key.Nama {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
func linearSearchMahasiswa(data *Mahasiswa, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].Nama) >= len(prefix) && data[i].Nama[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}
func binarySearchPrefixMahasiswa(data *Mahasiswa, n int, prefix string) []int {
    temp := make([]DataMahasiswa, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
    bawah, atas := 0, n-1
    var hasil []int
    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].Nama) >= len(prefix) && temp[tengah].Nama[:len(prefix)] == prefix {
            // Cari ke kiri
            i := tengah
            for i >= 0 && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--
            }
            // Cari ke kanan
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

// Dosen
func selectionSortingDosen(data *Dosen, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data[j].NIDN < data[minIndex].NIDN {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
func insertionSortingDosen(data *Dosen, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Nama > key.Nama {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
func linearSearchDosen(data *Dosen, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].Nama) >= len(prefix) && data[i].Nama[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}
func binarySearchDosen(data *Dosen, n int, prefix string) []int {
    temp := make([]DataDosen, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
    bawah, atas := 0, n-1
    var hasil []int
    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].Nama) >= len(prefix) && temp[tengah].Nama[:len(prefix)] == prefix {
            // Kiri
            i := tengah
            for i >= 0 && len(temp[i].Nama) >= len(prefix) && temp[i].Nama[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--
            }
            // Kanan
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

// Kelas
func selectionSortingKelas(data *Kelas, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data[j].Ruangan < data[minIndex].Ruangan {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
func insertionSortingKelas(data *Kelas, n int) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].NamaKelas > key.NamaKelas {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
func linearSearchKelas(data *Kelas, n int, prefix string) []int {
    var hasil []int
    for i := 0; i < n; i++ {
        if len(data[i].NamaKelas) >= len(prefix) && data[i].NamaKelas[:len(prefix)] == prefix {
            hasil = append(hasil, i)
        }
    }
    return hasil
}
func binarySearchKelas(data *Kelas, n int, prefix string) []int {
    temp := make([]DataKelas, n)
    copy(temp, data[:n])
    sort.Slice(temp, func(i, j int) bool { return temp[i].NamaKelas < temp[j].NamaKelas })
    bawah, atas := 0, n-1
    var hasil []int
    for bawah <= atas {
        tengah := (bawah + atas) / 2
        if len(temp[tengah].NamaKelas) >= len(prefix) && temp[tengah].NamaKelas[:len(prefix)] == prefix {
            // Kiri
            i := tengah
            for i >= 0 && len(temp[i].NamaKelas) >= len(prefix) && temp[i].NamaKelas[:len(prefix)] == prefix {
                hasil = append(hasil, i)
                i--
            }
            // Kanan
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

// --- Fungsi Input Data ---

func inputDataMahasiswa(data *Mahasiswa, n *int) {
	var tambah int
	fmt.Print("Masukkan jumlah mahasiswa yang ingin ditambahkan: ")
	for {
		if _, err := fmt.Scan(&tambah); err != nil || tambah <= 0 || *n+tambah > 100 {
			fmt.Println("Input jumlah mahasiswa harus berupa angka dan total tidak lebih dari 100!")
			fmt.Print("Masukkan jumlah mahasiswa yang ingin ditambahkan: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := *n; i < *n+tambah; i++ {
		for {
			fmt.Printf("Nama mahasiswa ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].Nama); err != nil || data[i].Nama == "" {
				fmt.Println("Nama tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[i].Nama {
				if c < '0' || c > '9' {
					isAngka = false
					break
				}
			}
			if isAngka {
				fmt.Println("Nama tidak boleh berupa angka!")
				continue
			}
			break
		}
		for {
			fmt.Printf("NIM mahasiswa ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].NIM); err != nil || data[i].NIM == "" {
				fmt.Println("NIM tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[i].NIM {
				if c < '0' || c > '9' {
					isAngka = false
					break
				}
			}
			if !isAngka {
				fmt.Println("NIM harus berupa angka!")
				continue
			}
			break
		}
	}
	*n += tambah // tambah jumlah mahasiswa
}

func inputDataDosen(data *Dosen, n *int) {
	var tambah int
	fmt.Print("Masukkan jumlah dosen yang ingin ditambah: ")
	for {
		if _, err := fmt.Scan(&tambah); err != nil || tambah <= 0 || *n+tambah > 100 {
			fmt.Println("Input jumlah dosen harus berupa angka dan total tidak boleh lebih dari 100!")
			fmt.Print("Masukkan jumlah dosen yang ingin ditambah: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := 0; i < tambah; i++ {
		idx := *n + i
		// Validasi Nama: tidak boleh kosong & tidak boleh angka
		for {
			fmt.Printf("Nama dosen ke-%d: ", idx+1)
			if _, err := fmt.Scan(&data[idx].Nama); err != nil || data[idx].Nama == "" {
				fmt.Println("Nama tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[idx].Nama {
				if c < '0' || c > '9' {
					isAngka = false
					break
				}
			}
			if isAngka {
				fmt.Println("Nama tidak boleh berupa angka!")
				continue
			}
			break
		}
		// Validasi NIDN: tidak boleh kosong & harus angka
		for {
			fmt.Printf("NIDN dosen ke-%d: ", idx+1)
			if _, err := fmt.Scan(&data[idx].NIDN); err != nil || data[idx].NIDN == "" {
				fmt.Println("NIDN tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[idx].NIDN {
				if c < '0' || c > '9' {
					isAngka = false
					break
				}
			}
			if !isAngka {
				fmt.Println("NIDN harus berupa angka!")
				continue
			}
			break
		}
	}
	*n += tambah
}

func inputDataKelas(data *Kelas, n *int) {
	var tambah int
	fmt.Print("Masukkan jumlah kelas yang ingin ditambah: ")
	for {
		if _, err := fmt.Scan(&tambah); err != nil || tambah <= 0 || *n+tambah > 100 {
			fmt.Println("Input jumlah kelas harus berupa angka dan total tidak boleh lebih dari 100!")
			fmt.Print("Masukkan jumlah kelas yang ingin ditambah: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := 0; i < tambah; i++ {
		idx := *n + i
		// Validasi NamaKelas: tidak boleh kosong & tidak boleh angka
		for {
			fmt.Printf("Nama kelas ke-%d: ", idx+1)
			if _, err := fmt.Scan(&data[idx].NamaKelas); err != nil || data[idx].NamaKelas == "" {
				fmt.Println("Nama kelas tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[idx].NamaKelas {
				if c < '0' || c > '9' {
					isAngka = false
					break
				}
			}
			if isAngka {
				fmt.Println("Nama kelas tidak boleh berupa angka!")
				continue
			}
			break
		}
		// Validasi Ruangan: tidak boleh kosong
		for {
			fmt.Printf("Ruangan kelas ke-%d: ", idx+1)
			if _, err := fmt.Scan(&data[idx].Ruangan); err != nil || data[idx].Ruangan == "" {
				fmt.Println("Ruangan tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			break
		}
	}
	*n += tambah
}

// --- Fungsi Tampilkan Data ---

func tampilkanDataMahasiswa(data *Mahasiswa, n int) {
	fmt.Println("Daftar Mahasiswa:")
	if n == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s, NIM: %s\n", i+1, data[i].Nama, data[i].NIM)
	}
}

func tampilkanDataDosen(data *Dosen, n int) {
	fmt.Println("Daftar Dosen:")
	if n == 0 {
		fmt.Println("Belum ada data dosen.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama: %s, NIDN: %s\n", i+1, data[i].Nama, data[i].NIDN)
	}
}

func tampilkanDataKelas(data *Kelas, n int) {
	fmt.Println("Daftar Kelas:")
	if n == 0 {
		fmt.Println("Belum ada data kelas.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Nama Kelas: %s, Ruangan: %s\n", i+1, data[i].NamaKelas, data[i].Ruangan)
	}
}

// --- Fungsi Menu ---
func menuSortMahasiswa(data *Mahasiswa, n int) {
	var pilih int
	fmt.Print("1. Selection Sort Mahasiswa\n2. Insertion Sort Mahasiswa\nPilih: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSortingMahasiswa(data, n)
		fmt.Println("Setelah selection sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].Nama, data[i].NIM)
		}
	} else if pilih == 2 {
		insertionSortingMahasiswa(data, n)
		fmt.Println("Setelah insertion sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].Nama, data[i].NIM)
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}
}
func menuSortDosen(data *Dosen, n int) {
	var pilih int
	fmt.Print("1. Selection Sort Dosen\n2. Insertion Sort Dosen\nPilih: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSortingDosen(data, n)
		fmt.Println("Setelah selection sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].Nama, data[i].NIDN)
		}
	} else if pilih == 2 {
		insertionSortingDosen(data, n)
		fmt.Println("Setelah insertion sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].Nama, data[i].NIDN)
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}
}
func menuSortKelas(data *Kelas, n int) {
	var pilih int
	fmt.Print("1. Selection Sort Kelas\n2. Insertion Sort Kelas\nPilih: ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSortingKelas(data, n)
		fmt.Println("Setelah selection sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].NamaKelas, data[i].Ruangan)
		}
	} else if pilih == 2 {
		insertionSortingKelas(data, n)
		fmt.Println("Setelah insertion sort:")
		for i := 0; i < n; i++ {
			fmt.Println(data[i].NamaKelas, data[i].Ruangan)
		}
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}
}
func menuSearchMahasiswa(data *Mahasiswa, n int) {
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
func menuSearchDosen(data *Dosen, n int) {
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
func menuSearchKelas(data *Kelas, n int) {
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

// --- Main Program ---
func main() {
	var mahasiswa Mahasiswa
	var dosen Dosen
	var kelas Kelas
	var nMahasiswa, nDosen, nKelas int
	var pilih, subpilih int
	for {
		fmt.Println("-----Kemahasiswaan-----")
		fmt.Println("1. Input data Mahasiswa")
		fmt.Println("2. Input data Dosen")
		fmt.Println("3. Input data Kelas")
		fmt.Println("4. Sort Data")
		fmt.Println("5. Search Data")
		fmt.Println("6. Tampilkan Data")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		if _, err := fmt.Scan(&pilih); err != nil {
			fmt.Println("Input harus berupa angka!")
			var dummy string
			fmt.Scanln(&dummy) // bersihkan buffer
			fmt.Println()
			continue
		}
		if pilih == 0 {
			break
		}
		switch pilih {
		case 1:
			inputDataMahasiswa(&mahasiswa, &nMahasiswa)
		case 2:
			inputDataDosen(&dosen, &nDosen)
		case 3:
			inputDataKelas(&kelas, &nKelas)
		case 4:
			fmt.Println("Sort untuk: 1. Mahasiswa 2. Dosen 3. Kelas")
			fmt.Print("Pilih: ")
			if _, err := fmt.Scan(&subpilih); err != nil {
				fmt.Println("Input harus berupa angka!")
				var dummy string
				fmt.Scanln(&dummy)
				fmt.Println()
				continue
			}
			if subpilih == 1 {
				menuSortMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				menuSortDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				menuSortKelas(&kelas, nKelas)
			} else {
				fmt.Println("Pilihan tidak valid!")
				continue
			}
		case 5:
			fmt.Println("Search untuk: 1. Mahasiswa 2. Dosen 3. Kelas")
			fmt.Print("Pilih: ")
			if _, err := fmt.Scan(&subpilih); err != nil {
				fmt.Println("Input harus berupa angka!")
				var dummy string
				fmt.Scanln(&dummy)
				fmt.Println()
				continue
			}
			if subpilih == 1 {
				menuSearchMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				menuSearchDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				menuSearchKelas(&kelas, nKelas)
			} else {
				fmt.Println("Pilihan tidak valid!")
				continue
			}
		case 6:
			fmt.Println("Tampilkan data: 1. Mahasiswa 2. Dosen 3. Kelas")
			fmt.Print("Pilih: ")
			if _, err := fmt.Scan(&subpilih); err != nil {
				fmt.Println("Input harus berupa angka!")
				var dummy string
				fmt.Scanln(&dummy)
				fmt.Println()
				continue
			}
			if subpilih == 1 {
				tampilkanDataMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				tampilkanDataDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				tampilkanDataKelas(&kelas, nKelas)
			} else {
				fmt.Println("Pilihan tidak valid!")
				continue
			}
		default:
			fmt.Println("Pilihan menu tidak valid!")
			continue
		}
		fmt.Println()
	}
}
