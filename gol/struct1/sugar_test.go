package struct1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type People struct {
	Name string
}

func (p *People) ShowA() string { return "ShowA_" + p.ShowB() }

func (p *People) ShowB() string { return "ShowB" }

func (p *People) ShowC() string { return "ShowC+p.Name" + p.Name }

type Teacher struct {
	*People
}

func (t *Teacher) ShowB() string { return "teacher.ShowB" }

func Test_sugar(t *testing.T) {
	teacher := Teacher{}

	require.Equal(t, true, teacher.People == nil)

	// t.Log(teacher.Name) // panic
	// t.Log(teacher.People.Name) // panic

	require.Equal(t, teacher.ShowA(), "ShowA_ShowB")        // ok
	require.Equal(t, teacher.People.ShowA(), "ShowA_ShowB") // ok

	// teacher.People.ShowC() // panic
}
