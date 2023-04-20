package time1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Local问题

// time.Now()，，local time

// time.Unix,, local time

// time.Parse(),, utc time

func Test_default(t *testing.T) {

	now := time.Now()
	t.Log("now unix", now.Unix(), now.UnixNano()) // from 1970-01-01 00:00:00 开始的second, 可以为负数

	time1, _ := time.Parse("20060102150405", "19700101000000")                       // str 为UTC, loc指向UTC
	time2, _ := time.ParseInLocation("20060102150405", "19700101080000", time.Local) // str为本地时间, loc指向本地
	require.Equal(t, time1.Unix(), time2.Unix())

	time3, _ := time.ParseInLocation("20060102150405", "20230506110910", time.Local)
	time4, _ := time.Parse("20060102150405", "20230506030910")
	require.Equal(t, time3, time4.Local())

	var vartime time.Time
	time5, _ := time.Parse("20060102150405", "00010101000000")
	require.Equal(t, time5, vartime)

	time00 := time.Unix(0, 0) // local
	require.Equal(t, time1, time00.UTC())
}

func Test_add(t *testing.T) {
	d1 := 22 * time.Second
	d2 := 10 * time.Millisecond
	d12 := d1 + d2
	require.Equal(t, d12, 22010*time.Millisecond)
}
