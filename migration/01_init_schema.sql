-- Create enum types for preferences and units
CREATE TABLE IF NOT EXISTS public.users (
  id character varying(255) PRIMARY KEY DEFAULT gen_random_uuid(),
  email character varying(255) DEFAULT NULL,
  password character varying(255) NOT NULL,
  phone character varying(255) DEFAULT NULL,
  bank_account_name character varying(100) DEFAULT NULL,
  bank_account_holder character varying(100) DEFAULT NULL,
  bank_account_number character varying(100) DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the file types table
CREATE TABLE IF NOT EXISTS public.files (
  file_id SERIAL PRIMARY KEY,
  file_uri VARCHAR(255) NOT NULL,
  file_thumbnail_uri VARCHAR(255) NOT NULL
);

-- Create the product table without the generated column
CREATE TABLE IF NOT EXISTS public.products (
  product_id SERIAL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  category VARCHAR(50) NOT NULL,
  qty INTEGER NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  sku VARCHAR(32) NOT NULL,
  file_id INTEGER NOT NULL REFERENCES public.files(file_id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id character varying(255) NOT NULL
);

-- Create the purchase table without the generated column
CREATE TABLE IF NOT EXISTS public.purchases (
    purchase_id SERIAL PRIMARY KEY,
    total_price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    sender_name character varying(255),
    sender_contact_type character varying(255),
    sender_contact_detail character varying(255),
    user_ids INT[] NOT NULL DEFAULT array[]::INT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for storing purchased items
CREATE TABLE IF NOT EXISTS public.purchased_items (
    item_id SERIAL PRIMARY KEY,
    purchase_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    qty INTEGER NOT NULL,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0
);

-- Create function to update updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';


-- Create trigger in users to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- Create trigger in products to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON products
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- Create trigger in purchases to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON purchases
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- Create trigger in purchased_items to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON purchased_items
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();