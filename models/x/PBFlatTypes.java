package ir.ms.pb;

public class PBFlatTypes {

	public class PB_UserWithMe {
	   public int UserId;
	   public String UserName;
	   public String FirstName;
	   public String LastName;
	   public String About;
	   public String FullName;
	   public String AvatarUrl;
	   public int PrivacyProfile;
	   public String Phone;
	   public int UpdatedTime;
	   public int AppVersion;
	   public int FollowingType;
	}
	/*
	folding
	PBFlatTypes.PB_UserWithMe t = new PBFlatTypes.PB_UserWithMe();
    t.setUserId();
    t.setUserName();
    t.setFirstName();
    t.setLastName();
    t.setAbout();
    t.setFullName();
    t.setAvatarUrl();
    t.setPrivacyProfile();
    t.setPhone();
    t.setUpdatedTime();
    t.setAppVersion();
    t.setFollowingType();
	*/

	public class PB_Message {
	   public String MessageKey;
	   public String RoomKey;
	   public long UserId;
	   public long PeerId;
	   public int RoomTypeId;
	   public int MessageTypeId;
	   public String Text;
	   public String ExtraJson;
	   public int IsByMe;
	   public int IsStared;
	   public long CreatedMs;
	   public long CreatedDeviceMs;
	   public long SortId;
	   public long PeerSeenTime;
	   public long ServerReceivedTime;
	   public long ServerDeletedTime;
	   public long ISeenTime;
	   public int ToPush;
	   public String MsgFile_LocalSrc;
	   public int MsgFile_Status;
	   public PB_MsgFile File;
	}
	/*
	folding
	PBFlatTypes.PB_Message t = new PBFlatTypes.PB_Message();
    t.setMessageKey();
    t.setRoomKey();
    t.setUserId();
    t.setPeerId();
    t.setRoomTypeId();
    t.setMessageTypeId();
    t.setText();
    t.setExtraJson();
    t.setIsByMe();
    t.setIsStared();
    t.setCreatedMs();
    t.setCreatedDeviceMs();
    t.setSortId();
    t.setPeerSeenTime();
    t.setServerReceivedTime();
    t.setServerDeletedTime();
    t.setISeenTime();
    t.setToPush();
    t.setMsgFile_LocalSrc();
    t.setMsgFile_Status();
    t.setFile();
	*/

	public class PB_MsgFile {
	   public String Name;
	   public long Size;
	   public int FileType;
	   public String MimeType;
	   public int Width;
	   public int Height;
	   public int Duration;
	   public String Extension;
	   public byte[] ThumbData;
	   public byte[] Data;
	   public String ServerSrc;
	}
	/*
	folding
	PBFlatTypes.PB_MsgFile t = new PBFlatTypes.PB_MsgFile();
    t.setName();
    t.setSize();
    t.setFileType();
    t.setMimeType();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
    t.setThumbData();
    t.setData();
    t.setServerSrc();
	*/

	public class PB_Response {
	   public int Status;
	}
	/*
	folding
	PBFlatTypes.PB_Response t = new PBFlatTypes.PB_Response();
    t.setStatus();
	*/

	public class PB_Request {
	   public int Status;
	}
	/*
	folding
	PBFlatTypes.PB_Request t = new PBFlatTypes.PB_Request();
    t.setStatus();
	*/

	public class PB_RequestMsgAddMany {
	   public PB_Request Request;
	   public PB_Message Messages;
	}
	/*
	folding
	PBFlatTypes.PB_RequestMsgAddMany t = new PBFlatTypes.PB_RequestMsgAddMany();
    t.setRequest();
    t.setMessages();
	*/

	public class PB_ResponseMsgAddMany {
	   public PB_Response Response;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseMsgAddMany t = new PBFlatTypes.PB_ResponseMsgAddMany();
    t.setResponse();
	*/

	public class PB_Push {
	   public int Status;
	}
	/*
	folding
	PBFlatTypes.PB_Push t = new PBFlatTypes.PB_Push();
    t.setStatus();
	*/

	public class PB_Result {
	   public int Status;
	}
	/*
	folding
	PBFlatTypes.PB_Result t = new PBFlatTypes.PB_Result();
    t.setStatus();
	*/

	public class PB_PushMsgAddMany {
	   public PB_Push Push;
	   public PB_Message Messages;
	   public PB_UserWithMe Users;
	}
	/*
	folding
	PBFlatTypes.PB_PushMsgAddMany t = new PBFlatTypes.PB_PushMsgAddMany();
    t.setPush();
    t.setMessages();
    t.setUsers();
	*/

	public class PB_ResultMsgAddMany {
	   public PB_Result Result;
	}
	/*
	folding
	PBFlatTypes.PB_ResultMsgAddMany t = new PBFlatTypes.PB_ResultMsgAddMany();
    t.setResult();
	*/

