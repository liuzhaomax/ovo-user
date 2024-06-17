package schema

import (
	"github.com/liuzhaomax/ovo-user/internal/core"
	"github.com/liuzhaomax/ovo-user/src/api_user_rpc/model"
	"github.com/liuzhaomax/ovo-user/src/api_user_rpc/pb"
)

func MapUser2UserRes(user *model.User) *pb.UserRes {
	return &pb.UserRes{
		Status: &pb.Status{
			Code: int32(core.OK),
			Desc: "success",
		},
		Data: &pb.UserResData{
			Id:     int32(user.ID),
			UserID: user.UserID,
			Mobile: user.Mobile,
		},
	}
}
