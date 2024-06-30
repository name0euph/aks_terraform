From golang:1.23rc1-alpine AS builder

# gitインストール
RUN apk add --no-cache git

# 依存関係のインストール
WORKDIR .
COPY go.mod go.sum /go/src/work/
RUN go mod download

# ホスト上のソースコードをbuilderにコピー
COPY . /go/src/work/

# ビルドして/bin/workにバイナリ出力
RUN go build -o /bin/work

# マルチステージビルド
FROM scratch

# builderからバイナリをコピー
COPY --from=builder /bin/work /bin/work

ENTRYPOINT ["/bin/work"]
CMD ["-help"]