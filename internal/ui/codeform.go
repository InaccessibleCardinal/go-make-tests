package ui

import (
	"go-make-tests/internal/svc/chat"
	"log"

	"fyne.io/fyne/v2/widget"
)

type CodeFormDriver struct {
	queryTextArea *widget.Entry
	form          widget.Form
	chatService   chat.ChatServiceIface
}

func (cf *CodeFormDriver) handleSubmit() {
	log.Println("submitting", cf.queryTextArea.Text)
	msg, err := cf.chatService.InvokeUserQuery(cf.queryTextArea.Text)
	cf.queryTextArea.SetText("")
	cf.queryTextArea.Refresh()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(msg.Content)
}

func NewCodeFormDriver(chatService chat.ChatServiceIface) *CodeFormDriver {
	queryTextArea := widget.NewMultiLineEntry()
	queryTextArea.SetMinRowsVisible(20)
	queryTextArea.AcceptsTab()

	return &CodeFormDriver{chatService: chatService, queryTextArea: queryTextArea, form: widget.Form{}}
}

func CreateCodeForm() *widget.Form {
	cf := NewCodeFormDriver(chat.GetChatService())

	cf.form.Append("Enter your query", cf.queryTextArea)
	cf.form.OnSubmit = cf.handleSubmit
	return &cf.form
}
