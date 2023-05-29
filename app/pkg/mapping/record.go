package mapping

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/internal/adapters/repository/models/basis"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/dto"
)

// ManagerModelToDBO --
func ManagerModelToDBO(req *models.Manager) *dbo.RecordDBO {
	return &dbo.RecordDBO{
		URL:     req.URL,
		Service: req.Service,
		Status:  req.Status,
		Hash:    req.Hash,

		UpdatedAt: req.UpdatedAt,
		CreatedAt: req.CreatedAt,
	}
}

// CreateRecordDTOtoDDO --
func CreateRecordDTOtoDDO(req *dto.CreateRecordReqDTO) *ddo.CreateRecordRequestDDO {
	return &ddo.CreateRecordRequestDDO{
		URL:     req.URL,
		AuthURL: req.AuthURL,
		Service: req.Service,
	}
}

// CreateRecordDDOtoDBO --
func CreateRecordDDOtoDBO(ddo *ddo.CreateRecordResponseDDO) *dbo.RecordDBO {
	return &dbo.RecordDBO{
		UUID:      ddo.UUID,
		URL:       ddo.URL,
		AuthURL:   ddo.AuthURL,
		Service:   ddo.Service,
		Status:    ddo.Status,
		Hash:      ddo.Hash,
		UpdatedAt: ddo.UpdatedAt,
		CreatedAt: ddo.CreatedAt,
	}
}

// RecordDBOtoModel --
func RecordDBOtoModel(dbo *dbo.RecordDBO) *models.Manager {
	base := basis.BaseModel{}
	base.UUID = dbo.UUID
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.Manager{
		URL:       dbo.URL,
		AuthURL:   dbo.AuthURL,
		Service:   dbo.Service,
		Status:    dbo.Status,
		Hash:      dbo.Hash,
		BaseModel: base,
	}
}

// RecordDbOtoVkMachineDDO --
func RecordDbOtoVkMachineDDO(req *dbo.RecordDBO) *ddo.VKMachine {
	return &ddo.VKMachine{
		FileName: req.UUID,
		URL:      req.URL,
	}
}
