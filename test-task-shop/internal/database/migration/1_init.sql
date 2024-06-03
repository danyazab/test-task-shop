CREATE TABLE IF NOT EXISTS sellers (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     phone VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
      id SERIAL PRIMARY KEY,
      name VARCHAR(255) NOT NULL,
      description TEXT NOT NULL,
      price INT NOT NULL,
      seller_id INT NOT NULL,
      is_deleted BOOLEAN DEFAULT FALSE,
      FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
   id SERIAL PRIMARY KEY,
   product INT NOT NULL,
   quantity INT NOT NULL,
   price INT NOT NULL,
   FOREIGN KEY (product) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  seller_id INT NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers(id),
  FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

CREATE TABLE IF NOT EXISTS order_items (
     id SERIAL PRIMARY KEY,
     order_id INT NOT NULL,
     product INT NOT NULL,
     quantity INT NOT NULL,
     price INT NOT NULL,
     FOREIGN KEY (order_id) REFERENCES orders(id),
     FOREIGN KEY (product) REFERENCES products(id)
);