	public class PB_MsgEvent {
	   public String MessageKey;
	   public String RoomKey;
	   public long PeerUserId;
	   public int EventType;
	   public long AtTime;
	}
	/*
	folding
	PBFlatTypes.PB_MsgEvent t = new PBFlatTypes.PB_MsgEvent();
    t.setMessageKey();
    t.setRoomKey();
    t.setPeerUserId();
    t.setEventType();
    t.setAtTime();
	*/

	public class PB_PushMsgEvents {
	   public PB_Push Push;
	   public PB_MsgEvent Events;
	}
	/*
	folding
	PBFlatTypes.PB_PushMsgEvents t = new PBFlatTypes.PB_PushMsgEvents();
    t.setPush();
    t.setEvents();
	*/

	public class PB_ResultMsgEvents {
	   public PB_Result Result;
	}
	/*
	folding
	PBFlatTypes.PB_ResultMsgEvents t = new PBFlatTypes.PB_ResultMsgEvents();
    t.setResult();
	*/

	public class PB_MsgsSeenFromClient {
	   public String MessageKey;
	   public String RoomKey;
	   public long UserId;
	   public long AtTime;
	}
	/*
	folding
	PBFlatTypes.PB_MsgsSeenFromClient t = new PBFlatTypes.PB_MsgsSeenFromClient();
    t.setMessageKey();
    t.setRoomKey();
    t.setUserId();
    t.setAtTime();
	*/

	public class PB_MsgSeen {
	   public String MessageKey;
	   public String RoomKey;
	   public long UserId;
	   public long AtTime;
	}
	/*
	folding
	PBFlatTypes.PB_MsgSeen t = new PBFlatTypes.PB_MsgSeen();
    t.setMessageKey();
    t.setRoomKey();
    t.setUserId();
    t.setAtTime();
	*/

	public class PB_RequestMsgsSeen {
	   public PB_Request Request;
	   public PB_MsgSeen Seen;
	}
	/*
	folding
	PBFlatTypes.PB_RequestMsgsSeen t = new PBFlatTypes.PB_RequestMsgsSeen();
    t.setRequest();
    t.setSeen();
	*/

	public class PB_ResponseMsgsSeen {
	   public PB_Response Response;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseMsgsSeen t = new PBFlatTypes.PB_ResponseMsgsSeen();
    t.setResponse();
	*/

	public class PB_ReqRpcAddMsg {
	   public PB_Request Request;
	   public PB_Message Message;
	}
	/*
	folding
	PBFlatTypes.PB_ReqRpcAddMsg t = new PBFlatTypes.PB_ReqRpcAddMsg();
    t.setRequest();
    t.setMessage();
	*/

	public class PB_ResRpcAddMsg {
	   public PB_Response Response;
	   public String ServerSrc;
	}
	/*
	folding
	PBFlatTypes.PB_ResRpcAddMsg t = new PBFlatTypes.PB_ResRpcAddMsg();
    t.setResponse();
    t.setServerSrc();
	*/

	public class GeoLocation {
	   public double Lat;
	   public double Lon;
	}
	/*
	folding
	PBFlatTypes.GeoLocation t = new PBFlatTypes.GeoLocation();
    t.setLat();
    t.setLon();
	*/

	public class RoomMessageLog {
	   public RoomMessageLogEnum typ;
	   public long TargetUserId;
	   public long ByUserId;
	}
	/*
	folding
	PBFlatTypes.RoomMessageLog t = new PBFlatTypes.RoomMessageLog();
    t.settyp();
    t.setTargetUserId();
    t.setByUserId();
	*/

	public class RoomMessageForwardFrom {
	   public long RoomId;
	   public long MessageId;
	   public int RoomTypeEnum;
	}
	/*
	folding
	PBFlatTypes.RoomMessageForwardFrom t = new PBFlatTypes.RoomMessageForwardFrom();
    t.setRoomId();
    t.setMessageId();
    t.setRoomTypeEnum();
	*/

	public class RoomDraft {
	   public String Message;
	   public long ReplyTo;
	}
	/*
	folding
	PBFlatTypes.RoomDraft t = new PBFlatTypes.RoomDraft();
    t.setMessage();
    t.setReplyTo();
	*/

	public class ChatRoom {
	}
	/*
	folding
	PBFlatTypes.ChatRoom t = new PBFlatTypes.ChatRoom();
	*/

	public class Pagination {
	   public int Offset;
	   public int Limit;
	}
	/*
	folding
	PBFlatTypes.Pagination t = new PBFlatTypes.Pagination();
    t.setOffset();
    t.setLimit();
	*/

	public class PB_UserParam_CheckUserName2 {
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_CheckUserName2 t = new PBFlatTypes.PB_UserParam_CheckUserName2();
	*/

	public class PB_UserResponse_CheckUserName2 {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_CheckUserName2 t = new PBFlatTypes.PB_UserResponse_CheckUserName2();
	*/

