package algorithms

func MFU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	freq := make(map[int]int)       // Tần suất truy cập
	inPage := make(map[int]int)     // Vị trí trang trong frames
	insertTime := make(map[int]int) // Thời điểm trang được thêm vào
	timeCounter := 0                // Dùng cho FIFO khi tie-break
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
				// Tìm trang có freq cao nhất, nếu bằng thì chọn trang vào sớm nhất (FIFO)
				mfuPage := -1
				mfuIdx := -1
				maxFreq := -1
				oldestTime := int(1e9)

				for p := range inPage {
					if freq[p] > maxFreq || (freq[p] == maxFreq && insertTime[p] < oldestTime) {
						maxFreq = freq[p]
						oldestTime = insertTime[p]
						mfuPage = p
						mfuIdx = inPage[p]
					}
				}

				idx := mfuIdx
				delete(inPage, mfuPage)
				delete(freq, mfuPage)
				delete(insertTime, mfuPage)

				frames[idx] = page
				inPage[page] = idx
				freq[page] = 1
				insertTime[page] = timeCounter
			}
		} else {
			freq[page]++
		}

		// Ghi lại bước mô phỏng
		step := createStep(page, frames, !found, append([]int{}, frames...))
		result.Steps = append(result.Steps, step)
		timeCounter++
	}

	return result
}
