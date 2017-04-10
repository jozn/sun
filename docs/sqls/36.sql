/*
SQLyog Community v12.2.4 (64 bit)
MySQL - 5.7.15-log : Database - ms
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ms` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `ms`;

/*Table structure for table `activity` */

DROP TABLE IF EXISTS `activity`;

CREATE TABLE `activity` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ActorUserId` int(11) NOT NULL,
  `ActionTypeId` int(11) NOT NULL,
  `TargetId` int(11) NOT NULL,
  `RefId` bigint(20) NOT NULL,
  `CreatedAt` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `ActorUserId` (`ActorUserId`,`Id`),
  KEY `RefId` (`RefId`)
) ENGINE=InnoDB AUTO_INCREMENT=25556 DEFAULT CHARSET=utf8;

/*Table structure for table `comments` */

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `Id` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `PostId` int(20) NOT NULL,
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `PostId` (`PostId`)
) ENGINE=InnoDB AUTO_INCREMENT=191881 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `following_list` */

DROP TABLE IF EXISTS `following_list`;

CREATE TABLE `following_list` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `ListType` int(11) NOT NULL DEFAULT '0',
  `Name` varchar(75) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `Count` int(11) NOT NULL DEFAULT '0',
  `IsAuto` tinyint(1) NOT NULL DEFAULT '0',
  `IsPimiry` tinyint(1) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`,`UserId`)
) ENGINE=MyISAM AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `following_list_member` */

DROP TABLE IF EXISTS `following_list_member`;

CREATE TABLE `following_list_member` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ListId` int(11) NOT NULL DEFAULT '0',
  `UserId` int(11) NOT NULL,
  `FollowedUserId` int(11) NOT NULL,
  `FollowType` tinyint(4) NOT NULL DEFAULT '1',
  `UpdatedTimeMs` bigint(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId` (`UserId`,`FollowedUserId`),
  KEY `FollowedUserId` (`FollowedUserId`,`UserId`),
  KEY `UserId_2` (`UserId`,`UpdatedTimeMs`)
) ENGINE=MyISAM AUTO_INCREMENT=4428 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `following_list_member_history` */

DROP TABLE IF EXISTS `following_list_member_history`;

CREATE TABLE `following_list_member_history` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ListId` int(11) NOT NULL DEFAULT '0',
  `UserId` int(11) NOT NULL DEFAULT '0',
  `FollowedUserId` int(11) NOT NULL DEFAULT '0',
  `FollowType` tinyint(4) NOT NULL DEFAULT '1',
  `UpdatedTimeMs` bigint(11) NOT NULL DEFAULT '0',
  `FollowId` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`,`UpdatedTimeMs`)
) ENGINE=MyISAM AUTO_INCREMENT=13458 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `likes` */

DROP TABLE IF EXISTS `likes`;

CREATE TABLE `likes` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `PostId` int(11) NOT NULL,
  `UserId` int(11) NOT NULL,
  `TypeId` tinyint(4) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `PostId` (`PostId`,`UserId`),
  KEY `Id` (`Id`),
  KEY `PostId_2` (`PostId`)
) ENGINE=MyISAM AUTO_INCREMENT=78699 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `media` */

DROP TABLE IF EXISTS `media`;

CREATE TABLE `media` (
  `Id` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `PostId` int(11) NOT NULL,
  `AlbumId` int(11) NOT NULL DEFAULT '0',
  `TypeId` smallint(4) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  `Src` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'json of sizes',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=272 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(11) NOT NULL DEFAULT '0',
  `RoomKey` varchar(40) NOT NULL,
  `MessageKey` varchar(250) NOT NULL DEFAULT '',
  `FromUserID` int(11) NOT NULL DEFAULT '0',
  `Data` text NOT NULL,
  `TimeMs` bigint(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `ToUserId` (`ToUserId`),
  KEY `ToUserId_2` (`ToUserId`,`TimeMs`)
) ENGINE=InnoDB AUTO_INCREMENT=2058 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `msg_deleted_from_server` */

DROP TABLE IF EXISTS `msg_deleted_from_server`;

CREATE TABLE `msg_deleted_from_server` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(11) NOT NULL DEFAULT '0',
  `MsgKey` varchar(40) NOT NULL DEFAULT '',
  `PeerUserId` int(11) NOT NULL DEFAULT '0',
  `RoomKey` varchar(40) NOT NULL DEFAULT '0',
  `AtTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `ToUserId` (`ToUserId`)
) ENGINE=MyISAM AUTO_INCREMENT=2920 DEFAULT CHARSET=utf8;

/*Table structure for table `msg_received_to_peer` */

DROP TABLE IF EXISTS `msg_received_to_peer`;

CREATE TABLE `msg_received_to_peer` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(11) NOT NULL DEFAULT '0',
  `MsgKey` varchar(40) NOT NULL DEFAULT '',
  `RoomKey` varchar(40) NOT NULL,
  `PeerUserId` int(11) NOT NULL,
  `AtTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `ToUserId` (`ToUserId`)
) ENGINE=MyISAM AUTO_INCREMENT=2921 DEFAULT CHARSET=utf8;

