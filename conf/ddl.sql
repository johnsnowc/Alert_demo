create table test.alert
(
    id         bigint auto_increment
        primary key,
    rule_code  varchar(255) not null,
    room_id    bigint       not null,
    info       text         not null,
    check_time bigint       not null
);

create index index_room_id
    on test.alert (room_id);

create table test.indicator
(
    id          bigint unsigned auto_increment
        primary key,
    code        varchar(100)      not null,
    name        varchar(100)      not null,
    type        tinyint           not null,
    left_child  varchar(100)      null,
    right_child varchar(100)      null,
    op          varchar(8)        null,
    expr        varchar(255)      null,
    time_range  bigint            null,
    create_time bigint            not null,
    update_time bigint            not null,
    is_deleted  tinyint default 0 null,
    constraint code_UNIQUE
        unique (code),
    constraint expr_UNIQUE
        unique (expr),
    constraint name_UNIQUE
        unique (name)
);

create index index_code
    on test.indicator (code);

create table test.rule
(
    id          bigint unsigned auto_increment
        primary key,
    code        varchar(255)      not null,
    name        varchar(100)      not null,
    room_id     bigint            not null,
    expr        longtext          not null,
    create_time bigint            not null,
    update_time bigint            not null,
    is_deleted  tinyint default 0 not null,
    constraint name_UNIQUE
        unique (name)
);

create index code_index
    on test.rule (code);

create index room_id_index
    on test.rule (room_id);

create table test.task
(
    id          bigint unsigned auto_increment
        primary key,
    name        varchar(45)       not null,
    room_id     bigint            not null,
    rule_code   varchar(255)      not null,
    frequency   bigint            not null,
    last_time   bigint            null,
    last_status tinyint           null,
    create_time bigint            not null,
    update_time bigint            not null,
    is_deleted  tinyint default 0 not null
);

create table test.transaction
(
    trading_id  bigint  not null
        primary key,
    room_id     bigint  not null,
    amount      double  not null,
    user_id     bigint  not null,
    item_id     bigint  not null,
    create_time bigint  not null,
    update_time bigint  not null,
    is_deleted  tinyint not null
);

create index idx_room_id
    on test.transaction (room_id);

