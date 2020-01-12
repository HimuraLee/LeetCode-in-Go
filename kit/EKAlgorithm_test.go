package kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EK(t *testing.T) {
	//M, c := EKRead()
	M := 4
	c := [][]int{
		{0, 40, 0, 20},
		{0, 0, 30, 20},
		{0, 0, 0, 10},
		{0, 0, 0, 0},
	}
	ast := assert.New(t)
	ast.Equal(50, EK2(c, M, 0, M-1), "Error")
}

func Test_EK_Multi_With_input(t *testing.T) {
	input := []string{
		"2 1 40",
		"2 3 20",
		"1 3 20",
		"2 4 30",
		"4 3 10",

		"5 6 40",
		"5 8 20",
		"6 8 20",
		"2 7 30",
		"7 8 10",
	}
	M := 8
	c, src, dst := genGraph(input, M)
	fmt.Println(EK3_Multi_Origin(c, M+2, src, dst))
}
