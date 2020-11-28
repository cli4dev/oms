package common

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

type ICommonLib interface {
	BuildXlsxData(list [][]string, header []string, fileName string) (err error)
	AddData(data []string, row *xlsx.Row)
	ReadDataFromXlsx(fileName string) (data [][]string, err error)
	BuildXlsxDataTabs(lists [][][]string, header []string, fileName string) (err error)
}

type CommonLib struct {
}

func (c *CommonLib) BuildXlsxData(list [][]string, header []string, fileName string) (err error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		err = fmt.Errorf("添加sheet失败:err:%+v", err)
		return
	}

	row := sheet.AddRow()
	// 添加头
	c.AddData(header, row)
	// 添加内容
	for _, item := range list {
		row = sheet.AddRow()
		c.AddData(item, row)
	}

	err = file.Save(fileName)
	if err != nil {
		err = fmt.Errorf("保存xlsx文件失败:err:%+v", err)
	}
	return
}

func (c *CommonLib) AddData(data []string, row *xlsx.Row) {
	var cell *xlsx.Cell
	for _, item := range data {
		cell = row.AddCell()
		cell.Value = item
	}
}

func (c *CommonLib) ReadDataFromXlsx(fileName string) (data [][]string, err error) {
	_, err = os.Stat(fileName)
	if !(err == nil || os.IsExist(err)) {
		err = fmt.Errorf("文件:%s不存在", fileName)
		return
	}

	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		err = fmt.Errorf("打开文件:%s失败:err:%+v", fileName, err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			list := []string{}
			for _, cell := range row.Cells {
				str := cell.String()
				if str != "" {
					list = append(list, cell.String())
				} else {
					break
				}
			}
			if len(list) > 0 {
				data = append(data, list)
			} else {
				break
			}
		}
	}
	return
}

//BuildXlsxDataTabs 生成多个tab的xlsx文件
func (c *CommonLib) BuildXlsxDataTabs(lists [][][]string, header []string, fileName string) (err error) {
	file := xlsx.NewFile()
	for k, list := range lists {

		sheet, err := file.AddSheet(fmt.Sprintf("附录%d", k+1))
		if err != nil {
			err = fmt.Errorf("添加sheet失败:err:%+v", err)
			break
		}

		row := sheet.AddRow()
		// 添加头
		c.AddData(header, row)
		// 添加内容
		for _, item := range list {
			row = sheet.AddRow()
			c.AddData(item, row)
		}
	}
	if err != nil {
		return
	}

	err = file.Save(fileName)
	if err != nil {
		err = fmt.Errorf("保存xlsx文件失败:err:%+v", err)
	}
	return
}
