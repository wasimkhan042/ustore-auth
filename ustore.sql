-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Oct 27, 2021 at 12:22 PM
-- Server version: 8.0.26-0ubuntu0.20.04.3
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ustore`
--

-- --------------------------------------------------------

--
-- Table structure for table `item`
--

CREATE TABLE `item` (
  `item_id` int NOT NULL,
  `item_name` varchar(25) NOT NULL,
  `item_details` varchar(255) NOT NULL,
  `monthly_price` decimal(10,2) NOT NULL,
  `yearly_price` decimal(10,2) NOT NULL,
  `available_items` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `subscription`
--

CREATE TABLE `subscription` (
  `id` char(40) NOT NULL,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `subs_price` decimal(8,2) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `user_id` char(40) NOT NULL,
  `item_name` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `subscription`
--

INSERT INTO `subscription` (`id`, `start_time`, `end_time`, `subs_price`, `status`, `user_id`, `item_name`) VALUES
('14c56e9f-470a-4002-901e-7117b6c9bad6', '2021-10-25 09:54:32', '2021-11-25 09:54:32', '1200.00', 1, '8a400fba-0f18-43f0-86c7-53f780753782', 'TV'),
('3894e673-646f-4099-a62f-820b38e545c6', '2021-10-25 11:37:49', '2021-11-25 11:37:49', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV'),
('4f3aeabe-c571-46b5-864e-18a7f3dcb82f', '2021-10-25 11:10:17', '2021-11-25 11:10:17', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV'),
('6d7de0c8-490b-42cf-8964-3cfa9da68892', '2021-10-25 10:23:46', '2021-11-25 10:23:46', '1200.00', 1, '8a400fba-0f18-43f0-86c7-53f780753782', 'TV'),
('72042521-84e2-4251-8b2c-6a56094dca02', '2021-10-25 10:24:35', '2021-11-25 10:24:35', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV'),
('75e95858-c2f6-4431-bf0a-a481aa93bbbc', '2021-10-27 02:09:10', '2021-11-27 02:09:10', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV'),
('77a079ba-db35-44d0-85dc-42f57a0eb7f9', '2021-10-25 09:55:19', '2021-11-25 09:55:19', '1200.00', 1, '8a400fba-0f18-43f0-86c7-53f780753782', 'TV'),
('a2d4e43f-18fc-4144-8aef-4e966a8de69d', '2021-10-25 11:35:00', '2021-11-25 11:35:00', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV'),
('b6bce02d-0d6a-4c53-bd8d-ec72fa4ccbf2', '2021-10-27 02:09:26', '2021-11-27 02:09:26', '1200.00', 1, '0c38f31b-d08a-470a-a1b5-9c7cff775934', 'TV');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `first_name` varchar(15) NOT NULL,
  `middle_name` varchar(15) NOT NULL,
  `last_name` varchar(15) NOT NULL,
  `email` varchar(50) NOT NULL,
  `username` varchar(25) NOT NULL,
  `password` char(60) NOT NULL,
  `profile_image` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `first_name`, `middle_name`, `last_name`, `email`, `username`, `password`, `profile_image`) VALUES
('0c38f31b-d08a-470a-a1b5-9c7cff775934', 'Wasim', '', 'Khan', 'wk@gmail.com', 'wk', 'password', 'image.png');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `item`
--
ALTER TABLE `item`
  ADD PRIMARY KEY (`item_id`),
  ADD UNIQUE KEY `item_name` (`item_name`);

--
-- Indexes for table `subscription`
--
ALTER TABLE `subscription`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `item`
--
ALTER TABLE `item`
  MODIFY `item_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
