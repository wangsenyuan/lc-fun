# 4067. 数对和受限的最长子数组

[Problem link](https://leetcode.cn/problems/longest-subarray-with-restricted-pair-sums/description/)

给你一个整数数组 `nums`。

如果不存在三个互不相同的下标 `i`、`j` 和 `k`，满足 `l <= i, j, k <= r` 且 `nums[i] + nums[j] == nums[k]`，则子数组 `nums[l..r]` 是有效子数组。

返回 `nums` 中有效子数组的最大长度。

子数组是数组中一个连续非空元素序列。

## Function signature

```go
func maxSubarray(nums []int) int
```

## Examples

### Example 1

```text
Input: nums = [2,3,5,3,2,1]
Output: 3
```

考虑子数组 `[3, 5, 3]`。由不同下标处的元素组成的数对，其元素和如下：

- `3 + 5 = 8`
- `3 + 3 = 6`，这里使用的是两个不同位置上的 3
- `5 + 3 = 8`

这些和都不等于剩余下标处的元素，因此该子数组是有效的。

每个长度为 4 的子数组都包含位于不同下标处的 2、3 和 5，并且 `2 + 3 = 5`。因此不存在更长的有效子数组，答案为 3。

### Example 2

```text
Input: nums = [3,4,5,6]
Output: 4
```

由不同下标处的任意两个元素相加，得到的和分别为 7、8、9、9、10 和 11。这些值都不等于剩余下标处的元素，因此整个数组都是有效的。

## Constraints

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 500`

## Solution

todo
