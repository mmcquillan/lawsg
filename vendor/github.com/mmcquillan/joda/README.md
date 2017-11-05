# Half-hearted Joda to Go Date Format Converter

The method of specifying dates in Go are based on a well-known date:

`Mon Jan 2 15:04:05 MST 2006`

Ref: https://golang.org/src/time/format.go

This library tries to use a more familiar method of specifying formats for dates:
```
Two digit year              y       06
Four digit year             Y       2006
Long month name             MMMM    January
Short month name            MMM     Jan
Two digit month             MM      01
Single digit month          M       1
Two digit day of month      dd      02
Single digit day of month   d       2
Long day of week            EE      Monday
Short day of week           E       Mon
Two digit 12 hours          hh      03
Single digit 12 hour        h       3
Two digit 24 hours          H       15
Two digit minutes           mm      04
Single digit minutes        m       4
Two digit seconds           ss      05
Single digit seconds        s       5
Milliseconds                S       .0
AM/PM                       a       pm
Time Zone                   z       MST
Time Zone offset            Z       -0700
```

Example:
```
package main

import(
  "fmt"
  "time"

  "github.com/mmcquillan/joda"
)

func main() {
  t := time.Now()
  f := joda.Format("Y-MM-dd H:mm:ssZ")
  fmt.Println(t.Format(f))
}
```
