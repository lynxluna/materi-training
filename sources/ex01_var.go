package main

import (
	"fmt"
	"time"
)

// variabel global
var sekarang int64 = time.Now().Unix()

func main() {

	// variabel lokal karena ada di dalam fungsi
	var name string
	var umur int = 30
	berat := 58.77

	name = "Brett"

	fmt.Println("Waktu sekarang sejak 1 Januari 1970: ", sekarang/1000, " detik")
	fmt.Printf("Nama: %s, umur %d, berat %.2f\n", name, umur, berat)
}