/*Table structure for table `msg_seen_by_peer` */

DROP TABLE IF EXISTS `msg_seen_by_peer`;

CREATE TABLE `msg_seen_by_peer` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(11) NOT NULL,
  `MsgKey` varchar(40) NOT NULL DEFAULT '',
  `RoomKey` varchar(40) NOT NULL,
  `PeerUserId` int(11) NOT NULL,
  `AtTime` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `ToUserId` (`ToUserId`)
) ENGINE=MyISAM AUTO_INCREMENT=1170 DEFAULT CHARSET=utf8;

/*Table structure for table `notification` */

DROP TABLE IF EXISTS `notification`;

CREATE TABLE `notification` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `ForUserId` int(10) NOT NULL DEFAULT '0',
  `ActorUserId` int(10) NOT NULL DEFAULT '0',
  `ActionTypeId` int(10) NOT NULL DEFAULT '0',
  `ObjectTypeId` int(10) NOT NULL DEFAULT '0',
  `TargetId` int(10) NOT NULL DEFAULT '0',
  `ObjectId` int(10) NOT NULL DEFAULT '0',
  `SeenStatus` int(10) NOT NULL DEFAULT '0',
  `CreatedTime` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `ForUserId` (`ForUserId`,`Id`),
  KEY `TargetId` (`TargetId`)
) ENGINE=InnoDB AUTO_INCREMENT=29995 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `notification_removed` */

DROP TABLE IF EXISTS `notification_removed`;

CREATE TABLE `notification_removed` (
  `NotificationId` int(11) NOT NULL,
  `ForUserId` int(11) NOT NULL,
  PRIMARY KEY (`NotificationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `phone_contacts` */

DROP TABLE IF EXISTS `phone_contacts`;

CREATE TABLE `phone_contacts` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `PhoneDisplayName` varchar(250) NOT NULL DEFAULT '',
  `PhoneFamilyName` varchar(250) NOT NULL DEFAULT '',
  `PhoneNumber` varchar(250) NOT NULL DEFAULT '',
  `PhoneNormalizedNumber` varchar(15) NOT NULL DEFAULT '',
  `PhoneContactRowId` int(10) NOT NULL DEFAULT '0',
  `UserId` int(10) NOT NULL DEFAULT '0',
  `DeviceUuidId` int(10) NOT NULL DEFAULT '0',
  `CreatedTime` int(10) NOT NULL DEFAULT '0',
  `UpdatedTime` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `PhoneContactRowId` (`PhoneContactRowId`,`UserId`),
  KEY `PhoneNumber` (`PhoneNumber`),
  KEY `UserId` (`UserId`),
  KEY `UserId_Time` (`UserId`,`CreatedTime`),
  KEY `PhoneNormalizedNumber` (`PhoneNormalizedNumber`)
) ENGINE=MyISAM AUTO_INCREMENT=918 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `post` */

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `Id` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `TypeId` int(4) NOT NULL DEFAULT '1',
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `FormatedText` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `MediaUrl` varchar(250) NOT NULL DEFAULT '""',
  `MediaServerId` int(11) NOT NULL DEFAULT '0',
  `Width` int(11) NOT NULL DEFAULT '0',
  `Height` int(11) NOT NULL DEFAULT '0',
  `SharedTo` int(11) NOT NULL DEFAULT '0',
  `HasTag` int(11) NOT NULL DEFAULT '0',
  `LikesCount` int(11) NOT NULL DEFAULT '0',
  `CommentsCount` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`)
) ENGINE=MyISAM AUTO_INCREMENT=3068 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `recommend_user` */

DROP TABLE IF EXISTS `recommend_user`;

