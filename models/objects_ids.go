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
const ACTION_TYPE_POST_LIKED = 200 //GENERAL
const ACTION_TYPE_POST_TEXT_LIKED = 201
const ACTION_TYPE_POST_PHOTO_LIKED = 202
const ACTION_TYPE_POST_VIDEO_LIKED = 203

//22x POSTS COMMENTED
const ACTION_TYPE_POST_COMMENTED = 220
const ACTION_TYPE_POST_TEXT_COMMENTED = 221
const ACTION_TYPE_POST_PHOTO_COMMENTED = 222
const ACTION_TYPE_POST_VIDEO_COMMENTED = 223

const ACTION_TYPE_FOLLOWED_USER = 250 //for notification USER is you

/////////////////////////////////////////////////////////////////
/*
const OBJECT_NOTIFICATION_POST_ADDED = 23
const OBJECT_NOTIFICATION_POST_COMMENTED = 23
const OBJECT_NOTIFICATION_POST_LIKED = 23
*/
