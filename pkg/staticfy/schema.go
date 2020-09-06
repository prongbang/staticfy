package staticfy

var PostgresQLSchema = `
CREATE SEQUENCE IF NOT EXISTS id_seq;

CREATE TABLE "public"."assets"
(
    "id"         numeric      NOT NULL DEFAULT nextval('id_seq'::regclass),
    "url"        text         NOT NULL,
    "user_id"    varchar(256) NOT NULL,
    "created_at" timestamp    NOT NULL DEFAULT now(),
    "updated_at" timestamp    NOT NULL DEFAULT now(),
    "file_path"  text         NOT NULL DEFAULT '':: text,
    PRIMARY KEY ("id")
);`

var MySQLSchema = `
CREATE TABLE assets
(
    id         bigint(20)   NOT NULL AUTO_INCREMENT,
    url        text         NOT NULL,
    user_id    varchar(256) NOT NULL DEFAULT '',
    file_path  text         NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;`
