# Project Ausgaben/ Haushaltsbuch App

DB:Has to be a multi tenant postrgesql dockerbased database where each client has his her own db instance 
https://dba.stackexchange.com/questions/1043/what-problems-will-i-get-creating-a-database-per-customer
Why is explained here 
postgresql is easy and I know how to use it more or less, dockerized so it can be run everywhere

Structs
    User wie jetzt
    Ausgaben id user wobei bei dieser archtekture nicht notwendig jeder hätte seine eigene db, Beschreibung, Höhe, id, Währung, Datum

Tables
    User
    user_id email firstname lastname address



Backend: A golang backend just like the server in backend server

Frontend: An angular Frontend with adaptability to small screens, flexbox etc

PSQL CMD
Connect to db
\c dbname

Questions to ask:
1. GORM PG OR Plain SQL and why?
2. How to migrate the db so set it up with all constraints not null serial etc.?
3. Convert Serial to int for ids?
4. CRUD Method for every struct?



## TODO
https://github.com/golang-migrate/migrate

SQL Code to Generate DB

CREATE TABLE IF NOT EXISTS users (
	user_id serial NOT NULL,
	email varchar NOT NULL,
	firstname varchar NOT NULL,
	lastname varchar NOT NULL,
	address varchar,
	PRIMARY KEY (user_Id)
);

CREATE TABLE IF NOT EXISTS transactions (
	tr_id serial NOT NULL,
	tr_userId int NOT NULL,
	tr_category varchar NOT NULL,
	tr_group varchar,
	tr_amount int,
	tr_date date,
	tr_currency varchar,
	PRIMARY KEY(tr_id),
	FOREIGN KEY(tr_userId) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS accounts (
	acc_id serial NOT NULL,
	acc_userId int NOT NULL,
	acc_balance int,
	acc_currency varchar,
	PRIMARY KEY(acc_id),
	FOREIGN KEY(acc_userId) REFERENCES users (user_id)
)

Drop Table accounts;
Drop Table transactions;
Drop Table users
