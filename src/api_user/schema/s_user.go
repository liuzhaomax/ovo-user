package schema

import (
	"errors"
	"fmt"
	"github.com/liuzhaomax/ovo-user/internal/core"
	"github.com/liuzhaomax/ovo-user/src/api_user/model"
	"github.com/liuzhaomax/ovo-user/src/utils"
)

type UserRes struct {
	UserID        string   `json:"userId"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
	DeletedAt     string   `json:"deletedAt"`
	Username      string   `json:"username"`
	Mobile        string   `json:"mobile"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"emailVerified"`
	Role          string   `json:"role"`
	Groups        []string `json:"groups"`
}

type RoleRes struct {
	Role       string   `json:"role"`
	Permission []string `json:"permission"`
}

type GroupRes struct {
	Group       string `json:"group"`
	ParentGroup string `json:"parentGroup"`
}

type PermissionRes struct {
	Permission string `json:"permission"`
}

func MapUser2UserRes(user *model.User) (*UserRes, error) {
	deletedAt := core.EmptyString
	if user.DeletedAt.Valid {
		deletedAt = user.DeletedAt.Time.String()
	}
	groupsAny, err := utils.Map(user.Groups, func(v any) (any, error) {
		group, ok := v.(model.Group)
		if !ok {
			return "", errors.New(fmt.Sprintf("提取信息失败: %v", v))
		}
		return group.Group, nil
	})
	if err != nil {
		return nil, fmt.Errorf("mapping错误: %v", err)
	}
	groupsStr, err := utils.Any2String(groupsAny)
	if err != nil {
		return nil, errors.New("mapping错误: any转string失败")
	}
	return &UserRes{
		UserID:        user.UserID,
		Username:      user.Username,
		Mobile:        user.Mobile,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		CreatedAt:     user.CreatedAt.String(),
		UpdatedAt:     user.UpdatedAt.String(),
		DeletedAt:     deletedAt,
		Role:          user.Role.Role,
		Groups:        groupsStr,
	}, nil
}

func MapRole2RoleRes(role *model.Role) (*RoleRes, error) {
	if role.DeletedAt.Valid {
		return nil, nil
	}
	permissionAny, err := utils.Map(role.Permissions, func(v any) (any, error) {
		permission, ok := v.(model.Permission)
		if !ok {
			return "", errors.New(fmt.Sprintf("mapping错误: 提取信息失败: %v", v))
		}
		return permission.Permission, nil
	})
	if err != nil {
		return nil, fmt.Errorf("mapping错误: %v", err)
	}
	permissionStr, err := utils.Any2String(permissionAny)
	if err != nil {
		return nil, errors.New("mapping错误: any转string失败")
	}
	return &RoleRes{
		Role:       role.Role,
		Permission: permissionStr,
	}, nil
}

func MapGroup2GroupRes(group *model.Group) *GroupRes {
	if group.DeletedAt.Valid {
		return nil
	}
	parentGroupStruct := group.ParentGroup
	parentGroup := ""
	if parentGroupStruct != nil {
		parentGroup = parentGroupStruct.Group
	}
	return &GroupRes{
		Group:       group.Group,
		ParentGroup: parentGroup,
	}
}

func MapPermission2PermissionRes(permission *model.Permission) *PermissionRes {
	if permission.DeletedAt.Valid {
		return nil
	}
	return &PermissionRes{
		Permission: permission.Permission,
	}
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenRes struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
}
