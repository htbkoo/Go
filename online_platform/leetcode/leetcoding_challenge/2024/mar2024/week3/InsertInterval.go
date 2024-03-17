package week3

func insert(intervals [][]int, newInterval []int) [][]int {
	var answer [][]int
	newStart := newInterval[0]
	newEnd := newInterval[1]

	inserted := false

	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]

		if end < newStart {
			answer = append(answer, interval)
			continue
		}

		if start > newEnd {
			if !inserted {
				answer = append(answer, []int{newStart, newEnd})
				inserted = true
			}

			answer = append(answer, interval)
			continue
		}

		newStart = min(newStart, start)
		newEnd = max(newEnd, end)
	}

	if !inserted {
		answer = append(answer, []int{newStart, newEnd})
		inserted = true
	}

	return answer
}
