/*
SQLyog Community v12.2.4 (64 bit)
MySQL - 10.1.12-MariaDB : Database - ms5
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ms5` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ms5`;

/*Table structure for table `_chat_room_list2` */

DROP TABLE IF EXISTS `_chat_room_list2`;

CREATE TABLE `_chat_room_list2` (
  `Id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `UpdatedTimestampMs` int(11) DEFAULT NULL,
  `CreatedTimestamp` int(11) DEFAULT NULL,
  `AvatarPath` varchar(50) DEFAULT NULL,
  `BackgroundId` int(11) DEFAULT NULL,
  `BackgroundPath` varchar(50) DEFAULT NULL,
  `LastSeenMessageType` varchar(50) DEFAULT NULL,
  `LastSeenMessagevarchar` int(50) DEFAULT NULL,
  `LastSeenMessageCount` int(11) DEFAULT NULL,
  `LastSeenMessageId` int(11) DEFAULT NULL,
  `CreatedTimestampMS` int(11) DEFAULT NULL,
  `LastVisitedTimestamp` int(11) DEFAULT NULL,
  `RoomCode` varchar(50) DEFAULT NULL,
  `RoomTypeId` int(11) DEFAULT NULL,
  `OnlineStatus` varchar(50) DEFAULT NULL,
  `OnlineStatusId` int(11) DEFAULT NULL,
  `AvatarUrl` varchar(500) DEFAULT NULL,
  `RoomName` varchar(50) DEFAULT NULL,
  `RoomType` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `abuse_report` */

DROP TABLE IF EXISTS `abuse_report`;

CREATE TABLE `abuse_report` (
  `Id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `ContentType` mediumint(9) DEFAULT NULL,
  `ContentId` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci,
  `ActorUserId` int(11) DEFAULT NULL,
  `AbuseType` mediumint(9) DEFAULT NULL,
  `Comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci,
  `IpAdress` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `activity` */

DROP TABLE IF EXISTS `activity`;

CREATE TABLE `activity` (
  `Id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ActorUserId` int(11) DEFAULT NULL,
  `ActivityType` tinyint(4) DEFAULT '0',
  `TargetActionId` bigint(20) DEFAULT '0',
  `Privacy` int(11) DEFAULT '0',
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `blocked` */

DROP TABLE IF EXISTS `blocked`;

CREATE TABLE `blocked` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `BlockedUserID` int(11) DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`UserId`),
  KEY `Id` (`Id`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `chat_group_info` */

DROP TABLE IF EXISTS `chat_group_info`;

CREATE TABLE `chat_group_info` (
  `RoomKey` varchar(250) DEFAULT '',
  `RoomName` varchar(250) DEFAULT '',
  `MembersCount` int(10) DEFAULT '0',
  `GroupPrivacyId` int(10) DEFAULT '0',
  `CreatorUserId` int(10) DEFAULT '0',
  `CreatedTimestamp` int(10) DEFAULT '0',
  `UpdatedTimestamp` int(10) DEFAULT '0',
  UNIQUE KEY `RoomKey` (`RoomKey`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `chat_group_member` */

DROP TABLE IF EXISTS `chat_group_member`;

CREATE TABLE `chat_group_member` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `RoomKey` varchar(50) DEFAULT '',
  `UserId` int(10) DEFAULT '0',
  `ByUserId` int(10) DEFAULT '0',
  `RoleId` int(10) DEFAULT '0',
  `IsAdmin` int(10) DEFAULT '0',
  `CreatedTimestamp` int(10) DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `RoomKey` (`RoomKey`,`UserId`)
) ENGINE=Aria AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `chat_message` */

DROP TABLE IF EXISTS `chat_message`;

CREATE TABLE `chat_message` (
  `MessageKey` varchar(250) DEFAULT '',
  `RoomKey` varchar(250) DEFAULT '',
  `UserId` int(10) DEFAULT '0',
  `MessageTypeId` int(10) DEFAULT '0',
  `Text` text,
  `MediaUrl` varchar(250) DEFAULT '',
  `ExtraJson` varchar(250) DEFAULT '',
  `CreatedTimestampMs` int(10) DEFAULT '0',
  KEY `RoomKey` (`RoomKey`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `comments` */

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `PostId` bigint(20) DEFAULT NULL,
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci,
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `PostId` (`PostId`)
) ENGINE=Aria AUTO_INCREMENT=99852 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `following_list` */

DROP TABLE IF EXISTS `following_list`;

CREATE TABLE `following_list` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `ListType` int(11) DEFAULT NULL,
  `Name` varchar(75) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci DEFAULT NULL,
  `Count` int(11) DEFAULT '0',
  `IsAuto` tinyint(1) DEFAULT NULL,
  `IsPimiry` tinyint(1) DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`Id`,`UserId`)
) ENGINE=Aria AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `following_list_member` */

DROP TABLE IF EXISTS `following_list_member`;

CREATE TABLE `following_list_member` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ListId` int(11) NOT NULL,
  `UserId` int(11) DEFAULT NULL,
  `FollowedUserId` int(11) DEFAULT NULL,
  `FollowType` tinyint(4) DEFAULT '1',
  `UpdatedTimeMs` bigint(11) DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId` (`UserId`,`FollowedUserId`),
  KEY `Time` (`UserId`,`UpdatedTimeMs`)
) ENGINE=Aria AUTO_INCREMENT=6439 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `likes` */

DROP TABLE IF EXISTS `likes`;

CREATE TABLE `likes` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `PostId` bigint(20) NOT NULL,
  `UserId` int(11) NOT NULL,
  `TypeId` tinyint(4) DEFAULT '0',
  `CreatedTime` int(11) DEFAULT '0',
  UNIQUE KEY `PostId` (`PostId`,`UserId`),
  KEY `Id` (`Id`),
  KEY `PostId_2` (`PostId`)
) ENGINE=Aria AUTO_INCREMENT=5003 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `media` */

DROP TABLE IF EXISTS `media`;

CREATE TABLE `media` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `PostId` int(11) DEFAULT NULL,
  `AlbumId` int(11) DEFAULT '0',
  `TypeId` smallint(4) DEFAULT '0',
  `CreatedTime` int(11) DEFAULT '0',
  `Src` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci COMMENT 'json of sizes',
  PRIMARY KEY (`Id`),
  KEY `Id` (`Id`)
) ENGINE=Aria AUTO_INCREMENT=272 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `phone_contacts` */

DROP TABLE IF EXISTS `phone_contacts`;

CREATE TABLE `phone_contacts` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `PhoneDisplayName` varchar(250) DEFAULT '',
  `PhoneFamilyName` varchar(250) DEFAULT '',
  `PhoneNumber` varchar(250) DEFAULT '',
  `PhoneNormalizedNumber` varchar(250) DEFAULT '',
  `PhoneContactRowId` int(10) DEFAULT '0',
  `UserId` int(10) DEFAULT '0',
  `DeviceUuidId` int(10) DEFAULT '0',
  `CreatedTimeStamp` int(10) DEFAULT '0',
  `UpdatedTimeStamp` int(10) DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=Aria AUTO_INCREMENT=4399 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `post` */

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `Id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `TypeId` smallint(4) NOT NULL DEFAULT '1',
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci,
  `FormatedText` text CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci,
  `MediaUrl` varchar(250) DEFAULT '""',
  `MediaServerId` int(11) DEFAULT '0',
  `Width` int(11) DEFAULT '0',
  `Height` int(11) DEFAULT '0',
  `SharedTo` int(11) NOT NULL DEFAULT '0',
  `HasTag` blob NOT NULL,
  `LikesCount` int(11) NOT NULL DEFAULT '0',
  `CommentsCount` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=Aria AUTO_INCREMENT=2637 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `session` */

DROP TABLE IF EXISTS `session`;

CREATE TABLE `session` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `SessionUuid` varchar(75) CHARACTER SET utf8 NOT NULL,
  `ClientUuid` varchar(75) CHARACTER SET utf8 DEFAULT NULL,
  `DeviceUuid` varchar(75) CHARACTER SET utf8 DEFAULT NULL,
  `LastActivityTime` int(11) DEFAULT '0',
  `LastIpAddress` varchar(75) CHARACTER SET utf8 DEFAULT NULL,
  `LastWifiMacAddress` varchar(75) CHARACTER SET utf8 DEFAULT NULL,
  `LastNetworkType` varchar(75) CHARACTER SET utf8 DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`SessionUuid`),
  KEY `Id` (`Id`)
) ENGINE=Aria AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `tags` */

