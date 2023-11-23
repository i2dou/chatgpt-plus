package admin

import (
	"chatplus/core"
	"chatplus/core/types"
	"chatplus/handler"
	"chatplus/store/model"
	"chatplus/store/vo"
	"chatplus/utils"
	"chatplus/utils/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ChatRoleHandler struct {
	handler.BaseHandler
	db *gorm.DB
}

func NewChatRoleHandler(app *core.AppServer, db *gorm.DB) *ChatRoleHandler {
	h := ChatRoleHandler{db: db}
	h.App = app
	return &h
}

// Save 创建或者更新某个角色
func (h *ChatRoleHandler) Save(c *gin.Context) {
	var data vo.ChatRole
	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}
	var role model.ChatRole
	err := utils.CopyObject(data, &role)
	if err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}
	role.Id = data.Id
	if data.CreatedAt > 0 {
		role.CreatedAt = time.Unix(data.CreatedAt, 0)
	}
	res := h.db.Save(&role)
	if res.Error != nil {
		resp.ERROR(c, "更新数据库失败！")
		return
	}
	// 填充 ID 数据
	data.Id = role.Id
	data.CreatedAt = role.CreatedAt.Unix()
	resp.SUCCESS(c, data)
}

func (h *ChatRoleHandler) List(c *gin.Context) {
	var items []model.ChatRole
	var roles = make([]vo.ChatRole, 0)
	res := h.db.Order("sort_num ASC").Find(&items)
	if res.Error != nil {
		resp.ERROR(c, "No data found")
		return
	}

	for _, v := range items {
		var role vo.ChatRole
		err := utils.CopyObject(v, &role)
		if err == nil {
			role.Id = v.Id
			role.CreatedAt = v.CreatedAt.Unix()
			role.UpdatedAt = v.UpdatedAt.Unix()
			roles = append(roles, role)
		}
	}

	resp.SUCCESS(c, roles)
}

// Sort 更新角色排序
func (h *ChatRoleHandler) Sort(c *gin.Context) {
	var data struct {
		Ids   []uint `json:"ids"`
		Sorts []int  `json:"sorts"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	for index, id := range data.Ids {
		res := h.db.Model(&model.ChatRole{}).Where("id = ?", id).Update("sort_num", data.Sorts[index])
		if res.Error != nil {
			resp.ERROR(c, "更新数据库失败！")
			return
		}
	}

	resp.SUCCESS(c)
}

func (h *ChatRoleHandler) Remove(c *gin.Context) {
	id := h.GetInt(c, "id", 0)
	if id <= 0 {
		resp.ERROR(c, types.InvalidArgs)
		return
	}

	res := h.db.Where("id = ?", id).Delete(&model.ChatRole{})
	if res.Error != nil {
		resp.ERROR(c, "删除失败！")
		return
	}
	resp.SUCCESS(c)
}
