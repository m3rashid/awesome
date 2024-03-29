package auth

import (
	"awesome/controller"
	"awesome/models"
	"awesome/utils"
)

var AuthModule = utils.Module{
	Name: "auth",
	Models: []interface{}{
		&models.UserGroup{},
		&models.Profile{},
		&models.User{},
	},
	SchemaMap: utils.SchemaMap{
		models.USER_MODEL_NAME:       models.UserTableSchemaMap,
		models.PROFILE_MODEL_NAME:    models.ProfileTableSchemaMap,
		models.USER_GROUP_MODEL_NAME: models.UserGroupTableSchemaMap,
	},
	ProtectedRoutes: utils.ProtectedRoutes{
		"/users": {
			Description: "List all users",
			Controller: controller.List[models.User](
				models.USER_MODEL_NAME,
				controller.ListOptions{},
			),
			Permissions: utils.RoutePermissions{
				"user": utils.LIST,
			},
		},
		"/users/search": {
			Description: "Search Users",
			Controller:  SearchUsers,
		},
		"/users/get": {
			Description: "Get any user",
			Controller: controller.Get[models.User](
				controller.GetOptions[models.User]{},
			),
		},
		"/users/update": {
			Description: "Update user",
			Controller: controller.Update[models.User](
				models.USER_MODEL_NAME,
				controller.UpdateOptions[models.User]{},
			),
			Permissions: utils.RoutePermissions{
				"user": utils.EDIT,
			},
		},

		"/init": {
			Description: "Init automatic auth at refresh",
			Controller: utils.GetInitialUser(utils.AuthControllerOptions{
				IsTenant: true,
			}),
			Permissions: utils.RoutePermissions{},
		},
		"/profiles": {
			Description: "List all profiles",
			Controller: controller.List[models.Profile](
				models.PROFILE_MODEL_NAME,
				controller.ListOptions{},
			),
		},
		"/profile": {
			Description: "Get profile",
			Controller: controller.Get[models.Profile](
				controller.GetOptions[models.Profile]{},
			),
		},
		"/profile/update": {
			Description: "Update profile",
			Controller: controller.Update[models.Profile](
				models.PROFILE_MODEL_NAME,
				controller.UpdateOptions[models.Profile]{},
			),
		},
	},
	AnonymousRoutes: utils.AnonymousRoutes{
		"/login": {
			Description: "Login",
			Controller: utils.Login(utils.AuthControllerOptions{
				IsTenant: true,
			}),
		},
		"/register": {
			Description: "Register",
			Controller: utils.Register(utils.AuthControllerOptions{
				IsTenant: true,
			}),
		},
	},
}
