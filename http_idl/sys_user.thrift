namespace go sys

include "./base/base.thrift"

struct UserInfo {
    /** 主键ID */
    1: optional i64 id;
    /** 用户id */
    2: optional string user_id;
    /** 姓名 */
    5: string name;
    /** 年龄 */
    6: i32 age;
    /** 地址 */
    7: string address;
    /** 密码 */
    8: string password;
}

struct SaveUserReq {
    /** 姓名 */
    2: UserInfo userInfo;
}

struct SaveUserResp {
    /** 状态嘛 */
    1: i32 code;
    /** 提示信息 */
    2: string msg;
    /** 响应结果 */
    3: string data;
}

struct GetUserReq {
    /** 用户id */
    1: required string user_ids;
}

struct GetUserResp {
    /** 状态嘛 */
    1: i32 code;
    /** 提示信息 */
    2: string msg;
    /** 用户信息 */
    3: list<UserInfo> data;
}

struct UpdateUserResp {
    /** 状态嘛 */
    1: i32 code;
    /** 提示信息 */
    2: string msg;
    /** 响应结果 */
    3: string data;
}

struct UpdateUserReq {
    /** 用户ID */
    1: string user_id (api.path="id");
    /** 用户信息 */
    2: UserInfo userInfo;
}

service UserService {
    /** 保存用户 */
    SaveUserResp SaveUser(1: SaveUserReq req) (api.post="/api/v1/sys/user")
    /** 获取用户 */
    GetUserResp GetUser(1: GetUserReq req) (api.get="/api/v1/sys/user")
    /** 更新用户 */
    UpdateUserResp UpdateUser(1: UpdateUserReq req) (api.put="/api/v1/sys/user/:id")
}
