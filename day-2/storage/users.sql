CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext,
  `phone_number` longtext,
  `status` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`, `phone_number`, `status`)
VALUES
	(1, '2022-09-14 16:40:55.000', '2022-09-14 16:41:12.000', NULL, 'Andy Nur', 'andynur.id@gmail.com', '$2a$10$f1jNJdnu47votYfDnupFXupPya/lp.f9mn/DCAZjfRvUNIQMhBN1q', '085161318191', 1),
	(2, '2022-09-14 16:41:11.000', '2022-09-14 16:41:12.000', NULL, 'Bayu Pratama', 'bayu@gmail.com', '$2a$10$f1jNJdnu47votYfDnupFXupPya/lp.f9mn/DCAZjfRvUNIQMhBN1q', '081242875660', 1),
	(3, '2022-09-14 16:41:11.000', '2022-09-14 16:41:12.000', NULL, 'Siti Aisyah', 'siti@gmail.com', '$2a$10$f1jNJdnu47votYfDnupFXupPya/lp.f9mn/DCAZjfRvUNIQMhBN1q', '082244112311', 0),
	(4, '2022-09-14 09:09:22.624', '2022-09-14 09:23:07.272', '2022-09-14 19:27:12.000', 'Miftah Hanif', 'miftah@gmail.com', '$2a$10$f1jNJdnu47votYfDnupFXupPya/lp.f9mn/DCAZjfRvUNIQMhBN1q', '0851614317892', 0);
