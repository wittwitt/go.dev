package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx/v3"
)

var monthFontColor = []string{
	"000000", "000000", "000000", "000000", "000000", "FFFFFF",
	"000000", "000000", "FFFFFF", "FFFFFF", "FFFFFF", "FFFFFF",
}

var monthColor = []string{
	"C1D1A4", "EECFB4", "98D1D4", "EEF7EF", "FEFBE0", "926EA1",
	"F8FBF0", "CDEBDC", "C88286", "B099BF", "8BD2DF", "54678B",
}

func main() {
	log.Printf("%v", ag())
}

func ag() error {
	year := 2022

	wb := xlsx.NewFile()

	yearSheetName := fmt.Sprintf("%d日历", year)
	var yearSheet *xlsx.Sheet
	if sh, ok := wb.Sheet[yearSheetName]; !ok {
		if newSheet, err := wb.AddSheet(yearSheetName); err != nil {
			return err
		} else {
			yearSheet = newSheet
		}
	} else {
		yearSheet = sh
	}

	yearTimeStart, err := time.ParseInLocation("20060102150405", fmt.Sprintf("%d0101000000", year), time.Local)
	if err != nil {
		return err
	}
	yearTimeEnd, err := time.ParseInLocation("20060102150405", fmt.Sprintf("%d0101000000", year+1), time.Local)
	if err != nil {
		return err
	}

	log.Printf("%v,%v", yearTimeStart.Format("2006-01-02"), yearTimeEnd.Format("2006-01-02"))

	for i := 1; i <= 7; i++ {
		newColumn := xlsx.NewColForRange(i, i)
		newColumn.SetWidth(30)
		yearSheet.SetColParameters(newColumn)
	}
	if err := writeWeekHeader(yearSheet); err != nil {
		return err
	}

	startNumber := 1
	rowNumber := 1
	for dayTime := yearTimeStart; dayTime.Before(yearTimeEnd); dayTime = dayTime.Add(24 * time.Hour) {

		weekNumber := int(dayTime.Weekday())
		if weekNumber == 0 {
			weekNumber = 7
		}
		weekNumber = weekNumber - 1

		weekYear, _ := dayTime.ISOWeek()
		if weekYear < year {
			rowNumber = startNumber
		} else {
			if weekNumber == 0 {
				rowNumber = rowNumber + 2
			}
		}

		monthNumber := int(dayTime.Month())

		dayCell, err := yearSheet.Cell(rowNumber, weekNumber)
		if err != nil {
			return err
		}
		writeDayCell(dayCell, dayTime, monthNumber)

		contentCell, err := yearSheet.Cell(rowNumber+1, weekNumber)
		if err != nil {
			return err
		}
		writeDayContent(contentCell)

		rowSheet, err := yearSheet.Row(rowNumber)
		if err != nil {
			return err
		}
		rowSheet.SetHeight(16)

		contentSheet, err := yearSheet.Row(rowNumber + 1)
		if err != nil {
			return err
		}
		contentSheet.SetHeight(10)

	}

	wb.Save("./2022.xlsx")

	return nil
}

func writeWeekHeader(yearSheet *xlsx.Sheet) error {
	sl := xlsx.NewStyle()
	sl.Alignment.Horizontal = "center"
	sl.Alignment.Vertical = "center"
	sl.Font.Size = 14
	sl.Font.Color = "FFFFFF"
	sl.Font.Bold = true
	sl.Fill.PatternType = "solid"
	sl.Fill.FgColor = "418DDB"

	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	border.LeftColor = "C4C4C4"
	border.RightColor = "C4C4C4"
	border.TopColor = "C4C4C4"
	border.BottomColor = "C4C4C4"
	sl.Border = border
	sl.ApplyBorder = true

	sl.ApplyAlignment = true
	sl.ApplyFill = true
	sl.ApplyFont = true

	z := []string{"一", "二", "三", "四", "五", "六", "日"}

	for i := 0; i < 7; i++ {
		c, err := yearSheet.Cell(0, i)
		if err != nil {
			return err
		}
		c.SetStyle(sl)
		c.SetValue(z[i])
	}

	rowSheet, err := yearSheet.Row(0)
	if err != nil {
		return err
	}
	rowSheet.SetHeight(24)
	return nil
}

func writeDayCell(cell *xlsx.Cell, dayTime time.Time, monthNumber int) error {
	sl := xlsx.NewStyle()
	sl.Alignment.Horizontal = "center"
	sl.Alignment.Vertical = "center"
	sl.Font.Size = 10
	sl.Font.Color = monthFontColor[monthNumber-1] // "FFFFFF"
	sl.Fill.PatternType = "solid"
	sl.Fill.FgColor = monthColor[monthNumber-1] // "FFFFFF00"
	// myStyle.Font.Bold = true

	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	border.LeftColor = "C4C4C4"
	border.RightColor = "C4C4C4"
	border.TopColor = "C4C4C4"
	border.BottomColor = "C4C4C4"
	sl.Border = border
	sl.ApplyBorder = true

	sl.ApplyAlignment = true
	sl.ApplyFill = true
	sl.ApplyFont = true

	cell.SetStyle(sl)

	if dayTime.Day() == 1 {
		cell.SetValue(fmt.Sprintf("(%s月)1", dayTime.Format("1")))
	} else {
		cell.SetValue(dayTime.Format("2"))
	}
	return nil
}

func writeDayContent(cell *xlsx.Cell) error {

	sl2 := xlsx.NewStyle()
	sl2.Alignment.Horizontal = "left"
	sl2.Alignment.Vertical = "center"
	sl2.Border = border()

	sl2.ApplyBorder = true
	sl2.ApplyAlignment = true
	cell.SetStyle(sl2)
	return nil
}

func border() xlsx.Border {
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	border.LeftColor = "C4C4C4"
	border.RightColor = "C4C4C4"
	border.TopColor = "C4C4C4"
	border.BottomColor = "C4C4C4"

	return border
}
