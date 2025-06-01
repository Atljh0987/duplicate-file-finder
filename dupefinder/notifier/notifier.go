package notifier

import "fmt"

type Notifier interface {
	NotifyWarn(message string)
	NotifySuccess(message string)
	NotifyErrorMessage(message string, err error)
	NotifyError(err error)
}

type ConsoleNotifier struct{}

func (c ConsoleNotifier) NotifyWarn(message string) {
	fmt.Println("WARNING:", message)
}

func (c ConsoleNotifier) NotifySuccess(message string) {
	fmt.Println(message)
}

func (c ConsoleNotifier) NotifyErrorMessage(message string, err error) {
	fmt.Println("ERROR:", message, err)
}

func (c ConsoleNotifier) NotifyError(err error) {
	fmt.Println("ERROR:", err)
}
