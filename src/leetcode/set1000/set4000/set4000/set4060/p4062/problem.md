# 4062. 成对操作转化数组

[Problem link](https://leetcode.cn/problems/transform-array-using-pair-operations/description/)

给你两个整数数组 `source` 和 `target`。

在一次操作中，你可以选择 `source` 中两个不同的下标 `i` 和 `j`，以及任意整数 `delta`。然后按如下方式更新 `source`：

- `source[i] = source[i] + source[j] - delta`
- `source[j] = delta`

如果在执行该操作任意次（包括零次）后能够使 `source` 等于 `target`，则返回 `true`。否则，返回 `false`。

## Function signature

```go
func canTransform(source []int, target []int) bool
```

## Examples

### Example 1

```text
Input: source = [1,2,3], target = [0,2,4]
Output: true
```

选择下标 `i = 0` 和 `j = 2`，并设置 `delta = 4`。操作前，`source[0] = 1` 且 `source[2] = 3`。操作后，`source[0] = 1 + 3 - 4 = 0`，`source[2] = 4`。因此 `source` 变为 `[0, 2, 4]`，与 `target` 相等。

### Example 2

```text
Input: source = [-5,-5], target = [-15,5]
Output: true
```

选择下标 `i = 1` 和 `j = 0`，并设置 `delta = -15`。操作前，`source[1] = -5` 且 `source[0] = -5`。操作后，`source[1] = -5 + (-5) - (-15) = 5`，`source[0] = -15`。因此 `source` 变为 `[-15, 5]`，与 `target` 相等。

### Example 3

```text
Input: source = [1,2,1], target = [0,2,5]
Output: false
```

无论执行什么操作，都无法使 `source` 等于 `target`。

## Constraints

- `2 <= source.length == target.length <= 10^5`
- `-10^9 <= source[i], target[i] <= 10^9`

## Solution

todo
