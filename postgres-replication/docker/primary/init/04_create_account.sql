-- テスト用のテーブルを作成
CREATE TABLE accounts (
  id uuid PRIMARY KEY,
  username text NOT NULL UNIQUE,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX index_accounts_username ON accounts(username);
