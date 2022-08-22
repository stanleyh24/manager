CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "email" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL
);

CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "users_roles" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "role_id" int NOT NULL
);

CREATE TABLE "services" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "price" int NOT NULL,
  "rate" varchar(255) NOT NULL
);

CREATE TABLE "router" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "ip" varchar(20) UNIQUE NOT NULL,
  "name" varchar(255) NOT NULL,
  "username" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL
);

CREATE TABLE "clients" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "payment_date" int NOT NULL,
  "service_id" int NOT NULL,
  "router_id" int NOT NULL
);

CREATE TABLE "invoices" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL,
  "last_name" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "plan" varchar(255) NOT NULL,
  "month" varchar(25) NOT NULL,
  "amount" int NOT NULL,
  "creation_date" timestamptz NOT NULL,
  "limit_date" timestamptz NOT NULL,
  "payment_date" timestamptz,
  "paid" boolean
);

ALTER TABLE "users_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "clients" ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");

ALTER TABLE "clients" ADD FOREIGN KEY ("router_id") REFERENCES "router" ("id");
