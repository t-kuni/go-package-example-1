# go-package-example-1

プライベートリポジトリにある自作パッケージを使うサンプル

## ローカルの設定

ローカルでプライベートなパッケージを使用する際に必要な設定

```
git config --global url."git@github.com:t-kuni/go-package-example-2.git".insteadOf "https://github.com/t-kuni/go-package-example-2"
export GOPRIVATE=github.com/t-kuni/go-package-example-2
```

## 特定のパッケージに依存しているリポジトリを探す方法

https://github.com/search?q=path:go.mod%20user:t-kuni%20github.com/t-kuni/go-package-example-2&type=Code

## Renovateの設定画面

https://developer.mend.io/github/t-kuni