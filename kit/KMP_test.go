package kit

import (
	"fmt"
	"testing"
)

func Test_kmp(t *testing.T) {
	s, p := "DACBAACADC", "ACA"
	fmt.Println(KMP(s, p))
}
