DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

CREATE TABLE categories (
    id VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE products (
    id VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
