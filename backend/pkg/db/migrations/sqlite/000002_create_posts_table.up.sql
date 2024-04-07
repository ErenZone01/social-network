-- 001_initial.sql
CREATE TABLE Post (
    ID_Post INTEGER PRIMARY KEY AUTOINCREMENT,
    Content TEXT NOT NULL,
    Title TEXT NOT NULL,
    Images TEXT NOT NULL,
    ID_User INTEGER NOT NULL,
    Privacy TEXT NOT NULL,
    ID_Group INTEGER NOT NULL,
    Types   TEXT NOT NULL,
    FOREIGN KEY (ID_User) REFERENCES User(ID_User)
);

INSERT INTO Post (Content, Title, Images, ID_User, Privacy, ID_Group, Types) VALUES 
('Découvrez les dernières avancées en intelligence artificielle !', 'Nouvelles avancées en IA', 'chemin_vers_images/ia.jpg', 1, 'Public', 1, 'Article'),
('Conseils pour maintenir un mode de vie sain et actif !', 'Conseils de santé', 'chemin_vers_images/sante.jpg', 2, 'Public', 1, 'Conseils'),
('Explorez les merveilles du monde avec nos guides de voyage !', 'Guides de voyage', 'chemin_vers_images/voyage.jpg', 3, 'Public', 2, 'Guide'),
('Recette facile et délicieuse de pizza maison !', 'Recette de pizza', 'chemin_vers_images/pizza.jpg', 4, 'Public', 2, 'Recette'),
('Les dernières nouvelles du monde politique et économique.', 'Actualités mondiales', 'chemin_vers_images/actualites.jpg', 5, 'Public', 3, 'Actualités'),
('Découvrez les chefs-d''œuvre de la Renaissance italienne !', 'Art de la Renaissance', 'chemin_vers_images/renaissance.jpg', 6, 'Public', 3, 'Article'),
('Conseils pour améliorer votre endurance et votre force musculaire !', 'Entraînement sportif', 'chemin_vers_images/sport.jpg', 7, 'Public', 4, 'Conseils'),
('Les tendances mode de cette saison à ne pas manquer !', 'Tendances mode', 'chemin_vers_images/mode.jpg', 8, 'Public', 4, 'Tendances'),
('Conseils pour réussir vos études et atteindre vos objectifs académiques !', 'Conseils d''éducation', 'chemin_vers_images/education.jpg', 9, 'Public', 5, 'Conseils'),
('Découvrez les artistes émergents et les dernières chansons à écouter !', 'Découvertes musicales', 'chemin_vers_images/musique.jpg', 10, 'Public', 5, 'Découvertes');
