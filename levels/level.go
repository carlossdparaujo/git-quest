package levels

type Level interface {
	Check() (result bool, output string)
}