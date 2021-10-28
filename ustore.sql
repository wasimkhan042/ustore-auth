CREATE
DATABASE ustore;

use ustore;

CREATE TABLE user
(
    id            VARCHAR(50) NOT NULL,
    first_name    varchar(15) NOT NULL,
    middle_name   varchar(15) NOT NULL,
    last_name     varchar(15) NOT NULL,
    email         varchar(50) NOT NULL,
    username      varchar(25) NOT NULL,
    password      char(60)    NOT NULL,
    profile_image varchar(25)
);

CREATE TABLE subscription
(
    id         varchar(40)   NOT NULL,
    start_time datetime      NOT NULL,
    end_time   datetime      NOT NULL,
    subs_price decimal(8, 2) NOT NULL,
    status     tinyint(1) NOT NULL,
    user_id    char(40)      NOT NULL,
    item_name  varchar(25)   NOT NULL
);