package algorithms

// LRU implements Least Recently Used page replacement
func LRU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}

	// Keep track of last used time for each frame
	lastUsed := make([]int, frameCount)
	for i := range lastUsed {
		lastUsed[i] = -1
	}

	for i, page := range pages {
		found := false
		foundIdx := -1
		for j := range frames {
			if frames[j] == page {
				found = true
				foundIdx = j
				break
			}
		}

		if !found {
			result.PageFaults++
			if i < frameCount {
				frames[i] = page
				lastUsed[i] = i
			} else {
				// Find least recently used frame
				lruIdx := 0
				minTime := lastUsed[0]
				for j := 1; j < frameCount; j++ {
					if lastUsed[j] < minTime {
						minTime = lastUsed[j]
						lruIdx = j
					}
				}
				frames[lruIdx] = page
				lastUsed[lruIdx] = i
			}
		} else {
			// Update last used time for the found page
			lastUsed[foundIdx] = i
		}

		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
	}

	return result
}
