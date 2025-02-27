CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  age INT
);

insert into users (name, age) values ('isaachx', 20);

select * from users;

drop table customers;
drop table contacts;

--Create table contacts relation to customers
--Each customer has zero or many contacts and each contact belongs to zero or one customer.
CREATE TABLE customers(
   customer_id SERIAL primary key,
   customer_name VARCHAR(255) NOT NULL
 
);

create table contacts(
	contact_id serial primary key,
	customer_id int,
	contact_name varchar(255) not null,
	constraint fk_customer
		foreign key(customer_id)
			references customers(customer_id)
			on delete set null 
)

--INSERT TABLE customers
insert into customers(customer_name)
VALUES('Blue Corp INC'),('Nice Hut');

--INSERT TABLE contacts
insert into contacts(customer_id,contact_name)
values(1,'Andrean ilyas yahya'),(2,'Hilman nautilus kobra'),(1,'Diemas to be here');

--SELECT data customers
select * from customers;

--select data contacts
select * from contacts;

--DELETe data from customers
delete from customers where customer_id=1

--Add a foreign key constraint to an existing table 
alter table contacts
add constraint fk_customer
foreign key (customer_id)
references customers(customer_id)
on delete cascade;

--rename column table books
alter table books
rename column author to author_id;

-- update type column books
alter table books
alter column author_id type int using author_id::int;


-- delete column author_id
alter table books
drop column author_id;

--add column author_id
alter table books
add column author_id int;

-- add foreign key and constraint to table books
alter table books
add constraint fk_author
foreign key (author_id)
references author(author_id)
on delete cascade;

select * from books;

create table author(
author_id serial primary key,
name varchar(255) not null
);


update books
set author_id=1
where id= 12;

-- Inner join select table reference from author_id
select * from books inner join authors on books.author_id = authors.author_id; 


SELECT * from books ORDER BY id ASC;