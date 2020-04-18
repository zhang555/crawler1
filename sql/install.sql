create table wiki
(
    id          varchar(200)
        primary key,
    content     text                                                              not null,
    status      varchar(10) default ''                                            not null,
    create_time datetime    default CURRENT_TIMESTAMP                             not null,
    update_time datetime    default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null
);

