-- user table
create table users (
  id uuid primary key,
  name text not null unique
);

-- note table
create table notes (
  id uuid primary key,
  owner_id uuid not null,
  title text not null,
  body text not null,
  foreign key (owner_id) references users(id)
);
