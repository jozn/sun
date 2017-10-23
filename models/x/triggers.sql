
################################ Activity ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS activity_OnCreateLogger $$
CREATE TRIGGER activity_OnCreateLogger AFTER INSERT ON activity
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Activity","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS activity_OnUpdateLogger $$
CREATE TRIGGER activity_OnUpdateLogger AFTER UPDATE ON activity
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Activity","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS activity_OnDeleteLogger $$
CREATE TRIGGER activity_OnDeleteLogger AFTER DELETE ON activity
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Activity","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Bucket ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS bucket_OnCreateLogger $$
CREATE TRIGGER bucket_OnCreateLogger AFTER INSERT ON bucket
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Bucket","INSERT",NEW.BucketId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS bucket_OnUpdateLogger $$
CREATE TRIGGER bucket_OnUpdateLogger AFTER UPDATE ON bucket
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Bucket","UPDATE",NEW.BucketId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS bucket_OnDeleteLogger $$
CREATE TRIGGER bucket_OnDeleteLogger AFTER DELETE ON bucket
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Bucket","DELETE",OLD.BucketId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Chat ######################################

delimiter $$
DROP TRIGGER IF EXISTS chat_OnCreateLogger $$
CREATE TRIGGER chat_OnCreateLogger AFTER INSERT ON chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","INSERT",NEW.ChatKey, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS chat_OnUpdateLogger $$
CREATE TRIGGER chat_OnUpdateLogger AFTER UPDATE ON chat
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","UPDATE",NEW.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS chat_OnDeleteLogger $$
CREATE TRIGGER chat_OnDeleteLogger AFTER DELETE ON chat
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetStr,CreatedSe) VALUES ("Chat","DELETE",OLD.ChatKey, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ Comment ######################################

delimiter $$
DROP TRIGGER IF EXISTS comment_OnCreateLogger $$
CREATE TRIGGER comment_OnCreateLogger AFTER INSERT ON comment
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS comment_OnUpdateLogger $$
CREATE TRIGGER comment_OnUpdateLogger AFTER UPDATE ON comment
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS comment_OnDeleteLogger $$
CREATE TRIGGER comment_OnDeleteLogger AFTER DELETE ON comment
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Comment","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ DirectMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS direct_message_OnCreateLogger $$
CREATE TRIGGER direct_message_OnCreateLogger AFTER INSERT ON direct_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","INSERT",NEW.MessageId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS direct_message_OnUpdateLogger $$
CREATE TRIGGER direct_message_OnUpdateLogger AFTER UPDATE ON direct_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","UPDATE",NEW.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS direct_message_OnDeleteLogger $$
CREATE TRIGGER direct_message_OnDeleteLogger AFTER DELETE ON direct_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectMessage","DELETE",OLD.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ DirectToMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS direct_to_message_OnCreateLogger $$
CREATE TRIGGER direct_to_message_OnCreateLogger AFTER INSERT ON direct_to_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectToMessage","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS direct_to_message_OnUpdateLogger $$
CREATE TRIGGER direct_to_message_OnUpdateLogger AFTER UPDATE ON direct_to_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectToMessage","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS direct_to_message_OnDeleteLogger $$
CREATE TRIGGER direct_to_message_OnDeleteLogger AFTER DELETE ON direct_to_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectToMessage","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ DirectUpdate ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS direct_update_OnCreateLogger $$
CREATE TRIGGER direct_update_OnCreateLogger AFTER INSERT ON direct_update
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectUpdate","INSERT",NEW.DirectUpdateId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS direct_update_OnUpdateLogger $$
CREATE TRIGGER direct_update_OnUpdateLogger AFTER UPDATE ON direct_update
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectUpdate","UPDATE",NEW.DirectUpdateId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS direct_update_OnDeleteLogger $$
CREATE TRIGGER direct_update_OnDeleteLogger AFTER DELETE ON direct_update
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("DirectUpdate","DELETE",OLD.DirectUpdateId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingList ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_OnCreateLogger $$
CREATE TRIGGER following_list_OnCreateLogger AFTER INSERT ON following_list
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_OnUpdateLogger $$
CREATE TRIGGER following_list_OnUpdateLogger AFTER UPDATE ON following_list
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_OnDeleteLogger $$
CREATE TRIGGER following_list_OnDeleteLogger AFTER DELETE ON following_list
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingList","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingListMember ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_member_OnCreateLogger $$
CREATE TRIGGER following_list_member_OnCreateLogger AFTER INSERT ON following_list_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_OnUpdateLogger $$
CREATE TRIGGER following_list_member_OnUpdateLogger AFTER UPDATE ON following_list_member
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_OnDeleteLogger $$
CREATE TRIGGER following_list_member_OnDeleteLogger AFTER DELETE ON following_list_member
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMember","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ FollowingListMemberHistory ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS following_list_member_history_OnCreateLogger $$
CREATE TRIGGER following_list_member_history_OnCreateLogger AFTER INSERT ON following_list_member_history
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberHistory","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_history_OnUpdateLogger $$
CREATE TRIGGER following_list_member_history_OnUpdateLogger AFTER UPDATE ON following_list_member_history
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberHistory","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS following_list_member_history_OnDeleteLogger $$
CREATE TRIGGER following_list_member_history_OnDeleteLogger AFTER DELETE ON following_list_member_history
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("FollowingListMemberHistory","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GeneralLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS general_log_OnCreateLogger $$
CREATE TRIGGER general_log_OnCreateLogger AFTER INSERT ON general_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GeneralLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS general_log_OnUpdateLogger $$
CREATE TRIGGER general_log_OnUpdateLogger AFTER UPDATE ON general_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GeneralLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS general_log_OnDeleteLogger $$
CREATE TRIGGER general_log_OnDeleteLogger AFTER DELETE ON general_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GeneralLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Group ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_OnCreateLogger $$
CREATE TRIGGER group_OnCreateLogger AFTER INSERT ON group
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","INSERT",NEW.GroupId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_OnUpdateLogger $$
CREATE TRIGGER group_OnUpdateLogger AFTER UPDATE ON group
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","UPDATE",NEW.GroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_OnDeleteLogger $$
CREATE TRIGGER group_OnDeleteLogger AFTER DELETE ON group
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Group","DELETE",OLD.GroupId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupMember ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_member_OnCreateLogger $$
CREATE TRIGGER group_member_OnCreateLogger AFTER INSERT ON group_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnUpdateLogger $$
CREATE TRIGGER group_member_OnUpdateLogger AFTER UPDATE ON group_member
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_member_OnDeleteLogger $$
CREATE TRIGGER group_member_OnDeleteLogger AFTER DELETE ON group_member
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMember","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_message_OnCreateLogger $$
CREATE TRIGGER group_message_OnCreateLogger AFTER INSERT ON group_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","INSERT",NEW.MessageId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_message_OnUpdateLogger $$
CREATE TRIGGER group_message_OnUpdateLogger AFTER UPDATE ON group_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","UPDATE",NEW.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_message_OnDeleteLogger $$
CREATE TRIGGER group_message_OnDeleteLogger AFTER DELETE ON group_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupMessage","DELETE",OLD.MessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ GroupToMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS group_to_message_OnCreateLogger $$
CREATE TRIGGER group_to_message_OnCreateLogger AFTER INSERT ON group_to_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupToMessage","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS group_to_message_OnUpdateLogger $$
CREATE TRIGGER group_to_message_OnUpdateLogger AFTER UPDATE ON group_to_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupToMessage","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS group_to_message_OnDeleteLogger $$
CREATE TRIGGER group_to_message_OnDeleteLogger AFTER DELETE ON group_to_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("GroupToMessage","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Like ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS likes_OnCreateLogger $$
CREATE TRIGGER likes_OnCreateLogger AFTER INSERT ON likes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS likes_OnUpdateLogger $$
CREATE TRIGGER likes_OnUpdateLogger AFTER UPDATE ON likes
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS likes_OnDeleteLogger $$
CREATE TRIGGER likes_OnDeleteLogger AFTER DELETE ON likes
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Like","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ LogChange ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS log_changes_OnCreateLogger $$
CREATE TRIGGER log_changes_OnCreateLogger AFTER INSERT ON log_changes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("LogChange","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS log_changes_OnUpdateLogger $$
CREATE TRIGGER log_changes_OnUpdateLogger AFTER UPDATE ON log_changes
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("LogChange","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS log_changes_OnDeleteLogger $$
CREATE TRIGGER log_changes_OnDeleteLogger AFTER DELETE ON log_changes
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("LogChange","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Media ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS media_OnCreateLogger $$
CREATE TRIGGER media_OnCreateLogger AFTER INSERT ON media
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS media_OnUpdateLogger $$
CREATE TRIGGER media_OnUpdateLogger AFTER UPDATE ON media
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS media_OnDeleteLogger $$
CREATE TRIGGER media_OnDeleteLogger AFTER DELETE ON media
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Media","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ MessageFile ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS message_file_OnCreateLogger $$
CREATE TRIGGER message_file_OnCreateLogger AFTER INSERT ON message_file
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","INSERT",NEW.MessageFileId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS message_file_OnUpdateLogger $$
CREATE TRIGGER message_file_OnUpdateLogger AFTER UPDATE ON message_file
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","UPDATE",NEW.MessageFileId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS message_file_OnDeleteLogger $$
CREATE TRIGGER message_file_OnDeleteLogger AFTER DELETE ON message_file
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("MessageFile","DELETE",OLD.MessageFileId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Notification ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS notification_OnCreateLogger $$
CREATE TRIGGER notification_OnCreateLogger AFTER INSERT ON notification
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notification","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS notification_OnUpdateLogger $$
CREATE TRIGGER notification_OnUpdateLogger AFTER UPDATE ON notification
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notification","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS notification_OnDeleteLogger $$
CREATE TRIGGER notification_OnDeleteLogger AFTER DELETE ON notification
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Notification","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ NotificationRemoved ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS notification_removed_OnCreateLogger $$
CREATE TRIGGER notification_removed_OnCreateLogger AFTER INSERT ON notification_removed
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotificationRemoved","INSERT",NEW.NotificationId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS notification_removed_OnUpdateLogger $$
CREATE TRIGGER notification_removed_OnUpdateLogger AFTER UPDATE ON notification_removed
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotificationRemoved","UPDATE",NEW.NotificationId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS notification_removed_OnDeleteLogger $$
CREATE TRIGGER notification_removed_OnDeleteLogger AFTER DELETE ON notification_removed
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("NotificationRemoved","DELETE",OLD.NotificationId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Offline ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS offline_OnCreateLogger $$
CREATE TRIGGER offline_OnCreateLogger AFTER INSERT ON offline
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Offline","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS offline_OnUpdateLogger $$
CREATE TRIGGER offline_OnUpdateLogger AFTER UPDATE ON offline
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Offline","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS offline_OnDeleteLogger $$
CREATE TRIGGER offline_OnDeleteLogger AFTER DELETE ON offline
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Offline","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ OldMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS old_messages_OnCreateLogger $$
CREATE TRIGGER old_messages_OnCreateLogger AFTER INSERT ON old_messages
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMessage","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS old_messages_OnUpdateLogger $$
CREATE TRIGGER old_messages_OnUpdateLogger AFTER UPDATE ON old_messages
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMessage","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS old_messages_OnDeleteLogger $$
CREATE TRIGGER old_messages_OnDeleteLogger AFTER DELETE ON old_messages
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMessage","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ OldMsgFile ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS old_msg_file_OnCreateLogger $$
CREATE TRIGGER old_msg_file_OnCreateLogger AFTER INSERT ON old_msg_file
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgFile","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS old_msg_file_OnUpdateLogger $$
CREATE TRIGGER old_msg_file_OnUpdateLogger AFTER UPDATE ON old_msg_file
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgFile","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS old_msg_file_OnDeleteLogger $$
CREATE TRIGGER old_msg_file_OnDeleteLogger AFTER DELETE ON old_msg_file
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgFile","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ OldMsgPush ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS old_msg_push_OnCreateLogger $$
CREATE TRIGGER old_msg_push_OnCreateLogger AFTER INSERT ON old_msg_push
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPush","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS old_msg_push_OnUpdateLogger $$
CREATE TRIGGER old_msg_push_OnUpdateLogger AFTER UPDATE ON old_msg_push
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPush","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS old_msg_push_OnDeleteLogger $$
CREATE TRIGGER old_msg_push_OnDeleteLogger AFTER DELETE ON old_msg_push
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPush","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ OldMsgPushEvent ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS old_msg_push_event_OnCreateLogger $$
CREATE TRIGGER old_msg_push_event_OnCreateLogger AFTER INSERT ON old_msg_push_event
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPushEvent","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS old_msg_push_event_OnUpdateLogger $$
CREATE TRIGGER old_msg_push_event_OnUpdateLogger AFTER UPDATE ON old_msg_push_event
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPushEvent","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS old_msg_push_event_OnDeleteLogger $$
CREATE TRIGGER old_msg_push_event_OnDeleteLogger AFTER DELETE ON old_msg_push_event
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("OldMsgPushEvent","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PhoneContact ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger $$
CREATE TRIGGER phone_contacts_OnCreateLogger AFTER INSERT ON phone_contacts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger $$
CREATE TRIGGER phone_contacts_OnUpdateLogger AFTER UPDATE ON phone_contacts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger $$
CREATE TRIGGER phone_contacts_OnDeleteLogger AFTER DELETE ON phone_contacts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PhoneContact","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Photo ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS photo_OnCreateLogger $$
CREATE TRIGGER photo_OnCreateLogger AFTER INSERT ON photo
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Photo","INSERT",NEW.PhotoId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS photo_OnUpdateLogger $$
CREATE TRIGGER photo_OnUpdateLogger AFTER UPDATE ON photo
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Photo","UPDATE",NEW.PhotoId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS photo_OnDeleteLogger $$
CREATE TRIGGER photo_OnDeleteLogger AFTER DELETE ON photo
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Photo","DELETE",OLD.PhotoId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Post ######################################

delimiter $$
DROP TRIGGER IF EXISTS post_OnCreateLogger $$
CREATE TRIGGER post_OnCreateLogger AFTER INSERT ON post
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS post_OnUpdateLogger $$
CREATE TRIGGER post_OnUpdateLogger AFTER UPDATE ON post
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS post_OnDeleteLogger $$
CREATE TRIGGER post_OnDeleteLogger AFTER DELETE ON post
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Post","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ PushEvent ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS push_event_OnCreateLogger $$
CREATE TRIGGER push_event_OnCreateLogger AFTER INSERT ON push_event
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushEvent","INSERT",NEW.PushEventId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS push_event_OnUpdateLogger $$
CREATE TRIGGER push_event_OnUpdateLogger AFTER UPDATE ON push_event
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushEvent","UPDATE",NEW.PushEventId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS push_event_OnDeleteLogger $$
CREATE TRIGGER push_event_OnDeleteLogger AFTER DELETE ON push_event
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushEvent","DELETE",OLD.PushEventId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ PushMessage ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS push_message_OnCreateLogger $$
CREATE TRIGGER push_message_OnCreateLogger AFTER INSERT ON push_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushMessage","INSERT",NEW.PushMessageId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS push_message_OnUpdateLogger $$
CREATE TRIGGER push_message_OnUpdateLogger AFTER UPDATE ON push_message
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushMessage","UPDATE",NEW.PushMessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS push_message_OnDeleteLogger $$
CREATE TRIGGER push_message_OnDeleteLogger AFTER DELETE ON push_message
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("PushMessage","DELETE",OLD.PushMessageId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ RecommendUser ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS recommend_user_OnCreateLogger $$
CREATE TRIGGER recommend_user_OnCreateLogger AFTER INSERT ON recommend_user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("RecommendUser","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS recommend_user_OnUpdateLogger $$
CREATE TRIGGER recommend_user_OnUpdateLogger AFTER UPDATE ON recommend_user
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("RecommendUser","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS recommend_user_OnDeleteLogger $$
CREATE TRIGGER recommend_user_OnDeleteLogger AFTER DELETE ON recommend_user
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("RecommendUser","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Room ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS room_OnCreateLogger $$
CREATE TRIGGER room_OnCreateLogger AFTER INSERT ON room
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Room","INSERT",NEW.RoomId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS room_OnUpdateLogger $$
CREATE TRIGGER room_OnUpdateLogger AFTER UPDATE ON room
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Room","UPDATE",NEW.RoomId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS room_OnDeleteLogger $$
CREATE TRIGGER room_OnDeleteLogger AFTER DELETE ON room
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Room","DELETE",OLD.RoomId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SearchClicked ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS search_clicked_OnCreateLogger $$
CREATE TRIGGER search_clicked_OnCreateLogger AFTER INSERT ON search_clicked
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS search_clicked_OnUpdateLogger $$
CREATE TRIGGER search_clicked_OnUpdateLogger AFTER UPDATE ON search_clicked
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS search_clicked_OnDeleteLogger $$
CREATE TRIGGER search_clicked_OnDeleteLogger AFTER DELETE ON search_clicked
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SearchClicked","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Session ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS session_OnCreateLogger $$
CREATE TRIGGER session_OnCreateLogger AFTER INSERT ON session
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS session_OnUpdateLogger $$
CREATE TRIGGER session_OnUpdateLogger AFTER UPDATE ON session
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS session_OnDeleteLogger $$
CREATE TRIGGER session_OnDeleteLogger AFTER DELETE ON session
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Session","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SettingClient ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS setting_client_OnCreateLogger $$
CREATE TRIGGER setting_client_OnCreateLogger AFTER INSERT ON setting_client
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS setting_client_OnUpdateLogger $$
CREATE TRIGGER setting_client_OnUpdateLogger AFTER UPDATE ON setting_client
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS setting_client_OnDeleteLogger $$
CREATE TRIGGER setting_client_OnDeleteLogger AFTER DELETE ON setting_client
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingClient","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ SettingNotification ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger $$
CREATE TRIGGER setting_notifications_OnCreateLogger AFTER INSERT ON setting_notifications
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger $$
CREATE TRIGGER setting_notifications_OnUpdateLogger AFTER UPDATE ON setting_notifications
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger $$
CREATE TRIGGER setting_notifications_OnDeleteLogger AFTER DELETE ON setting_notifications
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("SettingNotification","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ Tag ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS tag_OnCreateLogger $$
CREATE TRIGGER tag_OnCreateLogger AFTER INSERT ON tag
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS tag_OnUpdateLogger $$
CREATE TRIGGER tag_OnUpdateLogger AFTER UPDATE ON tag
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS tag_OnDeleteLogger $$
CREATE TRIGGER tag_OnDeleteLogger AFTER DELETE ON tag
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("Tag","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ TagsPost ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS tags_posts_OnCreateLogger $$
CREATE TRIGGER tags_posts_OnCreateLogger AFTER INSERT ON tags_posts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS tags_posts_OnUpdateLogger $$
CREATE TRIGGER tags_posts_OnUpdateLogger AFTER UPDATE ON tags_posts
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS tags_posts_OnDeleteLogger $$
CREATE TRIGGER tags_posts_OnDeleteLogger AFTER DELETE ON tags_posts
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TagsPost","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ TestChat ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS test_chat_OnCreateLogger $$
CREATE TRIGGER test_chat_OnCreateLogger AFTER INSERT ON test_chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TestChat","INSERT",NEW.Id4, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS test_chat_OnUpdateLogger $$
CREATE TRIGGER test_chat_OnUpdateLogger AFTER UPDATE ON test_chat
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TestChat","UPDATE",NEW.Id4, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS test_chat_OnDeleteLogger $$
CREATE TRIGGER test_chat_OnDeleteLogger AFTER DELETE ON test_chat
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TestChat","DELETE",OLD.Id4, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ TriggerLog ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS trigger_log_OnCreateLogger $$
CREATE TRIGGER trigger_log_OnCreateLogger AFTER INSERT ON trigger_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS trigger_log_OnUpdateLogger $$
CREATE TRIGGER trigger_log_OnUpdateLogger AFTER UPDATE ON trigger_log
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS trigger_log_OnDeleteLogger $$
CREATE TRIGGER trigger_log_OnDeleteLogger AFTER DELETE ON trigger_log
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("TriggerLog","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ User ######################################

delimiter $$
DROP TRIGGER IF EXISTS user_OnCreateLogger $$
CREATE TRIGGER user_OnCreateLogger AFTER INSERT ON user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_OnUpdateLogger $$
CREATE TRIGGER user_OnUpdateLogger AFTER UPDATE ON user
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_OnDeleteLogger $$
CREATE TRIGGER user_OnDeleteLogger AFTER DELETE ON user
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("User","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


delimiter ;
################################ UserMetaInfo ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS user_meta_info_OnCreateLogger $$
CREATE TRIGGER user_meta_info_OnCreateLogger AFTER INSERT ON user_meta_info
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","INSERT",NEW.Id, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_meta_info_OnUpdateLogger $$
CREATE TRIGGER user_meta_info_OnUpdateLogger AFTER UPDATE ON user_meta_info
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","UPDATE",NEW.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_meta_info_OnDeleteLogger $$
CREATE TRIGGER user_meta_info_OnDeleteLogger AFTER DELETE ON user_meta_info
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserMetaInfo","DELETE",OLD.Id, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/
################################ UserPassword ######################################

/* #### delimiter $$
DROP TRIGGER IF EXISTS user_password_OnCreateLogger $$
CREATE TRIGGER user_password_OnCreateLogger AFTER INSERT ON user_password
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","INSERT",NEW.UserId, UNIX_TIMESTAMP(NOW()) );
  END;
$$

DROP TRIGGER IF EXISTS user_password_OnUpdateLogger $$
CREATE TRIGGER user_password_OnUpdateLogger AFTER UPDATE ON user_password
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","UPDATE",NEW.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$

DROP TRIGGER IF EXISTS user_password_OnDeleteLogger $$
CREATE TRIGGER user_password_OnDeleteLogger AFTER DELETE ON user_password
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,TargetId,CreatedSe) VALUES ("UserPassword","DELETE",OLD.UserId, UNIX_TIMESTAMP(NOW()));
  END;
$$


 #### delimiter ;*/

###############################################################################################
################################## Delete of all triggers #####################################
/*

### Activity ##
DROP TRIGGER IF EXISTS activity_OnCreateLogger ;
DROP TRIGGER IF EXISTS activity_OnUpdateLogger ;
DROP TRIGGER IF EXISTS activity_OnDeleteLogger ;
### Bucket ##
DROP TRIGGER IF EXISTS bucket_OnCreateLogger ;
DROP TRIGGER IF EXISTS bucket_OnUpdateLogger ;
DROP TRIGGER IF EXISTS bucket_OnDeleteLogger ;
### Chat ##
DROP TRIGGER IF EXISTS chat_OnCreateLogger ;
DROP TRIGGER IF EXISTS chat_OnUpdateLogger ;
DROP TRIGGER IF EXISTS chat_OnDeleteLogger ;
### Comment ##
DROP TRIGGER IF EXISTS comment_OnCreateLogger ;
DROP TRIGGER IF EXISTS comment_OnUpdateLogger ;
DROP TRIGGER IF EXISTS comment_OnDeleteLogger ;
### DirectMessage ##
DROP TRIGGER IF EXISTS direct_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS direct_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS direct_message_OnDeleteLogger ;
### DirectToMessage ##
DROP TRIGGER IF EXISTS direct_to_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS direct_to_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS direct_to_message_OnDeleteLogger ;
### DirectUpdate ##
DROP TRIGGER IF EXISTS direct_update_OnCreateLogger ;
DROP TRIGGER IF EXISTS direct_update_OnUpdateLogger ;
DROP TRIGGER IF EXISTS direct_update_OnDeleteLogger ;
### FollowingList ##
DROP TRIGGER IF EXISTS following_list_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_OnDeleteLogger ;
### FollowingListMember ##
DROP TRIGGER IF EXISTS following_list_member_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_member_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_member_OnDeleteLogger ;
### FollowingListMemberHistory ##
DROP TRIGGER IF EXISTS following_list_member_history_OnCreateLogger ;
DROP TRIGGER IF EXISTS following_list_member_history_OnUpdateLogger ;
DROP TRIGGER IF EXISTS following_list_member_history_OnDeleteLogger ;
### GeneralLog ##
DROP TRIGGER IF EXISTS general_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS general_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS general_log_OnDeleteLogger ;
### Group ##
DROP TRIGGER IF EXISTS group_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_OnDeleteLogger ;
### GroupMember ##
DROP TRIGGER IF EXISTS group_member_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_member_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_member_OnDeleteLogger ;
### GroupMessage ##
DROP TRIGGER IF EXISTS group_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_message_OnDeleteLogger ;
### GroupToMessage ##
DROP TRIGGER IF EXISTS group_to_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS group_to_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS group_to_message_OnDeleteLogger ;
### Like ##
DROP TRIGGER IF EXISTS likes_OnCreateLogger ;
DROP TRIGGER IF EXISTS likes_OnUpdateLogger ;
DROP TRIGGER IF EXISTS likes_OnDeleteLogger ;
### LogChange ##
DROP TRIGGER IF EXISTS log_changes_OnCreateLogger ;
DROP TRIGGER IF EXISTS log_changes_OnUpdateLogger ;
DROP TRIGGER IF EXISTS log_changes_OnDeleteLogger ;
### Media ##
DROP TRIGGER IF EXISTS media_OnCreateLogger ;
DROP TRIGGER IF EXISTS media_OnUpdateLogger ;
DROP TRIGGER IF EXISTS media_OnDeleteLogger ;
### MessageFile ##
DROP TRIGGER IF EXISTS message_file_OnCreateLogger ;
DROP TRIGGER IF EXISTS message_file_OnUpdateLogger ;
DROP TRIGGER IF EXISTS message_file_OnDeleteLogger ;
### Notification ##
DROP TRIGGER IF EXISTS notification_OnCreateLogger ;
DROP TRIGGER IF EXISTS notification_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notification_OnDeleteLogger ;
### NotificationRemoved ##
DROP TRIGGER IF EXISTS notification_removed_OnCreateLogger ;
DROP TRIGGER IF EXISTS notification_removed_OnUpdateLogger ;
DROP TRIGGER IF EXISTS notification_removed_OnDeleteLogger ;
### Offline ##
DROP TRIGGER IF EXISTS offline_OnCreateLogger ;
DROP TRIGGER IF EXISTS offline_OnUpdateLogger ;
DROP TRIGGER IF EXISTS offline_OnDeleteLogger ;
### OldMessage ##
DROP TRIGGER IF EXISTS old_messages_OnCreateLogger ;
DROP TRIGGER IF EXISTS old_messages_OnUpdateLogger ;
DROP TRIGGER IF EXISTS old_messages_OnDeleteLogger ;
### OldMsgFile ##
DROP TRIGGER IF EXISTS old_msg_file_OnCreateLogger ;
DROP TRIGGER IF EXISTS old_msg_file_OnUpdateLogger ;
DROP TRIGGER IF EXISTS old_msg_file_OnDeleteLogger ;
### OldMsgPush ##
DROP TRIGGER IF EXISTS old_msg_push_OnCreateLogger ;
DROP TRIGGER IF EXISTS old_msg_push_OnUpdateLogger ;
DROP TRIGGER IF EXISTS old_msg_push_OnDeleteLogger ;
### OldMsgPushEvent ##
DROP TRIGGER IF EXISTS old_msg_push_event_OnCreateLogger ;
DROP TRIGGER IF EXISTS old_msg_push_event_OnUpdateLogger ;
DROP TRIGGER IF EXISTS old_msg_push_event_OnDeleteLogger ;
### PhoneContact ##
DROP TRIGGER IF EXISTS phone_contacts_OnCreateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS phone_contacts_OnDeleteLogger ;
### Photo ##
DROP TRIGGER IF EXISTS photo_OnCreateLogger ;
DROP TRIGGER IF EXISTS photo_OnUpdateLogger ;
DROP TRIGGER IF EXISTS photo_OnDeleteLogger ;
### Post ##
DROP TRIGGER IF EXISTS post_OnCreateLogger ;
DROP TRIGGER IF EXISTS post_OnUpdateLogger ;
DROP TRIGGER IF EXISTS post_OnDeleteLogger ;
### PushEvent ##
DROP TRIGGER IF EXISTS push_event_OnCreateLogger ;
DROP TRIGGER IF EXISTS push_event_OnUpdateLogger ;
DROP TRIGGER IF EXISTS push_event_OnDeleteLogger ;
### PushMessage ##
DROP TRIGGER IF EXISTS push_message_OnCreateLogger ;
DROP TRIGGER IF EXISTS push_message_OnUpdateLogger ;
DROP TRIGGER IF EXISTS push_message_OnDeleteLogger ;
### RecommendUser ##
DROP TRIGGER IF EXISTS recommend_user_OnCreateLogger ;
DROP TRIGGER IF EXISTS recommend_user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS recommend_user_OnDeleteLogger ;
### Room ##
DROP TRIGGER IF EXISTS room_OnCreateLogger ;
DROP TRIGGER IF EXISTS room_OnUpdateLogger ;
DROP TRIGGER IF EXISTS room_OnDeleteLogger ;
### SearchClicked ##
DROP TRIGGER IF EXISTS search_clicked_OnCreateLogger ;
DROP TRIGGER IF EXISTS search_clicked_OnUpdateLogger ;
DROP TRIGGER IF EXISTS search_clicked_OnDeleteLogger ;
### Session ##
DROP TRIGGER IF EXISTS session_OnCreateLogger ;
DROP TRIGGER IF EXISTS session_OnUpdateLogger ;
DROP TRIGGER IF EXISTS session_OnDeleteLogger ;
### SettingClient ##
DROP TRIGGER IF EXISTS setting_client_OnCreateLogger ;
DROP TRIGGER IF EXISTS setting_client_OnUpdateLogger ;
DROP TRIGGER IF EXISTS setting_client_OnDeleteLogger ;
### SettingNotification ##
DROP TRIGGER IF EXISTS setting_notifications_OnCreateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnUpdateLogger ;
DROP TRIGGER IF EXISTS setting_notifications_OnDeleteLogger ;
### Tag ##
DROP TRIGGER IF EXISTS tag_OnCreateLogger ;
DROP TRIGGER IF EXISTS tag_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tag_OnDeleteLogger ;
### TagsPost ##
DROP TRIGGER IF EXISTS tags_posts_OnCreateLogger ;
DROP TRIGGER IF EXISTS tags_posts_OnUpdateLogger ;
DROP TRIGGER IF EXISTS tags_posts_OnDeleteLogger ;
### TestChat ##
DROP TRIGGER IF EXISTS test_chat_OnCreateLogger ;
DROP TRIGGER IF EXISTS test_chat_OnUpdateLogger ;
DROP TRIGGER IF EXISTS test_chat_OnDeleteLogger ;
### TriggerLog ##
DROP TRIGGER IF EXISTS trigger_log_OnCreateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnUpdateLogger ;
DROP TRIGGER IF EXISTS trigger_log_OnDeleteLogger ;
### User ##
DROP TRIGGER IF EXISTS user_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_OnDeleteLogger ;
### UserMetaInfo ##
DROP TRIGGER IF EXISTS user_meta_info_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_meta_info_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_meta_info_OnDeleteLogger ;
### UserPassword ##
DROP TRIGGER IF EXISTS user_password_OnCreateLogger ;
DROP TRIGGER IF EXISTS user_password_OnUpdateLogger ;
DROP TRIGGER IF EXISTS user_password_OnDeleteLogger ;
*/