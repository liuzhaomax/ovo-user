# ovo-user ovo用户微服务

模板使用[go-maxms微服务脚手架](https://github.com/liuzhaomax/go-maxms)

用户自行注册，由其他用户加入到组织中，并赋予角色

## TODO
+ ~~model~~
+ ~~schema~~
+ api
  + 用户注册 PUT /login
  + 用户登录 POST /login
  + 用户登出 DELETE /login
  + 获取puk GET /login
  + 获取用户列表 GET /users
  + 获取角色列表 GET /roles
  + 获取组织列表 GET /groups
  + 获取权限列表 GET /permissions
  + 通过ID，获取单个用户 GET /user/{userID}
  + 新增修改单个用户 PUT /user/{userID}
  + 通过ID，修改单个用户的角色 PUT /user/{userID}/role/{roleID}
  + 通过ID，修改单个用户的组织 PUT /user/{userID}/group/{groupID}
  + 新增修改单个角色及其对应权限 PUT /role/{roleID}
  + 新增修改单个组织 PUT /group/{groupID}
  + 新增修改单个权限 PUT /permission/{permissionID}
  + 删除单个用户 DELETE /user/{userID}
  + 删除单个角色 DELETE /role/{roleID}
  + 删除单个组织，及其下游组织 DELETE /group/{groupID}
  + 删除单个权限 DELETE /permission/{permissionID}
