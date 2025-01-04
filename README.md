## Overview

簡易 DDD とクリーンアーキテクチャの兼ね合いを実験的に試すリポジトリ。

WEB フレームワークには gin、 ORM には試験的に sqlc を採用している。

簡易 DDD ＝エンティティと値オブジェクトではなく、モデルにする

## Develop

Docker DB の起動

```sh
docker compose up -d
```

WEB サーバの起動

```sh
air
```

sqlc の反映

```sh
sqlc generate
```

### Docs

- [WEB フレームワーク Gin](https://gin-gonic.com/)
- [sqlc](https://sqlc.dev/)
- [pgx](https://github.com/jackc/pgx)
