package sandbox

import (
	"fmt"
	"io"
	"strings"

	// "golang.org/x/text/message"
)

func writeMessage(w io.Writer) {
	fmt.Fprintln(w, "Привет, интерфейсы!")
}

func readMessage(r io.Reader) {
	buf := make([]byte, 32)
	n, _ := r.Read(buf)
	fmt.Println("Прочитано:", string(buf[:n]))
}

func Run13() {
	var builder strings.Builder
	writeMessage(&builder)
	fmt.Println("Результат:", builder.String())
	reader := strings.NewReader("Текст из строки")
	readMessage(reader)
}

type Notifier interface {
	Notify(message string)
}

type EmailSender struct {
	Greeting string
}

func (e EmailSender) Notify(message string) {
	fmt.Println("Отправка email: ", message)
}

func Send(n Notifier) {
	n.Notify("Добро пожаловать!")
}
