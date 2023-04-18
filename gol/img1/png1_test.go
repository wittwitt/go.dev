package img1

import (
	"bufio"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPng(t *testing.T) {
	f, err := os.Open("../../testdata/imgs/googlelogo_color_272x92dp.png")
	require.NoError(t, err)
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

func TestAddInfo(t *testing.T) {

	commStr := uuid.New().String()
	newPngPath := "/tmp/go.dev.TestAddInfo.png"
	err := AddExtToPng("../../testdata/imgs/googlelogo_color_272x92dp.png", newPngPath, commStr)
	require.NoError(t, err)

	commStr2, err := ReadExtFromPng(newPngPath, "")
	require.NoError(t, err)
	t.Log(commStr2)
}

func AddExtToPng(pngPath string, savePath string, commentStr string) error {
	pngF, err := os.Open(pngPath)
	if err != nil {
		return err
	}
	defer pngF.Close()

	img, err := png.Decode(pngF)
	if err != nil {
		return err
	}

	newPngF, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer newPngF.Close()
	newPngWriter := bufio.NewWriter(newPngF)
	pngEncoder := png.Encoder{CompressionLevel: png.NoCompression}
	pngEncoder.Encode(newPngWriter, img)

	tEXtChunk := pngChunk{
		"tEXt",
		[]byte(fmt.Sprintf("Comment\x00%s", commentStr)),
	}
	newPngWriter.Write(tEXtChunk.Bytes())

	newPngWriter.Flush()
	return nil
}

func ReadExtFromPng(pngPath string, savePathWithoutExt string) (string, error) {
	pngF, err := os.Open(pngPath)
	if err != nil {
		return "", err
	}
	defer pngF.Close()

	img, err := png.Decode(pngF)
	if err != nil {
		return "", err
	}

	meta, ok := img.(interface {
		Metadata() map[string]string
	})
	if !ok {
		log.Fatal("PNG元数据不可用")
	}

	// 打印PNG元数据
	for k, v := range meta.Metadata() {
		if k == "tEXt" {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	return "", nil
}

type pngChunk struct {
	Name string
	Data []byte
}

func (c pngChunk) Bytes() []byte {
	length := uint32(len(c.Data))
	buf := make([]byte, 8+int(length))
	copy(buf[0:4], []byte(fmt.Sprintf("%04X", length)))
	copy(buf[4:8], []byte(c.Name))
	copy(buf[8:], c.Data)
	return buf
}
