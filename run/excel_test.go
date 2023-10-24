package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ichaly/go-next/pkg/util"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func zhToUnicode(raw string) (string, error) {
	return strconv.Unquote(strings.Replace(strconv.Quote(raw), `\\u`, `\u`, -1))
}

func TestExcel(t *testing.T) {
	source, _ := os.Open("/Users/Chaly/Desktop/20231023测试输入数据5000条_副本.xlsx")
	excel, _ := excelize.OpenReader(source)
	sheets := excel.GetSheetList()
	for _, s := range sheets {
		rows, _ := excel.GetRows(s)
		for i, row := range rows {
			body, _ := util.MarshalJson(map[string]any{
				"userId": "1004544051", "inputText": row[0], "inputType": row[1],
			})
			res, err := resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).
				Post("http://c-commercialize-impl.zpidc.com/commercialize/userAiResume/1.0.0/useAiResume")
			if err != nil {
				t.Error(err)
			}
			fmt.Println(i, res.String())
		}
		t.Log("全部操作完成")
		break
	}
}

func TestDelayDelivery(t *testing.T) {
	source, _ := os.Open("/Users/Chaly/Desktop/代投需要延期订单（更新订单编号和uid）_副本2.xlsx")
	excel, _ := excelize.OpenReader(source)
	sheets := excel.GetSheetList()
	for _, s := range sheets {
		rows, _ := excel.GetRows(s)
		for i, row := range rows {
			if row[3] == "不存在" || i == 0 {
				continue
			}
			body, _ := util.MarshalJson(map[string]any{
				"uid": row[3], "oid": row[2], "days": 8,
			})
			res, err := resty.New().R().SetHeader("Content-Type", "application/json").SetBody(body).
				Post("http://c-commercialize-impl.zpidc.com/commercialize/resumeProxyDelivery/1.0.0/delayProxyDelivery")
			if err != nil {
				t.Log(err)
			}
			fmt.Printf("%d:%s %s %s %s %s\n", i, row[0], row[1], row[3], row[2], res.String())
			time.Sleep(1 * time.Second)
		}
	}
}

func flattenDetail(kind string, list []gjson.Result) string {
	data := map[string]string{}
	for _, item := range list {
		data[item.Get("type").String()] = item.Get("entity").String()
	}
	res := strings.Builder{}
	switch kind {
	case "1":
		if len(data["review_detail"]) > 0 {
			res.WriteString(data["review_detail"])
		}
	case "2":
		if len(data["work_company"]) > 0 {
			res.WriteString(data["work_company"])
		}
		if len(data["work_position"]) > 0 {
			res.WriteString("\n")
			res.WriteString(data["work_position"])
		}
		if len(data["work_time"]) > 0 {
			res.WriteString("\n")
			res.WriteString(data["work_time"])
		}
		if len(data["work_detail"]) > 0 {
			res.WriteString("\n")
			res.WriteString(data["work_detail"])
		}
	case "3":
		if len(data["project_name"]) > 0 {
			res.WriteString(data["project_name"])
		}
		if len(data["project_time"]) > 0 {
			res.WriteString("\n")
			res.WriteString(data["project_time"])
		}
		if len(data["project_desc"]) > 0 {
			res.WriteString("\n")
			res.WriteString(data["project_desc"])
		}
	}
	return res.String()
}

func flattenEntirety(kind string, list []gjson.Result) string {
	if kind == "4" {
		data := map[string]string{}
		for _, item := range list {
			chunkType := item.Get("chunk_type").String()
			switch chunkType {
			case "review":
				data["review"] = flattenDetail("1", item.Get("detail").Array())
			case "work":
				res := strings.Builder{}
				subList := item.Get("detail").Array()
				for _, sub := range subList {
					res.WriteString(flattenDetail("2", sub.Array()))
				}
				data["work"] = res.String()
			case "project":
				res := strings.Builder{}
				subList := item.Get("detail").Array()
				for _, sub := range subList {
					res.WriteString(flattenDetail("3", sub.Array()))
				}
				data["project"] = res.String()
			}
		}
		result := strings.Builder{}
		if len(data["review"]) > 0 {
			result.WriteString(data["review"])
		}
		if len(data["work"]) > 0 {
			result.WriteString("\n\n")
			result.WriteString(data["work"])
		}
		if len(data["project"]) > 0 {
			result.WriteString("\n\n")
			result.WriteString(data["project"])
		}
		return result.String()
	} else {
		result := strings.Builder{}
		for _, sub := range list {
			text := flattenDetail(kind, sub.Array())
			result.WriteString(text)
			result.WriteString("\n\n")
		}
		return result.String()
	}
}

func TestDataCleaning(t *testing.T) {
	source, _ := os.Open("/Users/Chaly/Desktop/20231023输出数据5000条.xlsx")
	excel, _ := excelize.OpenReader(source)
	sheets := excel.GetSheetList()

	target := excelize.NewFile()
	sheet, _ := target.NewSheet("Sheet1")
	count := 0
	for _, s := range sheets {
		rows, _ := excel.GetRows(s)
		for j, row := range rows {
			if j == 0 {
				continue
			}
			res, kind, input := "", row[2], row[3]
			if len(row) > 4 {
				text := row[4]
				code := gjson.Get(text, "code").Int()
				if code == 200 {
					list := gjson.Get(text, "data.choices.0.message.format_content").Array()
					res = flattenEntirety(kind, list)
				}
			}

			if len(res) == 0 {
				res = "很抱歉，您的问题可能含有小智无法理解的词汇，暂时无法回答，请换个问题，小智将尽力解答。"
			}

			_ = target.SetCellValue("Sheet1", fmt.Sprintf("A%d", count), input)
			_ = target.SetCellValue("Sheet1", fmt.Sprintf("B%d", count), res)
			_ = target.SetCellValue("Sheet1", fmt.Sprintf("C%d", count), kind)
			count++
			//t.Log(i, j, input, res, kind)
		}
	}
	target.SetActiveSheet(sheet)
	if err := target.SaveAs("/Users/Chaly/Desktop/export.xlsx"); err != nil {
		t.Error(err)
	}
	t.Log("全部操作完成")
}
