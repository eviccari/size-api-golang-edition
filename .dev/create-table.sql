create schema if not exists products;

use products;

create table if not exists products.sizes (
    id varchar(100) not null, 
    description varchar(250) not null,
    short_description varchar(20) not null, 
    size_range varchar(20) not null, 
    primary key(id)
);

CREATE USER 'size_user_service'@'%' IDENTIFIED BY '123';
GRANT ALL on products.* to size_user_service;  