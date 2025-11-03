package examples

import (
	"github.com/AdityaShakya-Treeleaf/go-examples/internal/dsa"
	"log"
	"testing"
)

func TestSmallestMissingNonNegativeInteger(t *testing.T) {
	smnni := dsa.SmallestMissingNonNegativeInteger{}
	log.Println("TestCase1.................")
	nums := []int{0, -3}
	value := 4
	expected := 2
	actual := smnni.Execute(nums, value)
	if expected != actual {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected, actual)
	}

}

func TestMaxDistinctElements(t *testing.T) {
	maxDistinct := dsa.MaxDistinctElementsAfterOperations{}
	log.Println("TestCase1.................")
	nums := []int{1, 2, 2, 3, 3, 4}
	value := 2
	expected := 6
	actual := maxDistinct.MaxDistinctElements(nums, value)
	if expected != actual {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected, actual)
	}

}
