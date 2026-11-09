# 4106. 找出两个字符串的不匹配片段

[Problem link](https://leetcode.cn/problems/find-mismatched-segments-between-two-strings/description/)

给你两个长度相同的字符串 `s1` 和 `s2`。

不匹配片段是一个极大连续下标区间，其中每个下标处 `s1` 和 `s2` 的字符都不相同。

返回一个二维数组，其中包含每个不匹配片段对应的 `[sub1, sub2]`，`sub1` 和 `sub2` 分别是 `s1` 和 `s2` 中对应的子字符串。请按片段起始下标的递增顺序返回。如果没有不匹配片段，则返回一个空数组。

子字符串是字符串中一个连续的非空字符序列。

## Function signature

```go
func findMismatchedSegments(s1 string, s2 string) [][]string
```

## Examples

### Example 1

```text
Input: s1 = "earth", s2 = "empty"
Output: [["ar","mp"],["h","y"]]
```

下标 1 和 2 处的字符不同，形成不匹配片段 `["ar","mp"]`。下标 4 处的字符不同，形成不匹配片段 `["h","y"]`。

### Example 2

```text
Input: s1 = "action", s2 = "nation"
Output: [["ac","na"]]
```

下标 0 和 1 处的字符不同，形成不匹配片段 `["ac","na"]`。两个字符串中其余位置的字符均相同。

### Example 3

```text
Input: s1 = "hot", s2 = "hot"
Output: []
```

两个字符串完全相同，因此没有不匹配片段。

## Constraints

- `1 <= s1.length == s2.length <= 100`
- `s1` 和 `s2` 仅由小写英文字母组成

## Solution

todo
