package main

import (
	"fyne.io/fyne/v2/data/binding"
	"ghar/cmd"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

func main() {

	m := cmd.Har{
		Width:          800,
		Height:         600,
		TabSelectIndex: 0,
		DBLog:          binding.NewString(),
	}

	m.Initialize()

	paths := findfont.List()
	for _, path := range paths {
		// msyh.ttc
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	m.MainWindow()

	os.Unsetenv("FYNE_FONT")

}
