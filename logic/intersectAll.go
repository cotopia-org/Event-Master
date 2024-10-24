package logic

import "math"


// Intersect function returns the intersection of two line segments
func intersect(seg1, seg2 LineSegment) (LineSegment, bool) {
	// Calculate the intersection
	start := math.Max(seg1.Start, seg2.Start)
	end := math.Min(seg1.End, seg2.End)
	
	// If there is no overlap, return an invalid segment
	if start > end {
		return LineSegment{}, false
	}
	
	return LineSegment{Start: start, End: end}, true
}

// IntersectAll function returns the intersection of many line segments
func IntersectAll(segments []LineSegment) []LineSegment {
	if len(segments) == 0 {
		return []LineSegment{}
	}

	// Start by assuming the first segment is the intersection
	currentIntersection := segments[0]

	for i := 1; i < len(segments); i++ {
		inter, ok := intersect(currentIntersection, segments[i])
		if !ok {
			// If there is no intersection, return an empty array
			return []LineSegment{}
		}
		// Update the current intersection
		currentIntersection = inter
	}

	// If we reach here, it means all segments overlap, so return the final intersection
	return []LineSegment{currentIntersection}
}