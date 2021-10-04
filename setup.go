package gopio

/*
Functions in this file: New, Close
*/

import (
	"fmt"
	"io/ioutil"
)

func New(pnum int, dir DIR) (pin Pin, err error) {
	if pnum == 27 {
		pnum = 0
	} else {
		pnum = PIN_NUMS[pnum] // converts the physical pin number to BCM numbering
		if pnum == 0 {
			return Pin{}, fmt.Errorf("INVALID PIN NUMBER")
		}
	}

	file := fmt.Sprintf("%s/export", basedir)

	// export the pin
	err = ioutil.WriteFile(file, []byte(fmt.Sprintf("%d", pnum)), 0666)
	if err != nil {
		return Pin{}, err
	}

	// specify the direction
	file = fmt.Sprintf("%s/gpio%d/direction", basedir, pnum)
	err = ioutil.WriteFile(file, []byte(dir), 0666)
	if err != nil {
		return Pin{}, err
	}
	pin = Pin{Num: pnum, Dir: dir}
	return pin, err
}

func (p *Pin) Close() error {
	file := fmt.Sprintf("%s/unexport", basedir)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%d", p.Num)), 0666)
	if err != nil {
		return err
	}
	return nil
}
