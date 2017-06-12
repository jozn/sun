package models

const OBJECT_USER = 1
const OBJECT_POST = 2
const OBJECT_COMMENT = 3
const OBJECT_FOLLOWING = 4
const OBJECT_LIKE = 5

const OBJECT_MSG = 10 //+ 1XX ACTIONS

//POSTS SUB TYPES
const OBJECT_POST_TEXT = 20
const OBJECT_POST_IMAGE = 21
const OBJECT_POST_VIDEO = 22
const OBJECT_POST_URL = 23

////////////////////////////////////////////////////////////////
/// Action_Types For Notifications
// ALL Notifications ACTIONS 2XX GLOBAL OBJECTS TYPES
//SPEC: all actions un-actions have negative ex: 200:like => -200:unlike
//20x POSTS LIKED
const ACTION_TYPE_POST_LIKED = 11 //GENERAL
const ACTION_TYPE_POST_TEXT_LIKED = 201
const ACTION_TYPE_POST_PHOTO_LIKED = 202
const ACTION_TYPE_POST_VIDEO_LIKED = 203

//22x POSTS COMMENTED
const ACTION_TYPE_POST_COMMENTED = 21
const ACTION_TYPE_POST_TEXT_COMMENTED = 221
const ACTION_TYPE_POST_PHOTO_COMMENTED = 222
const ACTION_TYPE_POST_VIDEO_COMMENTED = 223

const ACTION_TYPE_FOLLOWED_USER = 1 //for notification USER is you

/////////////////////////////////////////////////////////////////
/*
const OBJECT_NOTIFICATION_POST_ADDED = 23
const OBJECT_NOTIFICATION_POST_COMMENTED = 23
const OBJECT_NOTIFICATION_POST_LIKED = 23
*/

//Messages
const MESSAGE_TEXT = 10
const MESSAGE_CONTACT = 12
const MESSAGE_lOCCASION = 14
const MESSAGE_STICKER = 18
const MESSAGE_POST = 30
const MESSAGE_IMAGE = 40
const MESSAGE_VIDEO = 42
const MESSAGE_AUDIO = 44
const MESSAGE_FILE = 46

//PUSH EVENTS
const MESSAGE_PUSH_EVENT_RECEIVED_TO_PEER = 1
const MESSAGE_PUSH_EVENT_DELETED_FROM_SERVER = 2
const MESSAGE_PUSH_EVENT_SEEN_BY_PEER = 3

const PB_CommandReceivedToServer = "PB_CommandReceivedToServer"
const PB_CommandReceivedToClient = "PB_CommandReceivedToClient"

const PB_RequestMsgAddMany = "PB_RequestMsgAddMany"
const PB_RequestMsgsSeen = "PB_RequestMsgsSeen"

const PB_PushMsgAddMany = "PB_PushMsgAddMany"
const PB_PushMsgEvents = "PB_PushMsgEvents"

//const PB_PushMsgsRemovedFromServer = "PB_PushMsgsRemovedFromServer"
