package math

// Menjumlahkan semua nomor yang dimasukkan dalam parameter
func Sum(nums ...int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	return sum
}
