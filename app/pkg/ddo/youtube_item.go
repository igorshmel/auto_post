package ddo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
)

// YoutubeItemDDO --
type YoutubeItemDDO struct {
	Title   string
	VideoID string
	Status  status.RecordStatusEnum
}
