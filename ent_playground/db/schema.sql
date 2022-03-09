CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `age` integer NOT NULL, `name` varchar(255) NOT NULL DEFAULT 'unknown');
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
-- Dbmate schema migrations
