DROP TABLE IF EXISTS Order_Book;

DROP TABLE IF EXISTS Depth_Order;

DROP TABLE IF EXISTS Order_History;

DROP TABLE IF EXISTS Client;

CREATE TABLE Client (
    client_id INT NOT NULL,
    client_name VARCHAR(255) NOT NULL,
    exchange_name VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL,
    pair VARCHAR(255) NOT NULL,
    PRIMARY KEY (client_id)
);