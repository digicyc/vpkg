package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Package struct {
	packageName string
	desc        string
	installed   bool
}

var packages = make([]Package, 0)

var app = tview.NewApplication()
var pages = tview.NewPages()

var packageText = tview.NewTextView()
var form = tview.NewForm()
var packagesList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()

// Info Grid
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("VPKG is a VoidLinux Package Manager. \n(Ctrl+C) to quit\n(s) to Search")

func main() {
	// Packages installed in this session.
	packagesList.SetSelectedFunc(
		func(index int, name string, second_name string, shortcut rune) {
			setConcatText(&packages[index])
			//InstallPkg(name)
			// Show install Window.
		})

	flex.SetDirection(tview.FlexColumn).AddItem(tview.NewFlex().
		AddItem(packageText, 0, 1, true).AddItem(
		tview.NewFlex().AddItem(packagesList, 0, 4, false),
		1, 6, false).AddItem(text, 0, 4, true),
		0, 8, false)

	// Capture Events on the Flex view.
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//if event.Rune() == 113 { // Press q?
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		} else if event.Rune() == 115 {
			// Reset Search
			log.Output(1, "Reset Search.")
			form.Clear(true)
			addPackageForm()
			pages.SwitchToPage("Search Package")
		}

		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Search Package", form, true, false)
	pages.AddPage("PKG Select", packagesList, true, false)

	form.Clear(true)
	addPackageForm()
	pages.SwitchToPage("Search Package")

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addPackageList(pkgMap map[string]string) {
	packagesList.Clear()
	for key, val := range pkgMap {
		packageObj := Package{}
		packageObj.packageName = key
		packageObj.desc = val
		packages = append(packages, packageObj)
		packagesList.AddItem(key, " ", rune(43), nil).
			SetBorder(true)
	}
}

func addPackageForm() *tview.Form {
	packageObj := Package{}

	form.AddInputField("Package Name: ", "", 20, nil, func(packageName string) {
		packageObj.packageName = packageName
	})

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			// Appending to a list to display for now.
			packages = append(packages, packageObj)
			pkgMap := SearchPkg(packageObj.packageName)
			addPackageList(pkgMap)
			pages.SwitchToPage("PKG Select")
			return nil
		}
		return event
	})

	return form
}

func setConcatText(packageObj *Package) {
	packageText.Clear()
	text := packageObj.packageName + " " + packageObj.desc + "\n"
	packageText.SetText(text)
}
