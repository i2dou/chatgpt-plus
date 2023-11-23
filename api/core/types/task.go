package types

// TaskType 任务类别
type TaskType string

func (t TaskType) String() string {
	return string(t)
}

const (
	TaskImage     = TaskType("image")
	TaskUpscale   = TaskType("upscale")
	TaskVariation = TaskType("variation")
	TaskTxt2Img   = TaskType("text2img")
)

// TaskSrc 任务来源
type TaskSrc string

const (
	TaskSrcChat = TaskSrc("chat") // 来自聊天页面
	TaskSrcImg  = TaskSrc("img")  // 专业绘画页面
)

// MjTask MidJourney 任务
type MjTask struct {
	Id          int      `json:"id"`
	SessionId   string   `json:"session_id"`
	Src         TaskSrc  `json:"src"`
	Type        TaskType `json:"type"`
	UserId      int      `json:"user_id"`
	Prompt      string   `json:"prompt,omitempty"`
	ChatId      string   `json:"chat_id,omitempty"`
	RoleId      int      `json:"role_id,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	Index       int      `json:"index,omitempty"`
	MessageId   string   `json:"message_id,omitempty"`
	MessageHash string   `json:"message_hash,omitempty"`
	RetryCount  int      `json:"retry_count"`
}

type SdTask struct {
	Id         int          `json:"id"` // job 数据库ID
	SessionId  string       `json:"session_id"`
	Src        TaskSrc      `json:"src"`
	Type       TaskType     `json:"type"`
	UserId     int          `json:"user_id"`
	Prompt     string       `json:"prompt,omitempty"`
	Params     SdTaskParams `json:"params"`
	RetryCount int          `json:"retry_count"`
}

type SdTaskParams struct {
	TaskId         string  `json:"task_id"`
	Prompt         string  `json:"prompt"`          // 提示词
	NegativePrompt string  `json:"negative_prompt"` // 反向提示词
	Steps          int     `json:"steps"`           // 迭代步数，默认20
	Sampler        string  `json:"sampler"`         // 采样器
	FaceFix        bool    `json:"face_fix"`        // 面部修复
	CfgScale       float32 `json:"cfg_scale"`       //引导系数，默认 7
	Seed           int64   `json:"seed"`            // 随机数种子
	Height         int     `json:"height"`
	Width          int     `json:"width"`
	HdFix          bool    `json:"hd_fix"`         // 启用高清修复
	HdRedrawRate   float32 `json:"hd_redraw_rate"` // 高清修复重绘幅度
	HdScale        int     `json:"hd_scale"`       // 放大倍数
	HdScaleAlg     string  `json:"hd_scale_alg"`   // 放大算法
	HdSteps        int     `json:"hd_steps"`       // 高清修复迭代步数
}
