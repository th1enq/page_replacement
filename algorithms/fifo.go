package algorithms

import ds "page_replacement/data_structure"

// FIFO implements First-In-First-Out page replacement
// O(nlogn)
func FIFO(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1 // Initialize with -1 to indicate empty frame
	}

	result := SimulationResult{}
	queue := ds.NewQueue(frameCount)

	// Map to store page positions in frames for O(1) lookup
	pagePositions := make(map[int]int)
	currentIndex := 0 // Track current position for empty frames

	for _, page := range pages {
		_, found := pagePositions[page]

		if !found {
			result.PageFaults++
			if currentIndex < frameCount {
				// Use currentIndex for empty frame
				frames[currentIndex] = page
				queue.Push(page)
				pagePositions[page] = currentIndex
				currentIndex++
			} else {
				// Replace oldest page
				oldestPage := queue.Pop()
				oldPosition := pagePositions[oldestPage]
				delete(pagePositions, oldestPage)

				// Replace at the old position
				frames[oldPosition] = page
				queue.Push(page)
				pagePositions[page] = oldPosition
			}
		}

		step := createStep(page, frames, !found, queue.ToSlice())
		result.Steps = append(result.Steps, step)
	}
	return result
}
