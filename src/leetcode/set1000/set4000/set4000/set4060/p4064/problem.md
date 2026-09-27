# 4064. 至多一次取反能被 K 整除的最长子数组 II

[Problem link](https://leetcode.cn/problems/longest-subarray-divisible-by-k-with-at-most-one-negation-ii/)

给你一个整数数组 `nums` 和一个整数 `k`。

如果一个子数组的和能够被 `k` 整除，或者在将该子数组中的一个元素取反后能使和被 `k` 整除，则称该子数组是有效的。

将一个元素取反意味着将其值 `x` 替换为 `-x`。

返回最长有效子数组的长度。如果不存在有效的子数组，则返回 `0`。

子数组是数组中一个连续且非空的元素序列。

## Function signature

```go
func longestSubarray(nums []int, k int) int
```

## Examples

### Example 1

```text
Input: nums = [4,1,2], k = 3
Output: 3
```

整个数组的和为 `7`，且 `7 % 3 = 1`，因此它不能被 `k = 3` 整除。将 `nums[2] = 2` 取反，其和变为 `4 + 1 - 2 = 3`，能够被 `k` 整除。因此整个数组是一个有效子数组，其长度为 `3`。

### Example 2

```text
Input: nums = [5,3,4], k = 7
Output: 2
```

整个数组的和为 `12`，将其任意一个元素取反都无法使和被 `7` 整除。子数组 `[3, 4]` 的和为 `7`，无需取反即可被 `k = 7` 整除。因此最长有效子数组的长度为 `2`。

### Example 3

```text
Input: nums = [2,2,5], k = 6
Output: 2
```

整个数组的和为 `9`，将其任意一个元素取反都无法使和被 `6` 整除。子数组 `[2, 2]` 的和为 `4`。将其中任意一个元素取反会使其变为 `[-2, 2]` 或 `[2, -2]`，两者的和皆为 `0`。因此最长有效子数组的长度为 `2`。

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^5 <= nums[i] <= 10^5`
- `1 <= k <= 3000`

## Solution

todo
