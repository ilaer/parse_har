package cmd

import "fyne.io/fyne/v2/data/binding"

type Har struct {
	Width          float32        //GUI width
	Height         float32        //GUI height
	TabSelectIndex int            `yaml:"tab_select_index"`
	DBLog          binding.String `json:"binding_log"` //滚动日志绑定的数据

	HarFilePath string `json:"HarFilePath"`
}

func (m *Har) Initialize() {
	m.DBLog = binding.NewString()
}
