package gopio

/*
#cgo LDFLAGS: -lwiringPi
#include <wiringPi.h>
#include "local_lib.h"
*/
import "C"

// initial function for wiringpi(call this function if you're going to use wiringpi)
func GopioSetUp() {
  C.wiringPiSetupPhys()
}

// define a wiringpi pin using this function
// pnum : physical pin number, d direction
func PinMode(pnum int, d DIR) WiringPiPin {
  if d == OUT {
    C.pinMode(C.int(pnum), C.OUTPUT)
  }else {
    C.pinMode(C.int(pnum), C.INPUT)
  }
  return WiringPiPin{Num: pnum, Dir: d}
}


// digital write accept two values HIGH and LOW
func (p *WiringPiPin) DigitalWrite(v VALUE) {
  if v == HIGH {
    C.digitalWrite(C.int(p.Num), C.HIGH)
  }else {
    C.digitalWrite(C.int(p.Num), C.LOW)
  }
}

// digitalread function returns the current value for a specific pin
func (p *WiringPiPin) DigitalRead() VALUE {
  v := int(C.digitalRead(C.int(p.Num)))
  if v == 1 {
    return HIGH
  }else {
    return LOW
  }
}

// this function takes a value(HIGH, LOW) and writes it to the pin on program interruption
func (p WiringPiPin) WriteOnInterrupt(v VALUE) {
  var val int
  if v == HIGH {
    val = 1
  }else {
    val = 0
  }
  C.close_pin(C.int(p.Num), C.int(val))
}
