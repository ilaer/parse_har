package cmd

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"ghar/util"
	"golang.org/x/image/colornames"
	"image/color"
	"path/filepath"
	"strings"
)

func (m *Har) HarContainer(nw fyne.Window) fyne.CanvasObject {
	var fileLabel *widget.Label
	var DBLogMultiLineEntry *widget.Entry

	nw.SetOnDropped(func(position fyne.Position, uris []fyne.URI) {
		if len(uris) > 0 {
			ret := strings.Split(uris[0].String(), "//")
			if len(ret) > 1 {
				m.HarFilePath = fmt.Sprintf("%v", ret[1])
			} else {
				m.HarFilePath = fmt.Sprintf("%v", ret[0])
			}

			dbLog, _ := m.DBLog.Get()
			m.DBLog.Set(fmt.Sprintf("%v\n%s", m.HarFilePath, dbLog))

			_, filename := filepath.Split(m.HarFilePath)
			fileLabel.SetText(filename)
		}
	})

	fileLabel = widget.NewLabel("")
	fileLabel.Move(fyne.NewPos(m.Width*0.05, m.Height*0.02))
	fileLabel.Resize(fyne.NewSize(m.Width*0.6, m.Height*0.06))

	fileButton := widget.NewButtonWithIcon("选择har", theme.FileIcon(), func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				util.XWarning(fmt.Sprintf("opening file error :%v", err))
				return
			}
			if reader == nil || len(reader.URI().String()) < 5 {
				return
			}
			ret := strings.Split(reader.URI().String(), "//")
			if len(ret) > 1 {
				m.HarFilePath = fmt.Sprintf("%v", ret[1])
			} else {
				m.HarFilePath = fmt.Sprintf("%v", ret[0])
			}
			reader.Close()

			dbLog, _ := m.DBLog.Get()
			m.DBLog.Set(fmt.Sprintf("%v\n%s", m.HarFilePath, dbLog))

			_, filename := filepath.Split(m.HarFilePath)
			fileLabel.SetText(filename)

		}, nw)

		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".har"}))

		fileDialog.Show()
	})
	fileButton.Move(fyne.NewPos(m.Width*0.62, m.Height*0.02))
	fileButton.Resize(fyne.NewSize(m.Width*0.1, m.Height*0.06))

	ParseButton := widget.NewButtonWithIcon("解析har", theme.DocumentSaveIcon(), nil)
	ParseButton.Move(fyne.NewPos(m.Width*0.83, m.Height*0.02))
	ParseButton.Resize(fyne.NewSize(m.Width*0.1, m.Height*0.06))

	ParseButton.OnTapped = func() {
		jsPath, err := m.ParseHar(m.HarFilePath)
		dbLog, _ := m.DBLog.Get()
		if err != nil {
			m.DBLog.Set(fmt.Sprintf("%v 解析失败\n%s", m.HarFilePath, dbLog))
		} else {
			util.Command("cmd.exe", []string{"/c", "explorer.exe", jsPath})
			m.DBLog.Set(fmt.Sprintf("%v 解析成功\n%s", m.HarFilePath, dbLog))
		}
	}

	// 交互组件和字段绑定组件的分割线
	Line1 := canvas.NewLine(color.Color(colornames.Black))
	Line1.StrokeWidth = 1

	Line1.Position1 = fyne.NewPos(m.Width*0.01, m.Height*0.10)
	Line1.Position2 = fyne.NewPos(m.Width*0.98, m.Height*0.10)

	DBLogMultiLineEntry = widget.NewMultiLineEntry()

	DBLogMultiLineEntry.TextStyle.Bold = true
	//DBLogMultiLineEntry.SetText("1234")
	DBLogMultiLineEntry.TextStyle.Bold = true
	DBLogMultiLineEntry.Bind(m.DBLog)
	DBLogMultiLineEntry.Move(fyne.NewPos(m.Width*0.05, m.Height*0.13))
	DBLogMultiLineEntry.Resize(fyne.NewSize(m.Width*0.88, m.Height*0.65))

	return container.NewVBox(
		container.NewWithoutLayout(
			fileLabel,
			fileButton,
			ParseButton,
			Line1,
			DBLogMultiLineEntry,
		))
}
