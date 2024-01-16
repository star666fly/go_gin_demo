
create table user
(
    id       int auto_increment
        primary key,
    username varchar(255) null,
    password varchar(255) null,
    salt     varchar(64)  null,
    constraint user_id_uindex
        unique (id)
);

create table role
(
    id   int auto_increment
        primary key,
    name varchar(64) not null,
    constraint role_id_uindex
        unique (id),
    constraint role_name_uindex
        unique (name)
);

create table user_role
(
    user_id int not null,
    role_id int not null
);

create index user_role_user_id_index
    on user_role (user_id);