# 4114. 统计一致子数组数目 II

[Problem link](https://leetcode.cn/problems/count-consistent-subarrays-ii/description/)

给你一个整数数组 `nums`。

如果一个子数组中所有元素的按位与运算结果等于该子数组的最小值，则称该子数组为一致子数组。

返回 `nums` 中一致子数组的个数。

子数组是数组中一个连续、非空的元素序列。

## Function signature

```go
func countConsistent(nums []int) int64
```

## Examples

### Example 1

```text
Input: nums = [5,3,1,2]
Output: 6
```

以下子数组是一致子数组：

| 子数组 | 包含的值 | 按位与 | 最小值 |
| --- | --- | --- | --- |
| nums[0..0] | [5] | 5 | 5 |
| nums[1..1] | [3] | 3 | 3 |
| nums[2..2] | [1] | 1 | 1 |
| nums[3..3] | [2] | 2 | 2 |
| nums[1..2] | [3, 1] | 3 & 1 = 1 | 1 |
| nums[0..2] | [5, 3, 1] | 5 & 3 & 1 = 1 | 1 |

因此，一致子数组的个数为 6。

### Example 2

```text
Input: nums = [1,1,1]
Output: 6
```

对于每个子数组，其按位与运算结果和最小值均等于 1。因此，答案为 6。

### Example 3

```text
Input: nums = [0]
Output: 1
```

唯一的子数组是 `[0]`，其按位与运算结果和最小值均为 0。因此，答案为 1。

## Constraints

- `1 <= nums.length <= 2 * 10^5`
- `0 <= nums[i] <= 10^6`

## Solution

todo
