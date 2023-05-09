CREATE TABLE IF NOT EXISTS question_groups(
   id VARCHAR (100) PRIMARY KEY,
   name VARCHAR (100) NOT NULL,
   description VARCHAR (300) NOT NULL,
   image_url VARCHAR (300) NOT NULL,
   sort_number INTEGER NOT NULL,
   created_at DATE NOT NULL,
   updated_at DATE NOT NULL
);