package controllers

import (
	"flag"
	"fmt"
	"snsall/config"
)

// Gitリポジトリのバージョン start.shで最新のバージョンに更新される
var Version = "1.0.0"

func CmdFlag() {
	// ポート設定のオプション
	flag.StringVar(&config.FlagPort, "port", config.Config.Port, "ポート設定が可能")
	flag.StringVar(&config.FlagPort, "p", config.Config.Port, "ポート設定が可能(short)")

	// Gitリポジトリのバージョン確認
	var showVersion bool
	// -v -versionが指定された場合にshowVerionが真になるよう定義
	flag.BoolVar(&showVersion, "version", false, "バージョン確認")
	flag.BoolVar(&showVersion, "v", false, "バージョン確認(short)")
	flag.Parse() //引数からオプションをパースする

	//ポート確認
	fmt.Println("port", config.FlagPort)
	if showVersion {
		// バージョン番号を表示する
		fmt.Println("version", Version)
	}
}
