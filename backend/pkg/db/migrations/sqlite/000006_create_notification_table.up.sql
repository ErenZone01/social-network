CREATE TABLE Notifications(
    ID_Notification INTEGER PRIMARY KEY AUTOINCREMENT,
    Messages TEXT NOT NULL,
    ID_Receiver INTEGER NOT NULL,
    Types TEXT NOT NULL,
    States TEXT NOT NULL,
    ID_User INTEGER NOT NULL,
    ID_Group INTEGER NOT NULL,
    FOREIGN KEY (ID_User) REFERENCES User(ID_User),
    FOREIGN KEY (ID_Receiver) REFERENCES User(ID_Receiver),
    FOREIGN KEY (ID_Group) REFERENCES Groupe(ID_Group)
);