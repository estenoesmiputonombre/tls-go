CREATE DATABASE IF NOT EXISTS google;
USE google;

CREATE TABLE IF NOT EXISTS languages (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500),
    PRIMARY KEY(id),
    KEY(name)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS employers (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    id_language INT NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(id_language) REFERENCES languages(id)
) ENGINE=INNODB;

/*****      Inserting data in languages        ***/
INSERT INTO 
    languages(name, description)
VALUES 
    ('go', 'Language oriented to the server part. It was created by google and it is a son of c programming language');

INSERT INTO 
    languages(name, description)
VALUES 
    ('c', 'Language oriented to the server part. It was created by someone and its funny');

INSERT INTO 
    languages(name, description)
VALUES 
    ('c++', 'Language oriented to the server part. It was created by you?');

/*****      Inserting data in employers        ***/
INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Pepe', 'Pepita', 1);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Jose', 'Manuel', 2);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Manuel', 'Cofrade', 3);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Noemi', 'Bouzo', 2);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Ivan', 'Martinez', 1);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Oscar', 'Dorrego', 3);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Pilar', 'Fernandez', 2);

INSERT INTO 
    employers(name, surname, id_language)
VALUES
    ('Pablo', 'Rodriguez', 1);

SELECT * FROM employers;
SELECT name FROM languages WHERE id IN (SELECT id_language FROM employers WHERE name LIKE 'Ivan');
SELECT e.name AS name, e.surname AS surname, l.name AS language
FROM employers e 
        INNER JOIN 
    languages l ON e.id_language = l.id;