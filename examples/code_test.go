package examples

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

func TestMaximumEnergyFromDungeon(t *testing.T) {
	case1 := []int{5, 2, -10, -5, 1}
	k1 := 3
	expectedOutput1 := 3
	actualOutput1 := maximumEnergy(case1, k1)

	fmt.Printf("\nTest case 1: Success = %t", expectedOutput1 == actualOutput1)

	case2 := []int{-2, -3, -1}
	k2 := 2
	expectedOutput2 := -1
	actualOutput2 := maximumEnergy(case2, k2)

	fmt.Printf("\nTest case 1: Success = %t", expectedOutput2 == actualOutput2)
}

func maximumEnergy(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	maxEnergy := math.MinInt

	for i := n - 1; i >= 0; i-- {
		if i+k < n {
			energy[i] += energy[i+k]
		}

		if energy[i] > maxEnergy {
			maxEnergy = energy[i]
		}
	}

	return maxEnergy
}

// Solution 2: Concurrent Processing with Goroutines
// Each chain (starting positions 0 to k-1) is processed in parallel
func maximumEnergyConcurrent(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	// Create a copy to avoid modifying original
	energyCopy := make([]int, n)
	copy(energyCopy, energy)

	// Channel to collect results from each chain
	results := make(chan int, k)
	var wg sync.WaitGroup

	// Process each chain concurrently
	// Chain i: positions i, i+k, i+2k, ...
	for start := 0; start < k && start < n; start++ {
		wg.Add(1)

		go func(chainStart int) {
			defer wg.Done()

			maxInChain := energyCopy[chainStart]

			// Process this chain backwards
			for i := chainStart; i < n; i += k {
				if i+k < n {
					energyCopy[i] += energyCopy[i+k]
				}
				if energyCopy[i] > maxInChain {
					maxInChain = energyCopy[i]
				}
			}

			results <- maxInChain
		}(start)
	}

	// Close results channel after all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Find maximum across all chains
	maxEnergy := -1 << 31 // Min int
	for chainMax := range results {
		if chainMax > maxEnergy {
			maxEnergy = chainMax
		}
	}

	return maxEnergy
}

// Solution 2 Alternative: Using sync.Mutex for shared max
func maximumEnergyConcurrentMutex(energy []int, k int) int {
	n := len(energy)

	if n == 0 {
		return 0
	} else if n == 1 {
		return energy[0]
	}

	energyCopy := make([]int, n)
	copy(energyCopy, energy)

	var (
		maxEnergy int = -1 << 31
		mu        sync.Mutex
		wg        sync.WaitGroup
	)

	// Process each chain concurrently
	for start := 0; start < k && start < n; start++ {
		wg.Add(1)

		go func(chainStart int) {
			defer wg.Done()

			localMax := energyCopy[chainStart]

			// Process chain backwards
			for i := chainStart; i < n; i += k {
				if i+k < n {
					energyCopy[i] += energyCopy[i+k]
				}
				if energyCopy[i] > localMax {
					localMax = energyCopy[i]
				}
			}

			// Update global max with mutex
			mu.Lock()
			if localMax > maxEnergy {
				maxEnergy = localMax
			}
			mu.Unlock()
		}(start)
	}

	wg.Wait()
	return maxEnergy
}
