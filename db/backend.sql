CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `token` varchar(100) DEFAULT NULL,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `businesses` (
  `id` varchar(100) NOT NULL,
  `location` varchar(100) DEFAULT NULL,
  `latitude` varchar(100) DEFAULT NULL,
  `longitude` varchar(100) DEFAULT NULL,
  `term` varchar(100) DEFAULT NULL,
  `radius` varchar(100) DEFAULT NULL,
  `categories` json DEFAULT NULL,
  `locale` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'locale',
  `price` json DEFAULT NULL,
  `open_now` tinyint(1) DEFAULT NULL,
  `open_at` int DEFAULT NULL,
  `attributes` json DEFAULT NULL,
  `sort_by` varchar(100) DEFAULT NULL,
  `device_platform` varchar(100) DEFAULT NULL,
  `reservation_date` varchar(100) DEFAULT NULL,
  `reservation_time` varchar(100) DEFAULT NULL,
  `reservation_covers` int DEFAULT NULL,
  `matches_party_size_param` tinyint(1) DEFAULT NULL,
  `limit` int DEFAULT NULL,
  `offset` int DEFAULT NULL,
  KEY `businesses_id_IDX` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
