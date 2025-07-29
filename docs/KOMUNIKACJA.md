# Komunikacja Między Aplikacjami

## 1. Schemat Komunikacji

```
FRONTEND KLIENTA                    FRONTEND KUCHNI
     │                                    │
     │ HTTP REST API                      │ HTTP + WebSocket
     │                                    │
     └──────────┬───────────────────────────┘
                │
        ┌───────▼────────┐
        │    BACKEND     │
        │    (Go API)    │
        └───────┬────────┘
                │ SQL
        ┌───────▼────────┐
        │   POSTGRESQL   │
        └────────────────┘
```

## 2. Główne Endpointy API

### Menu (dla klienta)
```
GET /api/categories          → Lista kategorii
GET /api/categories/1/items  → Pozycje w kategorii  
GET /api/items/1             → Szczegóły pozycji
```

### Zamówienia
```
POST /api/orders             → Złóż zamówienie
GET  /api/orders/ORD-001     → Status zamówienia
PUT  /api/orders/1/status    → Zmień status (kuchnia)
```

### Kuchnia
```
GET /api/kitchen/orders      → Lista zamówień do realizacji
```

## 3. Przykład Komunikacji

### Składanie Zamówienia
```json
POST /api/orders
{
  "customerName": "Jan Kowalski",
  "customerPhone": "123456789", 
  "items": [
    {
      "menuItemId": 1,
      "quantity": 2
    }
  ]
}

→ Response:
{
  "success": true,
  "data": {
    "orderNumber": "ORD-001",
    "totalAmount": 36.00,
    "status": "pending"
  }
}
```

### WebSocket dla Kuchni
```json
// Nowe zamówienie
{
  "type": "new_order",
  "payload": {
    "orderNumber": "ORD-001", 
    "customerName": "Jan Kowalski",
    "items": [...],
    "status": "pending"
  }
}

// Aktualizacja statusu
{
  "type": "status_update",
  "payload": {
    "orderNumber": "ORD-001",
    "newStatus": "in_progress"
  }
}
```

## 4. Przepływ Realizacji Zamówienia

```
1. KLIENT składa zamówienie
   │
   ▼
2. BACKEND zapisuje do bazy
   │
   ▼  
3. KUCHNIA otrzymuje powiadomienie (WebSocket)
   │
   ▼
4. KUCHNIA rozpoczyna realizację → zmienia status
   │
   ▼
5. BACKEND aktualizuje bazę
   │
   ▼
6. KLIENT może sprawdzić status zamówienia
```

## 5. Porty i Adresy

- **Frontend Klienta**: http://localhost:3000
- **Frontend Kuchni**: http://localhost:3001  
- **Backend API**: http://localhost:8080
- **WebSocket**: ws://localhost:8080/ws
- **PostgreSQL**: localhost:5432
