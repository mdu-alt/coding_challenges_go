package arrays

// TwoSum finds in slice the indices of the two numbers such that they add up
// to target.
// It returns a nil slice if no such pair exists.
//
//   Constraints:
//       - At most one pair exists.
//
// LeetCode: https://leetcode.com/problems/two-sum/
func TwoSum(slice []int, n int) []int {
	hashmap := make(map[int]int)

	for i, v := range slice {
		if j, ok := hashmap[v]; ok {
			return []int{j, i}
		}

		hashmap[n-v] = i
	}

	return nil
}
