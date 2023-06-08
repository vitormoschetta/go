INSERT INTO
    categories (id, `name`)
VALUES
    ('5c2e6cd9-30f5-4714-98de-f4e1139b817c', 'Category A'),
    ('e3c08c86-8046-474d-9b88-3786dbbdd226', 'Category B');    

INSERT INTO
    products (id, `name`, price, category_id)
VALUES
    ('5c2e6cd9-30f5-4714-98de-f4e1139b817c', 'Product 1', 10.00, '5c2e6cd9-30f5-4714-98de-f4e1139b817c'),
    ('e3c08c86-8046-474d-9b88-3786dbbdd226', 'Product 2', 20.00, '5c2e6cd9-30f5-4714-98de-f4e1139b817c'),
    ('5c2e6cd9-30f5-4714-98de-f4e1139b817d', 'Product 3', 30.00, 'e3c08c86-8046-474d-9b88-3786dbbdd226'),
    ('e3c08c86-8046-474d-9b88-3786dbbdd227', 'Product 4', 40.00, 'e3c08c86-8046-474d-9b88-3786dbbdd226');