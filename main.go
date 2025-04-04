package main

import (
    //"fmt"
    //"log"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Package struct {
    packageName   string
    desc          string
    installed     bool
}

const fileName = "sqlite.db"

var packages = make([]Package, 0)

var pages = tview.NewPages()
var packageText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var packagesList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
    SetTextColor(tcell.ColorGreen).
    SetText("(a) to Add a new Package. \n(q) to quit")


func main() {
    packagesList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
        setConcatText(&packages[index])
    })

    flex.SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().
            AddItem(packagesList, 0, 1, true).
            AddItem(packageText, 0, 4, false), 0, 6, false).
        AddItem(text, 0, 1, false)

    flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Rune() == 113 {
            app.Stop()
        } else if event.Rune() == 97 {
            form.Clear(true)
            addPackageForm()
            pages.SwitchToPage("Search Package")
        }
        return event
    })

    pages.AddPage("Menu", flex, true, true)
    pages.AddPage("Search Package", form, true, false)

    if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
        panic(err)
    }
}

func addPackageList() {
    packagesList.Clear()
    for index, packageObj := range packages {
        packagesList.AddItem(packageObj.packageName+" "+packageObj.desc, " ", rune(49+index), nil)
    }
}


func addPackageForm() *tview.Form {
    packageObj := Package{}

    form.AddInputField("Package Name: ", "", 20, nil, func(packageName string) {
        packageObj.packageName = packageName
    })

    form.AddCheckbox("Installed", false, func(installed bool) {
        packageObj.installed = installed 
    })

    form.AddButton("Search", func() {
        packages = append(packages, packageObj)
        addPackageList()
        pages.SwitchToPage("Menu")
    })

    return form
}

func setConcatText(packageObj *Package) {
    packageText.Clear()
    text := packageObj.packageName+ " " + packageObj.desc + "\n"
    packageText.SetText(text)
}
