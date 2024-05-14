CREATE TABLE
    Notifications (
        ID_Notification INTEGER PRIMARY KEY AUTOINCREMENT,
        Messages TEXT NOT NULL,
        Receiver TEXT NOT NULL,
        Types TEXT NOT NULL,
        States TEXT NOT NULL,
        Sender Text NOT NULL,
        AvatarSender TEXT,
        ID_Group INTEGER NOT NULL,
        ID_Event INTEGER NOT NULL,
        FOREIGN KEY (Sender) REFERENCES User (Nickname),
        FOREIGN KEY (Receiver) REFERENCES User (Nickname),
        FOREIGN KEY (ID_Group) REFERENCES Groupe (ID_Group),
        FOREIGN KEY (ID_Event) REFERENCES EventGroup (ID_Event)
    );