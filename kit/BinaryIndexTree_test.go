package kit

import (
	"fmt"
	"testing"
)
//https://www.cnblogs.com/KirinSB/p/9409096.html 基础题：求数组某个位置左边或右边比他小或大的元素个数
func Test_BIT(t *testing.T) {
	N := 5
	base := [][]int{
		{0, 1},
		{5, 1},
		{7, 1},
		{3, 3},
		{5, 5},

		//{5, 5},
		//{3, 3},
		//{7, 1},
		//{5, 1},
		//{0, 1},
	}
	bit := make([]int, 9)
	ans := make([]int, N)
	ans2 := make([]int, N)
	//var maxV int
	for i, v := range base {
		ans[bitQuery(bit, v[0]+1)]++
		//maxV = max(maxV, v[0]+1)
		//ans2[i] = bitQuery(bit, maxV) - bitQuery(bit, v[0]+1) //求左边大于当前的个数
		ans2[i] = bitQuery(bit, v[0]+1)	//求左边小于当前的个数
		bitUpdate(bit, 8, v[0]+1, 1)
	}
	//for i := N-1; i >= 0; i-- {
	//	ans[bitQuery(bit, base[i][0]+1)]++
	//	ans2[i] = bitQuery(bit, base[i][0]+1)  //求右边小于当前的个数
	//	bitUpdate(bit, 8, base[i][0]+1, 1)
	//}
	fmt.Println(ans)
	fmt.Println(ans2)
}
//https://blog.csdn.net/Codeblocksm/article/details/52411928求区间内出现至少出现两次/不相同的数的个数
func Test_BIT2(t *testing.T) {
	//至少出现两次
	//src := []int{1, 3, 3, 7, 1, 2, 4, 3, 5, 4, 15, 3, 1, 2}
	//N := 16
	//dup := make([]int, N)
	//xi := make([]int, N)
	//visited := make([]bool, N)
	//for _, v := range src {
	//	if xi[v] != 0 && !visited[v] {
	//		bitUpdate(dup, N-1, v, 1)	//如果重复了，就更新1，并置为访问过，下次再遇到直接pass
	//		visited[v] = true
	//	}
	//	xi[v]++
	//}
	//fmt.Println(dup)
	//fmt.Println(bitQuery(dup, N-1))

	//不相同的数
	src := []int{1, 3, 7, 4, 5, 15, 2}
	N := 16
	diff := make([]int, N)
	xi := make([]int, N)
	visited := make([]bool, N)
	for _, v := range src {
		if xi[v] != 0 && !visited[v] {
			bitUpdate(diff, N-1, v, -1)	//如果重复了，就更新-1，并置为访问过，下次再遇到直接pass
			visited[v] = true
		} else if xi[v] == 0 {
			bitUpdate(diff, N-1, v, 1)	//之前没有重复，就更新1，并更新重复数组
			xi[v]++
		}
	}
	fmt.Println(diff)
	fmt.Println(bitQuery(diff, N-1))
}
