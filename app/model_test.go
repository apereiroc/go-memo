package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel_CreateInitialModelNoDatabase(t *testing.T) {
	m, err := NewModel(nil)
	if err != nil {
		t.Fail()
	}
	assert.True(t, len(m.groups) == 0)

	// for _, g := range m.groups {
	// 	assert.True(t, len(g.cmds) > 0)
	//
	// 	for _, c := range g.cmds {
	// 		assert.NotEqual(t, c.cmd, "")
	// 	}
	// }

	assert.Equal(t, m.selectedGroup, index(0))
	assert.Equal(t, m.selectedCmd, index(0))
	assert.Equal(t, m.view, noDatabaseView{})
	assert.Equal(t, m.keys, groupKeys)
	assert.Equal(t, m.quitWithCmd, false)
}
