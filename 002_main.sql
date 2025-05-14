CREATE TABLE banks (
  id INTEGER PRIMARY KEY, -- This is the bank_id
  bank_name TEXT NOT NULL,
  bank_routing_number INTEGER NOT NULL
);

CREATE TABLE users_banks (
  user_id INTEGER,
  bank_id INTEGER,
  PRIMARY KEY (user_id, bank_id), -- Ensures no duplicate rows with the same user_id and bank_id
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (bank_id) REFERENCES banks(id)
);