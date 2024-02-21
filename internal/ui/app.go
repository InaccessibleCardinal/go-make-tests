package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func MakeTabs() *container.AppTabs {
	testForm, state := CreateTestForm()
	codeForm := CreateCodeForm()
	formContainer := container.NewVBox(testForm, state.loadingLabel)
	testTab := container.NewTabItem("Make a Test", formContainer)
	sourceTab := container.NewTabItem("Make Some Source Code", codeForm)
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
