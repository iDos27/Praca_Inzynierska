# GŁÓWNY SCHEMAT - System Restauracji

## 🎯 CEL
System do zarządzania zamówieniami w restauracji

## 🏗️ ARCHITEKTURA
```
Klient (React) ←→ Backend (Go) ←→ PostgreSQL
Kuchnia (React) ←↗
```

## 📊 BAZA DANYCH
```
CATEGORIES → MENU_ITEMS → ORDER_ITEMS ← ORDERS
    │            │            │          │
Kategorie    Pozycje      Pozycje    Zamówienia
(Wrapy)    (Wrap Ostry)  w zamów.   (ORD-001)
```

## 🔄 PRZEPŁYW ZAMÓWIENIA
```
1. Klient wybiera jedzenie
2. Składa zamówienie  
3. Kuchnia dostaje powiadomienie
4. Kuchnia gotuje i zmienia status
5. Klient może sprawdzić status
```

## 📱 INTERFEJSY

### Klient
- Menu z kategoriami
- Koszyk
- Formularz zamówienia

### Kuchnia  
- Lista zamówień
- Przyciski zmiany statusu
- Powiadomienia

## 🚀 TECHNOLOGIE
- **Frontend**: React
- **Backend**: Go (Gin framework)
- **Baza**: PostgreSQL
- **Real-time**: WebSocket
