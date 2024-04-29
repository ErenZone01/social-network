-- 001_initial.sql
CREATE TABLE User (
    ID_User INTEGER PRIMARY KEY AUTOINCREMENT,
    Email TEXT NOT NULL UNIQUE,
    Passwords TEXT NOT NULL,
    Firstname TEXT NOT NULL,
    Lastname TEXT NOT NULL,
    Nickname TEXT,
    Birth TEXT NOT NULL,
    Avatar  TEXT,
    About   TEXT,
    Privacy TEXT NOT NULL,
    Actif   TEXT
);