CREATE TABLE `recommend_user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL DEFAULT '0',
  `TargetId` int(11) NOT NULL DEFAULT '0',
  `Weight` float NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2290063 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `search_clicked` */

DROP TABLE IF EXISTS `search_clicked`;

CREATE TABLE `search_clicked` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Query` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `ClickType` int(11) NOT NULL DEFAULT '1',
  `TargetId` int(11) NOT NULL DEFAULT '0',
  `UserId` int(11) NOT NULL DEFAULT '0',
  `CreatedAt` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `session` */

DROP TABLE IF EXISTS `session`;

CREATE TABLE `session` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `SessionUuid` varchar(75) CHARACTER SET utf8 NOT NULL,
  `ClientUuid` varchar(75) CHARACTER SET utf8 NOT NULL,
  `DeviceUuid` varchar(75) CHARACTER SET utf8 NOT NULL,
  `LastActivityTime` int(11) NOT NULL DEFAULT '0',
  `LastIpAddress` varchar(75) CHARACTER SET utf8 NOT NULL,
  `LastWifiMacAddress` varchar(75) CHARACTER SET utf8 NOT NULL,
  `LastNetworkType` varchar(75) CHARACTER SET utf8 NOT NULL,
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `Id` (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `tags` */

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) NOT NULL,
  `Count` int(11) NOT NULL,
  `IsBlocked` int(1) NOT NULL,
  `CreatedTime` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name` (`Name`)
) ENGINE=MyISAM AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 DELAY_KEY_WRITE=1;

/*Table structure for table `tags_posts` */

DROP TABLE IF EXISTS `tags_posts`;

CREATE TABLE `tags_posts` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `TagId` int(11) NOT NULL,
  `PostId` int(11) NOT NULL,
  `TypeId` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `TagId` (`TagId`,`PostId`)
) ENGINE=MyISAM AUTO_INCREMENT=28218 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(250) NOT NULL DEFAULT '',
  `FirstName` varchar(250) NOT NULL DEFAULT '',
  `LastName` varchar(250) NOT NULL DEFAULT '',
  `About` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `FullName` varchar(250) NOT NULL DEFAULT '',
  `AvatarUrl` varchar(250) NOT NULL DEFAULT '',
  `PrivacyProfile` tinyint(10) NOT NULL DEFAULT '0',
  `Phone` varchar(250) NOT NULL DEFAULT '',
  `Email` varchar(250) NOT NULL DEFAULT '',
  `IsDeleted` int(10) NOT NULL DEFAULT '0',
  `PasswordHash` varchar(250) NOT NULL DEFAULT '',
  `PasswordSalt` varchar(250) NOT NULL DEFAULT '',
  `FollowersCount` int(10) NOT NULL DEFAULT '0',
  `FollowingCount` int(10) NOT NULL DEFAULT '0',
  `PostsCount` int(10) NOT NULL DEFAULT '0',
  `MediaCount` int(10) NOT NULL DEFAULT '0',
  `LikesCount` int(10) NOT NULL DEFAULT '0',
  `ResharedCount` int(10) NOT NULL DEFAULT '0',
  `LastActionTime` int(11) NOT NULL DEFAULT '0',
  `LastPostTime` int(11) NOT NULL DEFAULT '0',
  `PrimaryFollowingList` int(10) NOT NULL DEFAULT '0',
  `CreatedTime` int(10) NOT NULL DEFAULT '0',
  `UpdatedTime` int(10) NOT NULL DEFAULT '0',
  `SessionUuid` varchar(75) NOT NULL DEFAULT '',
  `DeviceUuid` varchar(75) NOT NULL DEFAULT '',
  `LastWifiMacAddress` varchar(75) NOT NULL DEFAULT '',
  `LastNetworkType` varchar(75) NOT NULL DEFAULT '',
  `AppVersion` int(11) NOT NULL DEFAULT '0',
  `LastActivityTime` int(11) NOT NULL DEFAULT '0',
  `LastLoginTime` int(10) NOT NULL DEFAULT '0',
  `LastIpAddress` varchar(75) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserName` (`UserName`),
  UNIQUE KEY `Email` (`Email`),
  KEY `SessionUuid` (`SessionUuid`),
  KEY `Phone` (`Phone`)
) ENGINE=MyISAM AUTO_INCREMENT=1084 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `user_meta_info` */

DROP TABLE IF EXISTS `user_meta_info`;

CREATE TABLE `user_meta_info` (
  `UserId` int(11) NOT NULL,
  `IsNotificationDirty` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `user_password` */

DROP TABLE IF EXISTS `user_password`;

CREATE TABLE `user_password` (
  `UserId` int(11) NOT NULL AUTO_INCREMENT,
  `Password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
