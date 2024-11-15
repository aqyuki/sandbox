-- create accounts table
CREATE TABLE accounts (
  id int PRIMARY KEY,
  username text UNIQUE NOT NULL
);


-- create notes table
CREATE TABLE notes (
  id int PRIMARY KEY,
  account_id int REFERENCES accounts(id),
  title text NOT NULL,
  body text NOT NULL
);

-- create tags table
CREATE TABLE tags (
  id int PRIMARY KEY,
  title text NOT NULL
);

-- create notes_tags table
CREATE TABLE notes_tags (
  note_id int REFERENCES notes(id),
  tag_id int REFERENCES tags(id),
  PRIMARY KEY (note_id, tag_id)
);

-- insert mock data
INSERT INTO accounts (id, username) VALUES
  (1, 'foo'),
  (2, 'bar'),
  (3, 'baz');

INSERT INTO notes (id, account_id, title, body) VALUES
  (1, 1, 'note 1', 'body 1'),
  (2, 1, 'note 2', 'body 2'),
  (3, 2, 'note 3', 'body 3'),
  (4, 2, 'note 4', 'body 4'),
  (5, 3, 'note 5', 'body 5'),
  (6, 3, 'note 6', 'body 6'),
  (7, 3, 'note 7', 'body 7'),
  (8, 3, 'note 8', 'body 8');

INSERT INTO tags (id, title) VALUES
  (1, 'tag 1'),
  (2, 'tag 2'),
  (3, 'tag 3'),
  (4, 'tag 4'),
  (5, 'tag 5'),
  (6, 'tag 6'),
  (7, 'tag 7'),
  (8, 'tag 8');

INSERT INTO notes_tags (note_id, tag_id) VALUES
  (1, 1),  (1, 2),  (1, 3),  (2, 4),  (2, 5),
  (2, 6),  (3, 7),  (3, 8),  (4, 1),  (4, 2),
  (4, 3),  (5, 4),  (5, 5),  (5, 6),  (6, 7),
  (6, 8),  (7, 1),  (7, 2),  (7, 3),  (8, 4);
