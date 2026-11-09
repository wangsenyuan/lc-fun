# 4109. 统计有效后缀排序的数量 II

[Problem link](https://leetcode.cn/problems/count-valid-suffix-sorts-ii/description/)

给你两个整数数组 `nums` 和 `target`，长度分别为 `n` 和 `m`，其中 `n <= m`。

对于每个下标 `i`，考虑将原数组的后缀 `nums[i..n - 1]` 按非递增顺序排序。

返回一个长度为 `m - n + 1` 的整数数组 `ans`，其中 `ans[k]` 表示使排序后的数组等于 `target[k..k + n - 1]` 的下标 `i` 的数量。

子数组是数组中一个连续的非空元素序列。

数组的后缀是从任意下标开始并延伸至数组末尾的子数组。

## Function signature

```go
func countSuffix(nums []int, target []int) []int
```

## Examples

### Example 1

```text
Input: nums = [2,2,3], target = [2,2,3,2,1,3,2,2]
Output: [1,1,0,0,0,1]
```

分别对从下标 0、1 和 2 开始的后缀排序，会依次得到 `[3, 2, 2]`、`[2, 3, 2]` 和 `[2, 2, 3]`。

- `k = 0`：选择 `i = 2` 可以得到 `[2, 2, 3]`。
- `k = 1`：选择 `i = 1` 可以得到 `[2, 3, 2]`。
- `k = 2, 3, 4`：均无法得到对应的子数组。
- `k = 5`：选择 `i = 0` 可以得到 `[3, 2, 2]`。

因此答案为 `[1, 1, 0, 0, 0, 1]`。

### Example 2

```text
Input: nums = [5,6,5], target = [5,6,5,5,6,5]
Output: [2,1,0,2]
```

对从下标 0 开始的后缀排序，会得到 `[6, 5, 5]`；对从下标 1 或 2 开始的后缀排序，都会得到 `[5, 6, 5]`。

- `k = 0`：选择 `i = 1` 或 `i = 2` 可以得到 `[5, 6, 5]`。
- `k = 1`：选择 `i = 0` 可以得到 `[6, 5, 5]`。
- `k = 2`：无法得到 `[5, 5, 6]`。
- `k = 3`：选择 `i = 1` 或 `i = 2` 可以得到 `[5, 6, 5]`。

因此答案为 `[2, 1, 0, 2]`。

## Constraints

- `n == nums.length`
- `m == target.length`
- `1 <= n <= m <= 10^5`
- `1 <= nums[i], target[i] <= 10^5`

## Solution

todo
