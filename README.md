# golang-algorithm

Merge three integer collections into one slice sorted in ascending order, without calling any sort function.

- `collection1` is already sorted from max to min
- `collection2` and `collection3` are already sorted from min to max

```go
func Merge(collection1, collection2, collection3 []int) []int
```

The implementation walks `collection1` from the end (so it is read in ascending order) and performs a 3-way merge.

## Setup

1. Install [Go 1.22+](https://go.dev/dl/).
2. Clone this repository and enter the project directory:

```bash
git clone https://github.com/honhonnhon/golang-algorithm.git
cd golang-algorithm
```

This project uses only the Go standard library. There are no extra dependencies to install.

```bash
go mod download
```

## Run the demo

```bash
go run .
```

Example output:

```text
[0 0 1 2 3 4 5 6 7 8 9 10]
```

## Run unit tests

```bash
go test ./...
```

Verbose output:

```bash
go test -v ./...
```
