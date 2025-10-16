package examples

import (
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/dsa"
	"log"
	"testing"
)

func TestSmallestMissingNonNegativeInteger(t *testing.T) {
	smnni := dsa.SmallestMissingNonNegativeInteger{}
	log.Println("TestCase1.................")
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 5
	expected := 4
	actual := smnni.Execute(nums, value)
	if expected != actual {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected, actual)
	}

}
