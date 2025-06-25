package main

import "fmt"

const NMAX int = 1000

type Sampah struct {
	Jenis1, Jenis2, Jenis3, Jenis4     string  // Organik / Anorganik / B3 / Berbahaya
	Berat1, Berat2, Berat3, Berat4     float64 // dalam kg
	Lokasi1, Lokasi2, Lokasi3, Lokasi4 string
}

type dataSampah [NMAX]Sampah

func main() {
	var data dataSampah
	var totalSampah int

	for {
		var pilih int

		fmt.Println("\n=== SELAMAT DATANG DI APLIKASI PENGELOLAAN SAMPAH ===")
		fmt.Println("1. Log In Sebagai Admin") // bisa semua fitur
		fmt.Println("2. Log In Sebagai Guest") // hanya bisa jual bahan daur ulang, dan melihat data sampah saja
		fmt.Println("3. Keluar")
		fmt.Println("=======================================================")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			loginadmin(&data, &totalSampah) // Pass by reference
		case 2:
			loginguest(&data, &totalSampah) // Pass by reference
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func loginadmin(x *dataSampah, y *int) {
	var usn, pw string

	fmt.Print("Username: ")
	fmt.Scan(&usn)
	fmt.Print("Password: ")
	fmt.Scan(&pw)

	if usn == "sampah" && pw == "sampah123" {
		fmt.Println("Log In Berhasil Sebagai Admin!")
		menuAdmin(x, y)
		return
	}

	fmt.Println("Username atau Password Salah!")
}

func menuAdmin(x *dataSampah, y *int) {
	for {
		var pilihan int
		var data Sampah

		fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH (ADMIN) ===")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Tambah Data Sampah Berbahaya")
		fmt.Println("3. Hapus Data Sampah")
		fmt.Println("4. Lihat Data Sampah")
		fmt.Println("5. Statistik Sampah")
		fmt.Println("6. Keluar")
		fmt.Println("=============================================")
		fmt.Print("Pilih Menu: ")

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahSampah(x, y, data)
		case 2:
			tambahSampahBerbahaya(x, y, data)
		case 3:
			hapusSampah(x, y)
		case 4:
			lihatSampah(*x, *y) // tanda(* pointer) agar mengubah jenis in/out -> in atau sebaliknya dari variabel procedure nya
		case 5:
			urutBeratSampah(x, *y) // tanda(* pointer) agar mengubah jenis in/out -> in atau sebaliknya dari variabel procedure nya
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func loginguest(x *dataSampah, y *int) {
	fmt.Println("Log In Sebagai Guest")
	menuGuest(x, y)
}

func menuGuest(x *dataSampah, y *int) {
	for {
		var pilihan int
		var data Sampah

		fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH (GUEST) ===")
		fmt.Println("1. Tambah Data Sampah")
		fmt.Println("2. Lihat Data Sampah")
		fmt.Println("3. Jual Daur Ulang Sampah")
		fmt.Println("4. Statistik Sampah")
		fmt.Println("5. Keluar")
		fmt.Println("=============================================")
		fmt.Print("Pilih Menu: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahSampah(x, y, data)
		case 2:
			lihatSampah(*x, *y) // tanda(* pointer) agar mengubah jenis in/out -> in atau sebaliknya dari si variabel nya
		case 3:
			tambahDaurUlang()
		case 4:
			urutBeratSampah(x, *y) // tanda(* pointer) agar mengubah jenis in/out -> in atau sebaliknya dari si variabel nya
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahSampah(x *dataSampah, y *int, z Sampah) { // z berfungsi sebagai variabel yang menampung ke dalam array

	for {
		var pilih int

		fmt.Println("=======================================================")
		fmt.Println("\n=== Tambah Jenis Sampah ===")
		fmt.Println("1. Organik")
		fmt.Println("2. Anorganik")
		fmt.Println("3. Bahan Berbahaya dan Beracun (B3)")
		fmt.Println("4. Keluar")
		fmt.Println("=======================================================")
		fmt.Print("Pilih jenis: ")

		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			z.Jenis1 = "Organik"

			fmt.Print("\nMasukkan Berat: ")
			fmt.Scan(&z.Berat1)

			fmt.Print("\nMasukkan Lokasi: ")
			fmt.Scan(&z.Lokasi1)

			// tidak menggunakan for looping karena data yang diinput satu per satu
			x[*y].Jenis1 = z.Jenis1
			x[*y].Berat1 = z.Berat1
			x[*y].Lokasi1 = z.Lokasi1

			*y++

			fmt.Println("Data Berhasil Ditambah!")

		case 2:
			z.Jenis2 = "Anorganik"

			fmt.Print("\nMasukkan Berat: ")
			fmt.Scan(&z.Berat2)

			fmt.Print("\nMasukkan Lokasi: ")
			fmt.Scan(&z.Lokasi2)

			// tidak menggunakan for looping karena data yang diinput satu per satu
			x[*y].Jenis2 = z.Jenis2
			x[*y].Berat2 = z.Berat2
			x[*y].Lokasi2 = z.Lokasi2

			*y++

			fmt.Println("Data Berhasil Ditambah!")

		case 3:
			z.Jenis3 = "Bahan Berbahaya dan Beracun (B3)"

			fmt.Print("\nMasukkan Berat: ")
			fmt.Scan(&z.Berat3)

			fmt.Print("\nMasukkan Lokasi: ")
			fmt.Scan(&z.Lokasi3)

			x[*y].Jenis3 = z.Jenis3
			x[*y].Berat3 = z.Berat3
			x[*y].Lokasi3 = z.Lokasi3

			*y++

			fmt.Println("Data Berhasil Ditambah!")

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
func tambahSampahBerbahaya(x *dataSampah, y *int, z Sampah) { // z berfungsi sebagai variabel yang menampung ke dalam array
	fmt.Println("=======================================================")
	fmt.Println("\n=== Tambah Jenis Sampah Berbahaya===")
	fmt.Println("=======================================================")
	z.Jenis4 = "Sampah Berbahaya"

	fmt.Print("\nMasukkan Berat: ")
	fmt.Scan(&z.Berat4)

	fmt.Print("\nMasukkan Lokasi: ")
	fmt.Scan(&z.Lokasi4)

	// tidak menggunakan for looping karena data yang diinput satu per satu
	x[*y].Jenis4 = z.Jenis4
	x[*y].Berat4 = z.Berat4
	x[*y].Lokasi4 = z.Lokasi4

	*y++

	fmt.Println("Data Berhasil Ditambah!")
}

func hapusSampah(x *dataSampah, y *int) { // y = totalSampah

	for {
		var i, pilih, pilih1 int

		fmt.Println("=======================================================")
		fmt.Println("\n=== Hapus Data Sampah ===")
		fmt.Println("1. Organik")
		fmt.Println("2. Anorganik")
		fmt.Println("3. Bahan Berbahaya dan Beracun (B3)")
		fmt.Println("4. Keluar")
		fmt.Println("=======================================================")
		fmt.Print("Pilih jenis sampah yang akan dihapus: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Hapus Data Sampah Organik ===")
			// tampilkan data per data dari setiap jenisnya nya terlebih dahulu
			for i = 0; i < *y; i++ {
				if x[i].Jenis1 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis1, x[i].Berat1, x[i].Lokasi1)
				}
			}
			fmt.Println("=======================================================")
			fmt.Print("Data berapa yang akan dihapus?: ")
			fmt.Scan(&pilih1)

			pilih1-- // supaya (posisi data yang dipilih - 1) => menjadi posisi data sebelum yang dipilih

			for i = pilih1; i < *y-1; i++ { //pilih1 nya menjadi posisi data sebelum yang dipilih
				x[i] = x[i+1] //sehingga data pilih1 nya menjadi posisi data yang dipilih tadi. Makanya data yang dipilihnya terhapus
			}
			*y--

			fmt.Println("=======================================================")
			fmt.Println("\nData berhasil dihapus!")

		} else if pilih == 2 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Hapus Data Sampah Anorganik ===")
			// tampilkan data per data dari setiap jenisnya nya terlebih dahulu
			for i = 0; i < *y; i++ {
				if x[i].Jenis2 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis2, x[i].Berat2, x[i].Lokasi2)
				}
			}
			fmt.Println("=======================================================")
			fmt.Print("Data berapa yang akan dihapus?: ")
			fmt.Scan(&pilih1)

			pilih1-- // supaya (posii data yang dipilih - 1) => menjadi posisi data sebelum yang dipilih

			for i = pilih1; i < *y-1; i++ { //pilih1 nya menjadi posisi data sebelum yang dipilih
				x[i] = x[i+1] //sehingga data pilih1 nya menjadi posisi data yang dipilih tadi. Makanya data yang dipilihnya terhapus
			}
			*y--

			fmt.Println("=======================================================")
			fmt.Println("\nData berhasil dihapus!")

		} else if pilih == 3 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Hapus Data Sampah Bahan Berbahaya dan Beracun (B3) ===")
			// tampilkan data per data dari setiap jenisnya nya terlebih dahulu
			for i = 0; i < *y; i++ {
				if x[i].Jenis3 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis3, x[i].Berat3, x[i].Lokasi3)
				}
			}
			fmt.Println("=======================================================")
			fmt.Print("Data berapa yang akan dihapus?: ")
			fmt.Scan(&pilih1)

			pilih1-- // supaya (posisi data yang dipilih - 1) => menjadi posisi data sebelum yang dipilih

			for i = pilih1; i < *y-1; i++ { //pilih1 nya menjadi posisi data sebelum yang dipilih
				x[i] = x[i+1] //sehingga data pilih1 nya menjadi posisi data yang dipilih tadi. Makanya data yang dipilihnya terhapus
			}
			*y--

			fmt.Println("=======================================================")
			fmt.Println("\nData berhasil dihapus!")

		} else if pilih == 4 {
			return
		} else {
			fmt.Println("Pilihan Tidak Valid!")
		}
	}

}

