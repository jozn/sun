/////////////////////////////////
UPDATE following_list_member SET ListId = UserId WHERE 1;`following_list_member`
////////////////////////////////
UPDATE `user` SET `FollowersCount`=0,`FollowingCount`=0 WHERE 1;
UPDATE `user` SET `PostsCount`=0 WHERE 1;
///////////////
//must be equal
SELECT SUM(FollowingCount) FROM `user` WHERE 1;
SELECT SUM(FollowersCount) FROM `user` WHERE 1;
//////////////////////


