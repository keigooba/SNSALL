package controllers

import (
	"fmt"
	"log"
	"net/http"
	"snsall/config"
	"text/template"

	"github.com/tcnksm/go-latest"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	// 最新版バージョンチェック
	upVersion := "0.1.0" //buildされるアプリケーションのバージョン
	json := &latest.JSON{
		// JSONを返すURL
		URL: config.Config.URL + fmt.Sprint(config.FlagPort) + "/json",
	}
	res, _ := latest.Check(json, upVersion)
	if res.Outdated {
		fmt.Printf("%s is not latest, you should upgrade to %s: %s\n", upVersion, res.Current, res.Meta.Message)
	}
	GenerateHTML(w, nil, "index")
}

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("view/%s.html", file))
	}
	// ヘッダー・フッターを追加
	files = append(files, "view/_header.html", "view/_footer.html")

	templates := template.Must(template.ParseFiles(files...))
	err := templates.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
