CREATE TABLE Groupe(
    ID_Group INTEGER PRIMARY KEY AUTOINCREMENT,
    GroupName TEXT NOT NULL,
    GroupImage TEXT NOT NULL,
    GroupDescription TEXT NOT NULL,
    IdCreator Integer NOT NULL
);