# 4066. 至多一次替换后的最大相邻相等元素对数

[Problem link](https://leetcode.cn/problems/maximum-equal-adjacent-pairs-after-at-most-one-replacement/)

给你一个下标从 1 开始的整数数组 `nums`。

你可以选择两个不同的值 `x` 和 `y`，并最多执行一次以下操作：

- 将 `nums` 中所有值为 `x` 的元素替换为 `y`。

返回执行操作后，相邻且相等的元素对数量的最大值。

## Function signature

```go
func maxEqualAdjacentPairs(nums []int) int
```

## Examples

### Example 1

```text
Input: nums = [1,2,3,2]
Output: 2
```

一种最优方案是选择 `x = 3` 和 `y = 2`。得到的数组为 `[1, 2, 2, 2]`。有 2 对相邻且相等的元素：`(nums[2], nums[3])` 和 `(nums[3], nums[4])`。

### Example 2

```text
Input: nums = [1,2,1,2,1]
Output: 4
```

一种最优方案是选择 `x = 1` 和 `y = 2`。得到的数组为 `[2, 2, 2, 2, 2]`。有 4 对相邻且相等的元素：`(nums[1], nums[2])`、`(nums[2], nums[3])`、`(nums[3], nums[4])` 和 `(nums[4], nums[5])`。

### Example 3

```text
Input: nums = [1,1,1]
Output: 2
```

一种最优方案是不执行任何操作。得到的数组仍为 `[1, 1, 1]`。有 2 对相邻且相等的元素：`(nums[1], nums[2])` 和 `(nums[2], nums[3])`。

## Constraints

- `2 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`

## Solution

todo
