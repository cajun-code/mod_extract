package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"

	//"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/cajun-code/bg3_mod_extractor/widgets"
	//"github.com/recogni/pakr/pak"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("BG3 Mod Extractor")
	// Destination entry
	destinationEntry := widgets.NewPathEntry(true, "Destination", &w)
	destinationEntryContainer := container.New(layout.NewFormLayout(), widget.NewLabel("Destination"), destinationEntry)
	// Source entry
	modEntry := widgets.NewPathEntry(false, "Mod", &w)

	modEntryContainer := container.New(layout.NewFormLayout(), widget.NewLabel("Mod"), modEntry)
	// Extract button
	btnExtract := widget.NewButton("Extract", func() {
		extracted, err := extract(modEntry.GetPath(), destinationEntry.GetPath())
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Extracted", *extracted, w)
	})
	btnExtractContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnExtract, layout.NewSpacer())
	w.SetContent(container.NewGridWithColumns(1, destinationEntryContainer, modEntryContainer, btnExtractContainer))
	w.Resize(fyne.NewSize(600, 300))
	w.CenterOnScreen()
	w.ShowAndRun()
	os.Exit(0)
}

func extract(source_zip, destination string) (*string, error) {
	r, err := zip.OpenReader(source_zip)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	out_file := ""
	for _, f := range r.File {
		if strings.Contains(f.Name, ".pak") {
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			defer rc.Close()

			// Create destination file
			destination := destination + "/" + f.Name
			outFile, err := os.Create(destination)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			defer outFile.Close()

			// Copy the contents
			_, err = io.Copy(outFile, rc)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			out_file = destination
			break
		}
	}
	return &out_file, nil
}
