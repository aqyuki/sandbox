-- 1. レプリカが利用するユーザーの作成
CREATE USER reader WITH PASSWORD 'reader_password';

-- 2. データベースへのアクセス権限を付与
GRANT CONNECT ON DATABASE demo TO reader;

-- 3. スキーマへのアクセス権限を付与
GRANT USAGE ON SCHEMA public TO reader;

-- 4. テーブルへのアクセス権限を付与
GRANT SELECT ON ALL TABLES IN SCHEMA public TO reader;

-- 5. 今後作成されるテーブルにもアクセス権限を自動付与する
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO reader;
