package mapping

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
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
		URL:         req.URL,
		AuthURL:     req.AuthURL,
		ImgURL:      req.ImgURL,
		Description: req.Description,
		Service:     req.Service,
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

// CreateMidRecordDDOtoDBO --
func CreateMidRecordDDOtoDBO(ddo *ddo.CreateRecordResponseDDO) *dbo.RecordMidDBO {
	return &dbo.RecordMidDBO{
		ImgURL:      ddo.ImgURL,
		Status:      ddo.Status,
		Hash:        ddo.Hash,
		Description: ddo.Description,
		UpdatedAt:   ddo.UpdatedAt,
		CreatedAt:   ddo.CreatedAt,
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

// RecordMidDBOtoModel --
func RecordMidDBOtoModel(dbo *dbo.RecordMidDBO) *models.Midjorney {
	base := basis.BaseModel{}
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.Midjorney{
		ImgURL:      dbo.ImgURL,
		Status:      dbo.Status,
		Hash:        dbo.Hash,
		Description: dbo.Description,
	}
}

// SetArtPublishCountDBOtoModel --
func SetArtPublishCountDBOtoModel(dbo *dbo.PublishCounterDBO) *models.PublishCounter {
	base := basis.BaseModel{}
	base.UUID = dbo.UUID
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.PublishCounter{
		Date:      dbo.Date,
		Count:     dbo.Count,
		Type:      status.RecordStatusEnum(dbo.Type),
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
