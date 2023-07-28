package main

import (
	"encoding/json"
	"testing"

	"github.com/filecoin-project/go-bitfield"
	"github.com/stretchr/testify/require"
)

func Test_bitfeild(t *testing.T) {

	c := bitfield.New()
	c.Set(33)
	c.Set(21)
	c.Set(100)

	all, err := c.All(100)
	require.NoError(t, err)
	require.Equal(t, []uint64{21, 33, 100}, all) // index order

	// c.ForEach(func(u uint64) error {
	// 	fmt.Println(u)
	// 	return nil
	// })

	data, err := json.Marshal(c)
	require.NoError(t, err)
	require.Equal(t, "[21,1,11,1,66,1]", string(data))

	{
		c2 := bitfield.New()
		c2.Set(0)
		c2.Set(5)
		c2.Set(6)
		c2.Set(7)
		c2data, err := json.Marshal(c2)
		require.NoError(t, err)
		require.Equal(t, "[0,1,4,3]", string(c2data))
	}

	{
		c2 := bitfield.New()
		c2data, err := json.Marshal(c2)
		require.NoError(t, err)
		require.Equal(t, "[0]", string(c2data))
	}

	{
		c2 := bitfield.New()
		c2.Set(0)
		c2data, err := json.Marshal(c2)
		require.NoError(t, err)
		require.Equal(t, "[0,1]", string(c2data))
	}
}
