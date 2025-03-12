# 一括設定ツール試作品のためのモックAPIサーバー

## 必要なもの
- docker
- docker compose
- Go : 1.22.2 <= version

## Usage

- .envを.env.exampleに沿って準備する
- シェルで以下を実行

```bash
# Postgresql ver 14のコンテナを立ち上げ
make up
# サーバー起動
make run &
# サーバー停止
make kill
```
