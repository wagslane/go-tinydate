# go-tinydate
A tiny date object in Go. Tinydate uses **4 bytes** of memory vs the **24 bytes** of a standard time.Time{}

A tinydate only has *day* precision. It has no knowledge of hours, minutes, seconds, or timezones.

[![](https://godoc.org/github.com/wagslane/go-tinydate?status.svg)](https://godoc.org/github.com/wagslane/go-tinydate)

## ⚙️ Installation

```bash
go get github.com/wagslane/go-tinydate
```

## 🚀 Quick Start

```go
package main

import (
    tinydate "github.com/wagslane/go-tinydate"
)

func main(){
    td, err := tinydate.New(2020, 04, 3)
	if err != nil {
		fmt.Println(err.Error())
    }
    
    td = td.Add(time.Hour * 48)
    fmt.Println(td)
    // prints 2020-04-05
}
```

## Need Second Precision? Go-TinyTime 

I've had people ask why go-tinydate doesn't use unix time as the underlying data, rather than day-month-year. Day-month-year supports dates from year
`0` to year `65535`. Unix timestamps only supports dates from `1970` to `2106`. If you need more precision, take a look at [go-tinytime](https://github.com/wagslane/go-tinytime) which
uses `unix` time underneath and supports second precision.

## Why?

If you don't have resource constraints, then don't use tinydate! Use the standard [time](https://golang.org/pkg/time/) pacakge.

go-tinydate works well in the following situations:

* You are working in embedded systems where memory is a luxury
* You are storing many dates and smaller memory footprint is desired
* You are 101% sure you don't need more than date precision

Example:

I needed to store many thousands of dates in memory in order to keep track of which IDs I had seen in the last *X* number of days. As such, I had a giant `map[int]time.Time`. Switching from `map[int]time.Time` to `map[int]tinydate.TinyDate` reduced my program's memory from an average ~1GB to ~200MB.

## 💬 Contact

[![Twitter Follow](https://img.shields.io/twitter/follow/wagslane.svg?label=Follow%20Wagslane&style=social)](https://twitter.com/intent/follow?screen_name=wagslane)

Submit an issue (above in the issues tab)

## API

Godoc: https://godoc.org/github.com/wagslane/go-tinydate

Tinydate mirrors the [time.Time](https://golang.org/pkg/time/) API for the most part. The only methods that are *not* included are the ones that makes no sense with only day precision, or without timezones apart from UTC.

## Formatting 

All formatting is done via the time.Time package's formatting capabilites, but anything besides date data will be zeroed out for obvious reasons.

## Transient Dependencies

None! And it will stay that way, except of course for the standard library.

Note: Currently the testify package is used **only** for testing, but that dependency will be removed in coming updates.

## 👏 Contributing

I love help! Contribute by forking the repo and opening pull requests. Please ensure that your code passes the existing tests and linting, and write tests to test your changes if applicable.

All pull requests should be submitted to the "master" branch.

```bash
go test
```

```bash
go fmt
```
