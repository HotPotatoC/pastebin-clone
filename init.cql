CREATE KEYSPACE pastebin
WITH replication = {
  'class': 'NetworkTopologyStrategy',
  'replication_factor': 2
};

CREATE TABLE IF NOT EXISTS pastebin.pastes (
  id UUID,
  user_id TEXT,
  paste TEXT,
  hash TEXT,
  created_at TIMESTAMP,
  PRIMARY KEY (id, hash, created_at)
);

CREATE TABLE IF NOT EXISTS pastebin.users (
  id UUID,
  name TEXT,
  email TEXT,
  password TEXT,
  created_at TIMESTAMP,
  PRIMARY KEY (id, email, created_at)
);