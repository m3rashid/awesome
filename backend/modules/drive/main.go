package drive

import (
	"awesome/controller"
	"awesome/models"
	"awesome/utils"
)

var DriveModule = utils.Module{
	Name: "drive",
	Models: []interface{}{
		&models.DriveFile{},
	},
	ProtectedRoutes: utils.ProtectedRouteConfig{
		// other drive routes are directly registered
		"/all": {
			Description: "List all files",
			Controller: controller.List[models.DriveFile](
				models.DRIVE_FILE_MODEL_NAME, controller.ListOptions{},
			),
			Permissions: utils.RoutePermissions{
				"driveFile": utils.LIST,
			},
		},
	},
	AnonymousRoutes: utils.AnonymousRouteConfig{},
}