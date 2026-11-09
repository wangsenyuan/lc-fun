# 4103. 按类型和颜色删除元素

[Problem link](https://leetcode.cn/problems/remove-elements-by-type-and-color/description/)

给你三个整数数组 `types`、`colors` 和 `groupSize`。`types` 和 `colors` 的长度是 `n`，`groupSize` 的长度是 `m`。

- `types[i]` 是下标 `i` 处元素的类型。
- `colors[i]` 是下标 `i` 处元素的颜色。
- `groupSize[t]` 是类型 `t` 的删除组需要的元素个数。

一次操作可以删掉恰好 `groupSize[t]` 个连续元素，前提是它们的类型都是 `t`，并且颜色相同。

每次删除后，剩下的元素重新相邻，相对顺序不变。

如果能通过一系列操作删掉所有元素，返回 `true`。否则返回 `false`。

## Function signature

```go
func canRemoveAll(types []int, colors []int, groupSize []int) bool
```

## Examples

### Example 1

```text
Input: types = [0,1,1,0], colors = [5,3,3,5], groupSize = [2,2]
Output: true
```

先删掉下标 1 和 2：类型都是 1，颜色都是 3，长度等于 `groupSize[1]`。剩下 `[0,0]` 和 `[5,5]`，再按 `groupSize[0] = 2` 删掉。

### Example 2

```text
Input: types = [0,1,0], colors = [9,9,9], groupSize = [3,2]
Output: false
```

没有任何一段连续元素满足对应类型的组大小。

### Example 3

```text
Input: types = [2,0,1,2], colors = [7,5,4,7], groupSize = [3,1,2]
Output: false
```

可以删掉下标 2（类型 1，颜色 4，`groupSize[1] = 1`）。剩下 `types = [2,0,2]`、`colors = [7,5,7]`，无法再删。

## Constraints

- `1 <= n == len(types) == len(colors) <= 10^5`
- `1 <= m == len(groupSize) <= 10^5`
- `0 <= types[i] <= m - 1`
- `1 <= colors[i] <= 10^5`
- `1 <= groupSize[t] <= n`

## Solution

todo
