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