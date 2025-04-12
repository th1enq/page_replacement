package algorithms

// MFU implements Most Frequently Used page replacement
func MFU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	prevFrames := make([]int, frameCount)
	copy(prevFrames, frames)

	// Keep track of frequency for each frame
	freq := make([]int, frameCount)
	for i := range freq {
		freq[i] = 0
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
				freq[i] = 1
			} else {
				// Find most frequently used frame
				mfuIdx := 0
				maxFreq := freq[0]
				for j := 1; j < frameCount; j++ {
					if freq[j] > maxFreq {
						maxFreq = freq[j]
						mfuIdx = j
					}
				}
				frames[mfuIdx] = page
				freq[mfuIdx] = 1
			}
		} else {
			// Increment frequency for the found page
			freq[foundIdx]++
		}

		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
	}

	return result
}
