// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "cloud_go/Disk/internal/handler/file"
	user "cloud_go/Disk/internal/handler/user"
	"cloud_go/Disk/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: registerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/EmailSend",
				Handler: EmailSendHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/updatedetail",
					Handler: user.UpdateDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/getdetail",
					Handler: user.GetDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/avatar",
					Handler: user.UpdateAvatarHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: file.GetFileDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/listfile",
					Handler: file.ListFileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/type/:type",
					Handler: file.ListfileByTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/folder-list/:parentFolderId",
					Handler: file.ListFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/move/:parentFolderId",
					Handler: file.ListFileMovableFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/folder-move",
					Handler: file.ListFolderMovableFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/move-file",
					Handler: file.MoveFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/move-folder",
					Handler: file.MoveFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/update-file",
					Handler: file.UpdateFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/update-folder",
					Handler: file.UpdateFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/copy-file",
					Handler: file.CopyFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delete-file",
					Handler: file.DeleteFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delete-folder",
					Handler: file.DeleteFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delete-file-truly",
					Handler: file.DeleteFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delete-folder-truly",
					Handler: file.DeleteFoldersTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/recover-file",
					Handler: file.RecoverFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/recover-folder",
					Handler: file.RecoverFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create-folder",
					Handler: file.CreateFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delete-item",
					Handler: file.ListDeleteItemsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)
}
