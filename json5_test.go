package json5

import (
	"testing"
)

const src = `
[
	{ key: 'EquipName', name: '设备名称', editable: true, }, // some comment
	{ key: 'PortNum', 'name': '端口数', "editable": true, },

	/*
	 * multiline comment
	 */
	{ key: 'CellHeight', name: '设备高度', editable: false, anotherField: 'anothor value' },
]`

func TestMap(t *testing.T) {
	var m interface{}
	if err := Unmarshal([]byte(src), &m); err != nil {
		t.Error(err)
	}
}

func TestStruct(t *testing.T) {
	var m []struct {
		Key      string `json:"key"`
		Name     string `json:"name"`
		Editable bool   `json:"editable"`
	}

	if err := Unmarshal([]byte(src), &m); err != nil {
		t.Error(err)
	}
}
