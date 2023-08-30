package string1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// 拼接字符串效率最高
// string.Builder

func Test_feild(t *testing.T) {

	s1 := "a b c d e"
	s2 := `a b 
	
	c 
	
	
	d e`

	fs := []string{"a", "b", "c", "d", "e"}

	require.Equal(t, fs, strings.Fields(s1))
	require.Equal(t, fs, strings.Split(s1, " "))
	//
	require.NotEqual(t, fs, strings.Split(s2, " "))
	require.Equal(t, fs, strings.Fields(s2)) // asciiSpace
}
