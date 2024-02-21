package svc

import (
	"errors"
	"fmt"
	"go-make-tests/internal/colors"
	"go-make-tests/internal/svc/chat"
	"go-make-tests/internal/svc/files"
	"log"
	"strings"
)

type AskForTestConfig struct {
	Language  string
	Framework string
	CodeInput string
	OutFile   string
}

func (a *AskForTestConfig) Set(i int, value string) {
	if i > 3 {
		panic(errors.New("trying to set ask for test config out of range"))
	}
	if i == 0 {
		a.Language = value
	}
	if i == 1 {
		a.Framework = value
	}
	if i == 2 {
		a.CodeInput = value
	}
	if i == 3 {
		a.OutFile = value
	}
}

type TestGenIface interface {
	AskForTest(AskForTestConfig) (string, error)
}

type TestGen struct {
	chatService chat.ChatServiceIface
}

func NewTestGen(chatService chat.ChatServiceIface, fileService files.FileServiceIface) TestGenIface {
	return &TestGen{chatService: chatService}
}

func (tg *TestGen) AskForTest(conf AskForTestConfig) (string, error) {
	prompt := getPrompt(conf.Language, conf.Framework)

	tg.chatService.Prompt(prompt)
	log.Println(colors.Blue("Asking AI..."))
	result, err := tg.chatService.InvokeUserQuery(conf.CodeInput)

	if err != nil {
		return "", err
	}

	processed := processContent(result.Content, conf.Language)
	resultToStdOut(processed)
	return processed, nil
}

func getPrompt(language, framework string) string {
	return fmt.Sprintf(
		`You are an expert %s programmer.
		Given the provided code, please write a unit test for it in the %s testing framework.
		Try to achieve total code coverage.
		Only return the code.`, language, framework)
}

func processContent(content, language string) string {
	cleanedContent := strings.Replace(content, fmt.Sprintf("```%s", strings.ToLower(language)), "// test generated by openai\n", 1)
	return strings.Replace(cleanedContent, "```", "", 1)
}

func resultToStdOut(result string) {
	log.Println(colors.Green("Success"))
	log.Println(result)
}
