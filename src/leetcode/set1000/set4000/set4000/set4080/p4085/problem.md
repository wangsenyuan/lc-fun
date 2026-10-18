# 4085. 找到颜色树的根

[Problem link](https://leetcode.cn/problems/find-root-of-a-colored-tree/description/)

给你一棵含有 `n` 个节点的无向树，节点编号从 `0` 到 `n - 1`，用一个长度为 `n - 1` 的二维整数数组 `edges` 表示，其中 `edges[i] = [u_i, v_i]` 表示节点 `u_i` 和 `v_i` 之间存在一条无向边。

同时给你一个长度为 `n` 的字符串 `colors`，其中 `colors[i]` 是 `'R'`、`'B'` 或 `'W'` 之一，表示节点 `i` 的颜色。

一个被选作根节点的节点，在以它为起点将树定向（即从根向外定向）之后，必须同时满足以下两个条件：

1. 二叉树性质：每个节点最多有两个子节点。
2. 分层颜色模式：处于相同深度的所有节点都具有相同的颜色。从深度 0 开始，这些颜色遵循以下三种循环序列之一：
   - `R → B → W → R → B → W → ...`
   - `B → W → R → B → W → R → ...`
   - `W → R → B → W → R → B → ...`

返回满足这两个条件的任意根节点的下标。如果不存在这样的根节点，则返回 `-1`。

树是一个无环的连通无向图。

节点的深度是从根节点到该节点的路径上的边数。根节点的深度为 0。一个节点的子节点是指与其相邻且深度大 1 的节点。

节点的度是指与该节点相连的边数。

## Function signature

```go
func findColoredRoot(edges [][]int, colors string) int
```

## Examples

### Example 1

```text
Input: edges = [[3,1],[3,2],[1,4],[2,0]], colors = "BRRWB"
Output: 3
```

以节点 3 为根节点时，每个节点最多有两个子节点，且每条边都向循环中的下一个颜色推进（`W → R → B`），因此返回 3。

### Example 2

```text
Input: edges = [[0,1],[1,2],[2,3]], colors = "RBRB"
Output: -1
```

没有一个根节点能在每条边上都满足所需的颜色递进，因此返回 -1。

## Constraints

- `2 <= n <= 10^5`
- `edges.length == n - 1`
- `edges[i] = [u_i, v_i]`
- `0 <= u_i, v_i < n`
- `colors.length == n`
- `colors[i]` 是 `'R'`、`'B'` 或 `'W'` 之一
- 输入图是一棵树，且每个节点的度最多为三

## Solution

todo
