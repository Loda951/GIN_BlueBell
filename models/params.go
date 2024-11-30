package models

// ParamsSignUp 定义注册请求的参数结构体
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogIn 定义登陆请求的参数结构体
type ParamsLogIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成1 反对-1 0取消
}

type ParamPostList struct {
	CommunityID int64  `form:"community_id" json:"community_id" binding:"omitempty"`
	Page        int64  `form:"page" json:"page"`
	Size        int64  `form:"size" json:"size"`
	Order       string `form:"order" json:"order"`
}

type ParamCommunityPostList struct {
	*ParamPostList
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
