CREATE TABLE "credentials" (
  "id" uuid PRIMARY KEY,
  "access_id" varchar,
  "secret_key" varchar
);

CREATE TABLE "profiles" (
  "id" uuid PRIMARY KEY,
  "description" varchar,
  "region" varchar,
  "cred_id" uuid,
  "username" varchar
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("cred_id") REFERENCES "credentials" ("id");

ALTER TABLE "profiles" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
