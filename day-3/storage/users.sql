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
	(1, '2022-09-14 16:40:55.000', '2022-09-14 16:41:12.000', NULL, 'Admin AGMC', 'agmc@gmail.com', '$2a$10$FfDfmmP3TmRDCb1kcXT1HOmDfutKM6X7XU27F2nDN0f9BvOyPp1jK', '085161318191', 1),
	(2, '2022-09-14 16:41:11.000', '2022-09-14 16:41:12.000', NULL, 'Bayu Pratama', 'bayu@gmail.com', '$2a$10$FfDfmmP3TmRDCb1kcXT1HOmDfutKM6X7XU27F2nDN0f9BvOyPp1jK', '081242875660', 1),
	(4, '2022-09-14 09:09:22.624', '2022-09-14 09:23:07.272', NULL, 'Miftah Hanif', 'miftah@gmail.com', '$2a$10$FfDfmmP3TmRDCb1kcXT1HOmDfutKM6X7XU27F2nDN0f9BvOyPp1jK', '0851614317892', 0),
	(3, '2022-09-14 16:41:11.000', '2022-09-14 16:41:12.000', NULL, 'Siti Aisyah', 'siti@gmail.com', '$2a$10$FfDfmmP3TmRDCb1kcXT1HOmDfutKM6X7XU27F2nDN0f9BvOyPp1jK', '082244112311', 0);
