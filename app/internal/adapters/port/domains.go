package port

import (
	"auto_post/app/pkg/ddo"
	"github.com/nuttech/bell/v2"
)

// ManagerDomain interface --
type ManagerDomain interface {
	CreateRecord(ddo *ddo.CreateRecordRequestDDO, events *bell.Events) *ddo.CreateRecordResponseDDO
}

// VkMachineDomain interface --
type VkMachineDomain interface {
	UploadPhotoToServer(ddo *ddo.VKMachineDDO)
}
