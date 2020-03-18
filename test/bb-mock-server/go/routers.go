/*
 * Bitbucket Server API
 *
 * Bitbucket Server API (former stash).
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/rest/",
		Index,
	},

	Route{
		"AddGroupToUser",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/users/add-group",
		AddGroupToUser,
	},

	Route{
		"AddUserToGroup",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/groups/add-user",
		AddUserToGroup,
	},

	Route{
		"AddUserToGroups",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/users/add-groups",
		AddUserToGroups,
	},

	Route{
		"AddUsersToGroup",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/groups/add-users",
		AddUsersToGroup,
	},

	Route{
		"Approve",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve",
		Approve,
	},

	Route{
		"AssignParticipantRole",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants",
		AssignParticipantRole,
	},

	Route{
		"CanMerge",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge",
		CanMerge,
	},

	Route{
		"ClearSenderAddress",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/mail-server/sender-address",
		ClearSenderAddress,
	},

	Route{
		"ClearUserCaptchaChallenge",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/users/captcha",
		ClearUserCaptchaChallenge,
	},

	Route{
		"CountPullRequestTasks",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/tasks/count",
		CountPullRequestTasks,
	},

	Route{
		"Create",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests",
		Create,
	},

	Route{
		"CreateBranch",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches",
		CreateBranch,
	},

	Route{
		"CreateComment",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments",
		CreateComment,
	},

	Route{
		"CreateCommentCommit",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments",
		CreateCommentCommit,
	},

	Route{
		"CreateGroup",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/groups",
		CreateGroup,
	},

	Route{
		"CreateProject",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects",
		CreateProject,
	},

	Route{
		"CreateRepository",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos",
		CreateRepository,
	},

	Route{
		"CreateTag",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags",
		CreateTag,
	},

	Route{
		"CreateTask",
		strings.ToUpper("Post"),
		"/rest/api/1.0/tasks",
		CreateTask,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/users",
		CreateUser,
	},

	Route{
		"CreateWebhook",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks",
		CreateWebhook,
	},

	Route{
		"Decline",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/decline",
		Decline,
	},

	Route{
		"Delete",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}",
		Delete,
	},

	Route{
		"DeleteAvatar",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/users/{userSlug}/avatar.png",
		DeleteAvatar,
	},

	Route{
		"DeleteComment",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId}",
		DeleteComment,
	},

	Route{
		"DeleteCommentCommit",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}",
		DeleteCommentCommit,
	},

	Route{
		"DeleteGroup",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/groups",
		DeleteGroup,
	},

	Route{
		"DeleteMailConfig",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/mail-server",
		DeleteMailConfig,
	},

	Route{
		"DeleteProject",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}",
		DeleteProject,
	},

	Route{
		"DeleteRepository",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}",
		DeleteRepository,
	},

	Route{
		"DeleteRepositoryHook",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}",
		DeleteRepositoryHook,
	},

	Route{
		"DeleteTask",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/tasks/{taskId}",
		DeleteTask,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/users",
		DeleteUser,
	},

	Route{
		"DeleteWebhook",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}",
		DeleteWebhook,
	},

	Route{
		"DisableHook",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/enabled",
		DisableHook,
	},

	Route{
		"DisableHookRepo",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled",
		DisableHookRepo,
	},

	Route{
		"EditFile",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path}",
		EditFile,
	},

	Route{
		"EnableHook",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/enabled",
		EnableHook,
	},

	Route{
		"EnableHookRepo",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled",
		EnableHookRepo,
	},

	Route{
		"FindGroupsForUser",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/users/more-members",
		FindGroupsForUser,
	},

	Route{
		"FindOtherGroupsForUser",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/users/more-non-members",
		FindOtherGroupsForUser,
	},

	Route{
		"FindUsersInGroup",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/groups/more-members",
		FindUsersInGroup,
	},

	Route{
		"FindUsersNotInGroup",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/groups/more-non-members",
		FindUsersNotInGroup,
	},

	Route{
		"FindWebhooks",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks",
		FindWebhooks,
	},

	Route{
		"ForkRepository",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}",
		ForkRepository,
	},

	Route{
		"Get",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}",
		Get,
	},

	Route{
		"GetActivities",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/activities",
		GetActivities,
	},

	Route{
		"GetApplicationProperties",
		strings.ToUpper("Get"),
		"/rest/api/1.0/application-properties",
		GetApplicationProperties,
	},

	Route{
		"GetArchive",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/archive",
		GetArchive,
	},

	Route{
		"GetAvatar",
		strings.ToUpper("Get"),
		"/rest/api/1.0/hooks/{hookKey}/avatar",
		GetAvatar,
	},

	Route{
		"GetBranches",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches",
		GetBranches,
	},

	Route{
		"GetChanges",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/changes",
		GetChanges,
	},

	Route{
		"GetChangesFile",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/changes",
		GetChangesFile,
	},

	Route{
		"GetComment",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId}",
		GetComment,
	},

	Route{
		"GetCommentCommit",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}",
		GetCommentCommit,
	},

	Route{
		"GetComments",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments",
		GetComments,
	},

	Route{
		"GetCommentsCommit",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments",
		GetCommentsCommit,
	},

	Route{
		"GetCommit",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}",
		GetCommit,
	},

	Route{
		"GetCommits",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits",
		GetCommits,
	},

	Route{
		"GetCommitsPR",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/commits",
		GetCommitsPR,
	},

	Route{
		"GetContentBrowse",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse",
		GetContentBrowse,
	},

	Route{
		"GetContentFile",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path}",
		GetContentFile,
	},

	Route{
		"GetContentRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/raw",
		GetContentRepository,
	},

	Route{
		"GetContentRepositoryPath",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/raw/{path}",
		GetContentRepositoryPath,
	},

	Route{
		"GetDefaultBranch",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default",
		GetDefaultBranch,
	},

	Route{
		"GetForkedRepositories",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/forks",
		GetForkedRepositories,
	},

	Route{
		"GetGroups",
		strings.ToUpper("Get"),
		"/rest/api/1.0/groups",
		GetGroups,
	},

	Route{
		"GetGroupsAdmin",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/groups",
		GetGroupsAdmin,
	},

	Route{
		"GetGroupsWithAnyPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/permissions/groups",
		GetGroupsWithAnyPermission,
	},

	Route{
		"GetGroupsWithAnyPermissionProject",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/permissions/groups",
		GetGroupsWithAnyPermissionProject,
	},

	Route{
		"GetGroupsWithAnyPermissionRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups",
		GetGroupsWithAnyPermissionRepository,
	},

	Route{
		"GetGroupsWithoutAnyPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/permissions/groups/none",
		GetGroupsWithoutAnyPermission,
	},

	Route{
		"GetGroupsWithoutAnyPermissionProject",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/permissions/groups/none",
		GetGroupsWithoutAnyPermissionProject,
	},

	Route{
		"GetGroupsWithoutAnyPermissionRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups/none",
		GetGroupsWithoutAnyPermissionRepository,
	},

	Route{
		"GetInformation",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/cluster",
		GetInformation,
	},

	Route{
		"GetLatestInvocation",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/latest",
		GetLatestInvocation,
	},

	Route{
		"GetLevel",
		strings.ToUpper("Get"),
		"/rest/api/1.0/logs/logger/{loggerName}",
		GetLevel,
	},

	Route{
		"GetLicense",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/license",
		GetLicense,
	},

	Route{
		"GetMailConfig",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/mail-server",
		GetMailConfig,
	},

	Route{
		"GetMergeConfig",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/pull-requests/{scmId}",
		GetMergeConfig,
	},

	Route{
		"GetPage",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests",
		GetPage,
	},

	Route{
		"GetProject",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}",
		GetProject,
	},

	Route{
		"GetProjectAvatar",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/avatar.png",
		GetProjectAvatar,
	},

	Route{
		"GetProjects",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects",
		GetProjects,
	},

	Route{
		"GetPullRequestCount",
		strings.ToUpper("Get"),
		"/rest/api/1.0/inbox/pull-requests/count",
		GetPullRequestCount,
	},

	Route{
		"GetPullRequestSettings",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests",
		GetPullRequestSettings,
	},

	Route{
		"GetPullRequestSettingsSCM",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/settings/pull-requests/{scmId}",
		GetPullRequestSettingsSCM,
	},

	Route{
		"GetPullRequestSuggestions",
		strings.ToUpper("Get"),
		"/rest/api/1.0/dashboard/pull-request-suggestions",
		GetPullRequestSuggestions,
	},

	Route{
		"GetPullRequestTasks",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/tasks",
		GetPullRequestTasks,
	},

	Route{
		"GetPullRequests",
		strings.ToUpper("Get"),
		"/rest/api/1.0/dashboard/pull-requests",
		GetPullRequests,
	},

	Route{
		"GetPullRequestsInbox",
		strings.ToUpper("Get"),
		"/rest/api/1.0/inbox/pull-requests",
		GetPullRequestsInbox,
	},

	Route{
		"GetRelatedRepositories",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/related",
		GetRelatedRepositories,
	},

	Route{
		"GetRepositories",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos",
		GetRepositories,
	},

	Route{
		"GetRepositoriesAll",
		strings.ToUpper("Get"),
		"/rest/api/1.0/repos",
		GetRepositoriesAll,
	},

	Route{
		"GetRepositoriesRecentlyAccessed",
		strings.ToUpper("Get"),
		"/rest/api/1.0/profile/recent/repos",
		GetRepositoriesRecentlyAccessed,
	},

	Route{
		"GetRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}",
		GetRepository,
	},

	Route{
		"GetRepositoryHook",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks/{hookKey}",
		GetRepositoryHook,
	},

	Route{
		"GetRepositoryHookSettings",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}",
		GetRepositoryHookSettings,
	},

	Route{
		"GetRepositoryHooks",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks",
		GetRepositoryHooks,
	},

	Route{
		"GetRepositoryHooksSettings",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks",
		GetRepositoryHooksSettings,
	},

	Route{
		"GetRootLevel",
		strings.ToUpper("Get"),
		"/rest/api/1.0/logs/rootLogger",
		GetRootLevel,
	},

	Route{
		"GetSenderAddress",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/mail-server/sender-address",
		GetSenderAddress,
	},

	Route{
		"GetSettings",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/settings",
		GetSettings,
	},

	Route{
		"GetSettingsHook",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings",
		GetSettingsHook,
	},

	Route{
		"GetStatistics",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics",
		GetStatistics,
	},

	Route{
		"GetStatisticsSummary",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics/summary",
		GetStatisticsSummary,
	},

	Route{
		"GetTag",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags/{name}",
		GetTag,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/tags",
		GetTags,
	},

	Route{
		"GetTask",
		strings.ToUpper("Get"),
		"/rest/api/1.0/tasks/{taskId}",
		GetTask,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/rest/api/1.0/users/{userSlug}",
		GetUser,
	},

	Route{
		"GetUserSettings",
		strings.ToUpper("Get"),
		"/rest/api/1.0/users/{userSlug}/settings",
		GetUserSettings,
	},

	Route{
		"GetUsers",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/users",
		GetUsers,
	},

	Route{
		"GetUsersAll",
		strings.ToUpper("Get"),
		"/rest/api/1.0/users",
		GetUsersAll,
	},

	Route{
		"GetUsersWithAnyPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users",
		GetUsersWithAnyPermission,
	},

	Route{
		"GetUsersWithAnyPermissionAll",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/permissions/users",
		GetUsersWithAnyPermissionAll,
	},

	Route{
		"GetUsersWithAnyPermissionProject",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/permissions/users",
		GetUsersWithAnyPermissionProject,
	},

	Route{
		"GetUsersWithoutAnyPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/admin/permissions/users/none",
		GetUsersWithoutAnyPermission,
	},

	Route{
		"GetUsersWithoutPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users/none",
		GetUsersWithoutPermission,
	},

	Route{
		"GetUsersWithoutPermissionProject",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/permissions/users/none",
		GetUsersWithoutPermissionProject,
	},

	Route{
		"GetWebhook",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}",
		GetWebhook,
	},

	Route{
		"HasAllUserPermission",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/permissions/{permission}/all",
		HasAllUserPermission,
	},

	Route{
		"ListParticipants",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants",
		ListParticipants,
	},

	Route{
		"Merge",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge",
		Merge,
	},

	Route{
		"ModifyAllUserPermission",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/permissions/{permission}/all",
		ModifyAllUserPermission,
	},

	Route{
		"Preview",
		strings.ToUpper("Post"),
		"/rest/api/1.0/markup/preview",
		Preview,
	},

	Route{
		"RemoveGroupFromUser",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/users/remove-group",
		RemoveGroupFromUser,
	},

	Route{
		"RemoveUserFromGroup",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/groups/remove-user",
		RemoveUserFromGroup,
	},

	Route{
		"RenameUser",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/users/rename",
		RenameUser,
	},

	Route{
		"Reopen",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/reopen",
		Reopen,
	},

	Route{
		"RetryCreateRepository",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/recreate",
		RetryCreateRepository,
	},

	Route{
		"RevokePermissionsForGroup",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups",
		RevokePermissionsForGroup,
	},

	Route{
		"RevokePermissionsForGroupAll",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/permissions/groups",
		RevokePermissionsForGroupAll,
	},

	Route{
		"RevokePermissionsForGroupProject",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/permissions/groups",
		RevokePermissionsForGroupProject,
	},

	Route{
		"RevokePermissionsForUser",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users",
		RevokePermissionsForUser,
	},

	Route{
		"RevokePermissionsForUserAll",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/admin/permissions/users",
		RevokePermissionsForUserAll,
	},

	Route{
		"RevokePermissionsForUserProject",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/permissions/users",
		RevokePermissionsForUserProject,
	},

	Route{
		"Search",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/participants",
		Search,
	},

	Route{
		"SetDefaultBranch",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default",
		SetDefaultBranch,
	},

	Route{
		"SetLevel",
		strings.ToUpper("Put"),
		"/rest/api/1.0/logs/logger/{loggerName}/{levelName}",
		SetLevel,
	},

	Route{
		"SetMailConfig",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/mail-server",
		SetMailConfig,
	},

	Route{
		"SetMergeConfig",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/pull-requests/{scmId}",
		SetMergeConfig,
	},

	Route{
		"SetPermissionForGroup",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/groups",
		SetPermissionForGroup,
	},

	Route{
		"SetPermissionForGroups",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/permissions/groups",
		SetPermissionForGroups,
	},

	Route{
		"SetPermissionForGroupsProject",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/permissions/groups",
		SetPermissionForGroupsProject,
	},

	Route{
		"SetPermissionForUser",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/permissions/users",
		SetPermissionForUser,
	},

	Route{
		"SetPermissionForUsers",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/permissions/users",
		SetPermissionForUsers,
	},

	Route{
		"SetPermissionForUsersProject",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/permissions/users",
		SetPermissionForUsersProject,
	},

	Route{
		"SetRootLevel",
		strings.ToUpper("Put"),
		"/rest/api/1.0/logs/rootLogger/{levelName}",
		SetRootLevel,
	},

	Route{
		"SetSenderAddress",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/mail-server/sender-address",
		SetSenderAddress,
	},

	Route{
		"SetSettings",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings",
		SetSettings,
	},

	Route{
		"SetSettingsProject",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/settings/hooks/{hookKey}/settings",
		SetSettingsProject,
	},

	Route{
		"Stream",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/last-modified",
		Stream,
	},

	Route{
		"StreamChanges",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/changes",
		StreamChanges,
	},

	Route{
		"StreamChangesCompare",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/changes",
		StreamChangesCompare,
	},

	Route{
		"StreamCommits",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/commits",
		StreamCommits,
	},

	Route{
		"StreamDiff",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff/{path}",
		StreamDiff,
	},

	Route{
		"StreamDiffCommit",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff",
		StreamDiffCommit,
	},

	Route{
		"StreamDiffCommits",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff/{path}",
		StreamDiffCommits,
	},

	Route{
		"StreamDiffCompare",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/compare/diff{path}",
		StreamDiffCompare,
	},

	Route{
		"StreamDiffPR",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff",
		StreamDiffPR,
	},

	Route{
		"StreamDiffRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/diff",
		StreamDiffRepository,
	},

	Route{
		"StreamDiffRepositoryFile",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/diff/{path}",
		StreamDiffRepositoryFile,
	},

	Route{
		"StreamFiles",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/files",
		StreamFiles,
	},

	Route{
		"StreamFilesLastMofied",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/last-modified/{path}",
		StreamFilesLastMofied,
	},

	Route{
		"StreamFilesRepository",
		strings.ToUpper("Get"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/files/{path}",
		StreamFilesRepository,
	},

	Route{
		"TestWebhook",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/test",
		TestWebhook,
	},

	Route{
		"UnassignParticipantRole",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug}",
		UnassignParticipantRole,
	},

	Route{
		"UnassignParticipantRolePR",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants",
		UnassignParticipantRolePR,
	},

	Route{
		"Unwatch",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch",
		Unwatch,
	},

	Route{
		"UnwatchPR",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch",
		UnwatchPR,
	},

	Route{
		"Update",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}",
		Update,
	},

	Route{
		"UpdateComment",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId}",
		UpdateComment,
	},

	Route{
		"UpdateCommentCommit",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}",
		UpdateCommentCommit,
	},

	Route{
		"UpdateLicense",
		strings.ToUpper("Post"),
		"/rest/api/1.0/admin/license",
		UpdateLicense,
	},

	Route{
		"UpdateProject",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}",
		UpdateProject,
	},

	Route{
		"UpdatePullRequestSettings",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests",
		UpdatePullRequestSettings,
	},

	Route{
		"UpdatePullRequestSettingsSCM",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/settings/pull-requests/{scmId}",
		UpdatePullRequestSettingsSCM,
	},

	Route{
		"UpdateRepository",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}",
		UpdateRepository,
	},

	Route{
		"UpdateSettings",
		strings.ToUpper("Post"),
		"/rest/api/1.0/users/{userSlug}/settings",
		UpdateSettings,
	},

	Route{
		"UpdateStatus",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug}",
		UpdateStatus,
	},

	Route{
		"UpdateTask",
		strings.ToUpper("Put"),
		"/rest/api/1.0/tasks/{taskId}",
		UpdateTask,
	},

	Route{
		"UpdateUserDetails",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/users",
		UpdateUserDetails,
	},

	Route{
		"UpdateUserDetailsAll",
		strings.ToUpper("Put"),
		"/rest/api/1.0/users",
		UpdateUserDetailsAll,
	},

	Route{
		"UpdateUserPassword",
		strings.ToUpper("Put"),
		"/rest/api/1.0/admin/users/credentials",
		UpdateUserPassword,
	},

	Route{
		"UpdateUserPasswordAll",
		strings.ToUpper("Put"),
		"/rest/api/1.0/users/credentials",
		UpdateUserPasswordAll,
	},

	Route{
		"UpdateWebhook",
		strings.ToUpper("Put"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}",
		UpdateWebhook,
	},

	Route{
		"UploadAvatar",
		strings.ToUpper("Post"),
		"/rest/api/1.0/users/{userSlug}/avatar.png",
		UploadAvatar,
	},

	Route{
		"UploadAvatarProject",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/avatar.png",
		UploadAvatarProject,
	},

	Route{
		"WatchCommit",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch",
		WatchCommit,
	},

	Route{
		"WatchPR",
		strings.ToUpper("Post"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch",
		WatchPR,
	},

	Route{
		"WithdrawApproval",
		strings.ToUpper("Delete"),
		"/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve",
		WithdrawApproval,
	},
}
