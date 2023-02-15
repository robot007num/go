DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
                       `id` bigint(20) NOT NULL AUTO_INCREMENT,
                       `user_id` bigint(20) NOT NULL,
                       `account` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                       `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                       `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                       `email` varchar(64) COLLATE  utf8mb4_general_ci DEFAULT 'NULL',
                       `type` tinyint(4) NOT NULL DEFAULT '0',
                       `enable` tinyint(4) NOT NULL DEFAULT '0',
                       `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                       `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`(
                                    `id` int(11) NOT NULL AUTO_INCREMENT,
                                    `section_name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                                    `introduction` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
                                    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `section`;
CREATE TABLE `section`(
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `class_id` int(11) NOT NULL,
                                `class_name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                                `introduction` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
                                `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`(
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `section_id` int(11) NOT NULL,
                          `post_id` bigint(20) NOT NULL,
                          `user_id` bigint(20) NOT NULL,
                          `type` tinyint(4) NOT NULL DEFAULT '0',
                          `status` tinyint(4) NOT NULL DEFAULT '0',
                          `comment_count` bigint(20) NOT NULL DEFAULT '0',
                          `score` bigint(20) NOT NULL DEFAULT '0',
                          `title` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                          `content` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
                          `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;