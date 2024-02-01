# golang-service


CREATE TABLE bh.user (
  id SERIAL PRIMARY KEY, 
  username VARCHAR (50) NOT NULL, 
  email VARCHAR (255) NOT NULL, 
  age int,
  created_at TIMESTAMP default current_timestamp NOT NULL
);