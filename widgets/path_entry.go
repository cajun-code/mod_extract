package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type PathEntry struct {
	widget.BaseWidget
	Entry        *widget.Entry
	Btn          *widget.Button
	FolderSelect bool
	parent       *fyne.Window
}

func NewPathEntry(folderSelect bool, placeholder string, window *fyne.Window) *PathEntry {
	p := &PathEntry{
		Entry:        widget.NewEntry(),
		Btn:          widget.NewButton("...", func() {}),
		FolderSelect: folderSelect,
		parent:       window,
	}
	p.Btn.OnTapped = p.Tapped
	p.Entry.SetPlaceHolder(placeholder)
	return p
}

func (p *PathEntry) Tapped() {
	if p.FolderSelect {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, *p.parent)
				return
			}
			if uri == nil {
				return
			}
			p.Entry.SetText(uri.Path())
		}, *p.parent)
		d.Show()
	} else {
		d := dialog.NewFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, *p.parent)
				return
			}
			if uri == nil {
				return
			}
			p.Entry.SetText(uri.URI().Path())
		}, *p.parent)
		d.Show()
	}
}
func (p *PathEntry) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, p.Btn, p.Entry)
	return widget.NewSimpleRenderer(c)
}

func (p *PathEntry) SetPath(path string) {
	p.Entry.SetText(path)
}

func (p *PathEntry) GetPath() string {
	return p.Entry.Text
}
