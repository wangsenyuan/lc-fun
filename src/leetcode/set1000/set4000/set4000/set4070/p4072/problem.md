# 4072. 一次删除后的最大交替子数组和

[Problem link](https://leetcode.cn/problems/maximum-alternating-subarray-sum-with-one-deletion/description/)

给你一个整数数组 `nums`。

你最多可以从 `nums` 中删除一个元素，然后在剩下数组里选一个子数组。

返回所选子数组的最大可能交替和。

子数组是数组中连续的非空元素序列。

数组的交替和是其偶数下标处元素之和减去奇数下标处元素之和。在计算其交替和之前，所选子数组会从 0 开始重新编下标。

## Function signature

```go
func maxAlternatingSum(nums []int) int64
```

## Examples

### Example 1

```text
Input: nums = [5,-5,1]
Output: 11
```

选择不删除元素，并选择整个数组。其交替和为 `5 - (-5) + 1 = 11`，这是最大可能的值。

### Example 2

```text
Input: nums = [10,-5,-100]
Output: 110
```

删除 `nums[1] = -5` 得到 `[10,-100]`，然后选择整个所得数组。其交替和为 `10 - (-100) = 110`，这是最大可能的值。

### Example 3

```text
Input: nums = [4,7]
Output: 7
```

选择不删除元素，并选择子数组 `[7]`。其交替和为 7，这是最大可能的值。

## Constraints

- `1 <= nums.length <= 10^5`
- `-10^5 <= nums[i] <= 10^5`

## Solution

todo
