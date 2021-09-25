# Blink

```go
import (
   "log"
   "time"

   "github.com/polarspetroll/gopio"
)

func main() {
   g8, err := gopio.New(8, gopio.OUT)
   if err != nil {
     log.Fatal(err)
   }
   defer g8.Close()

   for i := 0; i < 10; i++ {
     err = g8.SetHigh()
     if err != nil {
        log.Fatal(err)
     }
     time.Sleep(1000 * time.Millisecond)
     err = g8.ChangeValue(gopio.LOW)
     if err != nil {
        log.Fatal(err)
     }
     time.Sleep(1000 * time.Millisecond)
   }
}
```

##### Wiring

<div align="left">
<img src="wiring.png" width="380"/>
</div>
