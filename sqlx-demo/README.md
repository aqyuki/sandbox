# README

This directory has the demo program of the [sqlx](https://github.com/jmoiron/sqlx)

## Schema

### User table

| title  | type | description     |
| :----- | :--- | :-------------- |
| `id`   | UUID | User identifier |
| `name` | Text | User name       |

### Note table

| title      | type | description                  |
| :--------- | :--- | :--------------------------- |
| `id`       | UUID | Note identifier              |
| `owner_id` | UUID | Identifier of the note owner |
| `title`    | Text | Note title                   |
| `body`     | Text | Note body                    |
