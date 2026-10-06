# 4073. 统计好字符串数目

[Problem link](https://leetcode.cn/problems/count-good-strings/description/)

给你一个整数 `n`。

如果一个字符串仅由字符 `'a'` 和 `'b'` 组成，且满足以下条件之一，则该字符串被认为是好的：

- 它只包含一种字符，且其长度为奇数。
- 它可以写成 `s = s1 + s2` 的形式，其中 `s1` 和 `s2` 是非空的好字符串，且 `s1` 的最后一个字符与 `s2` 的第一个字符不同。

返回长度为 `n` 的好字符串的数量，模 `10^9 + 7`。

这里，`+` 表示字符串拼接。

## Function signature

```go
func countGoodStrings(n int64) int
```

## Examples

### Example 1

```text
Input: n = 4
Output: 6
```

好的字符串有 `"aaab"`、`"abbb"`、`"baaa"`、`"bbba"`、`"abab"` 和 `"baba"`。

例如，`"aaab" = "aaa" + "b"`。这两个部分都是好的，因为每个部分都只包含一种字符且长度为奇数，并且它们在边界处的字符不同。

同样地，`"ab" = "a" + "b"` 是好的，所以 `"abab" = "ab" + "ab"` 也是好的，因为它们的边界字符不同。

### Example 2

```text
Input: n = 3
Output: 4
```

好的字符串有 `"aaa"`、`"bbb"`、`"aba"` 和 `"bab"`。

### Example 3

```text
Input: n = 2
Output: 2
```

好的字符串有 `"ab"` 和 `"ba"`。

## Constraints

- `1 <= n <= 10^15`

## Solution

todo
