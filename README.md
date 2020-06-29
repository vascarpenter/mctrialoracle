# mctrialgo

## 多施設共同研究用サーバーをgoで書いてみた

node.jsで作っていましたが、
- コンパイラではない
- 型指定がいまいち
- ejsが助長
- ORMでない

などのいくつかの理由からやはり記述言語を変えるべきと思われました

候補にあがったのが swiftか goでしたが、ここはgoがdockerと親和性がよさそうだったので試してみました。

- webフレームワークには echo
- ORMには sqlboiler

go開発環境のインストール
```
brew install go
```

gomodを使ったので
```
go build
```
で必要なモジュールは自動的にダウンロードされるはず..

Sequel Proなどでデータベースを設計

sqlboiler.tomlファイルを作り
```
[mysql]
  dbname  = "studydb"
  host    = "localhost"
  port    = 3306
  user    = "oge"
  pass    = "hogehogeA00"
  sslmode = “false"
```

(mysqlのパスワードが@を含んでいるとsql URLに"@"が含まれていて問題だったのでやむなく変更しました)

データ構造ができたところで、`sqlboiler mysql`とすると
データベースの構造を読み込んでmodelsディレクトリ内に go ファイルを作ってくれる

現在の構造のデータ(+sample)を使うときの準備
```
mysqlにoge, hogehogeA00というユーザを準備する
make import  (SQLが読み込まれデータ構造とサンプルデータが入ります)
```


ディレクトリ構造
```
├── models
├── routes
├── static
│   ├── css
│   ├── img
│   └── javascript
└── views
```

実行
```
go run .
あるいは
make run
```

サーバが起動している状態で、ブラウザで `http://127.0.0.1:3000/`　へアクセスしてみてください
- サンプル病院1 : userid test pass test
- サンプル病院2 : userid test2 pass test2
- admin : user admin pass admin  （これは hospital_id=1である必要があります)
- logoutするとcookieは消えます

更新
- きちんと試験開始日を入力すると、経過日を計算してくれる
  - おおよそ30,60,90日後などに特定の検査を入れたりする場合にわかりやすいか
- admin/adminでログインするとadmin consoleが表示
  - [現在の日付でKM曲線を描く] をクリックすると、Rscriptを使ってデータベースから無理矢理Kaplan-Meier グラフを書きます
  - /usr/local/bin/Rscriptがあり、DBI/RMySQL/surviveがインストールされていることが前提
  - 詳しくはanalysis.Rを見て適宜変更を (現在は死亡イベントを解析)
  - 昔ダウンロードしてきた TatsukiRcodeKMplot.r を使用していますが、いまそこにアクセスできないし..
  - https://github.com/SteveLane/CodeLibrary 内にも TatsukiRcodeKMplot.r はありますね
  - 非同期計算にしました
  - Docker対応にした
    - 作成：docker-compose build
    - 起動：docker-compose up -d
    - 終了：docker-compose down
  - jqueryによる日付入力の簡便化
  
注意
- cookieのsecret keyがハードコードされているので変更を。
- VSCodeのlintはかなーり強力だが、html templateの中の typo までは指摘してくれず真っ白画面になる   html comment  {{/* */}} なんぞを使って切り分けしていくしかない go template lintは思ってたのとは違う

