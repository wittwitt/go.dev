package imgexif

import (
	"fmt"
	"os"

	exifv3 "github.com/dsoprea/go-exif/v3"
	exifv3common "github.com/dsoprea/go-exif/v3/common"
)

func LastReadCopyrightByExif(fPath string, newFPath string) (string, error) {
	opt := exifv3.ScanOptions{}
	dt, err := exifv3.SearchFileAndExtractExif(fPath)
	if err != nil {
		return "", err
	}

	if newFPath != "" {
		dtLen := len(dt)
		data, err := os.ReadFile(fPath)
		if err != nil {
			return "", err
		}
		if err := os.WriteFile(newFPath, data[0:len(data)-dtLen], 0644); err != nil {
			return "", err
		}
	}

	ets, _, err := exifv3.GetFlatExifData(dt, &opt)
	if err != nil {
		return "", err
	}
	for _, et := range ets {
		fmt.Println(et.TagId, et.TagName, et.TagTypeName, et.Value)

		if et.TagName == "Copyright" {
			return et.Value.(string), nil
		}
	}

	return "", nil
}

func AppedCopyrightByExif(fPath string, newFPath string, commentStr string) error {
	im, err := exifv3common.NewIfdMappingWithStandard()
	if err != nil {
		return err
	}
	ti := exifv3.NewTagIndex()
	ib := exifv3.NewIfdBuilder(im, ti, exifv3common.IfdStandardIfdIdentity, exifv3common.TestDefaultByteOrder)

	err = ib.AddStandardWithName("Copyright", commentStr)
	if err != nil {
		return err
	}

	ibe := exifv3.NewIfdByteEncoder()
	exifData, err := ibe.EncodeToExif(ib)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(fPath)
	if err != nil {
		return err
	}

	f, _ := os.OpenFile(newFPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if _, err = f.Write(data); err != nil {
		return err
	}
	if _, err = f.Write(exifData); err != nil {
		return err
	}
	return nil
}
