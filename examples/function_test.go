package examples

import (
	"fmt"
	utils "github.com/AdityaShakya-Treeleaf/go-examples/internal/utils"
	rand2 "math/rand/v2"
	"slices"
	"testing"
)

func TestTernaryOperation(t *testing.T) {

	string1 := "a"
	string2 := "b"
	i := rand2.IntN(10) + 1
	fmt.Println(i)
	fmt.Println(utils.TernaryOperation(i > 5, string1, string2))

	int1 := 2
	int2 := 4
	i2 := rand2.IntN(10) + 1
	fmt.Println(i2)
	fmt.Println(utils.TernaryOperation(i2 > 5, int1, int2))

}

func TestIsSortedFunc(t *testing.T) {
	nums := []int{4, 4, 4}
	fmt.Println(slices.IsSortedFunc(nums, func(a, b int) int {
		if a == b {
			return -1
		}

		return a - b
	}))
}
