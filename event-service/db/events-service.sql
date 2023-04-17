CREATE DATABASE IF NOT EXISTS events_service DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci; 

USE events_service;

CREATE TABLE IF NOT EXISTS events(
    id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci; 

INSERT INTO events(start_time, end_time)
VALUES(
    '2023-04-13 10:00:00',
    '2023-04-13 12:00:00'
),(
    '2023-04-13 14:00:00',
    '2023-04-13 16:00:00'
),(
    '2023-04-14 09:00:00',
    '2023-04-14 11:00:00'
),(
    '2023-04-15 11:00:00',
    '2023-04-15 13:00:00'
),(
    '2023-04-16 13:00:00',
    '2023-04-16 15:00:00'
);