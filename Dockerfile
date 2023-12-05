# Dockerfile
FROM golang:1.17.0-bullseye
# Visual Studio Code Remote Development 拡張機能がデフォルトで開くディレクトリーであるため
WORKDIR /workspace
# 2021-06-10 時点で Visual Studio Code でコードのフォーマットや Lint などの支援を受けるのに必要なモジュール
# モジュール未インストール時に Visual Studio Code を開き、次のコマンドで一覧表示されたものです:
# Ctrl + Shift + P -> Go: Install/Update Tools
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/cweill/gotests/...
RUN go get github.com/fatih/gomodifytags
RUN go get github.com/josharian/impl
RUN go get github.com/haya14busa/goplay/cmd/goplay
RUN go get github.com/go-delve/delve/cmd/dlv
# div-dap のインストール方法は次のドキュメントを参考にしました:
# https://github.com/golang/vscode-go/blob/v0.26.0/docs/dlv-dap.md#updating-dlv-dap
RUN GOBIN=/tmp/ go get github.com/go-delve/delve/cmd/dlv@master \
 && mv /tmp/dlv $GOPATH/bin/dlv-dap
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go get golang.org/x/tools/gopls