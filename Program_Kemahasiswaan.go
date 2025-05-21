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
func linearSearchMahasiswa(data *Mahasiswa, n int, target string) int {
	for i := 0; i < n; i++ {
		if data[i].Nama == target {
			return i
		}
	}
	return -1
}
func binarySearchMahasiswa(data *Mahasiswa, n int, target string) int {
	temp := make([]DataMahasiswa, n)
	copy(temp, data[:n])
	sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
	bawah, atas := 0, n-1
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if temp[tengah].Nama == target {
			return tengah
		} else if temp[tengah].Nama < target {
			bawah = tengah + 1
		} else {
			atas = tengah - 1
		}
	}
	return -1
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
func linearSearchDosen(data *Dosen, n int, target string) int {
	for i := 0; i < n; i++ {
		if data[i].Nama == target {
			return i
		}
	}
	return -1
}
func binarySearchDosen(data *Dosen, n int, target string) int {
	temp := make([]DataDosen, n)
	copy(temp, data[:n])
	sort.Slice(temp, func(i, j int) bool { return temp[i].Nama < temp[j].Nama })
	bawah, atas := 0, n-1
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if temp[tengah].Nama == target {
			return tengah
		} else if temp[tengah].Nama < target {
			bawah = tengah + 1
		} else {
			atas = tengah - 1
		}
	}
	return -1
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
func linearSearchKelas(data *Kelas, n int, target string) int {
	for i := 0; i < n; i++ {
		if data[i].NamaKelas == target {
			return i
		}
	}
	return -1
}
func binarySearchKelas(data *Kelas, n int, target string) int {
	temp := make([]DataKelas, n)
	copy(temp, data[:n])
	sort.Slice(temp, func(i, j int) bool { return temp[i].NamaKelas < temp[j].NamaKelas })
	bawah, atas := 0, n-1
	for bawah <= atas {
		tengah := (bawah + atas) / 2
		if temp[tengah].NamaKelas == target {
			return tengah
		} else if temp[tengah].NamaKelas < target {
			bawah = tengah + 1
		} else {
			atas = tengah - 1
		}
	}
	return -1
}

// --- Fungsi Input Data ---

