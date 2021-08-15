create database if  not exists user default  charset utf8mb4;

create table if not exists user2(
    id bigint primary  key atuo_increment,
    name varchar(32) not null default '',
    sex  boolean not null default true,
    addr  text not null default '',
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
)engine=innodb default charset utf8mb4;