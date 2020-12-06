CREATE TYPE "lul" AS ENUM (
  'credit',
  'debit'
);

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(255),
  "email" varchar(100),
  "password" varchar(100) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "auth" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer NOT NULL,
  "auth_uuid" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_balance" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "balance" integer DEFAULT 0,
  "balance_achieve" integer DEFAULT 0,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "user_balance_history" (
  "id" SERIAL PRIMARY KEY,
  "user_balance_id" integer,
  "balance_before" integer,
  "balance_after" integer,
  "activity" varchar(100),
  "type" lul,
  "ip" varchar(100),
  "location" varchar(100),
  "user_agent" varchar(100),
  "author" varchar(100),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "blance_bank" (
  "id" SERIAL PRIMARY KEY,
  "balance" integer DEFAULT 0,
  "balance_achieve" integer DEFAULT 0,
  "code" varchar(100),
  "enable" boolean,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "blance_bank_history" (
  "id" SERIAL PRIMARY KEY,
  "blance_bank_id" integer,
  "balance_before" integer,
  "balance_after" integer,
  "activity" varchar(100),
  "type" lul,
  "ip" varchar(50),
  "location" varchar(100),
  "user_agent" varchar(100),
  "author" varchar(100),
  "created_at" timestamp DEFAULT (now())
);


ALTER TABLE "user_balance" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_balance_history" ADD FOREIGN KEY ("user_balance_id") REFERENCES "user_balance" ("id");

ALTER TABLE "blance_bank_history" ADD FOREIGN KEY ("blance_bank_id") REFERENCES "blance_bank" ("id");

CREATE UNIQUE INDEX ON "user" ("username");

CREATE UNIQUE INDEX ON "user" ("email");