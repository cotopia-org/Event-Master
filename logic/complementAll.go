package logic

import "sort"

// ComplementAll calculates the complement of the union of given line segments in given boundaries
func ComplementAll(segments []LineSegment, minBound, maxBound float64) []LineSegment {
	// If there are no segments, return the whole range as complement
	if len(segments) == 0 {
		if minBound < maxBound {
			return []LineSegment{{minBound, maxBound}}
		}
		return []LineSegment{}
	}

	// Sort the segments by their starting point
	sort.Slice(segments, func(i, j int) bool {
		if segments[i].Start == segments[j].Start {
			return segments[i].End < segments[j].End
		}
		return segments[i].Start < segments[j].Start
	})

	// Merge overlapping or touching segments
	var merged []LineSegment
	current := segments[0]
	for i := 1; i < len(segments); i++ {
		if segments[i].Start <= current.End {
			// Merge overlapping or contiguous segments
			if segments[i].End > current.End {
				current.End = segments[i].End
			}
		} else {
			merged = append(merged, current)
			current = segments[i]
		}
	}
	merged = append(merged, current) // add the last segment

	// Calculate the complement by comparing gaps between merged segments and the bounds
	var complement []LineSegment
	start := minBound

	for _, seg := range merged {
		if start < seg.Start {
			complement = append(complement, LineSegment{start, seg.Start})
		}
		start = seg.End
	}

	// Handle any remaining space after the last segment
	if start < maxBound {
		complement = append(complement, LineSegment{start, maxBound})
	}

	return complement
}