	public class PB_UserParam_BlockUser {
	   public int UserId;
	   public String UserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_BlockUser t = new PBFlatTypes.PB_UserParam_BlockUser();
    t.setUserId();
    t.setUserName();
	*/

	public class PB_UserOfflineResponse_BlockUser {
	   public int ByUserId;
	   public int TargetUserId;
	   public String TargetUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserOfflineResponse_BlockUser t = new PBFlatTypes.PB_UserOfflineResponse_BlockUser();
    t.setByUserId();
    t.setTargetUserId();
    t.setTargetUserName();
	*/

	public class PB_UserParam_UnBlockUser {
	   public int Offset;
	   public int Limit;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UnBlockUser t = new PBFlatTypes.PB_UserParam_UnBlockUser();
    t.setOffset();
    t.setLimit();
	*/

	public class PB_UserOfflineResponse_UnBlockUser {
	   public UserView Users;
	}
	/*
	folding
	PBFlatTypes.PB_UserOfflineResponse_UnBlockUser t = new PBFlatTypes.PB_UserOfflineResponse_UnBlockUser();
    t.setUsers();
	*/

	public class PB_UserParam_BlockedList {
	   public int UserId;
	   public String UserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_BlockedList t = new PBFlatTypes.PB_UserParam_BlockedList();
    t.setUserId();
    t.setUserName();
	*/

	public class PB_UserResponse_BlockedList {
	   public int ByUserId;
	   public int TargetUserId;
	   public String TargetUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_BlockedList t = new PBFlatTypes.PB_UserResponse_BlockedList();
    t.setByUserId();
    t.setTargetUserId();
    t.setTargetUserName();
	*/

	public class PB_UserParam_UpdateAbout {
	   public String NewAbout;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UpdateAbout t = new PBFlatTypes.PB_UserParam_UpdateAbout();
    t.setNewAbout();
	*/

	public class PB_UserOfflineResponse_UpdateAbout {
	   public int UserId;
	   public String NewAbout;
	}
	/*
	folding
	PBFlatTypes.PB_UserOfflineResponse_UpdateAbout t = new PBFlatTypes.PB_UserOfflineResponse_UpdateAbout();
    t.setUserId();
    t.setNewAbout();
	*/

	public class PB_UserParam_UpdateUserName {
	   public String NewUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_UpdateUserName t = new PBFlatTypes.PB_UserParam_UpdateUserName();
    t.setNewUserName();
	*/

	public class PB_UserOfflineResponse_UpdateUserName {
	   public int UserId;
	   public String NewUserName;
	}
	/*
	folding
	PBFlatTypes.PB_UserOfflineResponse_UpdateUserName t = new PBFlatTypes.PB_UserOfflineResponse_UpdateUserName();
    t.setUserId();
    t.setNewUserName();
	*/

	public class PB_UserParam_ChangeAvatar {
	   public boolean None;
	   public byte[] ImageData2;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_ChangeAvatar t = new PBFlatTypes.PB_UserParam_ChangeAvatar();
    t.setNone();
    t.setImageData2();
	*/

	public class PB_UserOfflineResponse_ChangeAvatar {
	}
	/*
	folding
	PBFlatTypes.PB_UserOfflineResponse_ChangeAvatar t = new PBFlatTypes.PB_UserOfflineResponse_ChangeAvatar();
	*/

	public class PB_UserParam_ChangePrivacy {
	   public ProfilePrivacyLevelEnum Level;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_ChangePrivacy t = new PBFlatTypes.PB_UserParam_ChangePrivacy();
    t.setLevel();
	*/

	public class PB_UserResponseOffline_ChangePrivacy {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponseOffline_ChangePrivacy t = new PBFlatTypes.PB_UserResponseOffline_ChangePrivacy();
	*/

	public class PB_UserParam_CheckUserName {
	   public ProfilePrivacyLevelEnum Level;
	}
	/*
	folding
	PBFlatTypes.PB_UserParam_CheckUserName t = new PBFlatTypes.PB_UserParam_CheckUserName();
    t.setLevel();
	*/

	public class PB_UserResponse_CheckUserName {
	}
	/*
	folding
	PBFlatTypes.PB_UserResponse_CheckUserName t = new PBFlatTypes.PB_UserResponse_CheckUserName();
	*/

	public class UserView {
	}
	/*
	folding
	PBFlatTypes.UserView t = new PBFlatTypes.UserView();
	*/

