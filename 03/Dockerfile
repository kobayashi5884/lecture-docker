FROM golang:1.12.5-alpine
WORKDIR /go/src/app
# コンテナ起動前に実行しておいてほしいコマンドをRUNで指定します。
RUN apk add --no-cache git \
  && go get github.com/gorilla/mux
# 上記のコマンドを実行すると、/go/src/github.com/gorilla/mux/に
# ソースコードがダウンロードされます。
# 要するに、URLと/go/src/からの相対パスが一致します。
# apk addは、Linux Alpineのコマンドで、
# go getの実行に必要なgitというパッケージを、
# go getの実行前にインストールしています。