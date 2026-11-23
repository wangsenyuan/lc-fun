# 4120. 统计优势元素的数量

[Problem link](https://leetcode.cn/problems/count-dominant-elements/description/)

给你一个整数数组 `nums`。

如果 `nums` 中的某个元素大于或等于其相邻元素的平均值，则称它为优势元素。

`nums[i]` 的相邻元素是 `nums[i - 1]` 和 `nums[i + 1]`，仅在对应下标存在时计入。

没有相邻元素的单个元素视为优势元素。

返回 `nums` 中优势元素的数量。

## Function signature

```go
func countDominantElements(nums []int) int
```

## Examples

### Example 1

```text
Input: nums = [1,10,2]
Output: 1
```

- 对于 `nums[0] = 1`，它唯一的相邻元素是 10。由于 `1 < 10`，它不是优势元素。
- 对于 `nums[1] = 10`，它的相邻元素是 1 和 2，平均值为 `(1 + 2) / 2 = 1.5`。由于 `10 > 1.5`，它是优势元素。
- 对于 `nums[2] = 2`，它唯一的相邻元素是 10。由于 `2 < 10`，它不是优势元素。

因此，答案是 1。

### Example 2

```text
Input: nums = [6,4,2]
Output: 2
```

- 对于 `nums[0] = 6`，它唯一的相邻元素是 4。由于 `6 > 4`，它是优势元素。
- 对于 `nums[1] = 4`，它的相邻元素是 6 和 2，平均值为 `(6 + 2) / 2 = 4`。由于 `4 = 4`，它是优势元素。
- 对于 `nums[2] = 2`，它唯一的相邻元素是 4。由于 `2 < 4`，它不是优势元素。

因此，答案是 2。

### Example 3

```text
Input: nums = [7]
Output: 1
```

- 对于 `nums[0] = 7`，它没有相邻元素，因此它是优势元素。

因此，答案是 1。

## Constraints

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## Solution

todo