	public class PB_Chat {
	   public long ChatId;
	   public String ChatKey;
	   public int RoomTypeEnumId;
	   public int UserId;
	   public int LastSeqSeen;
	   public int LastSeqDelete;
	   public int PeerUserId;
	   public long GroupId;
	   public int CreatedTime;
	   public int CurrentSeq;
	   public long UpdatedMs;
	}
	/*
	folding
	PBFlatTypes.PB_Chat t = new PBFlatTypes.PB_Chat();
    t.setChatId();
    t.setChatKey();
    t.setRoomTypeEnumId();
    t.setUserId();
    t.setLastSeqSeen();
    t.setLastSeqDelete();
    t.setPeerUserId();
    t.setGroupId();
    t.setCreatedTime();
    t.setCurrentSeq();
    t.setUpdatedMs();
	*/

	public class PB_DirectMessage {
	   public long MessageId;
	   public String RoomKey;
	   public int UserId;
	   public long MessageFileId;
	   public int MessageTypeEnum;
	   public String Text;
	   public int Time;
	   public int PeerReceivedTime;
	   public int PeerSeenTime;
	   public int DeliviryStatusEnum;
	}
	/*
	folding
	PBFlatTypes.PB_DirectMessage t = new PBFlatTypes.PB_DirectMessage();
    t.setMessageId();
    t.setRoomKey();
    t.setUserId();
    t.setMessageFileId();
    t.setMessageTypeEnum();
    t.setText();
    t.setTime();
    t.setPeerReceivedTime();
    t.setPeerSeenTime();
    t.setDeliviryStatusEnum();
	*/

	public class PB_User {
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
	/*
	folding
	PBFlatTypes.PB_User t = new PBFlatTypes.PB_User();
    t.setId();
    t.setUserName();
    t.setUserNameLower();
    t.setFirstName();
    t.setLastName();
    t.setAbout();
    t.setFullName();
    t.setAvatarUrl();
    t.setPrivacyProfile();
    t.setPhone();
    t.setEmail();
    t.setIsDeleted();
    t.setPasswordHash();
    t.setPasswordSalt();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setMediaCount();
    t.setLikesCount();
    t.setResharedCount();
    t.setLastActionTime();
    t.setLastPostTime();
    t.setPrimaryFollowingList();
    t.setCreatedTime();
    t.setUpdatedTime();
    t.setSessionUuid();
    t.setDeviceUuid();
    t.setLastWifiMacAddress();
    t.setLastNetworkType();
    t.setAppVersion();
    t.setLastActivityTime();
    t.setLastLoginTime();
    t.setLastIpAddress();
	*/

	public class PB_CommandToServer {
	   public long ClientCallId;
	   public String Command;
	   public boolean RespondReached;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_CommandToServer t = new PBFlatTypes.PB_CommandToServer();
    t.setClientCallId();
    t.setCommand();
    t.setRespondReached();
    t.setData();
	*/

	public class PB_CommandToClient {
	   public long ServerCallId;
	   public String Command;
	   public boolean RespondReached;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_CommandToClient t = new PBFlatTypes.PB_CommandToClient();
    t.setServerCallId();
    t.setCommand();
    t.setRespondReached();
    t.setData();
	*/

	public class PB_CommandReachedToServer {
	   public long ClientCallId;
	}
	/*
	folding
	PBFlatTypes.PB_CommandReachedToServer t = new PBFlatTypes.PB_CommandReachedToServer();
    t.setClientCallId();
	*/

	public class PB_CommandReachedToClient {
	   public long ServerCallId;
	}
	/*
	folding
	PBFlatTypes.PB_CommandReachedToClient t = new PBFlatTypes.PB_CommandReachedToClient();
    t.setServerCallId();
	*/

	public class PB_ResponseToClient {
	   public long ClientCallId;
	   public String PBClass;
	   public byte[] Data;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseToClient t = new PBFlatTypes.PB_ResponseToClient();
    t.setClientCallId();
    t.setPBClass();
    t.setData();
	*/

	public class PB_ChatView {
	   public long ChatId;
	   public String ChatKey;
	   public int RoomTypeEnumId;
	   public int UserId;
	   public int LastSeqSeen;
	   public int LastSeqDelete;
	   public int PeerUserId;
	   public long GroupId;
	   public int CreatedTime;
	   public int CurrentSeq;
	   public long UpdatedMs;
	}
	/*
	folding
	PBFlatTypes.PB_ChatView t = new PBFlatTypes.PB_ChatView();
    t.setChatId();
    t.setChatKey();
    t.setRoomTypeEnumId();
    t.setUserId();
    t.setLastSeqSeen();
    t.setLastSeqDelete();
    t.setPeerUserId();
    t.setGroupId();
    t.setCreatedTime();
    t.setCurrentSeq();
    t.setUpdatedMs();
	*/

