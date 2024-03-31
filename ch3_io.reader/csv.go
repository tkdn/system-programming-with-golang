package ch3ioreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// wget --quiet https://www.post.japanpost.jp/zipcode/dl/utf/zip/utf_ken_all.zip && unzip utf_ken_all.zip
var csvData = `13101,"100  ","1000003","トウキョウト","チヨダク","ヒトツバシ（１チョウメ）","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101  ","1010003","トウキョウト","チヨダク","ヒトツバシ（２チョウメ）","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100  ","1000012","トウキョウト","チヨダク","ヒビヤコウエン","東京都","千代田区","日比谷公園",0,0,0,0,0,0
13101,"102  ","1020093","トウキョウト","チヨダク","ヒラカワチョウ","東京都","千代田区","平河町",0,0,1,0,0,0
13101,"102  ","1020071","トウキョウト","チヨダク","フジミ","東京都","千代田区","富士見",0,0,1,0,0,0`

func ReadCSV() {
	r := strings.NewReader(csvData)
	cr := csv.NewReader(r)
	for {
		l, err := cr.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(l[2], l[6:9])
	}
}
