-- Tabela kategorii menu
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    emoji VARCHAR(10),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela pozycji menu
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id),
    name VARCHAR(200) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    is_available BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela zamówień
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    table_number VARCHAR(50),
    total_amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela pozycji w zamówieniu
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id),
    menu_item_id INTEGER NOT NULL REFERENCES menu_items(id),
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL
);

-- Wstawianie przykładowych kategorii
INSERT INTO categories (name, emoji, description) VALUES
('Wrapy', '🌯', 'Świeże wrapy z różnymi nadzieniami'),
('Burgery', '🍔', 'Socziste burgery na każdy gust'),
('Sałatki', '🥗', 'Zdrowe i świeże sałatki');

-- Wstawianie przykładowych pozycji menu
INSERT INTO menu_items (category_id, name, description, price) VALUES
(1, 'Wrap Klasyczny', 'Kurczak, sałata, ogórek, sos czosnkowy', 18.00),
(1, 'Wrap Wege', 'Hummus, awokado, sałata, ogórek, papryka', 16.00),
(1, 'Wrap Ostry', 'Kurczak w ostrej marynacie, jalapeno, cebula, sos chipotle', 19.00),
(2, 'Burger Klasyczny', 'Wołowina, sałata, cebula, sos burger', 22.00),
(2, 'Burger Wege', 'Kotlet z quinoa, awokado, sałata, ', 20.00),
(3, 'Sałatka Cezar', 'Sałata rzymska, kurczak, parmezan, grzanki', 15.00),
(3, 'Sałatka Grecka', 'Pomidory, ogórki, oliwki, feta, czerwona cebula', 14.00);