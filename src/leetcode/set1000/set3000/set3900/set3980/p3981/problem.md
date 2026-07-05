# 3981. 统计交错字符的方案数

[Problem link](https://leetcode.cn/problems/count-ways-to-interleave-characters/description/)

给你三个字符串 `word1`、`word2` 和 `target`。

你的任务是计算从 `word1` 和 `word2` 中选择字符以形成 `target` 的方案数，需满足以下条件：

- 对于 `target` 中的每个字符，从 `word1` 或 `word2` 中选择一个匹配的字符。
- 从 `word1` 中选择的下标必须是**严格**递增的。
- 从 `word2` 中选择的下标必须是**严格**递增的。
- 必须从 `word1` 和 `word2` 两者中**至少**各选择一个字符。
- 如果对于 `target` 中的**至少**一个位置，选择的字符来自不同的字符串或不同的下标，则认为两种方案是不同的。

返回方案数。由于答案可能非常大，请返回它对 `10^9 + 7` 取余后的结果。

## Function signature

```go
func interleaveCharacters(word1 string, word2 string, target string) int
```

## Examples

### Example 1

```text
Input: word1 = "abc", word2 = "bac", target = "abc"
Output: 5
```

有 5 种形成 `target` 的方案：

- `word1[0] = 'a'`, `word1[1] = 'b'`, `word2[2] = 'c'`
- `word1[0] = 'a'`, `word2[0] = 'b'`, `word1[2] = 'c'`
- `word1[0] = 'a'`, `word2[0] = 'b'`, `word2[2] = 'c'`
- `word2[1] = 'a'`, `word1[1] = 'b'`, `word1[2] = 'c'`
- `word2[1] = 'a'`, `word1[1] = 'b'`, `word2[2] = 'c'`

所有方案都保持了每个字符串内部递增的下标顺序，并且从每个字符串中至少选择了一个字符。

### Example 2

```text
Input: word1 = "cd", word2 = "cd", target = "ccd"
Output: 4
```

有 4 种形成 `target` 的方案：

- `word1[0] = 'c'`, `word2[0] = 'c'`, `word1[1] = 'd'`
- `word1[0] = 'c'`, `word2[0] = 'c'`, `word2[1] = 'd'`
- `word2[0] = 'c'`, `word1[0] = 'c'`, `word1[1] = 'd'`
- `word2[0] = 'c'`, `word1[0] = 'c'`, `word2[1] = 'd'`

`target` 中的前两个 `'c'` 字符必须分别来自两个字符串。最后一个 `'d'` 可以从任意一个字符串中选择。

### Example 3

```text
Input: word1 = "xy", word2 = "xy", target = "xyxy"
Output: 2
```

有 2 种形成 `target` 的方案：

- `word1[0] = 'x'`, `word1[1] = 'y'`, `word2[0] = 'x'`, `word2[1] = 'y'`
- `word2[0] = 'x'`, `word2[1] = 'y'`, `word1[0] = 'x'`, `word1[1] = 'y'`

`target` 中的每个 `"xy"` 部分完全来自同一个字符串。

### Example 4

```text
Input: word1 = "ab", word2 = "cde", target = "ace"
Output: 1
```

唯一的方案是选择 `word1[0] = 'a'`、`word2[0] = 'c'` 和 `word2[2] = 'e'`。因此，答案为 1。

## Constraints

- `1 <= word1.length, word2.length, target.length <= 100`
- `word1`、`word2` 和 `target` 仅由小写英文字母组成。

## Solution Summary

把 `target` 从左到右处理。令 `dp[i][j]` 表示已经形成了当前前缀，并且 `word1` 中最后一次选择的位置是 `i - 1`，`word2` 中最后一次选择的位置是 `j - 1` 的方案数；当 `i = 0` 或 `j = 0` 时，表示还没有从对应字符串中选择过字符。

处理下一个字符 `c` 时，有两类转移：

- 如果选择 `word1[x] == c`，那么之前 `word1` 的最后选择位置必须小于 `x`，而 `word2` 的最后选择位置保持不变。因此新状态是 `ndp[x+1][j] += sum(dp[0..x][j])`。
- 如果选择 `word2[y] == c`，那么之前 `word2` 的最后选择位置必须小于 `y`，而 `word1` 的最后选择位置保持不变。因此新状态是 `ndp[i][y+1] += sum(dp[i][0..y])`。

直接求这些和会多一层循环。每处理一个 `target` 字符时，先从当前 `dp` 构造两个前缀和：

- `rowPrefix[i][j] = sum(dp[0..i][j])`，用于快速转移到 `word1` 的某个位置。
- `colPrefix[i][j] = sum(dp[i][0..j])`，用于快速转移到 `word2` 的某个位置。

这样每个 `target` 字符只需要 `O(nm + n + m)` 的时间。最后答案是所有 `i > 0 && j > 0` 的 `dp[i][j]` 之和，因为题目要求两个字符串都至少选过一个字符。

复杂度：

- 时间复杂度：`O(|target| * |word1| * |word2|)`。
- 空间复杂度：`O(|word1| * |word2|)`。

### Why the previous solution was wrong

原来的转移把 `dp[w][i][j]` 只转移到 `dp[w+1][i+1][j]` 或 `dp[w+1][i][j+1]`。这等价于要求下一个选择的字符必须正好来自当前下标 `i` 或 `j`，但题目要求的是子序列，只要下标严格递增即可，中间可以跳过任意字符。

例如 `word1 = "xa"`、`word2 = "b"`、`target = "ab"` 时，合法方案是跳过 `word1[0] = 'x'`，选择 `word1[1] = 'a'`，再选择 `word2[0] = 'b'`。旧转移只能从 `dp[0][0][0]` 尝试匹配 `word1[0]`，无法直接跳到 `word1[1]`，所以会漏掉这个方案。样例 1 中从 `word2[1] = 'a'` 开始、或从 `word2[0] = 'b'` 跳到 `word2[2] = 'c'` 的方案也是同一个问题。
