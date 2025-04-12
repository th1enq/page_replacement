package algorithms

// Optimal implements the optimal page replacement algorithm
func Optimal(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}

	for i, page := range pages {
		found := false
		for j := range frames {
			if frames[j] == page {
				found = true
				break
			}
		}

		if !found {
			result.PageFaults++
			if i < frameCount {
				frames[i] = page
			} else {
				// Find the page that will be used farthest in the future
				farthest := -1
				replaceIdx := 0

				for j := range frames {
					nextUse := -1
					for k := i + 1; k < len(pages); k++ {
						if pages[k] == frames[j] {
							nextUse = k
							break
						}
					}
					if nextUse == -1 {
						replaceIdx = j
						break
					}
					if nextUse > farthest {
						farthest = nextUse
						replaceIdx = j
					}
				}

				frames[replaceIdx] = page
			}
		}

		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
	}

	return result
}
