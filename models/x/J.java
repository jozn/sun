package com.mardomsara.social.json;

public class J {
	public static class Activity {//oridnal: 0
		public int Id;
		public int ActorUserId;
		public int ActionTypeId;
		public int RowId;
		public int RootId;
		public int RefId;
		public int CreatedAt;
	}

	public static class Bucket {//oridnal: 1
		public int BucketId;
		public String BucketName;
		public int Server1Id;
		public int Server2Id;
		public int BackupServerId;
		public int ContentObjectTypeId;
		public int ContentObjectCount;
		public int CreatedTime;
	}

	public static class Chat {//oridnal: 2
		public String ChatKey;
		public String RoomKey;
		public int RoomTypeEnumId;
		public int UserId;
		public int PeerUserId;
		public int GroupId;
		public int CreatedSe;
		public int StartMessageIdFrom;
		public int LastDeletedMessageId;
		public int LastSeenMessageId;
		public int LastMessageId;
		public int UpdatedMs;
	}

	public static class Comment {//oridnal: 3
		public int Id;
		public int UserId;
		public int PostId;
		public String Text;
		public int CreatedTime;
	}

	public static class DirectMessage {//oridnal: 4
		public int MessageId;
		public String MessageKey;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnumId;
		public String Text;
		public int CreatedSe;
		public int PeerReceivedTime;
		public int PeerSeenTime;
		public int DeliviryStatusEnumId;
	}

	public static class DirectOffline {//oridnal: 5
		public int DirectOfflineId;
		public int ToUserId;
		public String ChatKey;
		public int MessageId;
		public int MessageFileId;
		public String PBClass;
		public UNKNOWN DataPB;
		public String DataJson;
		public String DataTemp;
		public int AtTimeMs;
	}

	public static class DirectToMessage {//oridnal: 6
		public int Id;
		public String ChatKey;
		public int MessageId;
		public int SourceEnumId;
	}

	public static class FollowingList {//oridnal: 7
		public int Id;
		public int UserId;
		public int ListType;
		public String Name;
		public int Count;
		public int IsAuto;
		public int IsPimiry;
		public int CreatedTime;
	}

	public static class FollowingListMember {//oridnal: 8
		public int Id;
		public int ListId;
		public int UserId;
		public int FollowedUserId;
		public int FollowType;
		public int UpdatedTimeMs;
	}

	public static class FollowingListMemberHistory {//oridnal: 9
		public int Id;
		public int ListId;
		public int UserId;
		public int FollowedUserId;
		public int FollowType;
		public int UpdatedTimeMs;
		public int FollowId;
	}

	public static class GeneralLog {//oridnal: 10
		public int Id;
		public int ToUserId;
		public int TargetId;
		public int LogTypeId;
		public UNKNOWN ExtraPB;
		public String ExtraJson;
		public int CreatedMs;
	}

	public static class Group {//oridnal: 11
		public int GroupId;
		public String GroupName;
		public int MembersCount;
		public int GroupPrivacyEnum;
		public int CreatorUserId;
		public int CreatedTime;
		public int UpdatedMs;
		public int CurrentSeq;
	}

	public static class GroupMember {//oridnal: 12
		public int Id;
		public int GroupId;
		public String GroupKey;
		public int UserId;
		public int ByUserId;
		public int GroupRoleEnumId;
		public int CreatedTime;
	}

	public static class GroupMessage {//oridnal: 13
		public int MessageId;
		public String RoomKey;
		public int UserId;
		public int MessageFileId;
		public int MessageTypeEnum;
		public String Text;
		public int CreatedMs;
		public int DeliveryStatusEnum;
	}

	public static class GroupToMessage {//oridnal: 14
		public int Id;
		public int GroupId;
		public int MessageId;
		public int CreatedMs;
		public int StatusEnum;
	}

	public static class Like {//oridnal: 15
		public int Id;
		public int PostId;
		public int PostTypeId;
		public int UserId;
		public int TypeId;
		public int CreatedTime;
	}

