## サービス内容

Bitflyer APIを使って、データを取得し、指標に応じて売買をする。

バックテスト版と本番を切り替えることができる。

指標は、今のところ5つ作っていて、バックテストの成績トップ3を使う。(買ったときの指標を使って売る)

## 環境

[ローカル環境]
* golang
* sqlite3

[構築方法]

golangのtarファイルをダウンロード&インストール

$GOPATH/src/projectにプロジェクトを作成(git cloneする)

`example.config.ini`を`config.ini`にコピーして、bitflyerのAPIキーを入力する。

$GOPATHはデフォルトで、~/goに作られる。作られない場合は、$bash_profileなどでパスと通す。

```cassandraql
cd project

go build ./

go run gotrading
```
実行し、Webサーバを起動する。

`localhost:8080/chart/`にアクセス

※起動すると、売買が開始するので注意。