func inputDataMahasiswa(data *Mahasiswa, n *int) {
	fmt.Print("Masukkan jumlah mahasiswa: ")
	for {
		if _, err := fmt.Scan(n); err != nil || *n <= 0 || *n > 100 {
			fmt.Println("Input jumlah mahasiswa harus berupa angka 1-100!")
			fmt.Print("Masukkan jumlah mahasiswa: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := 0; i < *n; i++ {
		// Validasi Nama: tidak boleh kosong & tidak boleh angka
		for {
			fmt.Printf("Nama mahasiswa ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].Nama); err != nil || data[i].Nama == "" {
				fmt.Println("Nama tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			// Cek apakah nama hanya angka
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
		// Validasi NIM: tidak boleh kosong & harus angka
		for {
			fmt.Printf("NIM mahasiswa ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].NIM); err != nil || data[i].NIM == "" {
				fmt.Println("NIM tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			// Cek apakah NIM hanya angka
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
}
func inputDataDosen(data *Dosen, n *int) {
	fmt.Print("Masukkan jumlah dosen: ")
	for {
		if _, err := fmt.Scan(n); err != nil || *n <= 0 || *n > 100 {
			fmt.Println("Input jumlah dosen harus berupa angka 1-100!")
			fmt.Print("Masukkan jumlah dosen: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := 0; i < *n; i++ {
		// Validasi Nama: tidak boleh kosong & tidak boleh angka
		for {
			fmt.Printf("Nama dosen ke-%d: ", i+1)
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
		// Validasi NIDN: tidak boleh kosong & harus angka
		for {
			fmt.Printf("NIDN dosen ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].NIDN); err != nil || data[i].NIDN == "" {
				fmt.Println("NIDN tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[i].NIDN {
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
}

func inputDataKelas(data *Kelas, n *int) {
	fmt.Print("Masukkan jumlah kelas: ")
	for {
		if _, err := fmt.Scan(n); err != nil || *n <= 0 || *n > 100 {
			fmt.Println("Input jumlah kelas harus berupa angka 1-100!")
			fmt.Print("Masukkan jumlah kelas: ")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		break
	}
	for i := 0; i < *n; i++ {
		// Validasi NamaKelas: tidak boleh kosong & tidak boleh angka
		for {
			fmt.Printf("Nama kelas ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].NamaKelas); err != nil || data[i].NamaKelas == "" {
				fmt.Println("Nama kelas tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			isAngka := true
			for _, c := range data[i].NamaKelas {
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
			fmt.Printf("Ruangan kelas ke-%d: ", i+1)
			if _, err := fmt.Scan(&data[i].Ruangan); err != nil || data[i].Ruangan == "" {
				fmt.Println("Ruangan tidak boleh kosong!")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			break
		}
	}
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
	}
}
func menuSearchMahasiswa(data *Mahasiswa, n int) {
	var pilih int
	var target string
	fmt.Print("1. Linear Search Mahasiswa\n2. Binary Search Mahasiswa\nPilih: ")
	fmt.Scan(&pilih)
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&target)
	var index int
	if pilih == 1 {
		index = linearSearchMahasiswa(data, n, target)
	} else if pilih == 2 {
		index = binarySearchMahasiswa(data, n, target)
	}
	if index != -1 {
		fmt.Println("Ditemukan pada indeks:", index)
		fmt.Println(data[index].Nama, data[index].NIM)
	} else {
		fmt.Println("Tidak ditemukan")
	}
}
func menuSearchDosen(data *Dosen, n int) {
	var pilih int
	var target string
	fmt.Print("1. Linear Search Dosen\n2. Binary Search Dosen\nPilih: ")
	fmt.Scan(&pilih)
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&target)
	var index int
	if pilih == 1 {
		index = linearSearchDosen(data, n, target)
	} else if pilih == 2 {
		index = binarySearchDosen(data, n, target)
	}
	if index != -1 {
		fmt.Println("Ditemukan pada indeks:", index)
		fmt.Println(data[index].Nama, data[index].NIDN)
	} else {
		fmt.Println("Tidak ditemukan")
	}
}
func menuSearchKelas(data *Kelas, n int) {
	var pilih int
	var target string
	fmt.Print("1. Linear Search Kelas\n2. Binary Search Kelas\nPilih: ")
	fmt.Scan(&pilih)
	fmt.Print("Masukkan nama kelas yang dicari: ")
	fmt.Scan(&target)
	var index int
	if pilih == 1 {
		index = linearSearchKelas(data, n, target)
	} else if pilih == 2 {
		index = binarySearchKelas(data, n, target)
	}
	if index != -1 {
		fmt.Println("Ditemukan pada indeks:", index)
		fmt.Println(data[index].NamaKelas, data[index].Ruangan)
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
		fmt.Scan(&pilih)
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
			fmt.Scan(&subpilih)
			if subpilih == 1 {
				menuSortMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				menuSortDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				menuSortKelas(&kelas, nKelas)
			}
		case 5:
			fmt.Println("Search untuk: 1. Mahasiswa 2. Dosen 3. Kelas")
			fmt.Print("Pilih: ")
			fmt.Scan(&subpilih)
			if subpilih == 1 {
				menuSearchMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				menuSearchDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				menuSearchKelas(&kelas, nKelas)
			}

		case 6:
			fmt.Println("Tampilkan data: 1. Mahasiswa 2. Dosen 3. Kelas")
			fmt.Print("Pilih: ")
			fmt.Scan(&subpilih)
			if subpilih == 1 {
				tampilkanDataMahasiswa(&mahasiswa, nMahasiswa)
			} else if subpilih == 2 {
				tampilkanDataDosen(&dosen, nDosen)
			} else if subpilih == 3 {
				tampilkanDataKelas(&kelas, nKelas)
			}
		}
		fmt.Println()
	}
}
