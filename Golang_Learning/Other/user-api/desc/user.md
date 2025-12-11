### 1. "获取用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
	UserID int64 `json:"user_id"`
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserID int64 `json:"user_id"`
	Nickname string `json:"nickname"`
}
```

