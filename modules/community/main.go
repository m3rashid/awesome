package community

import (
	"awesome/controller"
	"awesome/models"
	"awesome/utils"
)

var CommunityModule = utils.Module{
	Name: "community",
	Models: []interface{}{
		&models.Post{},
		&models.Comment{},
		&models.Friend{},
		&models.FriendRequest{},
		&models.CommunityGroup{},
		&models.CommunityChatMessage{},
	},
	SchemaMap: utils.SchemaMap{
		models.POST_MODEL_NAME:    models.PostTableSchemaMap,
		models.COMMENT_MODEL_NAME: models.CommentTableSchemaMap,
	},
	ProtectedRoutes: utils.ProtectedRoutes{
		"/posts": {
			Description: "Get all posts",
			Controller: controller.List[models.Post](
				models.POST_MODEL_NAME,
				controller.ListOptions{},
			),
		},
		"/post/create": {
			Description: "Create a post",
			Controller: controller.Create[models.Post](
				models.POST_MODEL_NAME,
				controller.CreateOptions[models.Post]{},
			),
		},
		"/post/get": {
			Description: "Get a single post",
			Controller: controller.Get[models.Post](
				controller.GetOptions[models.Post]{},
			),
		},
		"/post/update": {
			Description: "Update a post",
			Controller: controller.Update[models.Post](
				models.POST_MODEL_NAME,
				controller.UpdateOptions[models.Post]{},
			),
		},

		"/comments": {
			Description: "List all comments",
			Controller: controller.List[models.Comment](
				models.COMMENT_MODEL_NAME,
				controller.ListOptions{},
			),
		},
		"/comments/create": {
			Description: "Create a comment",
			Controller: controller.Create[models.Comment](
				models.COMMENT_MODEL_NAME,
				controller.CreateOptions[models.Comment]{},
			),
		},
		"/comments/get": {
			Description: "Get a single comment",
			Controller: controller.Get[models.Comment](
				controller.GetOptions[models.Comment]{},
			),
		},
		"/comments/update": {
			Description: "Update a comment",
			Controller: controller.Update[models.Comment](
				models.COMMENT_MODEL_NAME,
				controller.UpdateOptions[models.Comment]{},
			),
		},

		"/friendships": {
			Description: "List all friendships",
			Controller: controller.List[models.Friend](
				models.FRIENDS_MODEL_NAME,
				controller.ListOptions{},
			),
		},
		"/friendships/send-request": {
			Description: "Send a friend request",
			Controller: controller.Create[models.FriendRequest](
				models.FRIENDS_MODEL_NAME,
				controller.CreateOptions[models.FriendRequest]{},
			),
		},
		"/friendships/update": {
			Description: "Accept a friend request",
			Controller: controller.Update[models.FriendRequest](
				models.FRIEND_REQUEST_MODEL_NAME,
				controller.UpdateOptions[models.FriendRequest]{},
			),
		},

		"/groups": {
			Description: "List of groups",
			Controller: controller.List[models.CommunityGroup](
				models.COMMUNITY_GROUP_MODEL_NAME,
				controller.ListOptions{
					ModifyDbCall: GetUserGroupsModifyDbCall,
				}),
		},
		"/groups/get": {
			Description: "Get a single group",
			Controller: controller.Get[models.CommunityGroup](
				controller.GetOptions[models.CommunityGroup]{},
			),
		},
		"/groups/create": {
			Description: "Create a group",
			Controller: controller.Create[models.CommunityGroup](
				models.COMMUNITY_GROUP_MODEL_NAME,
				controller.CreateOptions[models.CommunityGroup]{},
			),
		},
		"/groups/update": {
			Description: "Update a group",
			Controller: controller.Update[models.CommunityGroup](
				models.COMMUNITY_GROUP_MODEL_NAME,
				controller.UpdateOptions[models.CommunityGroup]{},
			),
		},

		"/chats": {
			Description: "List of chats",
			Controller: controller.List[models.CommunityChatMessage](
				models.COMMUNITY_CHAT_MESSAGE_MODEL_NAME,
				controller.ListOptions{},
			),
		},
		"/chats/create": {
			Description: "Create a chat",
			Controller: controller.Create[models.CommunityChatMessage](
				models.COMMUNITY_CHAT_MESSAGE_MODEL_NAME,
				controller.CreateOptions[models.CommunityChatMessage]{},
			),
		},
		"/chats/update": {
			Description: "Update a chat",
			Controller: controller.Update[models.CommunityChatMessage](
				models.COMMUNITY_CHAT_MESSAGE_MODEL_NAME,
				controller.UpdateOptions[models.CommunityChatMessage]{},
			),
		},
	},
	AnonymousRoutes: utils.AnonymousRoutes{},
}
