package mapping

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/dto"
	"git.fintechru.org/dfa/dfa_lib/models/basis"
)

// ParseImageModelToDBO --
func ParseImageModelToDBO(req *models.ParseImage) *dbo.ParseImageDBO {
	return &dbo.ParseImageDBO{
		URL:     req.URL,
		Service: req.Service,
		Status:  req.Status,
		Hash:    req.Hash,

		UpdatedAt: req.UpdatedAt,
		CreatedAt: req.CreatedAt,
	}
}

// ParseImageDTOtoDDO --
func ParseImageDTOtoDDO(req *dto.ParseImageReqDTO) *ddo.ParseImageReqDDO {
	return &ddo.ParseImageReqDDO{
		FileURL: req.FileURL,
		Service: req.Service,
	}
}

// ParseImageDDOtoDBO --
func ParseImageDDOtoDBO(ddo *ddo.ParseImageResDDO) *dbo.ParseImageDBO {
	return &dbo.ParseImageDBO{
		UUID:      ddo.FileUUID,
		URL:       ddo.FileURL,
		AuthURL:   ddo.AuthURL,
		Service:   ddo.Service,
		Status:    ddo.Status,
		Hash:      ddo.Hash,
		UpdatedAt: ddo.UpdatedAt,
		CreatedAt: ddo.CreatedAt,
	}
}

// ParseImageDBOtoModel --
func ParseImageDBOtoModel(dbo *dbo.ParseImageDBO) *models.ParseImage {
	base := basis.BaseModel{}
	base.UUID = dbo.UUID
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.ParseImage{
		URL:       dbo.URL,
		AuthURL:   dbo.AuthURL,
		Service:   dbo.Service,
		Status:    dbo.Status,
		Hash:      dbo.Hash,
		BaseModel: base,
	}
}
