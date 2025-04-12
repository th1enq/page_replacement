package algorithms

func MRU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	inPage := make(map[int]int) // page -> frame index

	lastUsed := 0

	for i, page := range pages {

		_, found := inPage[page]

		if !found {
			// Page fault
			result.PageFaults++

			if i < frameCount {
				frames[i] = page
				inPage[page] = i
			} else {
				oldPage := lastUsed
				mruIdx := inPage[oldPage]
				delete(inPage, oldPage)

				frames[mruIdx] = page
				inPage[page] = mruIdx
			}
		}
		lastUsed = page

		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
	}

	return result
}
