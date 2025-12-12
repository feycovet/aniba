package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

// var asciiArt = `
//   ／l、
// （ﾟ､ ｡ ７
//   l  ~ヽ
//   じしf_,)ノ
// `;

var arrLines = strings.Split(asciiArt, "\n");
var asciiSize = 0; // col limit, set 0 for dynamic or manually set a cutoff
var padding = 15;  // 3 is recommended but it is personal preference
var textPadding = 10; // 10 is recommended here, this is the desc and content gap

type Config struct {
	accentCol string
	lines []Line
};

type Line struct {
	desc string
	prependDesc bool
	
	lnFormat string
	formatSet func() []any
}

var PADDING = strings.Repeat(" ", textPadding);

func main() {
	if textPadding % 2 != 0 { textPadding++; }
	var conf = CONFIG;
	var lines []string;
	for _, l := range conf.lines {
		var ln = "";
		if l.prependDesc {
			if utf8.RuneCountInString(l.desc) >= textPadding {
				l.desc = l.desc[0:textPadding-1]; // trim to adjust
			}
			if utf8.RuneCountInString(l.desc) == 2 { l.desc = " " + l.desc; }
			ln = conf.accentCol + strings.Repeat(" ", (textPadding - utf8.RuneCountInString(l.desc))/2) + l.desc + TERM_RESET + strings.Repeat(" ", ((textPadding - utf8.RuneCountInString(l.desc))/2)+1);
		}
		lines = append(lines, ln + strings.ReplaceAll(fmt.Sprintf(l.lnFormat, l.formatSet()...), "\n", "\n" + strings.Repeat(" ", padding + utf8.RuneCountInString(ln))));
	}

	var count = max(len(arrLines), len(lines));

	refineArt();
	printFetch(count, arrLines, lines, conf);
}

func printFetch(c int, asciiArr, lnArr []string, cf Config) {
	for i := range c {
		if i >= len(asciiArr) {
			fmt.Print(strings.Repeat(" ", asciiSize));
		} else {
			fmt.Print(cf.accentCol + asciiArr[i] + TERM_RESET);
		}

		if i < len(lnArr) {
			fmt.Print(lnArr[i]);
		}

		fmt.Print("\n");
	}
}

func refineArt() {
	if asciiSize != 0 {
		return;
	}
	if padding % 2 != 0 { padding++; }
	var longestLen = 0;
	for i := range arrLines {
		arrLines[i] = strings.TrimRight(arrLines[i], " ");

		if runewidth.StringWidth(arrLines[i]) > longestLen {
			longestLen = runewidth.StringWidth(arrLines[i]) + 1;
		}
		
	}
	asciiSize = longestLen + padding;

	for i := range arrLines {
		if runewidth.StringWidth(arrLines[i]) < longestLen {
			arrLines[i] += strings.Repeat(" ", (longestLen - runewidth.StringWidth(arrLines[i]) + (padding/2)));
		}
		arrLines[i] = strings.Repeat(" ", (padding/2)) + arrLines[i];
	}

	asciiArt = strings.Join(arrLines, TERM_RESET + "\n")
}
