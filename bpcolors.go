package bpcolors

// Regular Color
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Reset   = "\033[0m"
	Clear   = "\033[H\033[2J"
)

// Background
const (
	BGBlack   = "\033[40m"
	BGRed     = "\033[41m"
	BGGreen   = "\033[42m"
	BGYellow  = "\033[43m"
	BGBlue    = "\033[44m"
	BGMagenta = "\033[45m"
	BGCyan    = "\033[46m"
	BGWhite   = "\033[47m"
)

// Bold
const (
	BReset   = "\033[1;30m"
	BRed     = "\033[1;31m"
	BGreen   = "\033[1;32m"
	BYellow  = "\033[1;33m"
	BBlue    = "\033[1;34m"
	BMagenta = "\033[1;35m"
	BCyan    = "\033[1;36m"
	BWhite   = "\033[1;37m"
)

// ref - https://gist.github.com/Prakasaka/219fe5695beeb4d6311583e79933a009
