package kit

import "math"
//线段树模板
type SegTree struct {
	l, r int
	sum, max, min int
	lazy int
}
func build(st []*SegTree, base []int, i, l, r int) {
	st[i] = &SegTree{l:l, r:r}
	if l == r {
		st[i].sum, st[i].max, st[i].min = base[l], base[l], base[l]
		return
	}
	mid := l + (r-l)>>1
	build(st, base, i<<1, l, mid)
	build(st, base, i<<1|1, mid+1, r)
	st[i].sum = st[i<<1].sum + st[i<<1|1].sum
	st[i].max = max(st[i<<1].max, st[i<<1|1].max)
	st[i].min = min(st[i<<1].min, st[i<<1|1].min)
}
func singleUpdate(st []*SegTree, i, p, v int) {
	if st[i].l == st[i].r && st[i].l == p {
		st[i].sum += v
		st[i].max += v
		st[i].min += v
		return
	}
	if p > st[i].r || p < st[i].l {
		return
	}
	if st[i<<1].r >= p {
		singleUpdate(st, i<<1, p, v)
	} else {
		singleUpdate(st, i<<1|1, p, v)
	}
	st[i].sum = st[i<<1].sum + st[i<<1|1].sum
	st[i].max = max(st[i<<1].max, st[i<<1|1].max)
	st[i].min = min(st[i<<1].min, st[i<<1|1].min)
}
func intervalUpdate(st []*SegTree, i, l, r, v int) {
	if st[i].l >= l && st[i].r <= r {
		st[i].sum += (st[i].r - st[i].l + 1) * v
		st[i].max += v
		st[i].min += v
		st[i].lazy += v
		return
	}
	if st[i].l > r || st[i].r < l {
		return
	}
	if st[i<<1].r >= l {
		intervalUpdate(st, i<<1, l, r, v)
	}
	if st[i<<1|1].l <= r {
		intervalUpdate(st, i<<1|1, l, r, v)
	}
	st[i].sum = st[i<<1].sum + st[i<<1|1].sum
	st[i].max = max(st[i<<1].max, st[i<<1|1].max)
	st[i].min = min(st[i<<1].min, st[i<<1|1].min)
}
func singleQuery(st []*SegTree, i, p int) int {
	if st[i].l == st[i].r && st[i].l == p {
		return st[i].sum
	}
	if st[i].l > p || st[i].r < p {
		return 0
	}
	pushDown(st, i)
	if st[i<<1].r >= p {
		return singleQuery(st, i<<1, p)
	} else {
		return singleQuery(st, i<<1|1, p)
	}
}
func intervalQuery(st []*SegTree, i, l, r int) (int, int, int) {
	if st[i].l >= l && st[i].r <= r {
		return st[i].sum, st[i].max, st[i].min
	}
	if st[i].l > r || st[i].r < l {
		return 0, 0, 0
	}
	pushDown(st, i)
	var tmp_sum, tmp_max, tmp_min int
	var tmp2_sum, tmp2_max, tmp2_min int
	tmp_max, tmp2_max = math.MinInt32, math.MinInt32
	tmp_min, tmp2_min = math.MaxInt32, math.MaxInt32
	if st[i<<1].r >= l {
		tmp_sum, tmp_max, tmp_min = intervalQuery(st, i<<1, l, r)
	}
	if st[i<<1|1].l <= r {
		tmp2_sum, tmp2_max, tmp2_min = intervalQuery(st, i<<1|1, l, r)
	}
	return tmp_sum + tmp2_sum, max(tmp_max, tmp2_max), min(tmp_min, tmp2_min)
}
func pushDown(st []*SegTree, i int) {
	if st[i].lazy > 0 {
		st[i<<1].lazy += st[i].lazy
		st[i<<1|1].lazy += st[i].lazy
		st[i<<1].sum += (st[i<<1].r-st[i<<1].l+1) * st[i].lazy
		st[i<<1].max += st[i].lazy
		st[i<<1].min += st[i].lazy
		st[i<<1|1].sum += (st[i<<1|1].r-st[i<<1|1].l+1) * st[i].lazy
		st[i<<1|1].max += st[i].lazy
		st[i<<1|1].min += st[i].lazy
		st[i].lazy = 0
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
