# LeetCode Practice In Go

[![Go CI](https://github.com/wangsenyuan/lc-fun/actions/workflows/go.yml/badge.svg)](https://github.com/wangsenyuan/lc-fun/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)

LeetCode and LeetCode China practice solutions in Go. This repository was split out from `learn-go` so LeetCode work can live in a smaller, focused checkout.

## Layout

```text
src/
  leetcode/
    set0000/
    set1000/
    contest/
    daily/
    util/
  leetcodecn/
    lcp13/
    lcr039/
    reversepairs/
```

Each problem directory is a self-contained Go package. Most directories contain:

```text
solution.go
solution_test.go
problem.md or readme.md, when available
```

## Run Tests

Run one problem package:

```bash
go test ./src/leetcode/set0000/set000/p025
```

Run one package from LeetCode China:

```bash
go test ./src/leetcodecn/reversepairs
```

Run a single test inside a package:

```bash
go test ./src/leetcode/set0000/set000/p025 -run TestSample1
```

`go test ./...` can be slow because the repository contains thousands of packages. Focused package tests are the normal workflow.

## GitHub Actions

The workflow in `.github/workflows/go.yml` runs on pushes and pull requests to `main`. It detects changed Go package directories and runs package-local tests for those directories.

## License

Apache License 2.0. See [LICENSE](LICENSE).
