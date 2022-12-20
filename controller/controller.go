package controller

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func IndexDisplayAction(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func Regit_Date(ctx *gin.Context) {

	// name := ctx.PostForm("name")
	file, header, err := ctx.Request.FormFile("CSV_FIle")

	// 文字コードのエンコード　ライブラリを入れる必要あり。
	// 初期時：「go get golang.org/x/text/encoding/japanese golang.org/x/text/transform」を実行
	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	// ダブルクオートを厳密にチェックしない
	reader.LazyQuotes = true

	// レコード読み取り (列数取得配列)
	var line []string
	//  二次元配列に格納
	var Csv_Result_2nd_ary [][]string

	for {
		// 行毎のデータ取得してlineに格納
		line, err = reader.Read()

		if err != nil {
			break
		}

		// 二次元配列に追加
		Csv_Result_2nd_ary = append(Csv_Result_2nd_ary, line)
	}

	// fmt.Printf("%v", array)
	// reader.LazyQuotes = true

	// ファイル名
	// ctx.HTML(200, "regit.html", gin.H{
	// 	"file_name": header.Filename,
	// 	"file_ary":  Csv_Result_2nd_ary,
	// })
	ctx.HTML(200, "index.html", gin.H{
		"file_name": header.Filename,
		"file_ary":  Csv_Result_2nd_ary,
	})
}
