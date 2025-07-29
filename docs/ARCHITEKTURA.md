# Prosty Schemat Architektury

## System składa się z 4 części:

```
┌─────────────────┐    ┌─────────────────┐
│  KLIENT         │    │  KUCHNIA        │
│  (React)        │    │  (React)        │
│  localhost:3000 │    │  localhost:3001 │
└─────────┬───────┘    └─────────┬───────┘
          │                      │
          │        HTTP          │ HTTP + WebSocket
          │                      │
          └──────┬─────────────────┘
                 │
         ┌───────▼────────┐
         │   BACKEND      │
         │   (Go)         │
         │   localhost:8080│
         └───────┬────────┘
                 │
                 │ SQL
                 │
         ┌───────▼────────┐
         │  POSTGRESQL    │
         │  localhost:5432│
         └────────────────┘
```

## Co robi każda część:

### 🛒 FRONTEND KLIENTA
- Pokazuje menu (Wrapy, Burgery, etc.)
- Koszyk na zamówienia
- Formularz z danymi klienta

### 👨‍🍳 FRONTEND KUCHNI  
- Lista zamówień do zrobienia
- Zmiana statusu: nowe → w trakcie → gotowe → wydane
- Powiadomienia o nowych zamówieniach

### ⚙️ BACKEND (GO)
- API do obsługi zamówień
- Łączenie z bazą danych
- WebSocket dla powiadomień

### 🗄️ BAZA DANYCH
- Menu i kategorie
- Zamówienia klientów
- Historia statusów

## Jak to działa:

1. **Klient** wybiera jedzenie i składa zamówienie
2. **Backend** zapisuje zamówienie do bazy
3. **Kuchnia** dostaje powiadomienie o nowym zamówieniu
4. **Kuchnia** zmienia status gdy gotuje/wydaje jedzenie
