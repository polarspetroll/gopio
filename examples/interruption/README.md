# Interruption
Writes a specific value on interruption signals.

```go
package main

import (
   "time"

   "github.com/polarspetroll/gopio"
)

func main() {
   gopio.GopioSetUp() // initialize
   g40 := gopio.PinMode(40, gopio.OUT) // physical pin numbering

   go g40.WriteOnInterrupt(gopio.LOW) // run another go routine for listener

   for true {
        g40.DigitalWrite(gopio.HIGH) //
        time.Sleep(1 * time.Second)  // a simple blink
        g40.DigitalWrite(gopio.LOW)  //
        time.Sleep(1 * time.Second)  //
   }
}
```

In this example if you run the program led starts to blink and on any interruption signal from os, led will be turned off.