func lihatSampah(x dataSampah, y int) { //tampilkan data per data

	for {
		var i, pilih int
		var left, right, mid, ketemu int
		var pilihBerbahaya float64

		fmt.Println("=======================================================")
		fmt.Println("\n=== Lihat Jenis Sampah ===")
		fmt.Println("1. Organik")
		fmt.Println("2. Anorganik")
		fmt.Println("3. Bahan Berbahaya dan Beracun (B3)")
		fmt.Println("4. Sampah Berbahaya")
		fmt.Println("5. Keluar")
		fmt.Println("=======================================================")
		fmt.Print("Pilih jenis sampah yang akan dilihat: ")

		fmt.Scan(&pilih)

		if y <= 0 {
			fmt.Println("=======================================================")
			fmt.Println("Saat ini data masih kosong")
			return
		}

		if pilih == 1 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Data Sampah Organik ===")
			for i = 0; i < y; i++ {
				if x[i].Jenis1 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis1, x[i].Berat1, x[i].Lokasi1)
				}
			}
		} else if pilih == 2 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Data Sampah Anorganik ===")
			for i = 0; i < y; i++ {
				if x[i].Jenis2 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis2, x[i].Berat2, x[i].Lokasi2)
				}
			}
		} else if pilih == 3 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Data Sampah Bahan Berbahaya dan Beracun (B3) ===")
			for i = 0; i < y; i++ {
				if x[i].Jenis3 != "" {
					fmt.Printf("%d. %s, %.2f kg, %s\n", i+1, x[i].Jenis3, x[i].Berat3, x[i].Lokasi3)
				}
			}
		} else if pilih == 4 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Data Sampah Berbahaya ===")
			urutSampahBerbahaya(&x, y) // memanggil prosedur khusus urut sampah berbahaya secara descending
			// Binary search DESCENDING
			fmt.Println("=======================================================")
			fmt.Print("Berapa berat spesifik yang dicari?: ")
			fmt.Scan(&pilihBerbahaya)

			left = 0
			right = y - 1
			ketemu = -1
			for left <= right && ketemu == -1 {
				mid = (left + right) / 2
				if pilihBerbahaya < x[mid].Berat4 {
					left = mid + 1
				} else if pilihBerbahaya > x[mid].Berat4 {
					right = mid - 1
				} else {
					ketemu = mid
				}
			}

			fmt.Println("=======================================================")
			fmt.Println("\n=== Hasil Pencarian Data Sampah Spesifik ===")
			if ketemu != -1 {
				fmt.Printf("%s, %.2f kg, %s\n", x[ketemu].Jenis4, x[ketemu].Berat4, x[ketemu].Lokasi4)
			} else {
				fmt.Println("Data dengan berat", pilihBerbahaya, "kg tidak ditemukan")
			}
		} else if pilih == 5 {
			return

		} else {
			fmt.Println("Pilihan Tidak Valid!")
		}
	}
}

