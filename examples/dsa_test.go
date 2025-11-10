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

func TestMinOperationConvertToZero(t *testing.T) {
	moctz := dsa.MinOperationConvertToZero{}
	log.Println("TestCase1.................")
	nums := []int{0, 2}
	expected := 1
	actual := moctz.MinOperations(nums)
	if expected != actual {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected, actual)
	}

	log.Println("TestCase2.................")
	nums2 := []int{3, 1, 2, 1}
	expected2 := 3
	actual2 := moctz.MinOperations(nums2)
	if expected2 != actual2 {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected2, actual2)
	}

}