	public class PB_MessageView {
	   public long MessageId;
	   public String RoomKey;
	   public int UserId;
	   public long MessageFileId;
	   public int MessageTypeEnumId;
	   public String Text;
	   public int Time;
	   public int PeerReceivedTime;
	   public int PeerSeenTime;
	   public int DeliviryStatusEnumId;
	   public long ChatId;
	}
	/*
	folding
	PBFlatTypes.PB_MessageView t = new PBFlatTypes.PB_MessageView();
    t.setMessageId();
    t.setRoomKey();
    t.setUserId();
    t.setMessageFileId();
    t.setMessageTypeEnumId();
    t.setText();
    t.setTime();
    t.setPeerReceivedTime();
    t.setPeerSeenTime();
    t.setDeliviryStatusEnumId();
    t.setChatId();
	*/

	public class PB_MessageFileView {
	   public long MessageFileId;
	   public String Name;
	   public int Size;
	   public int FileTypeEnum;
	   public String MimeType;
	   public int Width;
	   public int Height;
	   public int Duration;
	   public String Extension;
	   public String ThumbData64;
	   public String ServerSrc;
	   public String ServerPath;
	   public String ServerThumbPath;
	   public String BucketId;
	   public int ServerId;
	   public int CanDel;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_MessageFileView t = new PBFlatTypes.PB_MessageFileView();
    t.setMessageFileId();
    t.setName();
    t.setSize();
    t.setFileTypeEnum();
    t.setMimeType();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
    t.setThumbData64();
    t.setServerSrc();
    t.setServerPath();
    t.setServerThumbPath();
    t.setBucketId();
    t.setServerId();
    t.setCanDel();
    t.setCreatedTime();
	*/

	public class PB_UserView {
	   public int UserId;
	   public String UserName;
	   public String FirstName;
	   public String LastName;
	   public String About;
	   public String FullName;
	   public String AvatarUrl;
	   public int PrivacyProfile;
	   public int IsDeleted;
	   public int FollowersCount;
	   public int FollowingCount;
	   public int PostsCount;
	   public int UpdatedTime;
	   public int AppVersion;
	   public int LastActivityTime;
	   public int FollowingType;
	}
	/*
	folding
	PBFlatTypes.PB_UserView t = new PBFlatTypes.PB_UserView();
    t.setUserId();
    t.setUserName();
    t.setFirstName();
    t.setLastName();
    t.setAbout();
    t.setFullName();
    t.setAvatarUrl();
    t.setPrivacyProfile();
    t.setIsDeleted();
    t.setFollowersCount();
    t.setFollowingCount();
    t.setPostsCount();
    t.setUpdatedTime();
    t.setAppVersion();
    t.setLastActivityTime();
    t.setFollowingType();
	*/

	public class PB_RoomInlineView {
	   public long RoomId;
	   public RoomTypeEnum RoomTypeEnum;
	}
	/*
	folding
	PBFlatTypes.PB_RoomInlineView t = new PBFlatTypes.PB_RoomInlineView();
    t.setRoomId();
    t.setRoomTypeEnum();
	*/

	public class PB_MessageForwardedFrom {
	   public long RoomId;
	   public RoomTypeEnum RoomTypeEnum;
	   public long MessageId;
	   public int MessageSeq;
	}
	/*
	folding
	PBFlatTypes.PB_MessageForwardedFrom t = new PBFlatTypes.PB_MessageForwardedFrom();
    t.setRoomId();
    t.setRoomTypeEnum();
    t.setMessageId();
    t.setMessageSeq();
	*/

	public class PB_GroupView {
	   public long GroupId;
	   public String GroupName;
	   public int MembersCount;
	   public int GroupPrivacyEnum;
	   public int CreatorUserId;
	   public int CreatedTime;
	   public long UpdatedMs;
	   public int CurrentSeq;
	}
	/*
	folding
	PBFlatTypes.PB_GroupView t = new PBFlatTypes.PB_GroupView();
    t.setGroupId();
    t.setGroupName();
    t.setMembersCount();
    t.setGroupPrivacyEnum();
    t.setCreatorUserId();
    t.setCreatedTime();
    t.setUpdatedMs();
    t.setCurrentSeq();
	*/

	public class PB_GroupMemberView {
	   public long Id;
	   public long GroupId;
	   public int UserId;
	   public int ByUserId;
	   public int GroupRoleEnum;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_GroupMemberView t = new PBFlatTypes.PB_GroupMemberView();
    t.setId();
    t.setGroupId();
    t.setUserId();
    t.setByUserId();
    t.setGroupRoleEnum();
    t.setCreatedTime();
	*/

	public class PB_MessageFileView__DEp {
	   public long MessageFileId;
	   public String Name;
	   public int Size;
	   public int FileTypeEnum;
	   public String MimeType;
	   public int Width;
	   public int Height;
	   public int Duration;
	   public String Extension;
	   public String ThumbData64;
	   public String ServerSrc;
	   public String ServerPath;
	   public String ServerThumbPath;
	   public String BucketId;
	   public int ServerId;
	   public int CanDel;
	   public int CreatedTime;
	}
	/*
	folding
	PBFlatTypes.PB_MessageFileView__DEp t = new PBFlatTypes.PB_MessageFileView__DEp();
    t.setMessageFileId();
    t.setName();
    t.setSize();
    t.setFileTypeEnum();
    t.setMimeType();
    t.setWidth();
    t.setHeight();
    t.setDuration();
    t.setExtension();
    t.setThumbData64();
    t.setServerSrc();
    t.setServerPath();
    t.setServerThumbPath();
    t.setBucketId();
    t.setServerId();
    t.setCanDel();
    t.setCreatedTime();
	*/

