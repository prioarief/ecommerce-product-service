CREATE TABLE product_views (
  id INT PRIMARY KEY AUTO_INCREMENT,
  product_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (product_id) REFERENCES products(id)
);
