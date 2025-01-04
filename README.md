## Overview

簡易 DDD とクリーンアーキテクチャの兼ね合いを実験的に試すリポジトリ。

WEB フレームワークには gin、 ORM には試験的に sqlc を採用している。

簡易 DDD ＝エンティティと値オブジェクトではなく、モデルにする

## Develop

### Taskfile 実行コマンドのインストール

```sh
brew install go-task/tap/go-task
# or
go install github.com/go-task/task/v3/cmd/task@latest
```

- [参考](https://taskfile.dev/installation)

### 必要なツールのインストール

```sh
task install-tools
```

### サーバ起動

```sh
task dev
```

### マイグレーション

```sh
# マイグレーションファイルの作成
task migrate-create

# 実行
task migrate-up

# ロールバック
task migrate-down
```

### Docs

- [WEB フレームワーク Gin](https://gin-gonic.com/)
- [sqlc](https://sqlc.dev/)
- [pgx](https://github.com/jackc/pgx)
