package algorithms

// SecondChance implements Second Chance (Clock) page replacement
func SecondChance(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}

	// Keep track of reference bits for each frame
	refBits := make([]bool, frameCount)
	for i := range refBits {
		refBits[i] = false
	}

	// Keep track of the next frame to check
	nextFrame := 0

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
				refBits[i] = true
			} else {
				// Find a frame with reference bit = 0
				startIdx := nextFrame
				for {
					if !refBits[nextFrame] {
						break
					}
					refBits[nextFrame] = false
					nextFrame = (nextFrame + 1) % frameCount
					if nextFrame == startIdx {
						// All frames have been given a second chance
						nextFrame = (nextFrame + 1) % frameCount
						break
					}
				}
				frames[nextFrame] = page
				refBits[nextFrame] = true
				nextFrame = (nextFrame + 1) % frameCount
			}
		} else {
			// Set reference bit for the found page
			refBits[foundIdx] = true
		}

		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
	}
	return result
}
