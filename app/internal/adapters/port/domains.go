package port

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
)

// ManagerDomain interface --
type ManagerDomain interface {
	CreateRecord(ddo *ddo.CreateRecordRequestDDO) *ddo.CreateRecordResponseDDO
}

// VkMachineDomain interface --
type VkMachineDomain interface {
	GetPath(*ddo.VKMachine) string
	SaveWallPhoto(*ddo.ReqSaveWallPhoto) *ddo.ResSaveWallPhoto
	GetWallUploadServer() *ddo.GetWallUploadServer
	PostWallPhoto(*ddo.ReqPostWallPhoto) *ddo.ResPostWallPhoto
	GetUploadServer() *ddo.ResGetUploadServer
	PhotosSave(ddo.ReqPhotosSave) *ddo.ResPhotosSave
	GetTags() string
}

// YoutubeMachineDomain interface --
type YoutubeMachineDomain interface {
	KeepNextCursor(string)
	GetNextCursor() string
	KeepItems(title, videoID string)
	GetItems() []ddo.ResGetItems
}
