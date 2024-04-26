-- 001_initial.sql
CREATE TABLE User (
    ID_User INTEGER PRIMARY KEY AUTOINCREMENT,
    Email TEXT NOT NULL UNIQUE,
    Passwords TEXT NOT NULL,
    Firstname TEXT NOT NULL,
    Lastname TEXT NOT NULL,
    Nickname TEXT NOT NULL UNIQUE,
    Birth TEXT NOT NULL,
    Avatar  TEXT NOT NULL,
    About   TEXT NOT NULL,
    Privacy TEXT NOT NULL,
    Actif   TEXT
);
