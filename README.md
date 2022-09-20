# webapp-template

## モチベーション

- 自分がよく使う構成をまとめている
- 自分が使いたい技術を使ってる構成

## 今後

- [ ] console 追加
- [ ] admin 追加
- [ ] TimescaleDB を選択可能にする
- [ ] Litestream 追加
- [ ] Meilisearch 追加
- [ ] devcontainer 化
- [ ] Fluent Bit 追加
- [ ] VictoriaMetrics 追加

## webapp

- SQL から sqlc で Go 生成して、あとは API として返すだけ
    - 方針として JSON はまるっと返す
    - Remix 側でフィルタリングする

```
go 1.19

require (
	github.com/go-playground/validator/v10 v10.11.1
	github.com/goccy/go-yaml v1.9.5
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/labstack/echo-contrib v0.13.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/rs/zerolog v1.28.0
	github.com/shiguredo/lumberjack/v3 v3.0.0
	golang.org/x/sync v0.0.0-20220907140024-f12130a52804
)
```

## webapp-console

TBD

- コンソール(ダッシュボード)画面
- Cloudflare Workers
- Remix

## webapp-admin

TBD

- 管理画面
- Cloudflare Zero Trust でアクセス制御
- Cloudflare Workers
- Remix