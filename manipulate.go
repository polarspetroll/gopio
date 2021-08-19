package gopio

/*
Functions in this file: SetDirOut, SetDirIn, ChangeValue, SetHigh, SetLow
*/

import (
	"fmt"
	"io/ioutil"
)

func (p *Pin) SetDirOut() error {
	file := fmt.Sprintf("%s/gpio%d/direction", basedir, p.Num)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%s", OUT)), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetDirIn() error {
	file := fmt.Sprintf("%s/gpio%d/direction", basedir, p.Num)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%s", IN)), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) ChangeValue(val VALUE) error {
	if val != HIGH && val != LOW {
		return fmt.Errorf("INVALID VALUE FOR ChangeValue")
	}
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%s", val)), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetHigh() error {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%s", HIGH)), 0666)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetLow() error {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	err := ioutil.WriteFile(file, []byte(fmt.Sprintf("%s", LOW)), 0666)
	if err != nil {
		return err
	}
	return nil
}
