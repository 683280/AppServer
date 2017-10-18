package db

var SQL_LoginUser = "select * from user where user_name=? and user_password=?"

var SQL_RegisterUser = "insert into user (user_uuid,user_name,user_password) values (?,?,?)"

var SQL_LoginUserById = "select * from user where user_id=? and user_password=?"

var SQL_GetUserById = "select * from user where user_id=?"

var SQL_NameIsExis = "select count(user_name) from user where user_name=?"

var SQL_GetUUIDById = "select user_uuid from user where user_id=?"

var SQL_GetTopicForId = "select * from topic where t_id=?"

var SQL_GetTopicForTop = "select * from topic where id=?"

var SQL_GetAllFriends = "select * from user where user_id in (select friends.friend from friends where friends.me in(select user_id from user where user_uuid=?))"

var SQL_InsertMessage = "insert into message (msg_from,msg_to,msg_data,time) values (?,?,?,?)"

var SQL_GetUserAllMsg = "select * from message where msg_to=?"

var SQL_DelUserMsgByTime = "delete from message where time=?"