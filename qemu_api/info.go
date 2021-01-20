package qemu_api

import (
	"fmt"
	"bufio"
	"os"
)

func Info() {
	format := bufio.NewScanner(os.Stdin)

	fmt.Println(format)
}