	public class PB_ReqLastChangesForTheRoom {
	   public long RoomId;
	   public long LastLogId;
	   public int LastHaveSeq;
	}
	/*
	folding
	PBFlatTypes.PB_ReqLastChangesForTheRoom t = new PBFlatTypes.PB_ReqLastChangesForTheRoom();
    t.setRoomId();
    t.setLastLogId();
    t.setLastHaveSeq();
	*/

	public class PB_ResponseLastChangesForTheRoom {
	   public PB_MessageView Messages;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseLastChangesForTheRoom t = new PBFlatTypes.PB_ResponseLastChangesForTheRoom();
    t.setMessages();
	*/

	public class PB_RequestSetLastSeenMessages {
	   public long RoomId;
	   public long FromMessageId;
	   public int ToMessageId;
	   public long AtTimeMs;
	}
	/*
	folding
	PBFlatTypes.PB_RequestSetLastSeenMessages t = new PBFlatTypes.PB_RequestSetLastSeenMessages();
    t.setRoomId();
    t.setFromMessageId();
    t.setToMessageId();
    t.setAtTimeMs();
	*/

	public class PB_ResponseSetLastSeenMessages {
	   public PB_MessageView Messages;
	}
	/*
	folding
	PBFlatTypes.PB_ResponseSetLastSeenMessages t = new PBFlatTypes.PB_ResponseSetLastSeenMessages();
    t.setMessages();
	*/

	public class PB_MessagesCollections {
	   public long DirectMessagesIds;
	   public long GroupMessagesIds;
	   public long BroadCatMessagesIds;
	}
	/*
	folding
	PBFlatTypes.PB_MessagesCollections t = new PBFlatTypes.PB_MessagesCollections();
    t.setDirectMessagesIds();
    t.setGroupMessagesIds();
    t.setBroadCatMessagesIds();
	*/

	public class PB_DirectLogView {
	   public PB_DirectLog Row;
	   public PB_MessageView NewMessage;
	}
	/*
	folding
	PBFlatTypes.PB_DirectLogView t = new PBFlatTypes.PB_DirectLogView();
    t.setRow();
    t.setNewMessage();
	*/

	public class PB_DirectLog {
	   public long Id;
	   public int ToUserId;
	   public long MessageId;
	   public long ChatId;
	   public int PeerUserId;
	   public int EventType;
	   public int RoomLogTypeId;
	   public int FromSeq;
	   public int ToSeq;
	   public byte[] ExtraPB;
	   public String ExtraJson;
	   public long AtTimeMs;
	}
	/*
	folding
	PBFlatTypes.PB_DirectLog t = new PBFlatTypes.PB_DirectLog();
    t.setId();
    t.setToUserId();
    t.setMessageId();
    t.setChatId();
    t.setPeerUserId();
    t.setEventType();
    t.setRoomLogTypeId();
    t.setFromSeq();
    t.setToSeq();
    t.setExtraPB();
    t.setExtraJson();
    t.setAtTimeMs();
	*/

	public class PB_PushDirectLogViewsMany {
	   public PB_DirectLogView Rows;
	}
	/*
	folding
	PBFlatTypes.PB_PushDirectLogViewsMany t = new PBFlatTypes.PB_PushDirectLogViewsMany();
    t.setRows();
	*/

	public class PB_MsgParam_AddNewTextMessage {
	   public String Text;
	   public int PeerId;
	   public int Time;
	   public long ReplyToMessageId;
	   public PB_MessageForwardedFrom Forward;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_AddNewTextMessage t = new PBFlatTypes.PB_MsgParam_AddNewTextMessage();
    t.setText();
    t.setPeerId();
    t.setTime();
    t.setReplyToMessageId();
    t.setForward();
	*/

	public class PB_MsgResponse_AddNewTextMessage {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_AddNewTextMessage t = new PBFlatTypes.PB_MsgResponse_AddNewTextMessage();
	*/

	public class PB_MsgParam_SetRoomActionDoing {
	   public long GroupId;
	   public String DirectRoomKey;
	   public RoomActionDoingEnum ActionType;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_SetRoomActionDoing t = new PBFlatTypes.PB_MsgParam_SetRoomActionDoing();
    t.setGroupId();
    t.setDirectRoomKey();
    t.setActionType();
	*/

	public class PB_MsgResponse_SetRoomActionDoing {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_SetRoomActionDoing t = new PBFlatTypes.PB_MsgResponse_SetRoomActionDoing();
	*/

