package imgexif

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestXxx3(t *testing.T) {

	fPath := "../../testdata/imgs/duck.gif"

	newFPath := "/tmp/go.dev.Test.newf.gif"
	f2Path := "/tmp/go.dev.Test.fpath2.gif"

	ssid := uuid.New().String()

	err := AppedCopyrightByExif(fPath, newFPath, ssid)
	require.NoError(t, err)

	ssid2, err := LastReadCopyrightByExif(newFPath, f2Path)
	require.NoError(t, err)
	require.Equal(t, ssid, ssid2)

	f1md5 := ""
	f2md5 := ""
	{
		data2, err := os.ReadFile(fPath)
		require.NoError(t, err)
		hash := md5.New()
		hash.Write(data2[:])
		md5sum := hash.Sum(nil)
		f1md5 = hex.EncodeToString(md5sum)
	}
	{
		data2, err := os.ReadFile(newFPath)
		require.NoError(t, err)
		hash := md5.New()
		hash.Write(data2[:])
		md5sum := hash.Sum(nil)
		t.Log(newFPath, hex.EncodeToString(md5sum))
	}
	{
		data2, err := os.ReadFile(f2Path)
		require.NoError(t, err)
		hash := md5.New()
		hash.Write(data2[:])
		md5sum := hash.Sum(nil)
		f2md5 = hex.EncodeToString(md5sum)
	}

	t.Log(fPath, f1md5)
	t.Log(f2Path, f2md5)
	require.Equal(t, f1md5, f2md5)
}
