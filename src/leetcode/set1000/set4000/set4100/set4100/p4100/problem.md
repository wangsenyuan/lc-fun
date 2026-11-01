# 4100. 平衡异或路径的最小开销

[Problem link](https://leetcode.cn/problems/minimum-cost-path-with-even-xor-parity/description/)

给你一棵含有 `n` 个节点的树，节点编号从 0 到 `n - 1`，用一个二维整数数组 `edges` 表示。`edges[i] = [u_i, v_i, c_i]` 表示节点 `u_i` 和 `v_i` 之间的一条边，其开销为 `c_i`。

同时给你一个长度为 `n` 的整数数组 `values`，其中 `values[i]` 是节点 `i` 的值。

如果两个不同节点之间的路径上所有节点（包括两个端点）的值的按位异或结果在二进制表示中含有偶数个置位，则称该路径为平衡异或路径。

一条路径的开销是其所有边开销的总和。

返回一条平衡异或路径的最小开销。如果不存在这样的路径，返回 -1。

置位是指一个数字的二进制表示中值为 1 的位。数字 0 含有零个置位。

## Function signature

```go
func minActivationCost(edges [][]int, values []int) int
```

## Examples

### Example 1

```text
Input: edges = [[0,1,4],[0,2,6],[0,3,3],[3,4,5]], values = [3,1,7,1,3]
Output: 7
```

| 平衡异或路径 | 异或值 | 总开销 |
| --- | --- | --- |
| `1 - 0 - 2` | `1 XOR 3 XOR 7 = 5` | `4 + 6 = 10` |
| `1 - 0 - 3` | `1 XOR 3 XOR 1 = 3` | `4 + 3 = 7` |
| `1 - 0 - 3 - 4` | `1 XOR 3 XOR 1 XOR 3 = 0` | `4 + 3 + 5 = 12` |
| `2 - 0 - 3` | `7 XOR 3 XOR 1 = 5` | `6 + 3 = 9` |
| `2 - 0 - 3 - 4` | `7 XOR 3 XOR 1 XOR 3 = 6` | `6 + 3 + 5 = 14` |

最小开销为 7，由路径 `1 - 0 - 3` 达成。

### Example 2

```text
Input: edges = [[0,1,5],[1,2,10]], values = [1,7,3]
Output: 5
```

| 平衡异或路径 | 异或值 | 总开销 |
| --- | --- | --- |
| `0 - 1` | `1 XOR 7 = 6` | `5` |
| `0 - 1 - 2` | `1 XOR 7 XOR 3 = 5` | `5 + 10 = 15` |

最小开销为 5，由路径 `0 - 1` 达成。

### Example 3

```text
Input: edges = [[0,1,5],[0,2,7],[0,3,9]], values = [1,3,5,9]
Output: -1
```

任意两个不同节点之间的路径异或值在二进制下都含有奇数个置位。因此不存在平衡异或路径，答案为 -1。

## Constraints

- `2 <= n <= 10^5`
- `edges.length == n - 1`
- `edges[i] = [u_i, v_i, c_i]`
- `0 <= u_i, v_i <= n - 1`
- `1 <= c_i <= 10^9`
- `values.length == n`
- `1 <= values[i] <= 10^9`
- `edges` 表示一棵合法的树

## Solution

todo
