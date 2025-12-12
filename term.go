package main

const (
	TERM_RESET string = "\x1b[0;0m"

	TERM_BOLD       string = "\x1b[1m"
	TERM_UNDERLINE  string = "\x1b[4m"
	TERM_INVERSE    string = "\x1b[7m"
	TERM_BLINK_WARN string = "\x1b[5m" // warn means this might not work everywhere

	FORE_BLACK   string = "\x1b[30m"
	FORE_RED     string = "\x1b[31m"
	FORE_GREEN   string = "\x1b[32m"
	FORE_YELLOW  string = "\x1b[33m"
	FORE_BLUE    string = "\x1b[34m"
	FORE_MAGENTA string = "\x1b[35m"
	FORE_CYAN    string = "\x1b[36m"
	FORE_WHITE   string = "\x1b[37m"

	BACK_BLACK   string = "\x1b[40m"
	BACK_RED     string = "\x1b[41m"
	BACK_GREEN   string = "\x1b[42m"
	BACK_YELLOW  string = "\x1b[43m"
	BACK_BLUE    string = "\x1b[44m"
	BACK_MAGENTA string = "\x1b[45m"
	BACK_CYAN    string = "\x1b[46m"
	BACK_WHITE   string = "\x1b[47m"
)
