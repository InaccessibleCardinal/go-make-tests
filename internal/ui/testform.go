package ui

import (
	"go-make-tests/internal/svc"
	"go-make-tests/internal/svc/chat"
	"go-make-tests/internal/svc/files"
	"log"
	"os"

	"fyne.io/fyne/v2/widget"
)

var (
	labels = []string{"Language:", "Testing Framework", "In File", "Out File"}
	reader = os.ReadFile
	writer = os.WriteFile
)

func CreateInput(label string, entry *widget.Entry) *widget.FormItem {
	return &widget.FormItem{Text: label, Widget: entry}
}

type FormState struct {
	FormError      error
	languageEntry  *widget.Entry
	frameworkEntry *widget.Entry
	infileEntry    *widget.Entry
	outfileEntry   *widget.Entry

	Labels       []string
	loadingLabel *widget.Label
	testGenSvc   svc.TestGenIface
	fileService  files.FileServiceIface
}

func (fms *FormState) ReadCodeFromFile(codePath string) string {
	code, err := fms.fileService.ReadFile(codePath)
	if err != nil {
		log.Fatal(err)
	}
	return code
}

func (fms *FormState) SetLoading(value string) {
	fms.loadingLabel.Text = value
	fms.loadingLabel.Refresh()
}

func (fms *FormState) Submit() {
	fms.SetLoading("Loading...")
	testGenConfig := svc.AskForTestConfig{
		Language:  fms.languageEntry.Text,
		Framework: fms.frameworkEntry.Text,
		CodeInput: fms.infileEntry.Text,
		OutFile:   fms.outfileEntry.Text,
	}

	fms.fileService.SaveJson("sanity/sample-config.json", testGenConfig)

	testResult, err := fms.testGenSvc.AskForTest(testGenConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err := fms.fileService.SaveFile(testGenConfig.OutFile, testResult); err != nil {
		log.Fatal(err)
	}
	fms.SetLoading("")
}

func (fms *FormState) Clear() {
	fms.languageEntry.SetText("")
	fms.frameworkEntry.SetText("")
	fms.infileEntry.SetText("")
	fms.outfileEntry.SetText("")
}

func NewFormState() *FormState {
	chatService := chat.GetChatService()
	fileService := files.NewFileService(reader, writer)
	testGenSvc := svc.NewTestGen(chatService, fileService)
	loadingLabel := widget.NewLabel("")

	return &FormState{
		languageEntry:  widget.NewEntry(),
		frameworkEntry: widget.NewEntry(),
		infileEntry:    widget.NewEntry(),
		outfileEntry:   widget.NewEntry(),
		Labels:         labels,
		loadingLabel:   loadingLabel,
		testGenSvc:     testGenSvc,
		fileService:    fileService}
}

func CreateTestForm() (*widget.Form, *FormState) {
	formState := NewFormState()
	submitFunc := func() {
		formState.Submit()
		formState.Clear()
	}
	inputItems := []*widget.FormItem{
		CreateInput("Language", formState.languageEntry),
		CreateInput("Testing Framework", formState.frameworkEntry),
		CreateInput("In File", formState.infileEntry),
		CreateInput("Out File", formState.outfileEntry),
	}
	return &widget.Form{Items: inputItems, OnSubmit: submitFunc}, formState
}
