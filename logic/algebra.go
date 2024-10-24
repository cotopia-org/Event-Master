package logic

import "math"

// Intersection Algorithm
// Input: P1(a1), P2(a2), P3(b1), P4(b2)

// 1. Compute the intervals:
//    A1 = min(a1, a2)
//    A2 = max(a1, a2)
//    B1 = min(b1, b2)
//    B2 = max(b1, b2)

// 2. Check for intersection:
//    If max(A1, B1) <= min(A2, B2):
//       - The segments intersect.
//       - Compute the intersection interval:
//         Intersection = [max(A1, B1), min(A2, B2)]
//       - Return Intersection

// 3. Else:
//    - The segments do not intersect.
//    - Return "No intersection"

// Function to compute the intersection of two 1D segments
func Intersection(a1, a2, b1, b2 float64) (bool, float64, float64) {
	// Calculate the intervals for the two segments
	A1 := math.Min(a1, a2)
	A2 := math.Max(a1, a2)
	B1 := math.Min(b1, b2)
	B2 := math.Max(b1, b2)

	// Check if the segments intersect
	if math.Max(A1, B1) <= math.Min(A2, B2) {
		// If they intersect, return the intersection interval
		intersectionStart := math.Max(A1, B1)
		intersectionEnd := math.Min(A2, B2)
		return true, intersectionStart, intersectionEnd
	}

	// If no intersection, return false
	return false, 0, 0
}


// Union Algorithm
// Input: P1(a1), P2(a2), P3(b1), P4(b2)

// 1. Compute the intervals:
//    A1 = min(a1, a2)
//    A2 = max(a1, a2)
//    B1 = min(b1, b2)
//    B2 = max(b1, b2)

// 2. Check if the segments overlap or touch:
//    If max(A1, B1) <= min(A2, B2):
//       - The segments overlap or touch, so the union is a single interval.
//       - Compute the union interval:
//         Union = [min(A1, B1), max(A2, B2)]
//       - Return Union

// 3. Else:
//    - The segments do not overlap or touch.
//    - Return the two separate intervals:
//      Union = [A1, A2] and [B1, B2]

// Function to compute the union of two 1D segments
func Union(a1, a2, b1, b2 float64) (bool, float64, float64, float64, float64) {
	// Calculate the intervals for the two segments
	A1 := math.Min(a1, a2)
	A2 := math.Max(a1, a2)
	B1 := math.Min(b1, b2)
	B2 := math.Max(b1, b2)

	// Check if the segments overlap or touch
	if math.Max(A1, B1) <= math.Min(A2, B2) {
		// If they overlap or touch, return the combined union interval
		unionStart := math.Min(A1, B1)
		unionEnd := math.Max(A2, B2)
		return true, unionStart, unionEnd, 0, 0
	}

	// If they don't overlap, return the two separate intervals
	return false, A1, A2, B1, B2
}