package main

import (
    "log"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Package struct {
    packageName   string
    desc          string
    installed     bool
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
        SetText("VPKG is a VoidLinux Package Manager. \n(q) to quit\n(s) to Search")


func main() {
    // Packages installed in this session.
    packagesList.SetSelectedFunc(
        func(index int, name string, second_name string, shortcut rune) {
            setConcatText(&packages[index])
        })

    flex.SetDirection(tview.FlexRow).
        AddItem(tview.NewFlex().
            AddItem(packageText, 0, 1, true).
            AddItem(packagesList, 0, 4, false), 0, 6, false).
        AddItem(text, 0, 1, false)

    flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Rune() == 113 {
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

    form.Clear(true)
    addPackageForm()
    pages.SwitchToPage("Search Package")


    if err := app.SetRoot(pages, true).EnableMouse(false).Run(); err != nil {
        panic(err)
    }
}


func addPackageList(pkgMap map[string]string) {
    // This would call to something like xlocate
    packagesList.Clear()
    count := 1
    for key, val := range pkgMap {
        packageObj := Package{}
        packageObj.packageName = key
        packageObj.desc = val
        packages = append(packages, packageObj)
        packagesList.AddItem(
            //"["+key + "] = " + val, " ", rune(43), nil).
            key, " ", rune(64+count), nil).
        SetBorder(true)
        count++
    }
}


func addPackageForm() *tview.Form {
    packageObj := Package{}

    form.AddInputField("Package Name: ", "", 20, nil, func(packageName string) {
        packageObj.packageName = packageName
    })

    form.AddButton("Search", func() {
        // Appending to a list to display for now.
        packages = append(packages, packageObj)
        pkgMap := SearchPkg(packageObj.packageName)
        addPackageList(pkgMap)
        pages.SwitchToPage("Menu")
    })

    return form
}


func setConcatText(packageObj *Package) {
    packageText.Clear()
    text := packageObj.packageName+ " " + packageObj.desc + "\n"
    packageText.SetText(text)
}
