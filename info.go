package gopio

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
Functions in this file: Value, Direction
*/

// get the current value
func (p *Pin) Value() (VALUE, error) {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	v := string(b)
	if v == "0" {
		return LOW, nil
	} else if v == "1" {
		return HIGH, nil
	} else {
		return "", fmt.Errorf("UNDEFINED VALUE")
	}
}

// get the current direction
func (p *Pin) Direction() (DIR, error) {
	file := fmt.Sprintf("%s/gpio%d/direction", basedir, p.Num)
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	d := string(b)
	if d == "out" {
		return OUT, nil
	} else if d == "in" {
		return IN, nil
	} else {
		return "", fmt.Errorf("UNDEFINED DIRECTION")
	}

}
