# README

PostgreSQL のレプリケーションを Docker Compose 上で作成し、作成したレプリケーション環境に対して Go からアクセスを行った。

## PostgreSQL のレプリケーションについて

- Primary 1 台に Replica 2 台の構成で構築を行っている
- Replica 間のバランシングは haproxy で行っている。

> [!NOTE]
> Go の Bun ORMでも Replica を複数接続することが可能であるが今回間に Proxy を挟んだ。
> 理由としては、 Proxy を介して Replica にアクセスすることでアプリケーション側はどの Replica にクエリを発行するかを意識する必要がなくなり、インフラ側の変更に対してアプリケーション側のコードを変更しなくても良くなる。
