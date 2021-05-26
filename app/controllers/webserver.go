package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"snsall/config"
	"text/template"

	"github.com/tcnksm/go-latest"
)

func StartWebServer() error {
	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.FlagPort
		log.Printf("Defauting to port %s", port)
	}

	log.Printf("Linstening on port %s", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	// 最新版バージョンチェック
	upVersion := "0.1.0"
	json := &latest.JSON{
		// JSONを返すURL
		URL: config.Config.URL + fmt.Sprint(config.FlagPort) + "/json",
	}
	res, _ := latest.Check(json, upVersion)
	if res.Outdated {
		fmt.Printf("%sは最新ではない、これはアップグレードすべき %s: %s\n", upVersion, res.Current, res.Meta.Message)
	}
	generateHTML(w, nil, "index")
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
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
