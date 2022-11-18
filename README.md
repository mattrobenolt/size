# size

[![](https://pkg.go.dev/badge/go.withmatt.com/size)](https://pkg.go.dev/go.withmatt.com/size)

`size` is a package for working with byte capacities similar to `time.Duration`.

## Usage

```go
// Returns a Capacity which is a uint64 representing
// 5 gigabytes in bytes.
5*size.Gigabyte

// Turn the string "5G" into a Capacity
// Useful for parsing from flags.
size.ParseCapacity("5G")
```

## Installation

```bash
$ go get go.withmatt.com/size
```

## Resources
* [Documentation](http://pkg.go.dev/go.withmatt.com/size)
