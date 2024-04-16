package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 0 {
		return false
	}
	resultMap := make(map[int]struct{})

	low := 0
	for i, v := range nums {
		if (i - low) > k {
			delete(resultMap, low)
			low++
		}
		if _, ok := resultMap[v]; ok {
			return true
		}
		resultMap[v] = struct{}{}
	}
	return false
}
