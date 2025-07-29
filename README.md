# README - System Restauracji

## Opis Projektu
System do zarządzania zamówieniami w restauracji składający się z 4 głównych komponentów.

## Architektura
```
Frontend Klienta (React) ←→ Backend (Go) ←→ PostgreSQL
Frontend Kuchni (React)  ←↗
```

## Funkcjonalności

### Dla Klientów
- Przeglądanie menu według kategorii
- Dodawanie pozycji do koszyka  
- Składanie zamówienia z danymi kontaktowymi

### Dla Kuchni
- Dashboard z listą zamówień
- Zmiana statusów: pending → in_progress → ready → completed
- Powiadomienia o nowych zamówieniach (WebSocket)

## Technologie
- **Frontend**: React, Axios, React Router
- **Backend**: Go, Gin, GORM
- **Database**: PostgreSQL
- **Real-time**: WebSocket

## Struktura Projektu
```
frontend-klient/     # Aplikacja dla klientów
frontend-kuchnia/    # Aplikacja dla personelu kuchni
backend/             # API server w Go
docs/                # Dokumentacja architektury
```

## Uruchomienie
1. Skonfiguruj PostgreSQL
2. Uruchom backend Go na porcie 8080
3. Uruchom frontend klienta na porcie 3000  
4. Uruchom frontend kuchni na porcie 3001

## Dokumentacja
- `docs/ARCHITEKTURA_SYSTEMU.md` - Przegląd architektury
- `docs/DIAGRAM_BAZY_DANYCH.md` - Schemat bazy danych
- `docs/KOMUNIKACJA.md` - API i przepływ danych
- `docs/PLAN_IMPLEMENTACJI.md` - Etapy rozwoju
