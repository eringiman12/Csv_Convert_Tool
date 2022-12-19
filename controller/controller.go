package controller

import (
	"encoding/csv"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func IndexDisplayAction(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func Regit_Date(ctx *gin.Context) {
	name := ctx.PostForm("name")
	file, _, err := ctx.Request.FormFile("CSV_FIle")
	// file, header, _ := ctx.Request.FormFile("CSV_FIle")

	// ファイル読み取り
	// 文字コードのエンコード
	// 初期時：「go get golang.org/x/text/encoding/japanese golang.org/x/text/transform」を実行
	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	// ダブルクオートを厳密にチェックしない
	reader.LazyQuotes = true

	// レコード読み取り (列数取得配列)
	var line []string
	// var b [][]string
	var array [][]string

	for {
		// 行毎のデータ取得してlineに格納
		line, err = reader.Read()

		if err != nil {
			break
		}

		// 二次元配列に追加
		array = append(array, line)
	}

	fmt.Printf("%v", array)
	// reader.LazyQuotes = true

	ctx.HTML(200, "regit.html", gin.H{
		"text":     name,
		"csv_name": file,
	})
}
