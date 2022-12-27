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
	// var lenary []string
	
	// col_cnt := 0

	for {
		// 行毎のデータ取得してlineに格納
		line, err = reader.Read()
		
		
		if err != nil {
			break
		}
		
		
		// if col_cnt == 0 {
		// 	col_cnt = len(line)
			
		// 	for i := 1; i < len(line) + 1; i++ {
		// 		var in_str int = i
		// 		s := strconv.Itoa(in_str)
		// 		lenary = append(lenary, s)
		// 		// lenary = append(lenary, i)
		// 	}
		// }


		// 二次元配列に追加
		Csv_Result_2nd_ary = append(Csv_Result_2nd_ary, line)
	}
	
	ctx.HTML(200, "index.html", gin.H{
		"file_name": header.Filename,
		"file_ary":  Csv_Result_2nd_ary,
		// "col_cnt_ary":  lenary,
	})
}
