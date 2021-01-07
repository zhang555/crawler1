show databases;

use crawler;

# drop table wiki;
truncate wiki;


create table wiki
(
    id          varchar(200)
        primary key,
    content     text                                  not null,

    status      varchar(10) default ''                not null,
    page_type   varchar(10) default ''                not null,


    create_time datetime    default CURRENT_TIMESTAMP not null,
    update_time datetime    default CURRENT_TIMESTAMP
        on update CURRENT_TIMESTAMP                   not null,
    index page_type_index (page_type)
);

drop table wiki_image;

create table wiki_image
(
    id          int auto_increment
        primary key,

    page_type   varchar(10)  default ''                not null,
    image_url   varchar(200) default ''                not null,

    create_time datetime     default CURRENT_TIMESTAMP not null,
    update_time datetime     default CURRENT_TIMESTAMP
        on update CURRENT_TIMESTAMP                    not null,
    index page_type_index (page_type),
    unique index image_url_index (image_url)
);

alter table wiki_image
    add x int default 0 not null;

alter table wiki_image
    add y int default 0 not null;
