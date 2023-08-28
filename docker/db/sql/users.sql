CREATE TABLE users (
  id INT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO `users`
  (`id`, `name`, `email`, `created_at`)
VALUES
  (1, 'John Doe', 'example.com', NOW());
