# 4065. 移除不同值重排数组

[Problem link](https://leetcode.cn/problems/rearrange-array-by-removing-distinct-values/)

给你一个整数数组 `nums`。

初始时，你有一个空数组 `ans`。重复执行以下操作，直到 `nums` 变为空：

- 找出当前 `nums` 中所有不同的值。
- 将当前 `nums` 中每个不同的值各移除一个，并按升序将这些值依次添加到 `ans` 中。

返回数组 `ans`。

## Function signature

```go
func rearrangeArray(nums []int) []int
```

## Examples

### Example 1

```text
Input: nums = [3,1,3,2,1,3]
Output: [1,2,3,1,3,3]
```

| 操作 | 添加到 `ans` 的值 | 操作后的 `nums` | 操作后的 `ans` |
| --- | --- | --- | --- |
| 1 | 1, 2, 3 | `[3, 1, 3]` | `[1, 2, 3]` |
| 2 | 1, 3 | `[3]` | `[1, 2, 3, 1, 3]` |
| 3 | 3 | `[]` | `[1, 2, 3, 1, 3, 3]` |

此时 `nums` 已为空，因此答案为 `[1, 2, 3, 1, 3, 3]`。

### Example 2

```text
Input: nums = [7,7,4,4,4]
Output: [4,7,4,7,4]
```

| 操作 | 添加到 `ans` 的值 | 操作后的 `nums` | 操作后的 `ans` |
| --- | --- | --- | --- |
| 1 | 4, 7 | `[7, 4, 4]` | `[4, 7]` |
| 2 | 4, 7 | `[4]` | `[4, 7, 4, 7]` |
| 3 | 4 | `[]` | `[4, 7, 4, 7, 4]` |

此时 `nums` 已为空，因此答案为 `[4, 7, 4, 7, 4]`。

## Constraints

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## Solution

todo
