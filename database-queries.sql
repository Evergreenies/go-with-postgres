create database stocksdb;

create table stocks (
  stocksid serial primary key,
  name text,
  price int,
  company text
);

