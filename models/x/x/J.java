package com.mardomsara.social.json;

public class J {
	public static class Activity {
		public Integer Id;
		public Integer ActorUserId;
		public Integer ActionTypeId;
		public Integer RowId;
		public Integer RootId;
		public Integer RefId;
		public Integer CreatedAt;
	}

	public static class Bucket {
		public Integer BucketId;
		public String BucketName;
		public Integer Server1Id;
		public Integer Server2Id;
		public Integer BackupServerId;
		public Integer ContentObjectTypeId;
		public Integer ContentObjectCount;
		public Integer CreatedTime;
	}

	public static class Chat {
		public Integer ChatId;
		public String ChatKey;
		public Integer RoomTypeEnumId;
		public Integer UserId;
		public Integer LastSeqSeen;
		public Integer LastSeqDelete;
		public Integer PeerUserId;
		public Integer GroupId;
		public Integer CreatedTime;
		public Integer CurrentSeq;
		public Integer UpdatedMs;
	}

	public static class Comment {
		public Integer Id;
		public Integer UserId;
		public Integer PostId;
		public String Text;
		public Integer CreatedTime;
	}

	public static class DirectLog {
		public Integer Id;
		public Integer ToUserId;
		public Integer MessageId;
		public Integer ChatId;
		public Integer PeerUserId;
		public Integer EventType;
		public Integer RoomLogTypeId;
		public Integer FromSeq;
		public Integer ToSeq;
		public UNKNOWN ExtraPB;
		public String ExtraJson;
		public Integer AtTimeMs;
	}

	public static class DirectLog2 {
		public Integer Id;
		public UNKNOWN ToUserId;
		public UNKNOWN ChatId;
		public UNKNOWN LogTypeEnumId;
		public UNKNOWN DataPB;
		public UNKNOWN DataJson;
		public UNKNOWN Created;
	}

	public static class DirectMessage {
		public Integer MessageId;
		public String RoomKey;
		public Integer UserId;
		public Integer MessageFileId;
		public Integer MessageTypeEnum;
		public String Text;
		public Integer Time;
		public Integer PeerReceivedTime;
		public Integer PeerSeenTime;
		public Integer DeliviryStatusEnum;
	}

	public static class DirectToMessage {
		public Integer Id;
		public Integer ChatId;
		public Integer MessageId;
		public Integer Seq;
		public Integer SourceEnum;
	}

	public static class FollowingList {
		public Integer Id;
		public Integer UserId;
		public Integer ListType;
		public String Name;
		public Integer Count;
		public Integer IsAuto;
		public Integer IsPimiry;
		public Integer CreatedTime;
	}

	public static class FollowingListMember {
		public Integer Id;
		public Integer ListId;
		public Integer UserId;
		public Integer FollowedUserId;
		public Integer FollowType;
		public Integer UpdatedTimeMs;
	}

	public static class FollowingListMemberHistory {
		public Integer Id;
		public Integer ListId;
		public Integer UserId;
		public Integer FollowedUserId;
		public Integer FollowType;
		public Integer UpdatedTimeMs;
		public Integer FollowId;
	}

	public static class Group {
		public Integer GroupId;
		public String GroupName;
		public Integer MembersCount;
		public Integer GroupPrivacyEnum;
		public Integer CreatorUserId;
		public Integer CreatedTime;
		public Integer UpdatedMs;
		public Integer CurrentSeq;
	}

	public static class GroupMember {
		public Integer Id;
		public Integer GroupId;
		public String GroupKey;
		public Integer UserId;
		public Integer ByUserId;
		public Integer GroupRoleEnum;
		public Integer CreatedTime;
	}

	public static class GroupMessage {
		public Integer MessageId;
		public String RoomKey;
		public Integer UserId;
		public Integer MessageFileId;
		public Integer MessageTypeEnum;
		public String Text;
		public Integer CreatedMs;
		public Integer DeliveryStatusEnum;
	}

	public static class GroupToMessage {
		public Integer Id;
		public Integer GroupId;
		public Integer MessageId;
		public Integer CreatedMs;
		public Integer StatusEnum;
	}

	public static class Like {
		public Integer Id;
		public Integer PostId;
		public Integer PostTypeId;
		public Integer UserId;
		public Integer TypeId;
		public Integer CreatedTime;
	}

	public static class LogChange {
		public Integer Id;
		public String T;
	}

	public static class Media {
		public Integer Id;
		public Integer UserId;
		public Integer PostId;
		public Integer AlbumId;
		public Integer TypeId;
		public Integer CreatedTime;
		public String Src;
	}

	public static class MessageFile {
		public Integer MessageFileId;
		public String Name;
		public Integer Size;
		public Integer FileTypeEnum;
		public String MimeType;
		public Integer Width;
		public Integer Height;
		public Integer Duration;
		public String Extension;
		public UNKNOWN ThumbData;
		public String ThumbData64;
		public String ServerSrc;
		public String ServerPath;
		public String ServerThumbPath;
		public String BucketId;
		public Integer ServerId;
		public Integer CanDel;
		public Integer CreatedTime;
	}

	public static class Notification {
		public Integer Id;
		public Integer ForUserId;
		public Integer ActorUserId;
		public Integer ActionTypeId;
		public Integer ObjectTypeId;
		public Integer RowId;
		public Integer RootId;
		public Integer RefId;
		public Integer SeenStatus;
		public Integer CreatedTime;
	}

	public static class NotificationRemoved {
		public Integer NotificationId;
		public Integer ForUserId;
	}

	public static class OldMessage {
		public Integer Id;
		public Integer Uid;
		public Integer UserId;
		public String MessageKey;
		public String RoomKey;
		public Integer MessageType;
		public Integer RoomType;
		public Integer MsgFileId;
		public UNKNOWN DataPB;
		public String Data64;
		public String DataJson;
		public Integer CreatedTimeMs;
	}