func tambahDaurUlang() {

	for {
		var pilih, pilih1, pilih2 int
		var berat1, berat2 int
		var total1, total2 int

		fmt.Println("=======================================================")
		fmt.Println("\n=== Daur Ulang Jenis Sampah ===")
		fmt.Println("1. Organik")
		fmt.Println("2. Anorganik")
		fmt.Println("3. Bahan Berbahaya dan Beracun (B3)")
		fmt.Println("4. Keluar")
		fmt.Println("=======================================================")
		fmt.Print("Pilih jenis Sampah yang akan didaur ulang: ")

		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
			fmt.Println("1. Sisa Makanan")
			fmt.Println("2. Daun-Daunan")
			fmt.Println("3. Kotoran Hewan")
			fmt.Println("4. Ampas Tanaman")
			fmt.Println("=======================================================")
			fmt.Print("\nPilih jenis Sampah Organik yang akan didaur ulang: ")
			fmt.Scan(&pilih1)
			if pilih1 == 1 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Sisa Makanan seharga Rp5.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Sisa Makanan Anda berapa Kg?: ")
				fmt.Scan(&berat1)
				total1 = hargalimaribu(berat1)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Sisa Makanan Anda seharga: ", total1)
			} else if pilih1 == 2 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Daun-Daunan seharga Rp2.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Daun-Daunan Anda berapa Kg?: ")
				fmt.Scan(&berat1)
				total1 = hargaduaribu(berat1)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Daun-Daunan Anda seharga: ", total1)
			} else if pilih1 == 3 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Kotoran Hewan seharga Rp2.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Kotoran Hewan Anda berapa Kg?: ")
				fmt.Scan(&berat1)
				total1 = hargaduaribu(berat1)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Kotoran Hewan Anda seharga: ", total1)
			} else if pilih1 == 4 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Ampas Tanaman seharga Rp2.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Ampas Tanaman Anda berapa Kg?: ")
				fmt.Scan(&berat1)
				total1 = hargaduaribu(berat1)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Organik ===")
				fmt.Println("Untuk jenis Ampas Tanaman Anda seharga: ", total1)
			} else {
				fmt.Println("Pilihan Tidak Valid!")
			}
		} else if pilih == 2 {
			fmt.Println("=======================================================")
			fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
			fmt.Println("1. Plastik")
			fmt.Println("2. Kardus/Kertas")
			fmt.Println("3. Kaleng(Alumunium)")
			fmt.Println("4. Besi/Logam")
			fmt.Println("=======================================================")
			fmt.Print("\nPilih jenis Sampah Anorganik yang akan didaur ulang: ")
			fmt.Scan(&pilih2)
			if pilih2 == 1 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Plastik seharga Rp3.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Plastik Anda berapa Kg?: ")
				fmt.Scan(&berat2)
				total2 = hargatigaribu(berat2)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Plastik Anda seharga: ", total2)
			} else if pilih2 == 2 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Kardus/Kertas seharga Rp2.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Kardus/Kertas Anda berapa Kg?: ")
				fmt.Scan(&berat2)
				total2 = hargaduaribu(berat2)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Kardus/Kertas Anda seharga: ", total2)
			} else if pilih2 == 3 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Kaleng(Alumunium) seharga Rp2.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Kaleng(Alumunium) Anda berapa Kg?: ")
				fmt.Scan(&berat2)
				total2 = hargaduaribu(berat2)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Kaleng(Alumunium) Anda seharga: ", total2)
			} else if pilih2 == 4 {
				fmt.Println("=======================================================")
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Besi/Logam seharga Rp5.000/kg")
				fmt.Println("=======================================================")
				fmt.Print("Untuk jenis Besi/Logam Anda berapa Kg?: ")
				fmt.Scan(&berat2)
				total2 = hargalimaribu(berat2)
				fmt.Println("\n=== Daur Ulang Jenis Sampah Anorganik ===")
				fmt.Println("Untuk jenis Besi/Logam Anda seharga: ", total2)
			} else {
				fmt.Println("Pilihan Tidak Valid!")
			}

		} else if pilih == 3 {
			fmt.Println("=======================================================")
			fmt.Println("Bahan Berbahaya dan Beracun (B3) Tidak Dapat dilakukan Daur Ulang")

		} else if pilih == 4 {
			return

		} else {
			fmt.Println("Pilihan Tidak Valid!")
		}
	}
}