DROP TABLE IF EXISTS `tags`;

CREATE TABLE `tags` (
  `Id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) DEFAULT NULL,
  `Count` int(11) DEFAULT NULL,
  `IsBlocked` tinyint(1) DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=Aria AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `tags_posts` */

DROP TABLE IF EXISTS `tags_posts`;

CREATE TABLE `tags_posts` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `TagId` int(11) DEFAULT NULL,
  `PostId` int(11) DEFAULT NULL,
  `TypeId` int(11) DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=Aria AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `Id` int(10) NOT NULL,
  `UserName` varchar(250) DEFAULT '',
  `FirstName` varchar(250) DEFAULT '',
  `LastName` varchar(250) DEFAULT '',
  `About` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci DEFAULT '',
  `FullName` varchar(250) DEFAULT '',
  `AvatarSrc` varchar(250) DEFAULT '',
  `AvatarUrl` varchar(250) DEFAULT '',
  `IsProfilePrivate` int(10) DEFAULT '0',
  `Phone` varchar(250) DEFAULT '',
  `Email` varchar(250) DEFAULT '',
  `IsDeleted` int(10) DEFAULT '0',
  `PasswordHash` varchar(250) DEFAULT '',
  `PasswordSalt` varchar(250) DEFAULT '',
  `FollowersCount` int(10) DEFAULT '0',
  `FollowingCount` int(10) DEFAULT '0',
  `PostsCount` int(10) DEFAULT '0',
  `MediaCount` int(10) DEFAULT '0',
  `LikesCount` int(10) DEFAULT '0',
  `ResharedCount` int(10) DEFAULT '0',
  `LastActionTime` int(11) DEFAULT '0',
  `LastPostTime` int(11) DEFAULT '0',
  `PrimaryFollowingList` int(10) DEFAULT '0',
  `CreatedTime` int(10) DEFAULT '0',
  `UpdatedTime` int(10) DEFAULT '0',
  `SessionUuid` varchar(75) DEFAULT '',
  `DeviceUuid` varchar(75) DEFAULT '',
  `LastWifiMacAddress` varchar(75) DEFAULT '',
  `LastNetworkType` varchar(75) DEFAULT '',
  `AppVersion` int(11) DEFAULT '0',
  `LastActivityTime` int(11) DEFAULT '0',
  `LastLoginTime` int(10) DEFAULT '0',
  `LastIpAddress` varchar(75) DEFAULT '',
  PRIMARY KEY (`Id`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `user_info` */

DROP TABLE IF EXISTS `user_info`;

CREATE TABLE `user_info` (
  `UserId` int(10) unsigned NOT NULL,
  `FollowersCount` int(11) DEFAULT '0',
  `FollowingCount` int(11) DEFAULT '0',
  `PostsCount` int(11) DEFAULT '0',
  `MediaCount` int(11) DEFAULT '0',
  `LikesCount` int(11) DEFAULT '0',
  `ResharedCount` int(11) DEFAULT '0',
  `LastLoginTimestamp` int(11) DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*Table structure for table `user_passwod` */

DROP TABLE IF EXISTS `user_passwod`;

CREATE TABLE `user_passwod` (
  `UserId` int(11) NOT NULL,
  `Password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci DEFAULT NULL,
  `CreatedTime` int(11) DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=Aria DEFAULT CHARSET=utf8mb4 PAGE_CHECKSUM=1;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
