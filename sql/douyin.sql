/*
Navicat MySQL Data Transfer

Source Server         : test
Source Server Version : 50737
Source Host           : localhost:3306
Source Database       : douyin

Target Server Type    : MYSQL
Target Server Version : 50737
File Encoding         : 65001

Date: 2023-02-17 20:37:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comments`
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int(10) NOT NULL,
  `user_id` int(10) DEFAULT NULL,
  `content` varchar(40) DEFAULT NULL,
  `create_date` varchar(20) DEFAULT NULL,
  `deleted_at` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user` (`user_id`),
  CONSTRAINT `user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES ('1', '1', 'Test Comment', '05-01', null);

-- ----------------------------
-- Table structure for `fans`
-- ----------------------------
DROP TABLE IF EXISTS `fans`;
CREATE TABLE `fans` (
  `id` int(10) NOT NULL,
  `follower_id` int(10) DEFAULT NULL,
  `follow_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fans` (`follow_id`),
  KEY `interest` (`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of fans
-- ----------------------------
INSERT INTO `fans` VALUES ('1', '1', '2');
INSERT INTO `fans` VALUES ('2', '1', '3');
INSERT INTO `fans` VALUES ('3', '2', '1');
INSERT INTO `fans` VALUES ('4', '3', '2');
INSERT INTO `fans` VALUES ('5', '3', '1');
INSERT INTO `fans` VALUES ('6', '3', '3');

-- ----------------------------
-- Table structure for `logins`
-- ----------------------------
DROP TABLE IF EXISTS `logins`;
CREATE TABLE `logins` (
  `id` int(10) NOT NULL,
  `name` varchar(12) DEFAULT NULL,
  `password` varchar(18) DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `user_id` FOREIGN KEY (`id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of logins
-- ----------------------------
INSERT INTO `logins` VALUES ('1', 'TestUser', '123456');
INSERT INTO `logins` VALUES ('2', '123456', '123456');
INSERT INTO `logins` VALUES ('3', 'ccc', '123456');

-- ----------------------------
-- Table structure for `love_videos`
-- ----------------------------
DROP TABLE IF EXISTS `love_videos`;
CREATE TABLE `love_videos` (
  `id` int(10) NOT NULL,
  `user_id` int(10) DEFAULT NULL,
  `video_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fan` (`user_id`),
  KEY `love` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of love_videos
-- ----------------------------
INSERT INTO `love_videos` VALUES ('1', '3', '2');
INSERT INTO `love_videos` VALUES ('2', '2', '1');
INSERT INTO `love_videos` VALUES ('3', '2', '2');
INSERT INTO `love_videos` VALUES ('4', '2', '3');
INSERT INTO `love_videos` VALUES ('5', '1', '2');
INSERT INTO `love_videos` VALUES ('6', '1', '4');

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `name` varchar(12) DEFAULT NULL,
  `follow_count` int(10) DEFAULT NULL,
  `follower_count` int(10) DEFAULT NULL,
  `is_follow` tinyint(1) DEFAULT NULL,
  `deleted_at` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'TestUser', '2', '2', '0', null);
INSERT INTO `users` VALUES ('2', '123456', '1', '2', '0', null);
INSERT INTO `users` VALUES ('3', 'ccc', '3', '2', '0', null);

-- ----------------------------
-- Table structure for `videos`
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int(10) NOT NULL,
  `author_id` int(10) DEFAULT NULL,
  `play_url` varchar(40) DEFAULT NULL,
  `cover_url` varchar(40) DEFAULT NULL,
  `favorite_count` int(10) DEFAULT NULL,
  `comment_count` int(10) DEFAULT NULL,
  `is_favorite` tinyint(1) DEFAULT NULL,
  `deleted_at` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `author` (`author_id`),
  CONSTRAINT `author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES ('1', '1', 'bear.mp4', 'bear-1283347_1280.jpg', '1', '0', '0', null);
INSERT INTO `videos` VALUES ('2', '1', '2_VIDEO_20230126_200741242.mp4', 'end.png', '3', '0', '0', null);
INSERT INTO `videos` VALUES ('3', '3', '3_never.mp4', '3_never.jpg', '1', '0', '0', null);
INSERT INTO `videos` VALUES ('4', '1', '1_VIDEO_20230213_151754633.mp4', '1_VIDEO_20230213_151754633.jpg', '1', '0', '0', null);
