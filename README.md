## .env の作成

---

- pgadmin4
  env.example を複製して.env を作成し、中身の値を設定する
- postgresql
  env.example を複製して.env を作成し、中身の値を設定する

## docker image をビルドする

---

- ターミナルで下記コマンドを実行

```
make init
make up
```

下記にアクセスしてレスポンスが返ってくれば OK
http://localhost:8080/golang

## pgadmin に接続して DB サーバーを接続する

---

- 下記 URL にブラウザからアクセス
  http://localhost:5433

- .env で設定したメールとパスワードでログイン
- 新しいサーバを追加の接続タブのホスト名/アドレスとユーザー名・パスワードを.env で設定したもので埋めてから保存
