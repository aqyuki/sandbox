# README

SQLを練習するために作成したコードが含まれます．

## Schema

**accounts**

| カラム名   | 型   | 制約             |
| :--------- | :--- | :--------------- |
| `id`       | int  | primary key      |
| `username` | text | unique, not null |

**notes**

| カラム名     | 型   | 制約        |
| :----------- | :--- | :---------- |
| `id`         | int  | primary key |
| `account_id` | int  | foreign key |
| `title`      | text | not null    |
| `body`       | text | not null    |

**tags**

| カラム名 | 型   | 制約        |
| :------- | :--- | :---------- |
| `id`     | int  | primary key |
| `title`  | text | not null    |

**note_tag**

| カラム名  | 型   | 制約        |
| :-------- | :--- | :---------- |
| `note_id` | int  | foreign key |
| `tag_id`  | int  | foreign key |
