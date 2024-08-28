package slicify

// RemoveElement place all recurring val in s and
// return number of element that is not val
func RemoveElement(nums []int, val int) int {
	res := 0
	i := 0
	j := len(nums) - 1

	for _, v := range nums {
		if v != val {
			res++
		}
	}

	for i < j {
		for nums[j] == val && i < j {
			j--
		}
		for nums[i] != val && i < j {
			i++
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return res
}

func ProductExceptSelf(nums []int) []int {
	//left, right, l, r, j := make([]int, len(nums)), make([]int, len(nums)), 1, 1, len(nums)-1
	res, l, r, j := make([]int, len(nums)), 1, 1, len(nums)-1
	for i := range nums {
		res[i] = l
		l *= nums[i]
		res[j] *= r
		r *= nums[j]
		j--
		//left[i] = l
		//right[j] = r
		//l *= nums[i]
		//r *= nums[j]
		//j--
	}
	//for i, num := range left {
	//	left[i] = num * right[i]
	//}
	//return left
	return res
}
