package files

import (
	"os"
	"syscall"
)

//CloseOnExec makes sure closing the file on process forking.
func CloseOneExec(file *os.File) {
	if file != nil {
		syscall.CloseOnExec(syscall.Handle(int(file.Fd())))
	}
}
