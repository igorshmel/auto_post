package port

import (
	"auto_post/app/pkg/ddo"
)

// ManagerDomain interface --
type ManagerDomain interface {
	CreateRecord(ddo *ddo.CreateRecordRequestDDO) *ddo.CreateRecordResponseDDO
}

// VkMachineDomain interface --
type VkMachineDomain interface {
	//UploadPhotoToServer(*ddo.VKMachine)
	GetPath(*ddo.VKMachine) string
	SaveWallPhoto(*ddo.ReqSaveWallPhoto) *ddo.ResSaveWallPhoto
	GetWallUploadServer() *ddo.GetWallUploadServer
	PostWallPhoto(req *ddo.ReqPostWallPhoto) *ddo.ResPostWallPhoto
	GetUploadServer() *ddo.ResGetUploadServer
	PhotosSave(req ddo.ReqPhotosSave) *ddo.ResPhotosSave
	GetTags() string
}
