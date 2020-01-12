package kit

func bitUpdate(bit []int, n, i, v int) {
	for i <= n {
		bit[i] += v
		i += lowbit(i)
	}
}
func lowbit(a int) int {
	return a & -a
}
func bitQuery(bit []int, i int) int {
	var sum int
	for i != 0 {
		sum += bit[i]
		i -= lowbit(i)
	}
	return sum
}
