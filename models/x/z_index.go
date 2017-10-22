package x

import "github.com/jmoiron/sqlx"

//todo this code can be used for multi secondery coulmn -- but this one interfier with secondery_template - in futer merege this two temple to one unifed that uses both primiry keys and othe seconder options and with reloadings

// ActivityById Generated from index 'PRIMARY' -- retrieves a row from 'ms.activity' as a Activity.
func ActivityById(db *sqlx.DB, id int) (*Activity, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.activity ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	a := Activity{
		_exists: true,
	}

	err = db.Get(&a, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnActivity_LoadOne(&a)

	return &a, nil
}

// BucketByBucketId Generated from index 'PRIMARY' -- retrieves a row from 'ms.bucket' as a Bucket.
func BucketByBucketId(db *sqlx.DB, bucketId int) (*Bucket, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.bucket ` +
		`WHERE BucketId = ?`

	XOLog(sqlstr, bucketId)
	b := Bucket{
		_exists: true,
	}

	err = db.Get(&b, sqlstr, bucketId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnBucket_LoadOne(&b)

	return &b, nil
}

// ChatByChatKey Generated from index 'PRIMARY' -- retrieves a row from 'ms.chat' as a Chat.
func ChatByChatKey(db *sqlx.DB, chatKey string) (*Chat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.chat ` +
		`WHERE ChatKey = ?`

	XOLog(sqlstr, chatKey)
	c := Chat{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, chatKey)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnChat_LoadOne(&c)

	return &c, nil
}

// CommentById Generated from index 'PRIMARY' -- retrieves a row from 'ms.comment' as a Comment.
func CommentById(db *sqlx.DB, id int) (*Comment, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.comment ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	c := Comment{
		_exists: true,
	}

	err = db.Get(&c, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnComment_LoadOne(&c)

	return &c, nil
}

// DirectMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_message' as a DirectMessage.
func DirectMessageByMessageId(db *sqlx.DB, messageId int) (*DirectMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	dm := DirectMessage{
		_exists: true,
	}

	err = db.Get(&dm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectMessage_LoadOne(&dm)

	return &dm, nil
}

// DirectToMessageById Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_to_message' as a DirectToMessage.
func DirectToMessageById(db *sqlx.DB, id int) (*DirectToMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_to_message ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	dtm := DirectToMessage{
		_exists: true,
	}

	err = db.Get(&dtm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectToMessage_LoadOne(&dtm)

	return &dtm, nil
}

// DirectUpdateByDirectUpdateId Generated from index 'PRIMARY' -- retrieves a row from 'ms.direct_update' as a DirectUpdate.
func DirectUpdateByDirectUpdateId(db *sqlx.DB, directUpdateId int) (*DirectUpdate, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.direct_update ` +
		`WHERE DirectUpdateId = ?`

	XOLog(sqlstr, directUpdateId)
	du := DirectUpdate{
		_exists: true,
	}

	err = db.Get(&du, sqlstr, directUpdateId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnDirectUpdate_LoadOne(&du)

	return &du, nil
}

// FollowingListById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list' as a FollowingList.
func FollowingListById(db *sqlx.DB, id int) (*FollowingList, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	fl := FollowingList{
		_exists: true,
	}

	err = db.Get(&fl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingList_LoadOne(&fl)

	return &fl, nil
}

// FollowingListMemberById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list_member' as a FollowingListMember.
func FollowingListMemberById(db *sqlx.DB, id int) (*FollowingListMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flm := FollowingListMember{
		_exists: true,
	}

	err = db.Get(&flm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMember_LoadOne(&flm)

	return &flm, nil
}

// FollowingListMemberHistoryById Generated from index 'PRIMARY' -- retrieves a row from 'ms.following_list_member_history' as a FollowingListMemberHistory.
func FollowingListMemberHistoryById(db *sqlx.DB, id int) (*FollowingListMemberHistory, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.following_list_member_history ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	flmh := FollowingListMemberHistory{
		_exists: true,
	}

	err = db.Get(&flmh, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnFollowingListMemberHistory_LoadOne(&flmh)

	return &flmh, nil
}

// GeneralLogById Generated from index 'PRIMARY' -- retrieves a row from 'ms.general_log' as a GeneralLog.
func GeneralLogById(db *sqlx.DB, id int) (*GeneralLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.general_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gl := GeneralLog{
		_exists: true,
	}

	err = db.Get(&gl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGeneralLog_LoadOne(&gl)

	return &gl, nil
}

// GroupByGroupId Generated from index 'PRIMARY' -- retrieves a row from 'ms.group' as a Group.
func GroupByGroupId(db *sqlx.DB, groupId int) (*Group, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group ` +
		`WHERE GroupId = ?`

	XOLog(sqlstr, groupId)
	g := Group{
		_exists: true,
	}

	err = db.Get(&g, sqlstr, groupId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroup_LoadOne(&g)

	return &g, nil
}

// GroupMemberById Generated from index 'PRIMARY' -- retrieves a row from 'ms.group_member' as a GroupMember.
func GroupMemberById(db *sqlx.DB, id int) (*GroupMember, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group_member ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gm := GroupMember{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMember_LoadOne(&gm)

	return &gm, nil
}

// GroupMessageByMessageId Generated from index 'PRIMARY' -- retrieves a row from 'ms.group_message' as a GroupMessage.
func GroupMessageByMessageId(db *sqlx.DB, messageId int) (*GroupMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group_message ` +
		`WHERE MessageId = ?`

	XOLog(sqlstr, messageId)
	gm := GroupMessage{
		_exists: true,
	}

	err = db.Get(&gm, sqlstr, messageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupMessage_LoadOne(&gm)

	return &gm, nil
}

// GroupToMessageById Generated from index 'PRIMARY' -- retrieves a row from 'ms.group_to_message' as a GroupToMessage.
func GroupToMessageById(db *sqlx.DB, id int) (*GroupToMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.group_to_message ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	gtm := GroupToMessage{
		_exists: true,
	}

	err = db.Get(&gtm, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnGroupToMessage_LoadOne(&gtm)

	return &gtm, nil
}

// LikeById Generated from index 'PRIMARY' -- retrieves a row from 'ms.likes' as a Like.
func LikeById(db *sqlx.DB, id int) (*Like, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.likes ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	l := Like{
		_exists: true,
	}

	err = db.Get(&l, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnLike_LoadOne(&l)

	return &l, nil
}

// LogChangeById Generated from index 'PRIMARY' -- retrieves a row from 'ms.log_changes' as a LogChange.
func LogChangeById(db *sqlx.DB, id int) (*LogChange, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.log_changes ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	lc := LogChange{
		_exists: true,
	}

	err = db.Get(&lc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnLogChange_LoadOne(&lc)

	return &lc, nil
}

// MediaById Generated from index 'PRIMARY' -- retrieves a row from 'ms.media' as a Media.
func MediaById(db *sqlx.DB, id int) (*Media, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.media ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	m := Media{
		_exists: true,
	}

	err = db.Get(&m, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnMedia_LoadOne(&m)

	return &m, nil
}

// MessageFileByMessageFileId Generated from index 'PRIMARY' -- retrieves a row from 'ms.message_file' as a MessageFile.
func MessageFileByMessageFileId(db *sqlx.DB, messageFileId int) (*MessageFile, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.message_file ` +
		`WHERE MessageFileId = ?`

	XOLog(sqlstr, messageFileId)
	mf := MessageFile{
		_exists: true,
	}

	err = db.Get(&mf, sqlstr, messageFileId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnMessageFile_LoadOne(&mf)

	return &mf, nil
}

// NotificationById Generated from index 'PRIMARY' -- retrieves a row from 'ms.notification' as a Notification.
func NotificationById(db *sqlx.DB, id int) (*Notification, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.notification ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	n := Notification{
		_exists: true,
	}

	err = db.Get(&n, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotification_LoadOne(&n)

	return &n, nil
}

// NotificationRemovedByNotificationId Generated from index 'PRIMARY' -- retrieves a row from 'ms.notification_removed' as a NotificationRemoved.
func NotificationRemovedByNotificationId(db *sqlx.DB, notificationId int) (*NotificationRemoved, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.notification_removed ` +
		`WHERE NotificationId = ?`

	XOLog(sqlstr, notificationId)
	nr := NotificationRemoved{
		_exists: true,
	}

	err = db.Get(&nr, sqlstr, notificationId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnNotificationRemoved_LoadOne(&nr)

	return &nr, nil
}

// OldMessageById Generated from index 'PRIMARY' -- retrieves a row from 'ms.old_messages' as a OldMessage.
func OldMessageById(db *sqlx.DB, id int) (*OldMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.old_messages ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	om := OldMessage{
		_exists: true,
	}

	err = db.Get(&om, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnOldMessage_LoadOne(&om)

	return &om, nil
}

// OldMsgFileById Generated from index 'PRIMARY' -- retrieves a row from 'ms.old_msg_file' as a OldMsgFile.
func OldMsgFileById(db *sqlx.DB, id int) (*OldMsgFile, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.old_msg_file ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	omf := OldMsgFile{
		_exists: true,
	}

	err = db.Get(&omf, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnOldMsgFile_LoadOne(&omf)

	return &omf, nil
}

// OldMsgPushById Generated from index 'PRIMARY' -- retrieves a row from 'ms.old_msg_push' as a OldMsgPush.
func OldMsgPushById(db *sqlx.DB, id int) (*OldMsgPush, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.old_msg_push ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	omp := OldMsgPush{
		_exists: true,
	}

	err = db.Get(&omp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnOldMsgPush_LoadOne(&omp)

	return &omp, nil
}

// OldMsgPushEventById Generated from index 'PRIMARY' -- retrieves a row from 'ms.old_msg_push_event' as a OldMsgPushEvent.
func OldMsgPushEventById(db *sqlx.DB, id int) (*OldMsgPushEvent, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.old_msg_push_event ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	ompe := OldMsgPushEvent{
		_exists: true,
	}

	err = db.Get(&ompe, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnOldMsgPushEvent_LoadOne(&ompe)

	return &ompe, nil
}

// PhoneContactById Generated from index 'PRIMARY' -- retrieves a row from 'ms.phone_contacts' as a PhoneContact.
func PhoneContactById(db *sqlx.DB, id int) (*PhoneContact, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.phone_contacts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	pc := PhoneContact{
		_exists: true,
	}

	err = db.Get(&pc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPhoneContact_LoadOne(&pc)

	return &pc, nil
}

// PhotoByPhotoId Generated from index 'PRIMARY' -- retrieves a row from 'ms.photo' as a Photo.
func PhotoByPhotoId(db *sqlx.DB, photoId int) (*Photo, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.photo ` +
		`WHERE PhotoId = ?`

	XOLog(sqlstr, photoId)
	p := Photo{
		_exists: true,
	}

	err = db.Get(&p, sqlstr, photoId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPhoto_LoadOne(&p)

	return &p, nil
}

// PostById Generated from index 'PRIMARY' -- retrieves a row from 'ms.post' as a Post.
func PostById(db *sqlx.DB, id int) (*Post, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.post ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	p := Post{
		_exists: true,
	}

	err = db.Get(&p, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPost_LoadOne(&p)

	return &p, nil
}

// PushEventByPushEventId Generated from index 'PRIMARY' -- retrieves a row from 'ms.push_event' as a PushEvent.
func PushEventByPushEventId(db *sqlx.DB, pushEventId int) (*PushEvent, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.push_event ` +
		`WHERE PushEventId = ?`

	XOLog(sqlstr, pushEventId)
	pe := PushEvent{
		_exists: true,
	}

	err = db.Get(&pe, sqlstr, pushEventId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPushEvent_LoadOne(&pe)

	return &pe, nil
}

// PushMessageByPushMessageId Generated from index 'PRIMARY' -- retrieves a row from 'ms.push_message' as a PushMessage.
func PushMessageByPushMessageId(db *sqlx.DB, pushMessageId int) (*PushMessage, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.push_message ` +
		`WHERE PushMessageId = ?`

	XOLog(sqlstr, pushMessageId)
	pm := PushMessage{
		_exists: true,
	}

	err = db.Get(&pm, sqlstr, pushMessageId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnPushMessage_LoadOne(&pm)

	return &pm, nil
}

// RecommendUserById Generated from index 'PRIMARY' -- retrieves a row from 'ms.recommend_user' as a RecommendUser.
func RecommendUserById(db *sqlx.DB, id int) (*RecommendUser, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.recommend_user ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	ru := RecommendUser{
		_exists: true,
	}

	err = db.Get(&ru, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnRecommendUser_LoadOne(&ru)

	return &ru, nil
}

// RoomByRoomId Generated from index 'PRIMARY' -- retrieves a row from 'ms.room' as a Room.
func RoomByRoomId(db *sqlx.DB, roomId int) (*Room, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.room ` +
		`WHERE RoomId = ?`

	XOLog(sqlstr, roomId)
	r := Room{
		_exists: true,
	}

	err = db.Get(&r, sqlstr, roomId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnRoom_LoadOne(&r)

	return &r, nil
}

// SearchClickedById Generated from index 'PRIMARY' -- retrieves a row from 'ms.search_clicked' as a SearchClicked.
func SearchClickedById(db *sqlx.DB, id int) (*SearchClicked, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.search_clicked ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	sc := SearchClicked{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSearchClicked_LoadOne(&sc)

	return &sc, nil
}

// SessionById Generated from index 'PRIMARY' -- retrieves a row from 'ms.session' as a Session.
func SessionById(db *sqlx.DB, id int) (*Session, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.session ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	s := Session{
		_exists: true,
	}

	err = db.Get(&s, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSession_LoadOne(&s)

	return &s, nil
}

// SettingClientByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.setting_client' as a SettingClient.
func SettingClientByUserId(db *sqlx.DB, userId int) (*SettingClient, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.setting_client ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sc := SettingClient{
		_exists: true,
	}

	err = db.Get(&sc, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingClient_LoadOne(&sc)

	return &sc, nil
}

// SettingNotificationByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.setting_notifications' as a SettingNotification.
func SettingNotificationByUserId(db *sqlx.DB, userId int) (*SettingNotification, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.setting_notifications ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	sn := SettingNotification{
		_exists: true,
	}

	err = db.Get(&sn, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnSettingNotification_LoadOne(&sn)

	return &sn, nil
}

// TagById Generated from index 'PRIMARY' -- retrieves a row from 'ms.tag' as a Tag.
func TagById(db *sqlx.DB, id int) (*Tag, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.tag ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	t := Tag{
		_exists: true,
	}

	err = db.Get(&t, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTag_LoadOne(&t)

	return &t, nil
}

// TagsPostById Generated from index 'PRIMARY' -- retrieves a row from 'ms.tags_posts' as a TagsPost.
func TagsPostById(db *sqlx.DB, id int) (*TagsPost, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.tags_posts ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tp := TagsPost{
		_exists: true,
	}

	err = db.Get(&tp, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTagsPost_LoadOne(&tp)

	return &tp, nil
}

// TestChatById4 Generated from index 'PRIMARY' -- retrieves a row from 'ms.test_chat' as a TestChat.
func TestChatById4(db *sqlx.DB, id4 int) (*TestChat, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.test_chat ` +
		`WHERE Id4 = ?`

	XOLog(sqlstr, id4)
	tc := TestChat{
		_exists: true,
	}

	err = db.Get(&tc, sqlstr, id4)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTestChat_LoadOne(&tc)

	return &tc, nil
}

// TriggerLogById Generated from index 'PRIMARY' -- retrieves a row from 'ms.trigger_log' as a TriggerLog.
func TriggerLogById(db *sqlx.DB, id int) (*TriggerLog, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.trigger_log ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	tl := TriggerLog{
		_exists: true,
	}

	err = db.Get(&tl, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnTriggerLog_LoadOne(&tl)

	return &tl, nil
}

// UserById Generated from index 'PRIMARY' -- retrieves a row from 'ms.user' as a User.
func UserById(db *sqlx.DB, id int) (*User, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.Get(&u, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUser_LoadOne(&u)

	return &u, nil
}

// UserMetaInfoById Generated from index 'PRIMARY' -- retrieves a row from 'ms.user_meta_info' as a UserMetaInfo.
func UserMetaInfoById(db *sqlx.DB, id int) (*UserMetaInfo, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user_meta_info ` +
		`WHERE Id = ?`

	XOLog(sqlstr, id)
	umi := UserMetaInfo{
		_exists: true,
	}

	err = db.Get(&umi, sqlstr, id)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserMetaInfo_LoadOne(&umi)

	return &umi, nil
}

// UserPasswordByUserId Generated from index 'PRIMARY' -- retrieves a row from 'ms.user_password' as a UserPassword.
func UserPasswordByUserId(db *sqlx.DB, userId int) (*UserPassword, error) {
	var err error

	const sqlstr = `SELECT * ` +
		`FROM ms.user_password ` +
		`WHERE UserId = ?`

	XOLog(sqlstr, userId)
	up := UserPassword{
		_exists: true,
	}

	err = db.Get(&up, sqlstr, userId)
	if err != nil {
		XOLogErr(err)
		return nil, err
	}

	OnUserPassword_LoadOne(&up)

	return &up, nil
}
