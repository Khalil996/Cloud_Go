// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	download "cloud_go/Disk/internal/handler/download"
	file "cloud_go/Disk/internal/handler/file"
	upload "cloud_go/Disk/internal/handler/upload"
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
					Path:    "/list/:parentFolderId",
					Handler: file.ListFileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/type/:fileType",
					Handler: file.ListFileByTypeHandler(serverCtx),
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
					Path:    "/move",
					Handler: file.MoveFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/move-folder",
					Handler: file.MoveFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/",
					Handler: file.UpdateFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder",
					Handler: file.UpdateFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/copy-file",
					Handler: file.CopyFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/delete",
					Handler: file.DeleteFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: file.DeleteFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder-delete",
					Handler: file.DeleteFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/clear",
					Handler: file.DeleteAllFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/recover",
					Handler: file.RecoverFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create-folder",
					Handler: file.CreateFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delete",
					Handler: file.ListDeleteFilesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check",
					Handler: upload.CheckFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: upload.UploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check_chunk",
					Handler: upload.CheckChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk",
					Handler: upload.UploadChunkHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/upload"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check_size",
					Handler: download.CheckSizeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: download.DownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk",
					Handler: download.ChunkDownloadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/download"),
	)
}