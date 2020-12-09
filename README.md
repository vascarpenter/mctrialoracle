# mctrialoracle

## 多施設共同研究用サーバーをgoで書いてみた の OCI 版

- webフレームワークには echo
- dbアクセスには godrorを使用
- ORMの採用を中止 (sqlboilerはそもそもoracleに対応なし、あとバージョンがおかしいため..)

ディレクトリ構造
```
├── routes
├── static
│   ├── css
│   ├── img
│   └── javascript
└── views
実行時にはstatic, viewsが必要
```

ビルド
```
go build -o mctrialgo
```

実行
- 環境変数の "OCISTRING" に 接続情報をセットしておくこと。
- 例： setenv OCISTRING admin/password@XXXXdb_tp
```
mctrialgo &
```

サーバが起動している状態で、ブラウザで `http://127.0.0.1:3000/`　へアクセスしてみてください
- サンプル病院1 : userid test pass test
- サンプル病院2 : userid test2 pass test2
- admin : user admin pass admin  （これは hospital_id=1である必要があります)
- logoutするとcookieは消えます
- 実行時にはstatic, viewsが必要で、この2つのディレクトリとdocker上のoracle linuxでビルドした実行ファイルを OCI Instanceにコピーすると動作します

更新
- きちんと試験開始日を入力すると、経過日を計算してくれる
  - おおよそ30,60,90日後などに特定の検査を入れたりする場合にわかりやすいか
- admin/adminでログインするとadmin consoleが表示
  - Rがインストールできないので計算は内部のみ
  - 非同期計算にしました
  - jqueryによる日付入力の簡便化
  
注意
- cookieのsecret keyがハードコードされているので変更を。

