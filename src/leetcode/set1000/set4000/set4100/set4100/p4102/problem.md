# 4102. 统计有效子字符串 I

[Problem link](https://leetcode.cn/problems/count-valid-substrings-i/description/)

给你一个仅由字符 `'a'`、`'b'` 和 `'c'` 组成的字符串 `s`。

如果一个字符串不包含 `"abc"` 作为子串，那么它被称为有效的。

返回 `s` 中有效子串的数量。起始或结束下标不同的子串分别计数。

子串是字符串内连续的非空字符序列。

## Function signature

```go
func countValidSubstrings(s string) int64
```

## Examples

### Example 1

```text
Input: s = "abac"
Output: 10
```

以下子串不包含 `"abc"` 作为子串：

- `s[0..0] = "a"`
- `s[0..1] = "ab"`
- `s[0..2] = "aba"`
- `s[0..3] = "abac"`
- `s[1..1] = "b"`
- `s[1..2] = "ba"`
- `s[1..3] = "bac"`
- `s[2..2] = "a"`
- `s[2..3] = "ac"`
- `s[3..3] = "c"`

因此答案是 10。

### Example 2

```text
Input: s = "abc"
Output: 5
```

以下子串不包含 `"abc"` 作为子串：

- `s[0..0] = "a"`
- `s[0..1] = "ab"`
- `s[1..1] = "b"`
- `s[1..2] = "bc"`
- `s[2..2] = "c"`

因此答案是 5。

### Example 3

```text
Input: s = "ac"
Output: 3
```

`s` 中的所有子串都不包含 `"abc"` 作为子串。因此答案是 3。

## Constraints

- `1 <= s.length <= 10^5`
- `s` 仅由字符 `'a'`、`'b'` 和 `'c'` 组成

## Solution

todo
