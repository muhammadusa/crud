drop table if exists categories cascade;
drop table if exists products cascade;

create table categories(
	category_id serial not null primary key,
	name character varying(32)
);

create table products(
	product_id serial not null primary key,
	name character varying(64) not null,
	price decimal (16, 2) not null,
	category_id int not null references categories(category_id)
);

insert into categories(name) values('Foods'), ('Drinks');

insert into products(name, price, category_id) values('Gamburger', 20000, 1), ('Coca-Cola', 10000, 2);
