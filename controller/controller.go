package controller

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

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
	
	ctx.HTML(200, "index.html", gin.H{
		"file_name": header.Filename,
		"file_ary":  Csv_Result_2nd_ary,
	})
}


func Create_Csv_Date(ctx *gin.Context) {	
	//  CSVの列数
	num, _ := strconv.Atoi(ctx.PostForm("Csv_Vals_cnt"))
	//  CSV書き込む値
	Csv_Vals_Data := strings.Split(ctx.PostForm("Csv_Vals"), ",")

	// CSVファイル名
	Csv_File_Name := "M:/社内共有/70暫定フォルダ/効率化/Csv_Tool/" + ctx.PostForm("csv_name")
		
	// 列数の初期化
	col := 1
	
	var line []string
	
	//  二次元配列に格納
	var Csv_Result_2nd_ary [][]string
	//  CSV作成
	file, err := os.Create(Csv_File_Name)
	
    if err != nil {
        log.Println(err)
    }
    defer file.Close()
	
	// CSV書き込み処理
	for _, s := range Csv_Vals_Data {
		
		line = append(line, s)
		if num == col {
			Csv_Result_2nd_ary = append(Csv_Result_2nd_ary, line)

			line = nil
			col = 1
		} else {
			col += 1
		}
	}
	
	//  CSVに書き込み
    writer := csv.NewWriter(file)
	//  すべt書き込み
    writer.WriteAll(Csv_Result_2nd_ary)          
    writer.Flush()    
	 
	ctx.HTML(200, "index.html", gin.H{
		"csv_file_name": Csv_File_Name,
	})
}