package algorithms

func LFU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	freq := make(map[int]int)       // Tần suất xuất hiện
	inPage := make(map[int]int)     // Vị trí trong frames
	insertTime := make(map[int]int) // Thời điểm được thêm vào
	timeCounter := 0                // Đếm thời gian để tie-break FIFO
	currentIndex := 0

	for _, page := range pages {
		_, found := inPage[page]

		if !found {
			result.PageFaults++
			if currentIndex < frameCount {
				frames[currentIndex] = page
				freq[page] = 1
				inPage[page] = currentIndex
				insertTime[page] = timeCounter
				currentIndex++
			} else {
				// Tìm trang có freq nhỏ nhất, nếu bằng thì theo FIFO
				lfuPage := -1
				lfuIdx := -1
				minFreq := int(1e9)
				oldestTime := int(1e9)

				for p := range inPage {
					if freq[p] < minFreq || (freq[p] == minFreq && insertTime[p] < oldestTime) {
						minFreq = freq[p]
						oldestTime = insertTime[p]
						lfuPage = p
						lfuIdx = inPage[p]
					}
				}

				idx := lfuIdx
				delete(inPage, lfuPage)
				delete(freq, lfuPage)
				delete(insertTime, lfuPage)

				frames[idx] = page
				inPage[page] = idx
				freq[page] = 1
				insertTime[page] = timeCounter
			}
		} else {
			freq[page]++
		}

		// Create a copy of the frequency map for the step
		freqCopy := make(map[int]int)
		for k, v := range freq {
			freqCopy[k] = v
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), freqCopy)
		result.Steps = append(result.Steps, step)
		timeCounter++
	}

	return result
}
