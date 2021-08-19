package gopio

/*
Functions in this file: SetDirOut, SetDirIn, ChangeValue, SetHigh, SetLow
*/

import (
	"fmt"
	"os"
)

func (p *Pin) SetDirOut() error {
	file := fmt.Sprintf("%s/gpio%d/direction", basedir, p.Num)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s", OUT))
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetDirIn() error {
	file := fmt.Sprintf("%s/gpio%d/direction", basedir, p.Num)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s", IN))
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
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s", val))
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetHigh() error {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s", HIGH))
	if err != nil {
		return err
	}
	return nil
}

func (p *Pin) SetLow() error {
	file := fmt.Sprintf("%s/gpio%d/value", basedir, p.Num)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%s", LOW))
	if err != nil {
		return err
	}
	return nil
}
