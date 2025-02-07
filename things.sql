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
