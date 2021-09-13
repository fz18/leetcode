/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */
package leetcode

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	// binary search midLeft and midRight
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid

		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			// move to left
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums2[nums2Mid-1] > nums1[nums1Mid] {
			// move to right
			low = nums1Mid + 1
		} else {
			// find it
			break
		}
	}
	// return by midLeft and midRight value

	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	// if odd return
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	// return
	return float64(midLeft+midRight) / 2
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func min(left, right int) int {
	if left > right {
		return right
	}
	return left
}

// @lc code=end
