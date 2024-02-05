begin;

create table bh.language (
    id serial primary key, 
    name varchar(255)
);

insert into bh.language (name) values ('Русский');

create table bh.user (
  id serial primary key, 
  username varchar (50) not null, 
  password varchar(255),
  email varchar (255) not null, 
  last_name varchar(255),
  first_name varchar(255),
  middle_name varchar(255),
  age int,
  language_id int default 1 not null,
  create_date timestamp default current_timestamp not null,
  update_date timestamp default current_timestamp not null,
  create_user varchar(255) default 'user' not null,
  update_user varchar(255) default 'user' not null
);

alter table bh.user add constraint language_id_fk foreign key ( language_id ) references bh.language ( id );

create table bh.book (
  id serial primary key,
  name text not null,
  description text not null,
  language_id int not null,
  create_date timestamp default current_timestamp not null,
  update_date timestamp default current_timestamp not null,
  create_user varchar(255) default 'user' not null,
  update_user varchar(255) default 'user' not null
);


create table bh.user_book (
  id serial primary key,
  user_id int,
  book_id int,
  is_deleted boolean default false not null,
  create_date timestamp default current_timestamp not null,
  update_date timestamp default current_timestamp not null,
  create_user varchar(255) default 'user' not null,
  update_user varchar(255) default 'user' not null
);

alter table bh.user_book add constraint user_id_fk foreign key ( user_id ) references bh.user ( id );
alter table bh.user_book add constraint book_id_fk foreign key ( book_id ) references bh.book ( id );

commit;