	public static class OldMsgFile {
		public Integer Id;
		public String Name;
		public Integer Size;
		public Integer FileType;
		public String MimeType;
		public Integer Width;
		public Integer Height;
		public Integer Duration;
		public String Extension;
		public UNKNOWN ThumbData;
		public String ThumbData64;
		public String ServerSrc;
		public String ServerPath;
		public Integer ServerId;
		public Integer CanDel;
	}

	public static class OldMsgPush {
		public Integer Id;
		public UNKNOWN Uid;
		public Integer ToUser;
		public Integer MsgUid;
		public Integer CreatedTimeMs;
	}

	public static class OldMsgPushEvent {
		public Integer Id;
		public Integer Uid;
		public Integer ToUserId;
		public Integer MsgUid;
		public String MsgKey;
		public String RoomKey;
		public Integer PeerUserId;
		public Integer EventType;
		public Integer AtTime;
	}

	public static class PhoneContact {
		public Integer Id;
		public String PhoneDisplayName;
		public String PhoneFamilyName;
		public String PhoneNumber;
		public String PhoneNormalizedNumber;
		public Integer PhoneContactRowId;
		public Integer UserId;
		public Integer DeviceUuidId;
		public Integer CreatedTime;
		public Integer UpdatedTime;
	}

	public static class Photo {
		public Integer PhotoId;
		public Integer UserId;
		public Integer PostId;
		public Integer AlbumId;
		public Integer ImageTypeId;
		public String Title;
		public String Src;
		public String PathSrc;
		public Integer BucketId;
		public Integer Width;
		public Integer Height;
		public Float Ratio;
		public String HashMd5;
		public String Color;
		public Integer CreatedTime;
		public Integer W1080;
		public Integer W720;
		public Integer W480;
		public Integer W320;
		public Integer W160;
		public Integer W80;
	}

	public static class Post {
		public Integer Id;
		public Integer UserId;
		public Integer TypeId;
		public String Text;
		public String FormatedText;
		public Integer MediaCount;
		public Integer SharedTo;
		public Integer DisableComment;
		public Integer HasTag;
		public Integer LikesCount;
		public Integer CommentsCount;
		public Integer EditedTime;
		public Integer CreatedTime;
	}

	public static class PushEvent {
		public Integer PushEventId;
		public Integer ToUserId;
		public Integer ToDeviceId;
		public Integer MessageId;
		public Integer RoomTypeEnum;
		public Integer RoomId;
		public Integer PeerUserId;
		public Integer PushEventTypeEnum;
		public Integer AtTime;
	}

	public static class PushMessage {
		public Integer PushMessageId;
		public Integer ToUserId;
		public Integer ToDeviceId;
		public Integer MessageId;
		public Integer RoomTypeEnum;
		public Integer CreatedMs;
	}

	public static class RecommendUser {
		public Integer Id;
		public Integer UserId;
		public Integer TargetId;
		public Float Weight;
		public Integer CreatedTime;
	}

	public static class Room {
		public Integer RoomId;
		public String RoomKey;
		public Integer RoomTypeEnum;
		public Integer UserId;
		public Integer LastSeqSeen;
		public Integer LastSeqDelete;
		public Integer PeerUserId;
		public Integer GroupId;
		public Integer CreatedTime;
		public Integer CurrentSeq;
	}

	public static class SearchClicked {
		public Integer Id;
		public String Query;
		public Integer ClickType;
		public Integer TargetId;
		public Integer UserId;
		public Integer CreatedAt;
	}

	public static class Session {
		public Integer Id;
		public Integer UserId;
		public String SessionUuid;
		public String ClientUuid;
		public String DeviceUuid;
		public Integer LastActivityTime;
		public String LastIpAddress;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public Integer LastNetworkTypeId;
		public Integer AppVersion;
		public Integer UpdatedTime;
		public Integer CreatedTime;
	}

	public static class Tag {
		public Integer Id;
		public String Name;
		public Integer Count;
		public Integer IsBlocked;
		public Integer CreatedTime;
	}

	public static class TagsPost {
		public Integer Id;
		public Integer TagId;
		public Integer PostId;
		public Integer TypeId;
		public Integer CreatedTime;
	}

	public static class TestChat {
		public Integer Id;
		public Integer Id4;
		public Integer TimeMs;
		public String Text;
		public String Name;
		public Integer UserId;
		public Integer C2;
		public Integer C3;
		public Integer C4;
		public Integer C5;
	}

	public static class User {
		public Integer Id;
		public String UserName;
		public String UserNameLower;
		public String FirstName;
		public String LastName;
		public String About;
		public String FullName;
		public String AvatarUrl;
		public Integer PrivacyProfile;
		public String Phone;
		public String Email;
		public Integer IsDeleted;
		public String PasswordHash;
		public String PasswordSalt;
		public Integer FollowersCount;
		public Integer FollowingCount;
		public Integer PostsCount;
		public Integer MediaCount;
		public Integer LikesCount;
		public Integer ResharedCount;
		public Integer LastActionTime;
		public Integer LastPostTime;
		public Integer PrimaryFollowingList;
		public Integer CreatedTime;
		public Integer UpdatedTime;
		public String SessionUuid;
		public String DeviceUuid;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public Integer AppVersion;
		public Integer LastActivityTime;
		public Integer LastLoginTime;
		public String LastIpAddress;
	}

	public static class UserMetaInfo {
		public Integer Id;
		public Integer UserId;
		public Integer IsNotificationDirty;
		public Integer LastUserRecGen;
	}

	public static class UserPassword {
		public Integer UserId;
		public String Password;
		public Integer CreatedTime;
	}

}