// +build darwin,amd64 windows linux,!android
// +build go1.1

package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"snsall/utils"

	"github.com/jinzhu/gorm"
)

type ConfigList struct {
	Port    int    `json:"port"`
	LogFile string `json:"log_file"`
	View    string `json:"view"`
	URL     string `json:"url"`
}

// Config Configの定義
var Config ConfigList

var Db *gorm.DB

// ポート手動変更させるためここで定義
var FlagPort int

func init() {
	// 設定ファイルconfigの読み込み
	err := LoadConfig()
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	// ログファイルの設定
	utils.LoggingSettings(Config.LogFile)

	// コマンドの実行
	err = utils.Command()
	if err != nil {
		log.Println(err) //本番のroot権限ではコマンド実行できないため、出力のみ
	}

}

// LoadConfig Configの設定
func LoadConfig() error {

	f, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer f.Close()

	//Configにconfig.jsonを読み込む
	err = json.NewDecoder(f).Decode(&Config)
	if err != nil {
		return err
	}

	// 環境変数の値の判定
	format := "Port: %d\nLogFile: %s\nView: %s\nURL: %s\n"
	_, err = fmt.Printf(format, Config.Port, Config.LogFile, Config.View, Config.URL)
	if err != nil {
		return err
	}

	return nil //自明であればnilにする
}
