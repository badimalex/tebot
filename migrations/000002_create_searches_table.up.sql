CREATE TABLE searches (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  track BOOLEAN NOT NULL,
  chat_id BIGINT NOT NULL,
  FOREIGN KEY (chat_id) REFERENCES chats(chat_id)
);
