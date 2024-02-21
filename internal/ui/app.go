package ui

import (
	"go-make-tests/internal/svc"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	labels = []string{"Language:", "Testing Framework", "In File", "Out File"}
	reader = os.ReadFile
	writer = os.WriteFile
)

func createInput(label string, entry *widget.Entry) *widget.FormItem {
	return &widget.FormItem{Text: label, Widget: entry}
}

type FormState struct {
	FormError    error
	Inputs       []*widget.Entry
	Labels       []string
	loadingLabel *widget.Label
	testGenSvc   svc.TestGenIface
	fileService  svc.FileServiceIface
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
	testGenConfig := svc.AskForTestConfig{}
	for i, input := range fms.Inputs {
		log.Printf("value for field %s: %s\n", fms.Labels[i], input.Text)
		if fms.Labels[i] == "In File" {
			code := fms.ReadCodeFromFile(input.Text)
			testGenConfig.Set(i, code)
		} else {
			testGenConfig.Set(i, input.Text)
		}
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

func (fs *FormState) Clear() {
	for _, input := range fs.Inputs {
		input.SetText("")
	}
}

func NewFormState(inputs []*widget.Entry, labels []string) *FormState {
	chatService := svc.GetChatService()
	fileService := svc.NewFileService(reader, writer)
	testGenSvc := svc.NewTestGen(chatService, fileService)
	loadingLabel := widget.NewLabel("")
	return &FormState{
		Inputs:       inputs,
		Labels:       labels,
		loadingLabel: loadingLabel,
		testGenSvc:   testGenSvc,
		fileService:  fileService}
}

func CreateForm() (*widget.Form, *FormState) {
	var entries []*widget.Entry
	var inputItems []*widget.FormItem
	for _, label := range labels {
		entry := widget.NewEntry()
		entries = append(entries, entry)
		inputItems = append(inputItems, createInput(label, entry))
	}
	formState := NewFormState(entries, labels)
	submitFunc := func() {
		formState.Submit()
		formState.Clear()
	}

	return &widget.Form{Items: inputItems, OnSubmit: submitFunc}, formState
}

func MakeTabs() *container.AppTabs {
	form, state := CreateForm()
	formContainer := container.NewVBox(form, state.loadingLabel)
	testTab := container.NewTabItem("Make a Test", formContainer)
	sourceTab := container.NewTabItem("Make Some Source Code", widget.NewLabel("Under construction"))
	return container.NewAppTabs(testTab, sourceTab)
}

func Run() {
	a := app.New()
	win := a.NewWindow("Make Some Code")

	tabs := MakeTabs()

	win.SetContent(tabs)
	win.Resize(fyne.NewSize(800, 600))
	win.ShowAndRun()
}
