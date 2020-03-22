# go-tinydate
A tiny date object in Go. Tinydate uses *4 bytes* of memory vs the *24 bytes* of a standard time.Time{}

A tinydate only has *day* precision. It has no knowledge of hours, minutes, seconds, or timezones.

## Why?

If you don't have resource constraints, then don't use tinydate! Use the standard [time](https://golang.org/pkg/time/) pacakge.

go-tinydate should only be used when you *know* you won't need more than just *date* precision, and when you are resoure constrained. Perhaps one of the following situations apply:

* Working in embedded systems where memory is a luxury
* Storing millions of dates in memory and looking to reduce on server costs
* Are sure you will never need more than date precision and just like being efficient

Example:

I needed to store many thousands of dates in memory in order to keep track of which IDs I had seen in the last *X* number of days. As such, I had a giant `map[int]time.Time`. Switching from time.Time to tinydate.TinyDate reduced my program's memory from an average ~1GB to ~200MB.

## API

Tinydate mirrors the time.Time API for the most part. The following are the only methods on the time.Time object that are *not* included on the tinydate.TinyDate object, because they make no sense without more than *day* precision

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

All formatting is done via the time.Time package's formatting capabilites, but anything besides date data will be zeroed out

## Transient Dependencies

None! Add it will stay that way, except of course for the standard library.
