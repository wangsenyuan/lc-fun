# 4097. 合并整数后的最小值

[Problem link](https://leetcode.cn/problems/minimum-value-of-merged-integers/description/)

给你一个长度为偶数 `n` 的整数数组 `nums`。

你需要将 `nums` 中的元素两两合并：将 `nums[0]` 与 `nums[1]` 合并，将 `nums[2]` 与 `nums[3]` 合并，依此类推，直到将 `nums[n - 2]` 与 `nums[n - 1]` 合并。

返回所有合并后得到的整数中的最小值。

合并两个整数，是指将第二个整数的所有十进制数位拼接在第一个整数的所有十进制数位之后。

## Function signature

```go
func minMergedInteger(nums []int) int
```

## Examples

### Example 1

```text
Input: nums = [23,55,42,21]
Output: 2355
```

将 `nums[0] = 23` 与 `nums[1] = 55` 合并，得到 2355。将 `nums[2] = 42` 与 `nums[3] = 21` 合并，得到 4221。合并后得到的整数中的最小值为 `min(2355, 4221) = 2355`。

### Example 2

```text
Input: nums = [34,7,11,11]
Output: 347
```

将 `nums[0] = 34` 与 `nums[1] = 7` 合并，得到 347。将 `nums[2] = 11` 与 `nums[3] = 11` 合并，得到 1111。合并后得到的整数中的最小值为 `min(347, 1111) = 347`。

### Example 3

```text
Input: nums = [9,20]
Output: 920
```

将 `nums[0] = 9` 与 `nums[1] = 20` 合并，得到 920。合并后只有一个整数，因此最小值为 920。

## Constraints

- `2 <= n == nums.length <= 100`
- `n` 为偶数
- `1 <= nums[i] <= 1000`

## Solution

todo
