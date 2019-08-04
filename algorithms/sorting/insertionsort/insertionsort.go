package insertionsort

// Sort - Elementary sort. Running time of insetion sort relies on the input. If it's already partially sorted, it's faster to sort than using selection sort.
// just like selection sort, this algorithm is quadratic.
func Sort(a []int) {
	len := len(a)
	for i := 1; i < len; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- { // the conditional here makes this algo faster than selection sort for partially sorted arrays
			exch(a, j, j-1)
		}
	}
}

func exch(a []int, ai int, bi int) {
	temp := a[ai]
	a[ai] = a[bi]
	a[bi] = temp
}