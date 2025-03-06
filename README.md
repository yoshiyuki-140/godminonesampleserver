# 一括設定ツール試作品のためのモックAPIサーバー

## 必要なもの
- docker
- docker compose
- Go : 1.26.6 <= version

## Usage

- .envを.env.exampleに沿って準備する

```bash
# Postgresql ver 14のコンテナを立ち上げ
make up
# サーバー起動
make run &
# サーバー停止
make kill
```