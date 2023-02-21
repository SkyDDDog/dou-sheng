/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.27 : Database - dousheng
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`dousheng` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `dousheng`;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
                           `message_id` bigint NOT NULL AUTO_INCREMENT,
                           `content` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                           `from_id` bigint DEFAULT NULL,
                           `to_id` bigint DEFAULT NULL,
                           `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                           `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                           `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                           `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                           PRIMARY KEY (`message_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1623647466282815489 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `message` */

insert  into `message`(`message_id`,`content`,`from_id`,`to_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1619372827817807872,'消息内容',1617455497353367552,1617452224676368384,'0','2023-01-28 16:32:23','2023-01-28 16:32:23','');
insert  into `message`(`message_id`,`content`,`from_id`,`to_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623644530723000320,'hello',1617455497353367552,1623641351151161344,'0','2023-02-09 11:26:37','2023-02-09 11:26:37','');
insert  into `message`(`message_id`,`content`,`from_id`,`to_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623644866183434240,'hi\n',1623641351151161344,1617455497353367552,'0','2023-02-09 11:27:57','2023-02-09 11:27:57','');
insert  into `message`(`message_id`,`content`,`from_id`,`to_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623647466282815488,'hihihi',1623641351151161344,1617455497353367552,'0','2023-02-09 11:38:17','2023-02-09 11:38:17','');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
                        `user_id` bigint NOT NULL AUTO_INCREMENT,
                        `username` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                        `nickname` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                        `password_digest` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                        `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                        `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                        `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                        `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                        PRIMARY KEY (`user_id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1623641351151161345 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `user` */

insert  into `user`(`user_id`,`username`,`nickname`,`password_digest`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1617452224676368384,'ErHuo','ErHuo','$2a$12$z4SqqXeEMYqe4RWh3m23W.CIlhlC5swDmaife5WAWro6LWDy0yq0S','0','2023-01-23 09:20:36','2023-01-23 09:20:36','');
insert  into `user`(`user_id`,`username`,`nickname`,`password_digest`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1617455497353367552,'362664609@qq.com','362664609@qq.com','$2a$12$hI3IkwixeP.ruufd7/99j.VvOYPkltJxHgMLJCYNp7dkR353tKgB.','0','2023-01-23 09:33:36','2023-01-23 09:33:36','');
insert  into `user`(`user_id`,`username`,`nickname`,`password_digest`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623641351151161344,'test@qq.com','test@qq.com','$2a$12$Y32CHsPXqlxcqGFdnEGOIOpjYtBhSRlp9QMl6lMrmktG3iMUslPGm','0','2023-02-09 11:13:59','2023-02-09 11:13:59','');

/*Table structure for table `user_favorite` */

DROP TABLE IF EXISTS `user_favorite`;

CREATE TABLE `user_favorite` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `user_id` bigint DEFAULT NULL,
                                 `video_id` bigint DEFAULT NULL,
                                 `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                                 `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                                 `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                                 `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1618266803459330049 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `user_favorite` */

insert  into `user_favorite`(`id`,`user_id`,`video_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1618266803459330048,1617455497353367552,1617794857957330944,'0','2023-01-25 15:17:27','2023-01-25 15:17:27','');

/*Table structure for table `user_follow` */

DROP TABLE IF EXISTS `user_follow`;

CREATE TABLE `user_follow` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `from_user_id` bigint DEFAULT NULL,
                               `to_user_id` bigint DEFAULT NULL,
                               `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                               `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                               `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                               `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1623642696163790849 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `user_follow` */

insert  into `user_follow`(`id`,`from_user_id`,`to_user_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1618589003735371776,1617455497353367552,1617452224676368384,'0','2023-01-26 12:37:45','2023-01-26 12:37:45','');
insert  into `user_follow`(`id`,`from_user_id`,`to_user_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623641399498903552,1623641351151161344,1617455497353367552,'0','2023-02-09 11:14:10','2023-02-09 11:14:10','');
insert  into `user_follow`(`id`,`from_user_id`,`to_user_id`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623642696163790848,1617455497353367552,1623641351151161344,'0','2023-02-09 11:19:19','2023-02-09 11:19:19','');

/*Table structure for table `video` */

DROP TABLE IF EXISTS `video`;

CREATE TABLE `video` (
                         `video_id` bigint NOT NULL AUTO_INCREMENT,
                         `author_id` bigint DEFAULT NULL,
                         `title` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                         `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                         `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                         `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                         `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                         `cover_url` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                         `video_url` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                         PRIMARY KEY (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1625870194675027969 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `video` */

insert  into `video`(`video_id`,`author_id`,`title`,`del_flag`,`create_date`,`update_date`,`remarks`,`cover_url`,`video_url`) values (1625870194675027968,1617455497353367552,'测试','0','2023-02-15 14:50:36','2023-02-15 14:50:36','','https://dousheng-1313040197.cos.ap-guangzhou.myqcloud.com/cover/1625870176782127104.jpg','https://dousheng-1313040197.cos.ap-guangzhou.myqcloud.com/video/1625870176782127104.mp4');

/*Table structure for table `video_comment` */

DROP TABLE IF EXISTS `video_comment`;

CREATE TABLE `video_comment` (
                                 `comment_id` bigint NOT NULL AUTO_INCREMENT,
                                 `video_id` bigint DEFAULT NULL,
                                 `author_id` bigint DEFAULT NULL,
                                 `comment_text` longtext CHARACTER SET utf8 COLLATE utf8_general_ci,
                                 `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '逻辑删除',
                                 `create_date` datetime DEFAULT NULL COMMENT '创建时间',
                                 `update_date` datetime DEFAULT NULL COMMENT '更新时间',
                                 `remarks` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
                                 PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1623643259953745921 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

/*Data for the table `video_comment` */

insert  into `video_comment`(`comment_id`,`video_id`,`author_id`,`comment_text`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1618284373985988608,1617794857957330944,1617455497353367552,'评论捏','0','2023-01-25 16:27:16','2023-01-25 16:27:16','');
insert  into `video_comment`(`comment_id`,`video_id`,`author_id`,`comment_text`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1619566633968209920,1617794857957330944,1617455497353367552,'哈哈哈','0','2023-01-29 05:22:30','2023-01-29 05:22:30','');
insert  into `video_comment`(`comment_id`,`video_id`,`author_id`,`comment_text`,`del_flag`,`create_date`,`update_date`,`remarks`) values (1623643259953745920,1617794857957330944,1623641351151161344,'评论一下','0','2023-02-09 11:21:34','2023-02-09 11:21:34','');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
