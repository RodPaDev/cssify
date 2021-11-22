package main

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
)

func progressBar(str string, total int) *pb.ProgressBar {
	tmpl := fmt.Sprintf(`{{ cyan "%v" }} {{ bar . (white "")  "█" "▓" "░" "░" (white "")}} {{percent . }} {{rtime .}}`, str)
	bar := pb.ProgressBarTemplate(tmpl).Start(total)
	return bar
}

func prettyPrint(str string, newLine bool) {
	fmt.Print("\033[2K\r")
	if newLine {
		fmt.Println("\033[32m" + str)
		return
	}
	fmt.Print("\033[36m" + str)
}
