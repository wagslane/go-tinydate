# go-tinydate
A tiny date object in Go. Tinydate uses **4 bytes** of memory vs the **24 bytes** of a standard time.Time{}

A tinydate only has *day* precision. It has no knowledge of hours, minutes, seconds, or timezones.

## ⚙️ Installation

```bash
go get github.com/lane-c-wagner/go-tinydate
```

## 🚀 Quick Start

```golang
package main

import (
    tinydate "github.com/lane-c-wagner/go-tinydate"
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

Tinydate mirrors the [time.Time](https://golang.org/pkg/time/) API for the most part. The following are the only methods on the time.Time object that are *not* included on the tinydate.TinyDate object, because they make no sense without more than *day* precision

* [Clock()](https://golang.org/pkg/time/#Time.Clock)
* [Hour()](https://golang.org/pkg/time/#Time.Hour)
* [In(location)](https://golang.org/pkg/time/#Time.In)
* [Local()](https://golang.org/pkg/time/#Time.Local)
* [Location()](https://golang.org/pkg/time/#Time.Location)
* [Minute()](https://golang.org/pkg/time/#Time.Minute)
* [Nanosecond()](https://golang.org/pkg/time/#Time.Nanosecond)
* [Round()](https://golang.org/pkg/time/#Time.Round)
* [Second()](https://golang.org/pkg/time/#Time.Second)
* [Truncate()](https://golang.org/pkg/time/#Time.Truncate)
* [UTC()](https://golang.org/pkg/time/#Time.UTC)
* [Zone()](https://golang.org/pkg/time/#Time.Zone)

## Formatting 

All formatting is done via the time.Time package's formatting capabilites, but anything besides date data will be zeroed out for obvious reasons.

## Transient Dependencies

None! Add it will stay that way, except of course for the standard library.

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
