namespace go common.request

struct PageInfo {
    /** 页码 */
    1: i32 page;
    /** 每页大小 */
    2: i32 page_size;
    /** 关键字 */
    3: string keyword;
}

struct IdReq {
    /** 主键ID */
    1: i64 id (api.path="id");
}

struct IdsReq {
    /** 主键ID 列表 */
    1: list<i64> ids (api.query="ids");
}

struct Empty {}