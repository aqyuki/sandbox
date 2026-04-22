-- 1. アプリケーションが利用するユーザーの作成
CREATE USER writer WITH PASSWORD 'writer_password';

-- 2. データベースへのアクセス権限を付与
GRANT CONNECT ON DATABASE demo TO writer;

-- 3. スキーマへのアクセス権限を付与
GRANT USAGE ON SCHEMA public TO writer;

-- 4. テーブルへのアクセス権限を付与
GRANT SELECT, INSERT, UPDATE, DELETE, TRUNCATE ON ALL TABLES IN SCHEMA public TO writer;

-- 5. 存在するシーケンスへのアクセス権限を付与
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO writer;

-- 6. 今後作成されるテーブルにもアクセス権限を自動付与する
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE, TRUNCATE ON TABLES TO writer;

-- 7. 今後作成されるシーケンスにもアクセス権限を自動付与する
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO writer;
