![Logo](http://svg.wiersma.co.za/hamba/project?title=slices&tag=Slice%20helper%20functions)

[![Go Report Card](https://goreportcard.com/badge/github.com/hamba/slices)](https://goreportcard.com/report/github.com/hamba/slices)
[![Build Status](https://travis-ci.com/hamba/slices.svg?branch=master)](https://travis-ci.com/hamba/slices)
[![Coverage Status](https://coveralls.io/repos/github/hamba/slices/badge.svg?branch=master)](https://coveralls.io/github/hamba/slices?branch=master)
[![GoDoc](https://godoc.org/github.com/hamba/slices?status.svg)](https://godoc.org/github.com/hamba/slices)
[![GitHub release](https://img.shields.io/github/release/hamba/slices.svg)](https://github.com/hamba/slices/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hamba/slices/master/LICENSE)

Go slice helper functions

## Overview

Install with:

```shell
go get github.com/hamba/slices
```

## Usage

### Contains

Contains is a generic slice contains function.

Supports: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```go
slice := []string{"foo", "bar"}
v := "bar"

fmt.Println(slices.Contains(slice, v))
// Outputs: true
```

### GreaterOf

GreaterOf returns a greater function for the given slice type for slice sorting.

Supports: string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```go
slice := []string{"foo", "bar"}
sort.Slice(slice, slices.GreaterOf(slice))

fmt.Println(slice)
// Outputs: [foo bar]
```

### LesserOf

LesserOf returns a lesser function for the given slice type for slice sorting.

Supports: string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```go
slice := []string{"foo", "bar"}
sort.Slice(slice, slices.LesserOf(slice))

fmt.Println(slice)
// Outputs: [bar foo]
```

### Intersect

Intersect returns the intersection of 2 slices.

Supports: string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```go
slice := []string{"foo", "bar", "baz", "bat"}
other := []string{"bar", "baz", "test"}

fmt.Println(slices.Intersect(slice, other))
// Outputs: [bar baz]
```

### Except

Except returns the all elements in the first slice that are not in the second.

Supports: string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```go
slice := []string{"foo", "bar", "baz", "bat"}
other := []string{"bar", "baz", "test"}

fmt.Println(slices.Except(slice, other))
// Outputs: [foo bat]
```

## Benchmark

```
BenchmarkContains-8          	35621572	        30.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsNative-8    	50106157	        23.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkExcept-8            	 6070610	       200 ns/op	      96 B/op	       2 allocs/op
BenchmarkExceptNative-8      	 5933550	       193 ns/op	      96 B/op	       2 allocs/op
BenchmarkGreaterOf-8         	 6290626	       189 ns/op	      80 B/op	       3 allocs/op
BenchmarkGreaterOfNative-8   	 8201284	       149 ns/op	      64 B/op	       2 allocs/op
BenchmarkIntersect-8         	 6012298	       196 ns/op	      96 B/op	       2 allocs/op
BenchmarkIntersectNative-8   	 6305799	       198 ns/op	      96 B/op	       2 allocs/op
BenchmarkLesserOf-8          	 6449050	       189 ns/op	      80 B/op	       3 allocs/op
BenchmarkLesserOfNative-8    	 8077785	       149 ns/op	      64 B/op	       2 allocs/op
```

Always benchmark with your own workload. The result depends heavily on the data input.
