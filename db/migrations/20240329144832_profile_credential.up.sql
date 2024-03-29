CREATE TABLE "credentials" (
  "id" uuid PRIMARY KEY,
  "access_id" varchar NOT NULL,
  "secret_key" varchar NOT NULL
);

CREATE TABLE "profiles" (
  "id" uuid PRIMARY KEY,
  "description" varchar NOT NULL,
  "region" varchar NOT NULL,
  "cred_id" uuid NOT NULL,
  "username" varchar NOT NULL
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("cred_id") REFERENCES "credentials" ("id");

ALTER TABLE "profiles" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