func urutBeratSampah(x *dataSampah, y int) { // pakai sorting -> selection & insertion
	var i, pilih int
	var idx, pass int
	var temp Sampah
	pass = 1

	fmt.Println("=======================================================")
	fmt.Println("\n=== Statistik Jenis Sampah ===")
	fmt.Println("Data mau diurutkan secara apa?")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("=======================================================")
	fmt.Print("Pilih cara urut data: ")
	fmt.Scan(&pilih)

	fmt.Println(y)
	if pilih == 1 {
		fmt.Println("=======================================================")
		fmt.Println("\n=== sorting menaik ===")
		for pass = 1; pass < y; pass++ { // insertion sort ascending
			i = pass
			temp = x[pass]

			for i > 0 && (x[i-1].Berat1 > temp.Berat1 ||
				(x[i-1].Berat1 == temp.Berat1 && x[i-1].Berat2 > temp.Berat2) ||
				(x[i-1].Berat1 == temp.Berat1 && x[i-1].Berat2 == temp.Berat2 && x[i-1].Berat3 > temp.Berat3)) {
				x[i] = x[i-1]
				i--
			}
			x[i] = temp
		}
	} else if pilih == 2 { // selection sort descending
		fmt.Println("=======================================================")
		fmt.Println("\n=== sorting menurun ===")
		for pass = 0; pass < y-1; pass++ {
			idx = pass
			for i = pass + 1; i < y; i++ {
				// Bandingkan berat1 pertama, lalu berat2, lalu berat3
				if x[i].Berat1 > x[idx].Berat1 ||
					(x[i].Berat1 == x[idx].Berat1 && x[i].Berat2 > x[idx].Berat2) ||
					(x[i].Berat1 == x[idx].Berat1 && x[i].Berat2 == x[idx].Berat2 && x[i].Berat3 > x[idx].Berat3) {
					idx = i
				}
			}
			// Tukar elemen
			temp = x[pass]
			x[pass] = x[idx]
			x[idx] = temp
		}
	} else {
		fmt.Println("Pilihan Tidak Valid!")
	}
}

func urutSampahBerbahaya(x *dataSampah, y int) { //menggunakan sorting descending selection
	var i int
	var idx, pass int
	var temp Sampah
	pass = 1
	for pass = 0; pass < y-1; pass++ {
		idx = pass
		for i = pass + 1; i < y; i++ {
			if x[i].Berat4 > x[idx].Berat4 {
				idx = i
			}
		}
		temp = x[pass]
		x[pass] = x[idx]
		x[idx] = temp
	}
}

/////////////////////////////////

func hargalimaribu(x int) int {
	total := x * 5000

	return total
}

func hargaduaribu(x int) int {
	total := x * 2000

	return total
}

func hargatigaribu(x int) int {
	total := x * 3000

	return total
}
