package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"snsall/app/controllers"
	"snsall/config"
)

func main() {
	cmdFlag()
	log.Println(controllers.StartWebServer())
}

// Gitリポジトリのバージョン start.sh実行後バージョン自動更新
var versions = "1.0.0"

func cmdFlag() {
	flag.StringVar(&config.FlagPort, "port", config.Config.Port, "ポート設定が可能")
	flag.StringVar(&config.FlagPort, "p", config.Config.Port, "ポート設定が可能(short)")

	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "バージョン確認")
	flag.BoolVar(&showVersion, "v", false, "バージョン確認(short)")
	flag.Parse()

	if showVersion {
		fmt.Println("version", versions)
		os.Exit(1)
	}
}
