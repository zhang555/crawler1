show databases ;

use crawler;

drop table wiki;


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

create table wiki_image
(
    id          int auto_increment
        primary key,

    page_type   varchar(10)  default ''                not null,
    image_url   varchar(200) default ''                not null,

    create_time datetime     default CURRENT_TIMESTAMP not null,
    update_time datetime     default CURRENT_TIMESTAMP
        on update CURRENT_TIMESTAMP                    not null,
    index page_type_index (page_type)
);

