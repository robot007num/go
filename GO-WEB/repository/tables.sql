DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE  utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community_section`;
CREATE TABLE `community_section`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `section_name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_section_name` (`section_name`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `section_class`;
CREATE TABLE `section_class`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `class_id` int(11) NOT NULL,
    `class_name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `new_post`;
CREATE TABLE `new_post`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子ID',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_name` bigint(20) NOT NULL COMMENT '所属用户',
    `section_class` bigint(20) NOT NULL COMMENT '所属社区',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;