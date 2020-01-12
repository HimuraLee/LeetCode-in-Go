package kit

import "fmt"

func KMP(s, p string) []int {
	n := len(s)
	pn := len(p)
	if n < pn {
		return nil
	}
	var res []int
	pmt := genPMT(p)
	maxLen := 0
	for i := 0; i < n; i++ {
		for maxLen > 0 && s[i] != p[maxLen] {
			maxLen = pmt[maxLen-1]
		}
		if s[i] == p[maxLen] {
			maxLen++
		}
		if maxLen == pn {
			res = append(res, i)
			maxLen = pmt[maxLen-2]
			fmt.Printf("from %d to %d match, maxLen turn to %d\n", i-pn+1, i, maxLen)
			i--
		}
	}
	return res
}
func genPMT(p string) []int {
	n := len(p)
	pmt := make([]int, n)
	var maxPattern int
	for i := 1; i < n; i++ {
		for maxPattern > 0 && p[maxPattern] != p[i] {
			maxPattern = pmt[maxPattern-1]
		}
		if p[maxPattern] == p[i] {
			maxPattern++

		}
		pmt[i] = maxPattern
	}
	return pmt
}
