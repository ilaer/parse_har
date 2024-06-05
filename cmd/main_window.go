package cmd

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (m *Har) MainWindow() {
	ap := app.New()

	nw := ap.NewWindow("la")

	tabs := container.NewAppTabs()
	tabs.Append(container.NewTabItemWithIcon("提取(拖曳har文件到本窗口)", theme.ComputerIcon(), m.HarContainer(nw)))

	tabs.SelectIndex(m.TabSelectIndex) //setting default tab
	nw.SetContent(tabs)

	nw.Resize(fyne.NewSize(m.Width, m.Height))
	nw.SetFixedSize(true)

	nw.ShowAndRun()
}