	public class PB_MsgParam_GetMessagesByIds {
	   public PB_MessagesCollections MessagesCollections;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_GetMessagesByIds t = new PBFlatTypes.PB_MsgParam_GetMessagesByIds();
    t.setMessagesCollections();
	*/

	public class PB_MsgResponse_GetMessagesByIds {
	   public PB_MessageView MessagesViews;
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_GetMessagesByIds t = new PBFlatTypes.PB_MsgResponse_GetMessagesByIds();
    t.setMessagesViews();
	*/

	public class PB_MsgParam_GetMessagesHistory {
	   public long ChatId;
	   public int FromSeq;
	   public int EndSeq;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_GetMessagesHistory t = new PBFlatTypes.PB_MsgParam_GetMessagesHistory();
    t.setChatId();
    t.setFromSeq();
    t.setEndSeq();
	*/

	public class PB_MsgResponse_GetMessagesHistory {
	   public PB_MessageView MessagesViews;
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_GetMessagesHistory t = new PBFlatTypes.PB_MsgResponse_GetMessagesHistory();
    t.setMessagesViews();
	*/

	public class PB_MsgParam_SetChatMessagesRangeAsSeen {
	   public long ChatId;
	   public int FromSeq;
	   public int EndSeq;
	   public long SeenTimeMs;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_MsgParam_SetChatMessagesRangeAsSeen();
    t.setChatId();
    t.setFromSeq();
    t.setEndSeq();
    t.setSeenTimeMs();
	*/

	public class PB_MsgResponse_SetChatMessagesRangeAsSeen {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_SetChatMessagesRangeAsSeen t = new PBFlatTypes.PB_MsgResponse_SetChatMessagesRangeAsSeen();
	*/

	public class PB_MsgParam_DeleteChatHistory {
	   public long ChatId;
	   public int ToSeq;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_DeleteChatHistory t = new PBFlatTypes.PB_MsgParam_DeleteChatHistory();
    t.setChatId();
    t.setToSeq();
	*/

	public class PB_MsgResponse_DeleteChatHistory {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_DeleteChatHistory t = new PBFlatTypes.PB_MsgResponse_DeleteChatHistory();
	*/

	public class PB_MsgParam_DeleteMessagesByIds {
	   public PB_MessagesCollections MessagesCollections;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_DeleteMessagesByIds t = new PBFlatTypes.PB_MsgParam_DeleteMessagesByIds();
    t.setMessagesCollections();
	*/

	public class PB_MsgResponse_DeleteMessagesByIds {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_DeleteMessagesByIds t = new PBFlatTypes.PB_MsgResponse_DeleteMessagesByIds();
	*/

	public class PB_MsgParam_SetMessagesAsReceived {
	   public PB_MessagesCollections MessagesCollections;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_SetMessagesAsReceived t = new PBFlatTypes.PB_MsgParam_SetMessagesAsReceived();
    t.setMessagesCollections();
	*/

	public class PB_MsgResponse_SetMessagesAsReceived {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_SetMessagesAsReceived t = new PBFlatTypes.PB_MsgResponse_SetMessagesAsReceived();
	*/

	public class PB_MsgParam_ForwardMessages {
	   public PB_MessagesCollections MessagesCollections;
	   public long ToDirectChatIds;
	   public long ToGroupChatIds;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_ForwardMessages t = new PBFlatTypes.PB_MsgParam_ForwardMessages();
    t.setMessagesCollections();
    t.setToDirectChatIds();
    t.setToGroupChatIds();
	*/

	public class PB_MsgResponse_ForwardMessages {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_ForwardMessages t = new PBFlatTypes.PB_MsgResponse_ForwardMessages();
	*/

	public class PB_MsgParam_EditMessage {
	   public long ChatId;
	   public RoomTypeEnum RoomType;
	   public long MessageId;
	   public String NewText;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_EditMessage t = new PBFlatTypes.PB_MsgParam_EditMessage();
    t.setChatId();
    t.setRoomType();
    t.setMessageId();
    t.setNewText();
	*/

	public class PB_MsgResponse_EditMessage {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_EditMessage t = new PBFlatTypes.PB_MsgResponse_EditMessage();
	*/

	public class PB_MsgParam_BroadcastNewMessage {
	   public String Text;
	   public int PeerId;
	   public int Time;
	   public PB_MessageForwardedFrom Forward;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_BroadcastNewMessage t = new PBFlatTypes.PB_MsgParam_BroadcastNewMessage();
    t.setText();
    t.setPeerId();
    t.setTime();
    t.setForward();
	*/

	public class PB_MsgResponse_BroadcastNewMessage {
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_BroadcastNewMessage t = new PBFlatTypes.PB_MsgResponse_BroadcastNewMessage();
	*/

