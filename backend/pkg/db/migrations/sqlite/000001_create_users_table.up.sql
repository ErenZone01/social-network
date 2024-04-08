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

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("max@gmail.com", "max", "1", "Matar", "ndaw", "1990-01-01", "avatar1.jpg", "Description de l'utilisateur 1", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user1@example.com", "user1", "password1", "John", "Doe", "1990-01-01", "avatar1.jpg", "Description de l'utilisateur 1", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user2@example.com", "user2", "password2", "Jane", "Smith", "1985-05-15", "avatar2.jpg", "Description de l'utilisateur 2", "private");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES("user3@example.com", "user3", "password3", "Alice", "Johnson", "1995-09-20", "avatar3.jpg", "Description de l'utilisateur 3", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user4@example.com", "user4", "password4", "Bob", "Brown", "1988-07-10", "avatar4.jpg", "Description de l'utilisateur 4", "private");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user5@example.com", "user5", "password5", "Emma", "Williams", "1993-03-25", "avatar5.jpg", "Description de l'utilisateur 5", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user6@example.com", "user6", "password6", "Michael", "Jones", "1982-11-12", "avatar6.jpg", "Description de l'utilisateur 6", "private");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user7@example.com", "user7", "password7", "Sophia", "Davis", "1998-04-30", "avatar7.jpg", "Description de l'utilisateur 7", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user8@example.com", "user8", "password8", "David", "Martinez", "1980-09-08", "avatar8.jpg", "Description de l'utilisateur 8", "private");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user9@example.com", "user9", "password9", "Olivia", "Garcia", "1997-06-17", "avatar9.jpg", "Description de l'utilisateur 9", "public");

INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES ("user10@example.com", "user10", "password10", "James", "Rodriguez", "1987-12-04", "avatar10.jpg", "Description de l'utilisateur 10", "private");
