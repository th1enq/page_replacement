package algorithms

// FIFO implements First-In-First-Out page replacement
func FIFO(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1 // Initialize with -1 to indicate empty frame
	}

	result := SimulationResult{}
	queue := NewQueue(frameCount)

	// Use map for O(1) lookup of pages in frames
	pagesInFrames := make(map[int]bool)

	for _, page := range pages {
		found := pagesInFrames[page]

		if !found {
			result.PageFaults++
			if queue.size < frameCount {
				// Find first empty frame
				for i := range frames {
					if frames[i] == -1 {
						frames[i] = page
						queue.Push(page)
						pagesInFrames[page] = true
						break
					}
				}
			} else {
				// Replace oldest page
				oldestPage := queue.Pop()
				delete(pagesInFrames, oldestPage)

				// Find and replace the oldest page in frames
				for i := range frames {
					if frames[i] == oldestPage {
						frames[i] = page
						queue.Push(page)
						pagesInFrames[page] = true
						break
					}
				}
			}
		}

		step := createStep(page, frames, !found, queue.ToSlice())
		result.Steps = append(result.Steps, step)
	}
	return result
}
