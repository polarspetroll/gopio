package gopio

// main data types and variables are declared here //

/*
Types And Variables :
	Types : Pin, DIR, VALUE
	Variables : OUT, IN, HIGH, LOW, PIN_NUMS, basedir
*/

// the object that defines a pin
type Pin struct {
	Num int
	Dir DIR
}

type DIR string
type VALUE string

const (
	OUT DIR = "out" //
	IN  DIR = "in"  //    directions

	HIGH VALUE = "1" //
	LOW  VALUE = "0" //    output values

	basedir = "/sys/class/gpio/" // base directory of gpio controller files
)

// physical pin numbering map(general purpose IO only)
var PIN_NUMS = map[int]int{
	3:  2,
	5:  3,
	7:  4,
	8:  14,
	10: 15,
	11: 17,
	12: 18,
	13: 27,
	15: 22,
	16: 23,
	18: 24,
	19: 10,
	21: 9,
	22: 25,
	23: 11,
	24: 8,
	26: 7,
	27: 0,
	28: 1,
	29: 5,
	31: 6,
	32: 12,
	33: 13,
	35: 19,
	36: 16,
	37: 26,
	38: 20,
	40: 21,
}
