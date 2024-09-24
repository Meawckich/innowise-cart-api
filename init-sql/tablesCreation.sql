
CREATE TABLE IF NOT EXISTS carts (
    id BIGSERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS items(
    id BIGSERIAL PRIMARY KEY,
    product VARCHAR(50) NOT NULL,
    quantity INT NOt NULL,
    cart_id INT NOT NULL,
    CONSTRAINT fk_cart FOREIGN KEY (cart_id)
        REFERENCES carts(id) 
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

BEGIN;
-- Insert mock data into the Cart table
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;
INSERT INTO Carts DEFAULT VALUES ;

-- Insert mock data into the Cart_item table
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Shoes', 10, 1);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Shoes', 10, 1);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Shoes', 10, 1);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Shirt', 5, 1);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Pants', 3, 2);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Hat', 7, 2);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Socks', 20, 3);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Jacket', 2, 4);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Gloves', 4, 5);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Belt', 6, 6);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Scarf', 8, 7);
INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, 'Backpack', 1, 8);
COMMIT;