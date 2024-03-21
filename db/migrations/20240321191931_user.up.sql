CREATE TABLE "users" (
  "username" varchar PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz DEFAULT '0001-01-01 00:00:00.00Z',
  "created_at" timestamp DEFAULT 'now()',
  "updated_at" timestamp DEFAULT 'now()'
);