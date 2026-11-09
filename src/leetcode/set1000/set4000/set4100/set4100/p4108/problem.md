# 4108. 使红色节点连通的最小代价

[Problem link](https://leetcode.cn/problems/minimum-cost-to-make-red-nodes-connected/description/)

给你一个整数 `n`，表示一棵有 `n` 个节点的树，节点编号从 0 到 `n - 1`，并且树以节点 0 为根节点。这棵树由一个长度为 `n - 1` 的二维数组 `edges` 表示，其中 `edges[i] = [u_i, v_i]` 表示节点 `u_i` 和 `v_i` 之间存在一条无向边。

另给你一个长度为 `n` 的字符串 `colors`，其中 `colors[i]` 表示节点 `i` 的颜色：`'R'` 表示红色，`'B'` 表示蓝色。

你可以翻转任意节点的颜色：从 `'R'` 变为 `'B'`，或从 `'B'` 变为 `'R'`。翻转节点 `u` 的代价等于以节点 `u` 为根的子树的大小。

你的目标是通过翻转颜色，使所有红色节点在树中形成一个连通分量。蓝色节点可以不连通。没有红色节点的配置也视为有效。

返回实现这一目标所需的最小总代价。

在有根树中，节点 `v` 的子树是由所有到根节点的路径经过 `v` 的节点组成的集合。

## Function signature

```go
func minCost(n int, colors string, edges [][]int) int64
```

## Examples

### Example 1

```text
Input: n = 3, colors = "RBR", edges = [[0,1],[1,2]]
Output: 1
```

有两个红色节点（0 和 2），它们彼此不连通。可以将节点 2 从 `'R'` 翻转为 `'B'`。节点 2 的子树大小为 1，因此代价为 1。翻转后，只有节点 0 仍为红色，恰好形成一个红色连通分量。

### Example 2

```text
Input: n = 3, colors = "RBR", edges = [[0,1],[0,2]]
Output: 0
```

有两个红色节点（0 和 2），它们已经彼此连通。所有红色节点已经形成一个连通分量，因此无须翻转。最小总代价为 0。

### Example 3

```text
Input: n = 4, colors = "BRRR", edges = [[0,1],[0,2],[0,3]]
Output: 2
```

有三个红色节点（1、2 和 3），它们彼此不连通。这三个红色节点都是叶节点。为了得到一个红色连通分量，保留一个红色叶节点，并将另外两个翻转为蓝色。每个叶节点的子树大小均为 1，因此翻转其中任意两个的总代价为 2。

### Example 4

```text
Input: n = 3, colors = "BBB", edges = [[0,1],[1,2]]
Output: 0
```

所有节点均为蓝色，因此没有红色节点。空的红色节点集合也视为有效，当前状态已经有效，代价为 0。

## Constraints

- `1 <= n <= 10^5`
- `n == colors.length`
- `colors[i]` 为 `'R'` 或 `'B'`
- `edges.length == n - 1`
- `edges[i] = [u_i, v_i]`
- `0 <= u_i, v_i <= n - 1`
- 输入保证 `edges` 表示一棵有效的树

## Solution

todo
