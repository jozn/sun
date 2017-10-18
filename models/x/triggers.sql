

delimiter |
CREATE TRIGGER activity_Create AFTER INSERT ON activity
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("activity","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER bucket_Create AFTER INSERT ON bucket
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("bucket","insert",NEW.BucketId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER chat_Create AFTER INSERT ON chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("chat","insert",NEW.ChatKey);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER comments_Create AFTER INSERT ON comments
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("comments","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER direct_message_Create AFTER INSERT ON direct_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("direct_message","insert",NEW.MessageId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER direct_to_message_Create AFTER INSERT ON direct_to_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("direct_to_message","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER direct_update_Create AFTER INSERT ON direct_update
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("direct_update","insert",NEW.DirectUpdateId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER following_list_Create AFTER INSERT ON following_list
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("following_list","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER following_list_member_Create AFTER INSERT ON following_list_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("following_list_member","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER following_list_member_history_Create AFTER INSERT ON following_list_member_history
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("following_list_member_history","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER general_log_Create AFTER INSERT ON general_log
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("general_log","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER group_Create AFTER INSERT ON group
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("group","insert",NEW.GroupId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER group_member_Create AFTER INSERT ON group_member
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("group_member","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER group_message_Create AFTER INSERT ON group_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("group_message","insert",NEW.MessageId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER group_to_message_Create AFTER INSERT ON group_to_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("group_to_message","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER likes_Create AFTER INSERT ON likes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("likes","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER log_changes_Create AFTER INSERT ON log_changes
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("log_changes","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER media_Create AFTER INSERT ON media
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("media","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER message_file_Create AFTER INSERT ON message_file
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("message_file","insert",NEW.MessageFileId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER notification_Create AFTER INSERT ON notification
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("notification","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER notification_removed_Create AFTER INSERT ON notification_removed
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("notification_removed","insert",NEW.NotificationId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER old_messages_Create AFTER INSERT ON old_messages
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("old_messages","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER old_msg_file_Create AFTER INSERT ON old_msg_file
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("old_msg_file","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER old_msg_push_Create AFTER INSERT ON old_msg_push
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("old_msg_push","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER old_msg_push_event_Create AFTER INSERT ON old_msg_push_event
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("old_msg_push_event","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER phone_contacts_Create AFTER INSERT ON phone_contacts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("phone_contacts","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER photo_Create AFTER INSERT ON photo
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("photo","insert",NEW.PhotoId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER post_Create AFTER INSERT ON post
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("post","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER push_event_Create AFTER INSERT ON push_event
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("push_event","insert",NEW.PushEventId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER push_message_Create AFTER INSERT ON push_message
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("push_message","insert",NEW.PushMessageId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER recommend_user_Create AFTER INSERT ON recommend_user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("recommend_user","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER room_Create AFTER INSERT ON room
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("room","insert",NEW.RoomId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER search_clicked_Create AFTER INSERT ON search_clicked
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("search_clicked","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER session_Create AFTER INSERT ON session
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("session","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER setting_client_Create AFTER INSERT ON setting_client
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("setting_client","insert",NEW.UserId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER setting_notifications_Create AFTER INSERT ON setting_notifications
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("setting_notifications","insert",NEW.UserId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER tags_Create AFTER INSERT ON tags
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("tags","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER tags_posts_Create AFTER INSERT ON tags_posts
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("tags_posts","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER test_chat_Create AFTER INSERT ON test_chat
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("test_chat","insert",NEW.Id4);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER user_Create AFTER INSERT ON user
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("user","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER user_meta_info_Create AFTER INSERT ON user_meta_info
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("user_meta_info","insert",NEW.Id);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;




delimiter |
CREATE TRIGGER user_password_Create AFTER INSERT ON user_password
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (TableName,ChangeType,TargetId) VALUES ("user_password","insert",NEW.UserId);
#     DELETE FROM test3 WHERE a3 = NEW.a1;
#     UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;


