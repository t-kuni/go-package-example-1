# go-package-example-1

プライベートリポジトリにある自作パッケージを使うサンプル

## ローカルの設定

ローカルでプライベートなパッケージを使用する際に必要な設定

```
git config --global url."git@github.com:t-kuni/go-package-example-2.git".insteadOf "https://github.com/t-kuni/go-package-example-2"
export GOPRIVATE=github.com/t-kuni/go-package-example-2
```

上記の設定を入れるとライブラリ側のリポジトリについて、GoLandからpushできなくなる。  
Shiftを2回おして `Manage Remotes` で検索しリモートリポジトリの設定を `git@github.com:t-kuni/go-package-example-2.git` に変更する。

## ライブラリにセマンティクスバージョンを与える

```
go tag v1.0.0
```

## ライブラリを更新する方法

再度go getすればOK  
セマンティクスバージョニングのタグがついている場合は、最新のタグのバージョンを取得する

```
go get -u github.com/t-kuni/go-package-example-2
```

特定のバージョンや、ブランチの最新を取りたい場合は以下のように指定する

``````
go get -u github.com/t-kuni/go-package-example-2@master
go get -u github.com/t-kuni/go-package-example-2@v1.0.0
```

## 特定のパッケージに依存しているリポジトリを探す方法

https://github.com/search?q=path:go.mod%20user:t-kuni%20github.com/t-kuni/go-package-example-2&type=Code

## Renovateの設定画面

https://developer.mend.io/github/t-kuni