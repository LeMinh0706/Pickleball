-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "gender" int NOT NULL,
  "avt" varchar NOT NULL,
  "lat"  float NOT NULL,
  "lng" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");
CREATE INDEX ON "users" ("lat");
CREATE INDEX ON "users" ("lng");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd
