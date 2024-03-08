package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_makeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.makeUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {
		t.Error("failed: did not find expected value in preview")
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config

	testApp := test.NewApp()
	testWindow := testApp.NewWindow("Test GoMarkdown")

	edit, preview := testCfg.makeUI()

	testCfg.createMenuItems(testWindow)

	testWindow.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()

	test.Type(edit, "Some text")

	if preview.String() != "Some text" {
		t.Error("failed: did not find expected value in preview")
	}
}
