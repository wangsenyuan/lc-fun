# 4083. 使相邻元素按位异或非零的最少删除次数

[Problem link](https://leetcode.cn/problems/minimum-deletions-to-make-array-sharp/)

给你一个整数数组 `nums`。

你可以从 `nums` 中删除任意数量的元素。剩余元素保持原有的相对顺序。

如果结果数组 `a` 中每一对相邻元素的按位异或结果都不为 0，则称数组 `a` 是相邻异或非零数组。也就是说，对于每个满足 `0 <= i < a.length - 1` 的下标 `i`，都有 `a[i] XOR a[i + 1] != 0`。长度不超过 1 的数组也被视为相邻异或非零数组。

返回使 `nums` 变为相邻异或非零数组所需的最少删除次数。

按位异或会逐位比较两个整数对应的二进制位。如果对应位不同，则结果中的该位为 1；否则为 0。

## Function signature

```go
func minDeletions(nums []int) int
```

## Examples

### Example 1

```text
Input: nums = [1,1,2,2,2,3]
Output: 3
```

下标对 `(0, 1)`、`(2, 3)` 和 `(3, 4)` 对应元素的异或结果均为 0。删除一个 1 和两个 2 后，得到 `[1, 2, 3]`，其中 `1 XOR 2 == 3` 且 `2 XOR 3 == 1`，两者都不为 0。因此答案为 3。

### Example 2

```text
Input: nums = [4,4,9,4]
Output: 1
```

下标对 `(0, 1)` 对应元素的异或结果为 0。删除这两个 4 中的任意一个后，得到 `[4, 9, 4]`，其中 `4 XOR 9 == 13` 且 `9 XOR 4 == 13`，两者都不为 0。

## Constraints

- `2 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## Solution

todo
