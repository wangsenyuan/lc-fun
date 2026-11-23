# 4123. 删除数组的最小代价

[Problem link](https://leetcode.cn/problems/minimum-cost-to-delete-array/description/)

给你一个整数数组 `nums`。

每次操作，你可以从当前数组中任选一个长度至少为 2 的子数组，并将其删除。该操作的代价为所选子数组的首尾元素之和。

每次操作后，将剩余元素按原有顺序拼接起来。

返回删除整个数组所需的最小总代价。

子数组是数组中一段连续且非空的元素序列。

## Function signature

```go
func minCost(nums []int) int64
```

## Examples

### Example 1

```text
Input: nums = [2,-7,3,-5]
Output: -7
```

- 删除 `[-7, 3]`，代价为 `-7 + 3 = -4`。数组变为 `[2, -5]`。
- 删除 `[2, -5]`，代价为 `2 + (-5) = -3`。
- 总代价为 `-4 + (-3) = -7`。

### Example 2

```text
Input: nums = [4,0,-5,2]
Output: 1
```

- 删除 `[0, -5]`，代价为 `0 + (-5) = -5`。数组变为 `[4, 2]`。
- 删除 `[4, 2]`，代价为 `4 + 2 = 6`。
- 总代价为 `-5 + 6 = 1`。

### Example 3

```text
Input: nums = [7,-8,1]
Output: 8
```

要删除整个数组，必须在一次操作中删除全部三个元素。如果只删除两个元素，剩下的一个元素将无法删除。删除 `[7, -8, 1]` 的代价为 `7 + 1 = 8`。

## Constraints

- `2 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

## Solution

todo
