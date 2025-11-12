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

func TestMinOperationConvertToOne(t *testing.T) {
	mocto := dsa.MinOperationsToOne{}
	log.Println("TestCase1.................")
	nums := []int{2, 6, 3, 4}
	expected := 4
	actual := mocto.MinOperations(nums)
	if expected != actual {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected, actual)
	}

	log.Println("TestCase2.................")
	nums2 := []int{2, 10, 6, 14}
	expected2 := -1
	actual2 := mocto.MinOperations(nums2)
	if expected2 != actual2 {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected2, actual2)
	}

	log.Println("TestCase2.................")
	nums3 := []int{6, 10, 15}
	expected3 := 4
	actual3 := mocto.MinOperations(nums3)
	if expected3 != actual3 {
		t.Fatalf("failed. Expected: %d, Actual: %d", expected3, actual3)
	}
}
