CREATE TABLE `circle` (
    `id` int NOT NULL AUTO_INCREMENT,
    `p_id` int NOT NULL,
    `p_name` varchar(45) NOT NULL,
    `content` varchar(45) NOT NULL,
    `release_time` timestamp NOT NULL,
    `by_time_age` int NOT NULL,
    `by_time_tall` float NOT NULL,
    `by_time_weight` float NOT NULL,
    `by_time_fat_rate` float NOT NULL,
    `visible` tinyint(1) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3