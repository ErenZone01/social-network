-- 001_initial.sql
CREATE TABLE Post (
    ID_Post INTEGER PRIMARY KEY AUTOINCREMENT,
    Names TEXT,
    Content TEXT NOT NULL,
    Title TEXT,
    Imagee TEXT,
    ID_User INTEGER NOT NULL,
    Privacy TEXT NOT NULL,
    ID_Group INTEGER,
    Types   TEXT NOT NULL,
    CreatedPost TEXT,
    FOREIGN KEY (ID_User) REFERENCES User(ID_User)
);
