CREATE TABLE IF NOT EXISTS alternatives(
   id VARCHAR (100) PRIMARY KEY,
   question_id VARCHAR (100) NOT NULL,
   value VARCHAR (225) NOT NULL,
   is_correct INT NOT NULL,
   created_at DATE NOT NULL,
   updated_at DATE NOT NULL,
   FOREIGN KEY(question_id) REFERENCES questions(id)
);