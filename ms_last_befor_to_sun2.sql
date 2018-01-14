/*
SQLyog Community v12.4.3 (32 bit)
MySQL - 5.7.19-log : Database - ms
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
  `ActionTypeId` int(11) NOT NULL DEFAULT '0',
  `RowId` int(11) NOT NULL DEFAULT '0',
  `RootId` int(11) NOT NULL DEFAULT '0',
  `RefId` bigint(20) NOT NULL DEFAULT '0',
  `CreatedAt` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `RefId` (`RefId`),
  KEY `ActorUserId` (`ActorUserId`,`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=20490 DEFAULT CHARSET=utf8;

/*Table structure for table `bucket` */

DROP TABLE IF EXISTS `bucket`;

CREATE TABLE `bucket` (
  `BucketId` int(11) NOT NULL AUTO_INCREMENT,
  `BucketName` varchar(100) NOT NULL,
  `Server1Id` int(11) NOT NULL DEFAULT '0',
  `Server2Id` int(11) NOT NULL DEFAULT '0',
  `BackupServerId` int(11) NOT NULL DEFAULT '0',
  `ContentObjectTypeId` smallint(6) NOT NULL DEFAULT '0',
  `ContentObjectCount` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`BucketId`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;

/*Table structure for table `chat` */

DROP TABLE IF EXISTS `chat`;

CREATE TABLE `chat` (
  `ChatKey` varchar(50) COLLATE utf8mb4_persian_ci NOT NULL,
  `RoomKey` varchar(50) COLLATE utf8mb4_persian_ci NOT NULL,
  `RoomTypeEnumId` smallint(6) NOT NULL,
  `UserId` int(11) NOT NULL,
  `PeerUserId` int(11) NOT NULL,
  `GroupId` bigint(20) NOT NULL,
  `CreatedSe` int(11) NOT NULL,
  `StartMessageIdFrom` bigint(20) NOT NULL,
  `LastDeletedMessageId` bigint(20) NOT NULL,
  `LastSeenMessageId` bigint(20) NOT NULL DEFAULT '0',
  `LastMessageId` bigint(20) NOT NULL DEFAULT '0' COMMENT 'just direct-not reliable- just increment not decremetnt',
  `UpdatedMs` bigint(20) NOT NULL,
  PRIMARY KEY (`ChatKey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `Id` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `PostId` int(20) NOT NULL,
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `PostId` (`PostId`)
) ENGINE=InnoDB AUTO_INCREMENT=9397 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `direct_message` */

DROP TABLE IF EXISTS `direct_message`;

CREATE TABLE `direct_message` (
  `MessageId` bigint(20) NOT NULL,
  `MessageKey` varchar(50) COLLATE utf8mb4_persian_ci NOT NULL,
  `RoomKey` varchar(70) COLLATE utf8mb4_persian_ci NOT NULL,
  `UserId` int(20) NOT NULL,
  `MessageFileId` bigint(20) NOT NULL,
  `MessageTypeEnumId` tinyint(4) NOT NULL,
  `Text` varchar(16000) COLLATE utf8mb4_persian_ci NOT NULL,
  `CreatedSe` int(11) NOT NULL,
  `PeerReceivedTime` int(11) NOT NULL,
  `PeerSeenTime` int(11) NOT NULL DEFAULT '0',
  `DeliviryStatusEnumId` tinyint(4) NOT NULL,
  PRIMARY KEY (`MessageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `direct_offline` */

DROP TABLE IF EXISTS `direct_offline`;

CREATE TABLE `direct_offline` (
  `DirectOfflineId` bigint(20) NOT NULL,
  `ToUserId` int(11) NOT NULL DEFAULT '0',
  `ChatKey` varchar(50) NOT NULL,
  `MessageId` bigint(20) NOT NULL,
  `MessageFileId` bigint(20) NOT NULL,
  `PBClass` varchar(250) NOT NULL,
  `DataPB` blob NOT NULL,
  `DataJson` text NOT NULL,
  `DataTemp` text NOT NULL COMMENT 'not real data',
  `AtTimeMs` bigint(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`DirectOfflineId`),
  KEY `ToUserId` (`ToUserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `direct_to_message` */

DROP TABLE IF EXISTS `direct_to_message`;

CREATE TABLE `direct_to_message` (
  `Id` bigint(20) NOT NULL,
  `ChatKey` varchar(50) COLLATE utf8mb4_persian_ci NOT NULL,
  `MessageId` bigint(20) NOT NULL,
  `SourceEnumId` tinyint(4) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `ChatKey` (`ChatKey`,`MessageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

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
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
) ENGINE=InnoDB AUTO_INCREMENT=429 DEFAULT CHARSET=utf8mb4;

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
) ENGINE=InnoDB AUTO_INCREMENT=13458 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `general_log` */

DROP TABLE IF EXISTS `general_log`;

CREATE TABLE `general_log` (
  `Id` bigint(20) NOT NULL,
  `ToUserId` int(11) NOT NULL,
  `TargetId` int(11) NOT NULL,
  `LogTypeId` int(11) NOT NULL,
  `ExtraPB` blob NOT NULL,
  `ExtraJson` varchar(5000) NOT NULL,
  `CreatedMs` bigint(20) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `group` */

DROP TABLE IF EXISTS `group`;

CREATE TABLE `group` (
  `GroupId` bigint(20) NOT NULL,
  `GroupName` varchar(200) COLLATE utf8mb4_persian_ci NOT NULL,
  `MembersCount` int(11) NOT NULL,
  `GroupPrivacyEnum` tinyint(4) NOT NULL,
  `CreatorUserId` int(20) NOT NULL,
  `CreatedTime` int(20) NOT NULL,
  `UpdatedMs` bigint(20) NOT NULL,
  `CurrentSeq` int(11) NOT NULL,
  PRIMARY KEY (`GroupId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `group_member` */

DROP TABLE IF EXISTS `group_member`;

CREATE TABLE `group_member` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `GroupId` bigint(20) NOT NULL,
  `GroupKey` varchar(100) COLLATE utf8mb4_persian_ci NOT NULL,
  `UserId` int(11) NOT NULL,
  `ByUserId` int(11) NOT NULL,
  `GroupRoleEnumId` tinyint(4) NOT NULL,
  `CreatedTime` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `Id` (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `group_message` */

DROP TABLE IF EXISTS `group_message`;

CREATE TABLE `group_message` (
  `MessageId` bigint(20) NOT NULL,
  `RoomKey` varchar(70) COLLATE utf8mb4_persian_ci NOT NULL,
  `UserId` int(20) NOT NULL,
  `MessageFileId` bigint(20) NOT NULL,
  `MessageTypeEnum` tinyint(4) NOT NULL,
  `Text` text COLLATE utf8mb4_persian_ci NOT NULL,
  `CreatedMs` bigint(20) NOT NULL,
  `DeliveryStatusEnum` tinyint(4) NOT NULL,
  PRIMARY KEY (`MessageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `group_to_message` */

DROP TABLE IF EXISTS `group_to_message`;

CREATE TABLE `group_to_message` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `GroupId` bigint(20) NOT NULL,
  `MessageId` bigint(20) NOT NULL,
  `CreatedMs` bigint(20) NOT NULL,
  `StatusEnum` tinyint(4) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

/*Table structure for table `likes` */

DROP TABLE IF EXISTS `likes`;

CREATE TABLE `likes` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `PostId` int(11) NOT NULL,
  `PostTypeId` tinyint(4) NOT NULL DEFAULT '0',
  `UserId` int(11) NOT NULL,
  `TypeId` tinyint(4) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `PostId` (`PostId`,`UserId`),
  KEY `Id` (`Id`),
  KEY `PostId_2` (`PostId`)
) ENGINE=InnoDB AUTO_INCREMENT=12192 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `log_changes` */

DROP TABLE IF EXISTS `log_changes`;

CREATE TABLE `log_changes` (
  `Id` int(11) NOT NULL,
  `T` varchar(50) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `message_file` */

DROP TABLE IF EXISTS `message_file`;

CREATE TABLE `message_file` (
  `MessageFileId` bigint(20) NOT NULL,
  `MessageFileKey` varchar(50) NOT NULL,
  `OriginalUserId` int(11) NOT NULL,
  `Name` varchar(150) NOT NULL,
  `Size` int(11) NOT NULL DEFAULT '0',
  `FileTypeEnumId` tinyint(4) NOT NULL DEFAULT '0',
  `Width` int(11) NOT NULL DEFAULT '0',
  `Height` int(11) NOT NULL DEFAULT '0',
  `Duration` int(11) NOT NULL DEFAULT '0',
  `Extension` varchar(50) NOT NULL DEFAULT '',
  `HashMd5` varchar(32) NOT NULL,
  `HashAccess` bigint(20) NOT NULL DEFAULT '0',
  `CreatedSe` int(11) NOT NULL,
  `ServerSrc` varchar(250) NOT NULL,
  `ServerPath` varchar(250) NOT NULL,
  `ServerThumbPath` varchar(250) NOT NULL,
  `BucketId` varchar(250) NOT NULL,
  `ServerId` int(11) NOT NULL DEFAULT '0',
  `CanDel` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`MessageFileId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `notification` */

DROP TABLE IF EXISTS `notification`;

CREATE TABLE `notification` (
  `Id` bigint(10) NOT NULL AUTO_INCREMENT,
  `ForUserId` int(10) NOT NULL DEFAULT '0',
  `ActorUserId` int(10) NOT NULL DEFAULT '0',
  `ActionTypeId` int(10) NOT NULL DEFAULT '0',
  `ObjectTypeId` int(10) NOT NULL DEFAULT '0',
  `RowId` int(10) NOT NULL DEFAULT '0',
  `RootId` int(10) NOT NULL DEFAULT '0',
  `RefId` int(10) NOT NULL DEFAULT '0',
  `SeenStatus` int(10) NOT NULL DEFAULT '0',
  `CreatedTime` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `ForUserId` (`ForUserId`,`Id`),
  KEY `TargetId` (`RowId`)
) ENGINE=InnoDB AUTO_INCREMENT=20243 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `notification_removed` */

DROP TABLE IF EXISTS `notification_removed`;

CREATE TABLE `notification_removed` (
  `NotificationId` int(11) NOT NULL,
  `ForUserId` int(11) NOT NULL,
  PRIMARY KEY (`NotificationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `offline` */

DROP TABLE IF EXISTS `offline`;

CREATE TABLE `offline` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `FromUserId` int(11) NOT NULL,
  `ToUserId` int(11) NOT NULL,
  `RpcName` varchar(200) NOT NULL,
  `PBClassName` varchar(200) NOT NULL,
  `Key` varchar(200) NOT NULL,
  `DataJson` varchar(6000) NOT NULL,
  `DataBlob` blob NOT NULL,
  `DataLength` int(11) NOT NULL,
  `CreatedMs` bigint(20) NOT NULL,
  PRIMARY KEY (`Id`)
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
) ENGINE=InnoDB AUTO_INCREMENT=60517 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `photo` */

DROP TABLE IF EXISTS `photo`;

CREATE TABLE `photo` (
  `PhotoId` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL DEFAULT '0',
  `PostId` int(11) NOT NULL DEFAULT '0',
  `AlbumId` int(11) NOT NULL DEFAULT '0',
  `ImageTypeId` smallint(4) NOT NULL DEFAULT '0',
  `Title` varchar(300) NOT NULL DEFAULT '',
  `Src` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'json of sizes',
  `PathSrc` varchar(300) NOT NULL DEFAULT '',
  `BucketId` int(11) NOT NULL DEFAULT '0',
  `Width` smallint(6) NOT NULL DEFAULT '0',
  `Height` smallint(6) NOT NULL DEFAULT '0',
  `Ratio` float NOT NULL DEFAULT '0',
  `HashMd5` char(32) NOT NULL DEFAULT '',
  `Color` varchar(10) NOT NULL DEFAULT '',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  `W1080` tinyint(1) NOT NULL DEFAULT '0',
  `W720` tinyint(1) NOT NULL DEFAULT '0',
  `W480` tinyint(1) NOT NULL DEFAULT '0',
  `W320` tinyint(1) NOT NULL DEFAULT '0',
  `W160` tinyint(1) NOT NULL DEFAULT '0',
  `W80` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'nojson',
  PRIMARY KEY (`PhotoId`),
  KEY `HashMd5` (`HashMd5`),
  KEY `CreatedTime` (`CreatedTime`),
  KEY `AlbumId` (`AlbumId`),
  KEY `PostId2` (`PostId`)
) ENGINE=InnoDB AUTO_INCREMENT=829 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `post` */

DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
  `Id` int(20) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `TypeId` int(4) NOT NULL DEFAULT '1',
  `Text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `FormatedText` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `MediaCount` tinyint(4) NOT NULL DEFAULT '0',
  `SharedTo` int(11) NOT NULL DEFAULT '0',
  `DisableComment` tinyint(1) NOT NULL DEFAULT '0',
  `HasTag` int(11) NOT NULL DEFAULT '0',
  `LikesCount` int(11) NOT NULL DEFAULT '0',
  `CommentsCount` int(11) NOT NULL DEFAULT '0',
  `EditedTime` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=1203 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `push_event` */

DROP TABLE IF EXISTS `push_event`;

CREATE TABLE `push_event` (
  `PushEventId` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(11) NOT NULL DEFAULT '0',
  `ToDeviceId` bigint(11) NOT NULL,
  `MessageId` bigint(20) NOT NULL,
  `RoomTypeEnum` smallint(6) NOT NULL,
  `RoomId` bigint(20) NOT NULL,
  `PeerUserId` int(11) NOT NULL DEFAULT '0',
  `PushEventTypeEnum` tinyint(4) NOT NULL DEFAULT '0',
  `AtTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`PushEventId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `push_message` */

DROP TABLE IF EXISTS `push_message`;

CREATE TABLE `push_message` (
  `PushMessageId` bigint(20) NOT NULL AUTO_INCREMENT,
  `ToUserId` int(20) NOT NULL,
  `ToDeviceId` bigint(20) NOT NULL,
  `MessageId` bigint(20) NOT NULL,
  `RoomTypeEnum` tinyint(4) NOT NULL,
  `CreatedMs` bigint(20) NOT NULL,
  PRIMARY KEY (`PushMessageId`),
  KEY `ToUser` (`ToUserId`,`CreatedMs`),
  KEY `Uid` (`PushMessageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `recommend_user` */

DROP TABLE IF EXISTS `recommend_user`;

CREATE TABLE `recommend_user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL DEFAULT '0',
  `TargetId` int(11) NOT NULL DEFAULT '0',
  `Weight` float NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=5365 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `room` */

DROP TABLE IF EXISTS `room`;

CREATE TABLE `room` (
  `RoomId` bigint(20) NOT NULL,
  `RoomKey` varchar(50) COLLATE utf8mb4_persian_ci NOT NULL,
  `RoomTypeEnum` smallint(6) NOT NULL,
  `UserId` int(11) NOT NULL,
  `LastSeqSeen` int(11) NOT NULL,
  `LastSeqDelete` int(11) NOT NULL,
  `PeerUserId` int(11) NOT NULL,
  `GroupId` bigint(20) NOT NULL,
  `CreatedTime` int(11) NOT NULL,
  `CurrentSeq` int(11) NOT NULL COMMENT 'just for peer to peer',
  PRIMARY KEY (`RoomId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_persian_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
  `LastNetworkTypeId` smallint(6) NOT NULL DEFAULT '0',
  `AppVersion` smallint(6) NOT NULL DEFAULT '0',
  `UpdatedTime` int(11) NOT NULL DEFAULT '0',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  `LastSyncDirectUpdateId` bigint(20) NOT NULL DEFAULT '0',
  `LastSyncGeneralUpdateId` bigint(20) NOT NULL DEFAULT '0',
  `LastSyncNotifyUpdateId` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `SessionUuid2` (`SessionUuid`),
  KEY `Id` (`Id`),
  KEY `UserId` (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `setting_client` */

DROP TABLE IF EXISTS `setting_client`;

CREATE TABLE `setting_client` (
  `UserId` int(11) NOT NULL,
  `AutoDownloadWifiVoice` int(11) NOT NULL,
  `AutoDownloadWifiImage` int(11) NOT NULL,
  `AutoDownloadWifiVideo` int(11) NOT NULL,
  `AutoDownloadWifiFile` int(11) NOT NULL,
  `AutoDownloadWifiMusic` int(11) NOT NULL,
  `AutoDownloadWifiGif` int(11) NOT NULL,
  `AutoDownloadCellularVoice` int(11) NOT NULL,
  `AutoDownloadCellularImage` int(11) NOT NULL,
  `AutoDownloadCellularVideo` int(11) NOT NULL,
  `AutoDownloadCellularFile` int(11) NOT NULL,
  `AutoDownloadCellularMusic` int(11) NOT NULL,
  `AutoDownloadCellularGif` int(11) NOT NULL,
  `AutoDownloadRoamingVoice` int(11) NOT NULL,
  `AutoDownloadRoamingImage` int(11) NOT NULL,
  `AutoDownloadRoamingVideo` int(11) NOT NULL,
  `AutoDownloadRoamingFile` int(11) NOT NULL,
  `AutoDownloadRoamingMusic` int(11) NOT NULL,
  `AutoDownloadRoamingGif` int(11) NOT NULL,
  `SaveToGallery` int(11) NOT NULL,
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `setting_notifications` */

DROP TABLE IF EXISTS `setting_notifications`;

CREATE TABLE `setting_notifications` (
  `UserId` int(11) NOT NULL,
  `SocialLedOn` int(11) NOT NULL,
  `SocialLedColor` varchar(30) NOT NULL,
  `ReqestToFollowYou` int(11) NOT NULL,
  `FollowedYou` int(11) NOT NULL,
  `AccptedYourFollowRequest` int(11) NOT NULL,
  `YourPostLiked` int(11) NOT NULL,
  `YourPostCommented` int(11) NOT NULL,
  `MenthenedYouInPost` int(11) NOT NULL,
  `MenthenedYouInComment` int(11) NOT NULL,
  `YourContactsJoined` int(11) NOT NULL,
  `DirectMessage` int(11) NOT NULL,
  `DirectAlert` int(11) NOT NULL,
  `DirectPerview` int(11) NOT NULL,
  `DirectLedOn` int(11) NOT NULL,
  `DirectLedColor` int(11) NOT NULL,
  `DirectVibrate` int(11) NOT NULL,
  `DirectPopup` int(11) NOT NULL,
  `DirectSound` int(11) NOT NULL,
  `DirectPriority` int(11) NOT NULL,
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tag` */

DROP TABLE IF EXISTS `tag`;

CREATE TABLE `tag` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(100) NOT NULL,
  `Count` int(11) NOT NULL,
  `IsBlocked` int(1) NOT NULL,
  `CreatedTime` int(11) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Name` (`Name`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 DELAY_KEY_WRITE=1;

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
) ENGINE=InnoDB AUTO_INCREMENT=9425 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `test_chat` */

DROP TABLE IF EXISTS `test_chat`;

CREATE TABLE `test_chat` (
  `Id` bigint(11) NOT NULL,
  `Id4` bigint(20) NOT NULL AUTO_INCREMENT,
  `TimeMs` bigint(20) NOT NULL,
  `Text` text NOT NULL,
  `Name` varchar(100) NOT NULL,
  `UserId` bigint(20) NOT NULL,
  `C2` int(11) NOT NULL,
  `C3` int(11) NOT NULL,
  `C4` int(11) NOT NULL,
  `C5` int(11) NOT NULL,
  PRIMARY KEY (`Id4`),
  KEY `UserId` (`UserId`,`TimeMs`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `trigger_log` */

DROP TABLE IF EXISTS `trigger_log`;

CREATE TABLE `trigger_log` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ModelName` varchar(50) NOT NULL DEFAULT '',
  `ChangeType` varchar(50) NOT NULL DEFAULT '',
  `TargetId` bigint(20) NOT NULL DEFAULT '0',
  `TargetStr` varchar(100) NOT NULL DEFAULT '',
  `CreatedSe` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `Id` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=5732 DEFAULT CHARSET=utf8;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `Id` int(10) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(250) NOT NULL DEFAULT '',
  `UserNameLower` varchar(250) NOT NULL DEFAULT '' COMMENT 'nojson',
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
  KEY `Phone` (`Phone`),
  KEY `Email_2` (`Email`,`IsDeleted`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `user_meta_info` */

DROP TABLE IF EXISTS `user_meta_info`;

CREATE TABLE `user_meta_info` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) NOT NULL,
  `IsNotificationDirty` tinyint(4) NOT NULL DEFAULT '0',
  `LastUserRecGen` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId2` (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Table structure for table `user_password` */

DROP TABLE IF EXISTS `user_password`;

CREATE TABLE `user_password` (
  `UserId` int(11) NOT NULL AUTO_INCREMENT,
  `Password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `CreatedTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/* Trigger structure for table `chat` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `chat_OnCreateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `chat_OnCreateLogger` AFTER INSERT ON `chat` FOR EACH ROW 
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","INSERT",NEW.ChatKey, UNIX_TIMESTAMP(NOW()) );
  END */$$


DELIMITER ;

/* Trigger structure for table `chat` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `chat_OnUpdateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `chat_OnUpdateLogger` AFTER UPDATE ON `chat` FOR EACH ROW 
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","UPDATE",NEW.ChatKey, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `chat` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `chat_OnDeleteLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `chat_OnDeleteLogger` AFTER DELETE ON `chat` FOR EACH ROW 
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","DELETE",OLD.ChatKey, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `comment` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `comment_OnCreateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `comment_OnCreateLogger` AFTER INSERT ON `comment` FOR EACH ROW 
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END */$$


DELIMITER ;

/* Trigger structure for table `comment` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `comment_OnUpdateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `comment_OnUpdateLogger` AFTER UPDATE ON `comment` FOR EACH ROW 
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `comment` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `comment_OnDeleteLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `comment_OnDeleteLogger` AFTER DELETE ON `comment` FOR EACH ROW 
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `post` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `post_OnCreateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `post_OnCreateLogger` AFTER INSERT ON `post` FOR EACH ROW 
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END */$$


DELIMITER ;

/* Trigger structure for table `post` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `post_OnUpdateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `post_OnUpdateLogger` AFTER UPDATE ON `post` FOR EACH ROW 
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `post` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `post_OnDeleteLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `post_OnDeleteLogger` AFTER DELETE ON `post` FOR EACH ROW 
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `tag` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `tags_OnCreateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `tags_OnCreateLogger` AFTER INSERT ON `tag` FOR EACH ROW 
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tags","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END */$$


DELIMITER ;

/* Trigger structure for table `tag` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `tags_OnUpdateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `tags_OnUpdateLogger` AFTER UPDATE ON `tag` FOR EACH ROW 
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tags","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `tag` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `tags_OnDeleteLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `tags_OnDeleteLogger` AFTER DELETE ON `tag` FOR EACH ROW 
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tags","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `user` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `user_OnCreateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `user_OnCreateLogger` AFTER INSERT ON `user` FOR EACH ROW 
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END */$$


DELIMITER ;

/* Trigger structure for table `user` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `user_OnUpdateLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `user_OnUpdateLogger` AFTER UPDATE ON `user` FOR EACH ROW 
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/* Trigger structure for table `user` */

DELIMITER $$

/*!50003 DROP TRIGGER*//*!50032 IF EXISTS */ /*!50003 `user_OnDeleteLogger` */$$

/*!50003 CREATE */ /*!50017 DEFINER = 'root'@'localhost' */ /*!50003 TRIGGER `user_OnDeleteLogger` AFTER DELETE ON `user` FOR EACH ROW 
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END */$$


DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