	public static class LogChange {//oridnal: 16
		public int Id;
		public String T;
	}

	public static class Media {//oridnal: 17
		public int Id;
		public int UserId;
		public int PostId;
		public int AlbumId;
		public int TypeId;
		public int CreatedTime;
		public String Src;
	}

	public static class MessageFile {//oridnal: 18
		public int MessageFileId;
		public String MessageFileKey;
		public int OriginalUserId;
		public String Name;
		public int Size;
		public int FileTypeEnumId;
		public int Width;
		public int Height;
		public int Duration;
		public String Extension;
		public String HashMd5;
		public int HashAccess;
		public int CreatedSe;
		public String ServerSrc;
		public String ServerPath;
		public String ServerThumbPath;
		public String BucketId;
		public int ServerId;
		public int CanDel;
	}

	public static class Notification {//oridnal: 19
		public int Id;
		public int ForUserId;
		public int ActorUserId;
		public int ActionTypeId;
		public int ObjectTypeId;
		public int RowId;
		public int RootId;
		public int RefId;
		public int SeenStatus;
		public int CreatedTime;
	}

	public static class NotificationRemoved {//oridnal: 20
		public int NotificationId;
		public int ForUserId;
	}

	public static class Offline {//oridnal: 21
		public int Id;
		public int FromUserId;
		public int ToUserId;
		public String RpcName;
		public String PBClassName;
		public String Key;
		public String DataJson;
		public UNKNOWN DataBlob;
		public int DataLength;
		public int CreatedMs;
	}

	public static class PhoneContact {//oridnal: 22
		public int Id;
		public String PhoneDisplayName;
		public String PhoneFamilyName;
		public String PhoneNumber;
		public String PhoneNormalizedNumber;
		public int PhoneContactRowId;
		public int UserId;
		public int DeviceUuidId;
		public int CreatedTime;
		public int UpdatedTime;
	}

	public static class Photo {//oridnal: 23
		public int PhotoId;
		public int UserId;
		public int PostId;
		public int AlbumId;
		public int ImageTypeId;
		public String Title;
		public String Src;
		public String PathSrc;
		public int BucketId;
		public int Width;
		public int Height;
		public float Ratio;
		public String HashMd5;
		public String Color;
		public int CreatedTime;
		public int W1080;
		public int W720;
		public int W480;
		public int W320;
		public int W160;
		public int W80;
	}

	public static class Post {//oridnal: 24
		public int Id;
		public int UserId;
		public int TypeId;
		public String Text;
		public String FormatedText;
		public int MediaCount;
		public int SharedTo;
		public int DisableComment;
		public int HasTag;
		public int LikesCount;
		public int CommentsCount;
		public int EditedTime;
		public int CreatedTime;
	}

	public static class PushEvent {//oridnal: 25
		public int PushEventId;
		public int ToUserId;
		public int ToDeviceId;
		public int MessageId;
		public int RoomTypeEnum;
		public int RoomId;
		public int PeerUserId;
		public int PushEventTypeEnum;
		public int AtTime;
	}

	public static class PushMessage {//oridnal: 26
		public int PushMessageId;
		public int ToUserId;
		public int ToDeviceId;
		public int MessageId;
		public int RoomTypeEnum;
		public int CreatedMs;
	}

	public static class RecommendUser {//oridnal: 27
		public int Id;
		public int UserId;
		public int TargetId;
		public float Weight;
		public int CreatedTime;
	}

	public static class Room {//oridnal: 28
		public int RoomId;
		public String RoomKey;
		public int RoomTypeEnum;
		public int UserId;
		public int LastSeqSeen;
		public int LastSeqDelete;
		public int PeerUserId;
		public int GroupId;
		public int CreatedTime;
		public int CurrentSeq;
	}

	public static class SearchClicked {//oridnal: 29
		public int Id;
		public String Query;
		public int ClickType;
		public int TargetId;
		public int UserId;
		public int CreatedAt;
	}

