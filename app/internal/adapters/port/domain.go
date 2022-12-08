package port

import (
	"auto_post/app/pkg/ddo"
)

// Filer interface --
type Filer interface {
	CreateFile(ddo *ddo.ReqFileDDO) *ddo.ResFileDDO
}
