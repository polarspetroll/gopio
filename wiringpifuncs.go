package gopio

/*
#cgo LDFLAGS: -lwiringPi
#include <wiringPi.h>
*/
import "C"

func GopioSetUp() {
  C.wiringPiSetupPhys()
}

func PinMode(pnum int, d DIR) WiringPiPin {
  if d == OUT {
    C.pinMode(C.int(pnum), C.OUTPUT)
  }else {
    C.pinMode(C.int(pnum), C.INPUT)
  }
  return WiringPiPin{Num: pnum, Dir: d}
}


func (p *WiringPiPin) DigitalWrite(v VALUE) {
  if v == HIGH {
    C.digitalWrite(C.int(p.Num), C.HIGH)
  }else {
    C.digitalWrite(C.int(p.Num), C.LOW)
  }
}

func (p *WiringPiPin) DigitalRead() int {
  return int(C.digitalRead(C.int(p.Num)))
}
