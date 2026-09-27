# 4061. 皇后到达目标格子的最少移动步数

[Problem link](https://leetcode.cn/problems/minimum-queen-moves-to-reach-target/)

一个 `8 x 8` 的空棋盘，其行和列的下标从 1 开始。

给你一个数组 `source = [sr, sc]` 表示皇后的初始位置，以及一个数组 `target = [tr, tc]` 表示目标位置。

在一步移动中，皇后可以在棋盘范围内，沿着单条行、列或对角线移动一个或多个方格。

返回皇后移动到恰好落在 `target` 位置所需的最小移动次数。

## Function signature

```go
func minQueenMoves(source []int, target []int) int
```

## Examples

### Example 1

```text
Input: source = [8,1], target = [1,8]
Output: 1
```

单次对角线移动即可让皇后直接从 `(8, 1)` 移动到 `(1, 8)`。

### Example 2

```text
Input: source = [4,2], target = [1,3]
Output: 2
```

皇后先从 `(4, 2)` 移动到 `(4, 3)`，然后从 `(4, 3)` 移动到 `(1, 3)`，共用 2 步移动到达目标位置。

### Example 3

```text
Input: source = [1,1], target = [1,1]
Output: 0
```

皇后已经处于目标位置，因此不需要任何移动。

## Constraints

- `source == [sr, sc]`
- `target == [tr, tc]`
- `1 <= sr, sc, tr, tc <= 8`

## Solution

todo
