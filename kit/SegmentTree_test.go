package kit

import (
	"fmt"
	"testing"
)

func Test_ST(t *testing.T) {
	base := []int{1,2,3,4,5,6}
	st := make([]*SegTree, 4*len(base))
	build(st, base, 1, 0, len(base)-1)
	val4 := singleQuery(st, 1, 4)
	sum14, max14, min14 := intervalQuery(st, 1, 1, 4)
	fmt.Printf("Query: index_4: %d; index_1~4: %d %d %d;\n", val4, sum14, max14, min14)
	singleUpdate(st, 1, 2, 4)
	fmt.Printf("Single Update: index_2: +4;\n")
	val2 := singleQuery(st, 1, 2)
	sum25, max25, min25 := intervalQuery(st, 1, 2, 5)
	fmt.Printf("Query: index_2: %d; index_2~5: %d %d %d;\n", val2, sum25, max25, min25)
	intervalUpdate(st, 1, 2, 5, 3)
	fmt.Printf("Interval Update: index_2~5: +3;\n")
	val5 := singleQuery(st, 1, 5)
	sum15, max15, min15 := intervalQuery(st, 1, 1, 5)
	fmt.Printf("Query: index_5: %d; index_1~5: %d %d %d;\n", val5, sum15, max15, min15)
}
