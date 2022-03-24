CREATE TABLE IF NOT EXISTS `movies` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(64) NOT NULL,
    `release_year` year DEFAULT NULL,
    `production` varchar(64) NOT NULL,
    `overview` text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
);