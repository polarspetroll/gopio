package gopio

/*
Functions in this file: New, Close
*/

import (
	"fmt"
	"os"
)

func New(pnum int, dir DIR) (pin Pin, err error) {
	if pnum == 27 {
		pnum = 0
	} else {
		pnum = PIN_NUMS[pnum] // converts the physical pin number to the virtual numbering
		if pnum == 0 {
			return Pin{}, fmt.Errorf("INVALID PIN NUMBER")
		}
	}

	file := fmt.Sprintf("%s/export", basedir)
	var f *os.File

	// export the pin
	f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Pin{}, err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%v", pnum))
	if err != nil {
		return Pin{}, err
	}

	// specify the direction
	file = fmt.Sprintf("%s/gpio%d/direction", basedir, pnum)
	f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Pin{}, err
	}
	_, err = f.WriteString(fmt.Sprintf("%v", dir))
	if err != nil {
		return Pin{}, err
	}
	pin = Pin{Num: pnum, Dir: dir}
	return pin, err
}

func (p *Pin) Close() error {
	file := fmt.Sprintf("%s/unexport", basedir)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d", p.Num))
	if err != nil {
		return err
	}
	return nil
}
