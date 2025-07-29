# Diagram Bazy Danych - ERD

## Schemat Relacji Tabel

```
┌─────────────────┐     ┌─────────────────┐
│   CATEGORIES    │     │   MENU_ITEMS    │
│                 │     │                 │
│ • id (PK)       │────▶│ • id (PK)       │
│ • name          │ 1:N │ • category_id   │
│ • description   │     │ • name          │
│ • is_active     │     │ • description   │
└─────────────────┘     │ • price         │
                        │ • is_available  │
                        │ • is_vegetarian │
                        │ • is_spicy      │
                        └─────────┬───────┘
                                  │
                                  │ 1:N
                                  ▼
┌─────────────────┐     ┌─────────────────┐
│     ORDERS      │     │   ORDER_ITEMS   │
│                 │     │                 │
│ • id (PK)       │────▶│ • id (PK)       │
│ • order_number  │ 1:N │ • order_id      │
│ • customer_name │     │ • menu_item_id  │◀──┘
│ • customer_phone│     │ • quantity      │
│ • total_amount  │     │ • unit_price    │
│ • status        │     │ • total_price   │
│ • notes         │     └─────────────────┘
│ • created_at    │
└─────────────────┘
```

## Główne Tabele

### 1. CATEGORIES (Kategorie Menu)
- Wrapy, Burgery, Sałatki, Napoje, Desery
- Każda kategoria może być aktywna/nieaktywna

### 2. MENU_ITEMS (Pozycje Menu) 
- Konkretne pozycje w każdej kategorii
- Cena, dostępność, informacje o alergenach
- Przykład: "Wrap Klasyczny", "Burger Wege"

### 3. ORDERS (Zamówienia)
- Główne zamówienie z danymi klienta
- Unikalny numer zamówienia (ORD-001)
- Status realizacji

### 4. ORDER_ITEMS (Pozycje w Zamówieniu)
- Konkretne pozycje zamówione przez klienta
- Ilość, cena jednostkowa, cena całkowita
- Specjalne życzenia klienta

## Przykładowe Dane

### Kategoria: Wrapy
- Wrap Klasyczny (18 zł)
- Wrap Wege (16 zł) 
- Wrap Ostry (19 zł)

### Zamówienie ORD-001
- Klient: Jan Kowalski
- 2x Wrap Klasyczny (36 zł)
- 1x Wrap Wege (16 zł)
- **Razem: 52 zł**
