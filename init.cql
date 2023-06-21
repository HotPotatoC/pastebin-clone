CREATE KEYSPACE pastebin
WITH replication = {
  'class': 'NetworkTopologyStrategy',
  'replication_factor': 2
};

CREATE TABLE IF NOT EXISTS pastebin.pastes (
  id UUID,
  user_id UUID,
  paste BLOB,
  short_link TEXT,
  hash TEXT,
  created_at TIMESTAMP,
  PRIMARY KEY (id, short_link, hash, created_at)
);

CREATE TABLE IF NOT EXISTS pastebin.users (
  id UUID,
  name TEXT,
  email TEXT,
  password TEXT,
  created_at TIMESTAMP,
  PRIMARY KEY (id, email, created_at)
);