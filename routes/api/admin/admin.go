package admin

import (
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin"
	"user_center/app/Http/Middleware"
)

func LoadAdmin(router *gin.Engine) {
	AuthAPI := router.Group("/api/admin", Middleware.Middleware.Api...)
	{
		// 用户相关
		user := AuthAPI.Group("/user")
		{
			// 保存用户数据
			user.POST("store", Admin.UserController{}.Store)
			// 更新用户数据
			user.PUT("update", Admin.UserController{}.Update)
			// 获取用户详情
			user.GET("detail", Admin.UserController{}.Detail)
			// 获取用户列表
			user.GET("list", Admin.UserController{}.GetList)
			// 禁用用户
			user.POST("forbidden", Admin.UserController{}.Forbidden)
			// 绑定角色
			user.POST("bind", Admin.UserController{}.BindRole)
			// 获取用户绑定的角色
			user.GET("get_bind_role", Admin.UserController{}.GetBindRole)
			// 获取用户绑定的权限
			user.GET("get_bind_permission", Admin.UserController{}.GetBindPermission)
		}

		// 角色相关
		role := AuthAPI.Group("/role")
		{
			// 保存角色数据
			role.POST("store", Admin.RoleController{}.Store)
			// 更新角色数据
			role.PUT("update", Admin.RoleController{}.Update)
			// 获取角色详情
			role.GET("detail", Admin.RoleController{}.Detail)
			// 获取角色列表
			role.GET("list", Admin.RoleController{}.GetList)
			// 删除角色
			role.DELETE("delete", Admin.RoleController{}.Delete)
			// 绑定权限
			role.POST("bind", Admin.RoleController{}.BindPermission)
		}

		// 权限相关
		permission := AuthAPI.Group("/permission")
		{
			// 保存数据
			permission.POST("store", Admin.PermissionController{}.Store)
			// 更新数据
			permission.PUT("update", Admin.PermissionController{}.Update)
			// 获取详情数据
			permission.GET("detail", Admin.PermissionController{}.Detail)
			// 获取列表数据
			permission.GET("list", Admin.PermissionController{}.GetList)
			// 删除数据
			permission.DELETE("delete", Admin.PermissionController{}.Delete)
			// 获取权限数
			permission.GET("tree", Admin.PermissionController{}.GetTree)
		}
	}
}
