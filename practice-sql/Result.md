# Result

試したSQLとその結果の記録

## 各アカウントの所持するノートを表示
```sql
select
  accounts.id, accounts.username, notes.id, as note_id, notes.title, notes.body
from
  accounts, notes
where
  accounts.id = notes.account_id;
```
| `id`  | `username` | `notes_id` | `title` | `body` |
| :---: | :--------- | :--------- | :-----: | :----: |
|   1   | foo        | 1          | note 1  | body 1 |
|   1   | foo        | 2          | note 2  | body 2 |
|   2   | bar        | 3          | note 3  | body 3 |
|   2   | bar        | 4          | note 4  | body 4 |
|   3   | baz        | 5          | note 5  | body 5 |
|   3   | baz        | 6          | note 6  | body 6 |
|   3   | baz        | 7          | note 7  | body 7 |
|   3   | baz        | 8          | note 8  | body 8 |

## tag 1がついているノートを抽出

```sql
select
  notes.*, tags.title as tag_title
from
  notes_tags, notes, tags
where
  notes_tags.note_id = notes.id and
  notes_tags.tag_id  = tags.id  and
  tags.id = 1;
```

| `id`  | `account_id` | `title` | `body` | `tag_title` |
| :---: | :----------: | :-----: | :----- | :---------- |
|   1   |      1       | note 1  | body 1 | tag 1       |
|   4   |      2       | note 4  | body 4 | tag 1       |
|   7   |      3       | note 7  | body 7 | tag 1       |
