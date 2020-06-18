![Logo](http://svg.wiersma.co.za/hamba/project?title=slices&tag=Slice%20helper%20functions)

[![Go Report Card](https://goreportcard.com/badge/github.com/hamba/slices)](https://goreportcard.com/report/github.com/hamba/slices)
[![Build Status](https://travis-ci.com/hamba/slices.svg?branch=master)](https://travis-ci.com/hamba/slices)
[![Coverage Status](https://coveralls.io/repos/github/hamba/slices/badge.svg?branch=master)](https://coveralls.io/github/hamba/slices?branch=master)
[![GoDoc](https://godoc.org/github.com/hamba/slices?status.svg)](https://godoc.org/github.com/hamba/slices)
[![GitHub release](https://img.shields.io/github/release/hamba/slices.svg)](https://github.com/hamba/slices/releases)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hamba/slices/master/LICENSE)

A fast Go slices codec

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

## Benchmark

```
BenchmarkContains-8          	36659943	        30.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkContainsNative-8    	48539482	        24.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGreaterOf-8         	 6257299	       193 ns/op	      80 B/op	       3 allocs/op
BenchmarkGreaterOfNative-8   	 7963461	       149 ns/op	      64 B/op	       2 allocs/op
BenchmarkLesserOf-8          	 6122317	       192 ns/op	      80 B/op	       3 allocs/op
BenchmarkLesserOfNative-8    	 7947242	       152 ns/op	      64 B/op	       2 allocs/op
```

Always benchmark with your own workload. The result depends heavily on the data input.
