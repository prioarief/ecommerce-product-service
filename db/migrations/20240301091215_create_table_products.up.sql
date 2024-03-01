CREATE TABLE products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  stock INT NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  minimum_order INT NOT NULL,
  description TEXT,
  category_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);