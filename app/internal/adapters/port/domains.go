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
	UploadPhotoToServer(ddo *ddo.VKMachineDDO)
}
