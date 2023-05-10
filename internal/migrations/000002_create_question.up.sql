CREATE TABLE IF NOT EXISTS questions(
   id VARCHAR (100) PRIMARY KEY,
   question_group_id VARCHAR (100) NOT NULL,
   title VARCHAR (225) NOT NULL,
   description TEXT,
   bible_text TEXT,
   image_url VARCHAR (300),
   sort_number INTEGER NOT NULL,
   created_at DATE NOT NULL,
   updated_at DATE NOT NULL,
   FOREIGN KEY(question_group_id) REFERENCES question_groups(id)
);