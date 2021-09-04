package gopio

import (
	"fmt"
	"io/ioutil"
	"os"
)

func (p *Pin) OnValueChange(cpin Pin, callback func(g Pin)) error {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	var listenvalue = LOW
	if string(b) == "0\n" {
		listenvalue = HIGH
	}
	var v VALUE

	for {
		v, err = p.Value()
		if err != nil {
			panic(err)
		}
		if v == listenvalue {
			callback(cpin)
		}
	}
}


func (p *Pin) OnValueHigh(cpin Pin, callback func(g Pin)) error {
	var v VALUE
	var err error
	for {
		v, err = p.Value()
		if err != nil {
			panic(err)
		}
		if v == HIGH {
			callback(cpin)
		}
	}
}

func (p *Pin) OnValueLow(cpin Pin, callback func(g Pin)) error {
	var v VALUE
	var err error
	for {
		v, err = p.Value()
		if err != nil {
			panic(err)
		}
		if v == LOW {
			callback(cpin)
		}
	}
}