	public class PB_MsgParam_Echo {
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_MsgParam_Echo t = new PBFlatTypes.PB_MsgParam_Echo();
    t.setText();
	*/

	public class PB_MsgResponse_PB_MsgParam_Echo {
	   public String Text;
	}
	/*
	folding
	PBFlatTypes.PB_MsgResponse_PB_MsgParam_Echo t = new PBFlatTypes.PB_MsgResponse_PB_MsgParam_Echo();
    t.setText();
	*/

	public class PB_UpdateGroupParticipants {
	}
	/*
	folding
	PBFlatTypes.PB_UpdateGroupParticipants t = new PBFlatTypes.PB_UpdateGroupParticipants();
	*/

	public class PB_UpdateNotifySettings {
	}
	/*
	folding
	PBFlatTypes.PB_UpdateNotifySettings t = new PBFlatTypes.PB_UpdateNotifySettings();
	*/

	public class PB_UpdateServiceNotification {
	}
	/*
	folding
	PBFlatTypes.PB_UpdateServiceNotification t = new PBFlatTypes.PB_UpdateServiceNotification();
	*/

	public class PB_UpdateMessageMeta {
	   public long MessageId;
	   public long AtTime;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateMessageMeta t = new PBFlatTypes.PB_UpdateMessageMeta();
    t.setMessageId();
    t.setAtTime();
	*/

	public class PB_UpdateMessageId {
	   public long OldMessageId;
	   public long NewMessageId;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateMessageId t = new PBFlatTypes.PB_UpdateMessageId();
    t.setOldMessageId();
    t.setNewMessageId();
	*/

	public class PB_UpdateMessageToEdit {
	   public long MessageId;
	   public String NewText;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateMessageToEdit t = new PBFlatTypes.PB_UpdateMessageToEdit();
    t.setMessageId();
    t.setNewText();
	*/

	public class PB_UpdateMessageToDelete {
	   public long MessageId;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateMessageToDelete t = new PBFlatTypes.PB_UpdateMessageToDelete();
    t.setMessageId();
	*/

	public class PB_UpdateRoomActionDoing {
	   public String RoomKey;
	   public RoomActionDoingEnum ActionType;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateRoomActionDoing t = new PBFlatTypes.PB_UpdateRoomActionDoing();
    t.setRoomKey();
    t.setActionType();
	*/

	public class PB_UpdateUserBlocked {
	   public int UserId;
	   public boolean Blocked;
	}
	/*
	folding
	PBFlatTypes.PB_UpdateUserBlocked t = new PBFlatTypes.PB_UpdateUserBlocked();
    t.setUserId();
    t.setBlocked();
	*/

	public class PB_ChangesHolderView {
	   public PB_MessageView NewMessages;
	   public PB_MessageFileView ChatFiles;
	   public PB_ChatView Chats;
	   public PB_UserView Users;
	   public PB_UpdateMessageId MessagesChangeIds;
	   public PB_UpdateMessageToEdit MessagesToUpdate;
	   public PB_UpdateMessageToDelete MessagesToDelete;
	   public PB_UpdateMessageMeta MessagesDelivierdToServer;
	   public PB_UpdateMessageMeta MessagesDelivierdToPeer;
	   public PB_UpdateMessageMeta MessagesSeenByPeer;
	   public PB_UpdateMessageMeta MessagesDeletedFromServer;
	   public PB_UpdateRoomActionDoing RoomActionDoing;
	   public PB_UpdateUserBlocked UserBlockedByMe;
	   public PB_UpdateUserBlocked UserBlockedMe;
	}
	/*
	folding
	PBFlatTypes.PB_ChangesHolderView t = new PBFlatTypes.PB_ChangesHolderView();
    t.setNewMessages();
    t.setChatFiles();
    t.setChats();
    t.setUsers();
    t.setMessagesChangeIds();
    t.setMessagesToUpdate();
    t.setMessagesToDelete();
    t.setMessagesDelivierdToServer();
    t.setMessagesDelivierdToPeer();
    t.setMessagesSeenByPeer();
    t.setMessagesDeletedFromServer();
    t.setRoomActionDoing();
    t.setUserBlockedByMe();
    t.setUserBlockedMe();
	*/

	
}

/*

RPC_INTERFACES.RpcMsgs RpcMsgs_Handeler = null;
RPC_INTERFACES.RPC_Auth RPC_Auth_Handeler = null;
RPC_INTERFACES.RPC_UserOffline RPC_UserOffline_Handeler = null;
RPC_INTERFACES.RPC_User RPC_User_Handeler = null;
RPC_INTERFACES.RPC_MessageReq RPC_MessageReq_Handeler = null;
RPC_INTERFACES.RPC_MessageReqOffline RPC_MessageReqOffline_Handeler = null;
RPC_INTERFACES.RPC_Msg RPC_Msg_Handeler = null;
	
*/