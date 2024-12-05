# DSA Tools

**DSA Tools** ist eine Sammlung von Werkzeugen für Spielleiter:innen des Rollenspiels *Das Schwarze Auge* (DSA). Es hilft dabei, Abenteuer, Charaktere und andere Spielressourcen effizient zu verwalten.

---

## 🛠 Funktionen

- **Charakter-Management**: Speichere und verwalte alle wichtigen Informationen zu Charakteren.
- **Ressourcenverwaltung**: Lade Karten, Handouts und andere Materialien hoch.
- **Würfelsimulation**: Einfaches und schnelles Würfeln für verschiedene Proben.
- **Integration mit PostgreSQL**: Speicherung von Daten in einer relationalen Datenbank.

---

## 🚀 Installation

### Voraussetzungen

- **Go**: Version 1.20 oder höher
- **Sass**: Installiert auf deinem System

- **PostgreSQL**: Ein laufender PostgreSQL-Server mit dem Supernutzer `postgres`
### Schritte

1. **Repository klonen**:
   ```bash
   git clone https://github.com/aureumapes/dsa-tools
   cd dsa-tools
   ```
2. **Abhängigkeiten installieren:**
    ```bash
   go get .
   go build .
    ```
3. **Postgres einrichten**
    * Datenbank dsa erschaffen:
    ```sql 
    CREATE DATABASE dsa;
   ```
4. **Umgebungsvariable setzen**
   * Stelle sicher, dass die Variable dein Passwort enthält: `export PASSWORT='dein_sql_passwort'`
5. **Starte den Webserver**
    `./dsa-tools`
6. Öffne http://localhost:4921/

## Pfade

# API-Routenübersicht

## POST-Routen

- **`/upload`**: Lädt eine Datei hoch.
- **`/save`**: Speichert Daten.
- **`/entry`**: Bearbeitet oder fügt einen neuen Eintrag hinzu.
- **`/newchar`**: Fügt einen neuen Charakter hinzu.

## PUT-Routen

- **`/save`**: Speichert Daten.
- **`/entry`**: Bearbeitet oder fügt einen neuen Eintrag hinzu.
- **`/newchar`**: Fügt einen neuen Charakter hinzu.

## GET-Routen

- **`/entries`**: Zeigt eine Liste aller Einträge an.
- **`/upload`**: Zeigt Informationen zu hochgeladenen Dateien.
- **`/`**: Lädt die Startseite (`index.html`).
- **`/time`**: Zeigt den aktuellen Kalender an.
- **`/newchar`**: Zeigt die Seite zum Erstellen eines neuen Charakters an.
- **`/favicon.ico`**: Leitet zum Favicon weiter.
- **`/dates`**: Zeigt eine Liste der Jahre an.

## Gruppenrouten

### `/adventure/`

- **`/`**: Zeigt eine Übersicht aller Abenteuer an.
- **`/:adv`**: Zeigt Details zu einem spezifischen Abenteuer an.

### `/dates`

- **`/`**: Zeigt eine Liste der Jahre an.
- **`/:year`**: Zeigt eine Liste der Monate für ein bestimmtes Jahr an.
- **`/:year/:month`**: Zeigt eine Liste der Tage für einen bestimmten Monat an.
- **`/:year/:month/:day`**: Zeigt Details zu einem bestimmten Tag an.

### `/chars`

- **`/`**: Zeigt eine Liste aller Charaktere an.
- **`/:name`**: Zeigt Details zu einem bestimmten Charakter an.

### `/files`

- **`/*file`**: Zeigt eine Datei an.

## ANY-Routen

- **`/files/*file`**: Erlaubt den Zugriff auf beliebige Datei-Routen.

