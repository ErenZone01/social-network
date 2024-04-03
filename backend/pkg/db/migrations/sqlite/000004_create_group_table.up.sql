CREATE TABLE Groupe(
    ID_Group INTEGER PRIMARY KEY AUTOINCREMENT,
    GroupName TEXT NOT NULL,
    GroupDescription TEXT NOT NULL,
    IdMember TEXT NOT NULL,
    IdCreator Integer NOT NULL
);