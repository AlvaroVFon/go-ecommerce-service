-- +migration no-transaction
CREATE TABLE IF NOT EXISTS cart_items (
    ID SERIAL PRIMARY KEY,
    cart_id INT NOT NULL,
    product_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    quantity int NOT NULL DEFAULT 1,
    snapshot_price NUMERIC(10,2) NOT NULL,
    discount_rate NUMERIC(5,2) DEFAULT 0,
    total_price NUMERIC(10,2) NOT NULL,
    image_url TEXT,
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_cart FOREIGN KEY(cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
);
