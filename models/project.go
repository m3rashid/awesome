package models

import "time"

const PROJECT_MODEL_NAME = "projects"
const PROJECT_TAG_MODEL_NAME = "project_tags"
const PROJECT_TASK_MODEL_NAME = "project_tasks"
const PROJECT_TASK_COMMENT_MODEL_NAME = "project_task_comments"

type TaskStatus string

const (
	TaskStatusBacklog    TaskStatus = "backlog"
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "inprogress"
	TaskStatusReview     TaskStatus = "review"
	TaskStatusDone       TaskStatus = "done"
)

type Project struct {
	BaseModel
	Name           string  `json:"name" gorm:"column:name;not null" validate:"required"`
	Description    string  `json:"description" gorm:"column:description" validate:""`
	Members        []*User `json:"members" gorm:"many2many:project_members_relation" validate:""`
	ProjectOwnerID uint    `json:"projectOwnerId" gorm:"column:projectOwnerId;not null" validate:"required"`
	ProjectOwner   *User   `json:"projectOwner" gorm:"column:projectOwnerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:""`
	Completed      bool    `json:"completed" gorm:"column:completed;default:false" validate:""`
}

var ProjectTableSchemaMap = map[string]string{
	"name":           "string",
	"projectOwnerId": "number",
	"completed":      "boolean",
	"createdAt":      "time",
}

type ProjectTask struct {
	BaseModel
	ProjectID    uint          `json:"projectId" gorm:"column:projectId;not null" validate:"required"`
	Project      *Project      `json:"project" gorm:"column:projectId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:""`
	Name         string        `json:"name" gorm:"column:name;not null" validate:"required"`
	Description  string        `json:"description" gorm:"column:description" validate:""`
	Deadline     time.Time     `json:"deadline" gorm:"column:deadline" validate:""`
	Completed    bool          `json:"completed" gorm:"column:completed;default:false" validate:""`
	TaskStatus   TaskStatus    `json:"taskStatus" gorm:"column:taskStatus;not null" validate:"required"`
	ReportedByID uint          `json:"reportedById" gorm:"column:reportedById;not null" validate:"required"`
	ReportedBy   *User         `json:"reportedBy" gorm:"column:reportedById;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:""`
	AssignedToID uint          `json:"assignedToId" gorm:"column:assignedToId;" validate:""`
	AssignedTo   *User         `json:"assignedTo" gorm:"column:assignedToId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:""`
	Tags         []*ProjectTag `json:"tags" gorm:"many2many:project_task_tags_relation" validate:""`
}

var ProjectTaskTableSchemaMap = map[string]string{
	"projectId":    "number",
	"name":         "string",
	"completed":    "boolean",
	"deadline":     "time",
	"taskStatus":   "string",
	"reportedById": "number",
	"assignedToId": "number",
	"createdAt":    "time",
}

type ProjectTag struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" validate:"required"`
}

type ProjectTaskComments struct {
	BaseModel
	TaskID  uint   `json:"taskId" gorm:"column:taskId;not null" validate:"required"`
	UserID  uint   `json:"userId" gorm:"column:userId;not null" validate:"required"`
	User    *User  `json:"user" gorm:"column:userId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:""`
	Comment string `json:"comment" gorm:"column:comment;not null" validate:"required"`
}

func (*Project) TableName() string {
	return PROJECT_MODEL_NAME
}

func (*ProjectTag) TableName() string {
	return PROJECT_TAG_MODEL_NAME
}

func (*ProjectTask) TableName() string {
	return PROJECT_TASK_MODEL_NAME
}

func (*ProjectTaskComments) TableName() string {
	return PROJECT_TASK_COMMENT_MODEL_NAME
}
