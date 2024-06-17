package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/ovo-user/src/api_user/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessUserSet,
)
