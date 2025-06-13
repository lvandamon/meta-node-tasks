-- create database blog
CREATE DATABASE `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
-- blog.users definition

CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `username` varchar(191) NOT NULL,
                         `password` longtext NOT NULL,
                         `email` varchar(191) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uni_users_username` (`username`),
                         UNIQUE KEY `uni_users_email` (`email`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- blog.posts definition

CREATE TABLE `posts` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `title` longtext NOT NULL,
                         `content` longtext NOT NULL,
                         `user_id` bigint unsigned DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_posts_deleted_at` (`deleted_at`),
                         KEY `fk_posts_user` (`user_id`),
                         CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- blog.comments definition

CREATE TABLE `comments` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `content` longtext NOT NULL,
                            `user_id` bigint unsigned DEFAULT NULL,
                            `post_id` bigint unsigned DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_comments_deleted_at` (`deleted_at`),
                            KEY `fk_comments_user` (`user_id`),
                            KEY `fk_comments_post` (`post_id`),
                            CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
                            CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;