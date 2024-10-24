package logic

import "sort"

// UnionAll takes an array of LineSegments and returns the union of all of them.
func UnionAll(segments []LineSegment) []LineSegment {
	if len(segments) == 0 {
		return []LineSegment{}
	}

	// Sort segments by starting point
	// func(i, j int) bool: This is a comparator function that takes two indices, i and j,
	// and compares the elements at those indices (segments[i] and segments[j]).
	// It returns true if segments[i] should appear before segments[j] in the sorted order,
	// and false otherwise.
	sort.Slice(segments, func(i, j int) bool {
		if segments[i].Start == segments[j].Start {
			return segments[i].End < segments[j].End
		}
		return segments[i].Start < segments[j].Start
	})

	// Initialize the result array with the first segment
	result := []LineSegment{segments[0]}

	// Iterate through the segments and merge overlapping or adjacent ones
	for _, current := range segments[1:] {
		last := &result[len(result)-1]
		if current.Start <= last.End {
			// Overlapping or adjacent segments, merge them
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			// No overlap, add current segment to the result
			result = append(result, current)
		}
	}

	return result
}
