package mapping

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/dto"
	"git.fintechru.org/dfa/dfa_lib/models/basis"
)

// FileModelToDBO --
func FileModelToDBO(req *models.FileModel) *dbo.FileDBO {
	return &dbo.FileDBO{
		FileURL: req.FileURL,
		Service: req.Service,
		Status:  req.Status,
		Hash:    req.Hash,

		UpdatedAt: req.UpdatedAt,
		CreatedAt: req.CreatedAt,
	}
}

// FileDTOtoDDO --
func FileDTOtoDDO(req *dto.ReqDownloadImage) *ddo.ReqFileDDO {
	return &ddo.ReqFileDDO{
		FileURL: req.FileURL,
		Service: req.Service,
	}
}

// FileDDOtoDBO --
func FileDDOtoDBO(ddo *ddo.ResFileDDO) *dbo.FileDBO {
	return &dbo.FileDBO{
		FileUUID:  ddo.FileUUID,
		FileURL:   ddo.FileURL,
		Service:   ddo.Service,
		Status:    ddo.Status,
		Hash:      ddo.Hash,
		UpdatedAt: ddo.UpdatedAt,
		CreatedAt: ddo.CreatedAt,
	}
}

// FileDBOtoModel --
func FileDBOtoModel(dbo *dbo.FileDBO) *models.FileModel {
	base := basis.BaseModel{}
	base.UUID = dbo.FileUUID
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt

	return &models.FileModel{
		FileURL:   dbo.FileURL,
		Service:   dbo.Service,
		Status:    dbo.Status,
		Hash:      dbo.Hash,
		BaseModel: base,
	}
}