	public static class Session {//oridnal: 30
		public int Id;
		public int UserId;
		public String SessionUuid;
		public String ClientUuid;
		public String DeviceUuid;
		public int LastActivityTime;
		public String LastIpAddress;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public int LastNetworkTypeId;
		public int AppVersion;
		public int UpdatedTime;
		public int CreatedTime;
		public int LastSyncDirectUpdateId;
		public int LastSyncGeneralUpdateId;
		public int LastSyncNotifyUpdateId;
	}

	public static class SettingClient {//oridnal: 31
		public int UserId;
		public int AutoDownloadWifiVoice;
		public int AutoDownloadWifiImage;
		public int AutoDownloadWifiVideo;
		public int AutoDownloadWifiFile;
		public int AutoDownloadWifiMusic;
		public int AutoDownloadWifiGif;
		public int AutoDownloadCellularVoice;
		public int AutoDownloadCellularImage;
		public int AutoDownloadCellularVideo;
		public int AutoDownloadCellularFile;
		public int AutoDownloadCellularMusic;
		public int AutoDownloadCellularGif;
		public int AutoDownloadRoamingVoice;
		public int AutoDownloadRoamingImage;
		public int AutoDownloadRoamingVideo;
		public int AutoDownloadRoamingFile;
		public int AutoDownloadRoamingMusic;
		public int AutoDownloadRoamingGif;
		public int SaveToGallery;
	}

	public static class SettingNotification {//oridnal: 32
		public int UserId;
		public int SocialLedOn;
		public String SocialLedColor;
		public int ReqestToFollowYou;
		public int FollowedYou;
		public int AccptedYourFollowRequest;
		public int YourPostLiked;
		public int YourPostCommented;
		public int MenthenedYouInPost;
		public int MenthenedYouInComment;
		public int YourContactsJoined;
		public int DirectMessage;
		public int DirectAlert;
		public int DirectPerview;
		public int DirectLedOn;
		public int DirectLedColor;
		public int DirectVibrate;
		public int DirectPopup;
		public int DirectSound;
		public int DirectPriority;
	}

	public static class Tag {//oridnal: 33
		public int Id;
		public String Name;
		public int Count;
		public int IsBlocked;
		public int CreatedTime;
	}

	public static class TagsPost {//oridnal: 34
		public int Id;
		public int TagId;
		public int PostId;
		public int TypeId;
		public int CreatedTime;
	}

	public static class TestChat {//oridnal: 35
		public int Id;
		public int Id4;
		public int TimeMs;
		public String Text;
		public String Name;
		public int UserId;
		public int C2;
		public int C3;
		public int C4;
		public int C5;
	}

	public static class TriggerLog {//oridnal: 36
		public int Id;
		public String ModelName;
		public String ChangeType;
		public int TargetId;
		public String TargetStr;
		public int CreatedSe;
	}

	public static class User {//oridnal: 37
		public int Id;
		public String UserName;
		public String UserNameLower;
		public String FirstName;
		public String LastName;
		public String About;
		public String FullName;
		public String AvatarUrl;
		public int PrivacyProfile;
		public String Phone;
		public String Email;
		public int IsDeleted;
		public String PasswordHash;
		public String PasswordSalt;
		public int FollowersCount;
		public int FollowingCount;
		public int PostsCount;
		public int MediaCount;
		public int LikesCount;
		public int ResharedCount;
		public int LastActionTime;
		public int LastPostTime;
		public int PrimaryFollowingList;
		public int CreatedTime;
		public int UpdatedTime;
		public String SessionUuid;
		public String DeviceUuid;
		public String LastWifiMacAddress;
		public String LastNetworkType;
		public int AppVersion;
		public int LastActivityTime;
		public int LastLoginTime;
		public String LastIpAddress;
	}

	public static class UserMetaInfo {//oridnal: 38
		public int Id;
		public int UserId;
		public int IsNotificationDirty;
		public int LastUserRecGen;
	}

	public static class UserPassword {//oridnal: 39
		public int UserId;
		public String Password;
		public int CreatedTime;
	}

}