--This table stores information about the customers.
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(54) NOT NULL,
  last_name VARCHAR(54) NOT NULL,
  email VARCHAR(54) NOT NULL,
  phone_number VARCHAR(54) NOT NULL,
  address VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

--This table stores information about the groceries available for purchase.
CREATE TABLE groceries (
  id SERIAL PRIMARY KEY,
  name VARCHAR(54) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  quantity INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

-- Define the ENUM type for order status
CREATE TYPE order_status AS ENUM ('PENDING', 'INPROGRESS', 'SHIPPED', 'DELIVERED', 'CANCELLED');

--This table stores information about the orders placed by customers.
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  customer_id INT REFERENCES customers(id),
  order_date TIMESTAMP NOT NULL DEFAULT NOW(),
  delivery_date TIMESTAMP NOT NULL DEFAULT NOW(),
  total_price FLOAT NOT NULL,
  status order_status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);


--This table stores information about the individual grocery items that are included in each order.
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL REFERENCES orders(id),
  grocery_id INT NOT NULL REFERENCES groceries(id),
  quantity INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP ,
  CONSTRAINT unique_order_item UNIQUE (order_id, grocery_id)
);
