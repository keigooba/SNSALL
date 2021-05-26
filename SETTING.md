# golang

# 初期設定

1. Go module導入
2. git導入
3. Makefileの作成
4. ログファイルの作成
5. header,footerの切り分け
6. configファイルの作成
7. golang-lintの導入
8. shファイルの作成 command.goで自動実行
9. バイナリファイルにGitのバージョン埋め込み設定・最新バージョンかのチェック
13. サーバー停止時のシグナルの作成
14. githubによるバイナリファイルリリースの設定
15. goxによるMac・Linux・Windowsに対応したそれぞれのバイナリファイルの作成
16. ポートの変更のコマンド作成

# 注意事項

1. runtime.GOOS or //+build を用いて windows でも同様の動作環境で動くようにすること
2. 各種メトリクス取得 API を利用して Munin や Zabbix 等のエージェント経由でメモリや GC の状況をモニタリングすること
3. ログレベルを分けて出力を見たい場合、logrusでログ出力する
4. リリース更新時、upVersion・info.jsonの更新を行うこと
5. golangci-lintでコードを検査し、可読性を担保する
