package kit

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

//最大流Edmond Karp算法。思想：反复从源点到汇点通过BFS寻找增光路径，如果找到则更新路径上的正向边与反向边可用容量值，并更新最大流。直到没有增光路经。
func BFS(c [][]int, M, s, t int) []int {
	que := []int{s}
	visited := make([]bool, M)
	visited[s] = true
	pre := make([]int, M)
	pre[s] = -1
	for len(que) > 0 {
		cur := que[0]
		for i := 0; i < M; i++ {
			if c[cur][i] > 0 && !visited[i] {
				pre[i] = cur
				visited[i] = true
				if i == t {
					return pre
				}
				que = append(que, i)
			}
		}
		que = que[1:]
	}
	return nil
}

func EK(c [][]int, M, s, t int) int {
	var maxFlow, d int
	var pre []int
	for pre = BFS(c, M, s, t); pre != nil ; pre = BFS(c, M, s, t) {
		d = math.MaxInt32
		for i := t; i != s; i = pre[i] {
			d = min(d, c[pre[i]][i])
		}
		for i := t; i != s; i = pre[i] {
			c[pre[i]][i] -= d
			c[i][pre[i]] += d
		}
		maxFlow += d
	}
	return maxFlow
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func EKRead() (int, [][]int) {
	sc := bufio.NewScanner(os.Stdin)
	var i int
	EM := make([]int, 2)
	for ; i < 2; i++ {
		sc.Scan()
		EM[i], _ = strconv.Atoi(sc.Text())
	}
	c := make([][]int, EM[1])
	for i := range c {
		c[i] = make([]int, EM[1])
	}
	for i = 0; i < EM[0]; i++ {
		sc.Scan()
		str := strings.Split(sc.Text(), " ")
		nums := make([]int, len(str))
		for i, v := range str {
			nums[i], _ = strconv.Atoi(v)
		}
		c[nums[0]-1][nums[1]-1] = nums[2]
	}
	return EM[1], c
}

func EK2(c [][]int, N, s, t int) int {
	var maxFlow int
	var pre []int
	var d int
	for pre = BFS2(c, N, s, t); pre != nil; pre = BFS2(c, N, s, t) {
		d = math.MaxInt32
		for i := t; i > s; i = pre[i] {
			d = min(d, c[pre[i]][i])
		}
		for i := t; i > s; i = pre[i] {
			c[pre[i]][i] -= d
			c[i][pre[i]] += d
		}
		maxFlow += d
	}
	return maxFlow
}
func BFS2(c [][]int, N, s, t int) []int {
	var pre []int
	var visited []bool
	var nodes []int
	pre = make([]int, N)
	pre[s] = -1
	visited = make([]bool, N)
	visited[s] = true
	nodes = []int{s}
	for len(nodes) > 0 {
		cur := nodes[0]
		for i := 0; i < N; i++ {
			if c[cur][i] > 0 && !visited[i] {
				pre[i] = cur
				if i == t {
					return pre
				}
				visited[i] = true
				nodes = append(nodes, i)
			}
		}
		nodes = nodes[1:]
	}
	return nil
}
func genGraph(input []string, M int) ([][]int, int, int){
	c := make([][]int, M+2)
	for i := range c {
		c[i] = make([]int, M+2)
	}
	var tmp []string
	var src, dst, flow int
	srcStatus, dstStatus := make([]bool, M+1), make([]bool, M+1)
	for _, s := range input {
		tmp = strings.Split(s, " ")
		src, _ = strconv.Atoi(tmp[0])
		dst, _ = strconv.Atoi(tmp[1])
		flow, _ = strconv.Atoi(tmp[2])
		c[src][dst] = flow
		srcStatus[src] = true
		dstStatus[dst] = true
	}
	var srcPoints, dstPoints []int
	for i := 1; i <= M; i++ {
		if srcStatus[i] == false {
			dstPoints = append(dstPoints, i)
		}
		if dstStatus[i] == false {
			srcPoints = append(srcPoints, i)
		}
	}
	for _, i := range srcPoints {
		c[0][i] = math.MaxInt32
	}
	for _, i := range dstPoints {
		c[i][M+1] = math.MaxInt32
	}
	return c, 0, M+1
}
//----------------------多源点多汇点
func EK3_Multi_Origin(c [][]int, M, src, dst int) int {
	var res int
	pre := BFS3_Multi_Origin(c, M, src, dst)
	for pre != nil {
		d := math.MaxInt32
		for i := dst; i != src; i = pre[i] {
			if c[pre[i]][i] < d {
				d = c[pre[i]][i]
			}
		}
		for i := dst; i != src; i = pre[i] {
			c[pre[i]][i] -= d
			c[i][pre[i]] += d
		}
		res += d
		pre = BFS3_Multi_Origin(c, M, src, dst)
	}
	return res
}
func BFS3_Multi_Origin(c [][]int, M, src, dst int) []int {
	pre := make([]int, M)
	vis := make([]bool, M)
	nodes := []int{src}
	for len(nodes) > 0 {
		node := nodes[0]
		vis[node] = true
		for i := 0; i < M; i++ {
			if c[node][i] != 0 && !vis[i] {
				pre[i] = node
				if i == dst {
					return pre
				}
				vis[i] = true
				nodes = append(nodes, i)
			}
		}
		nodes = nodes[1:]
	}
	return nil
}
