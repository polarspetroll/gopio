# Button listener



```go
package main

import (
        "time"

        "github.com/polarspetroll/gopio"
)

const (
    BUTTON = 19 //
                // Physical
    LED    = 8  //
)

func main() {
        gopio.GopioSetUp() //initial function for wiringpi

        btn := gopio.PinMode(BUTTON, gopio.IN)
        g8  := gopio.PinMode(LED, gopio.OUT)

        for {
            if btn.DigitalRead() == gopio.HIGH {
                g8.DigitalWrite(gopio.HIGH)
            }else{
                g8.DigitalWrite(gopio.LOW)
            }
            time.Sleep(20 * time.Millisecond) // 20 milliseconds is a reasonable delay. However, you can change this value depending on your usage.
        }
}

```

##### Wiring

<div align="left">
<img src="button.png" width="390"/>
</div>
