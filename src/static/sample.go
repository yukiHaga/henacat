package main

import (
	"fmt"
)

func main() {
	// 都道府県名と県庁所在地の対応マップを作成
	prefectureCapitalMap := map[string]string{
		"北海道":  "札幌",
		"青森県":  "青森",
		"岩手県":  "盛岡",
		"宮城県":  "仙台",
		"秋田県":  "秋田",
		"山形県":  "山形",
		"福島県":  "福島",
		"茨城県":  "水戸",
		"栃木県":  "宇都宮",
		"群馬県":  "前橋",
		"埼玉県":  "さいたま",
		"千葉県":  "千葉",
		"東京都":  "東京",
		"神奈川県": "横浜",
		// 他の都道府県も追加できます
	}

	// 都道府県名を入力として受け取る
	fmt.Print("都道府県名を入力してください: ")
	var inputPrefecture string
	fmt.Scanln(&inputPrefecture)

	// マップから県庁所在地を取得
	capital, found := prefectureCapitalMap[inputPrefecture]
	if found {
		fmt.Printf("%sの県庁所在地は%sです。\n", inputPrefecture, capital)
	} else {
		fmt.Println("該当する都道府県が見つかりませんでした。")
	}
}
