# 4105. 重复字符串匹配 II

[Problem link](https://leetcode.cn/problems/repeated-string-match-ii/description/)

给你两个由小写英文字母组成的字符串 `a` 和 `b`，以及两个整数 `x` 和 `y`。

如果将 `b` 重复恰好 `y` 次得到的字符串（`b * y`）是将 `a` 重复恰好 `x` 次得到的字符串（`a * x`）的子串，返回 `true`。否则返回 `false`。

子串是字符串内连续的非空字符序列。

## Function signature

```go
func repeatedStringMatch(a string, b string, x int, y int) bool
```

## Examples

### Example 1

```text
Input: a = "ab", b = "abab", x = 2, y = 1
Output: true
```

`a * x = "abab"`，`b * y = "abab"`，两者相同。

### Example 2

```text
Input: a = "abc", b = "bc", x = 2, y = 3
Output: false
```

`a * x = "abcabc"`，`b * y = "bcbcbc"`。`bcbcbc` 不是 `abcabc` 的子串。

### Example 3

```text
Input: a = "abcd", b = "cdab", x = 2, y = 1
Output: true
```

`a * x = "abcdabcd"`，`b * y = "cdab"` 出现在下标 `[2..5]`。

## Constraints

- `1 <= a.length, b.length <= 100`
- `1 <= x, y <= 10^9`
- `a` 和 `b` 仅由小写英文字母组成

## Solution

todo
