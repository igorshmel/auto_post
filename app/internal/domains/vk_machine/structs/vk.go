package structs

// PostPic --
type PostPic struct {
	ID         int
	Originlink string
}

// VkPicIni --
type VkPicIni struct {
	PreTimeOut  int
	MaxHashTags int
	VkHashTags  string
	VkToken     string
	VkGroupID   string
	VkAlbumID   string
}

// VkVideoGetErr --
type VkVideoGetErr struct {
	Error struct {
		ErrorCode     int    `json:"error_code"`
		ErrorMsg      string `json:"error_msg"`
		RequestParams []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"request_params"`
	} `json:"error"`
}

// VkVideoGet --
type VkVideoGet struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int    `json:"id"`
			OwnerID     int    `json:"owner_id"`
			Title       string `json:"title"`
			Duration    int    `json:"duration"`
			Description string `json:"description"`
			Date        int    `json:"date"`
			Comments    int    `json:"comments"`
			Views       int    `json:"views"`
			LocalViews  int    `json:"local_views"`
			Image       []struct {
				Height      int    `json:"height"`
				URL         string `json:"url"`
				Width       int    `json:"width"`
				WithPadding int    `json:"with_padding"`
			} `json:"image"`
			IsFavorite bool `json:"is_favorite"`
			AddingDate int  `json:"adding_date"`
			Files      struct {
				External string `json:"external"`
			} `json:"files"`
			Player        string `json:"player"`
			Platform      string `json:"platform"`
			CanEdit       int    `json:"can_edit"`
			Converting    int    `json:"converting"`
			CanAdd        int    `json:"can_add"`
			CanComment    int    `json:"can_comment"`
			CanRepost     int    `json:"can_repost"`
			CanLike       int    `json:"can_like"`
			CanAddToFaves int    `json:"can_add_to_faves"`
			Type          string `json:"type"`
		} `json:"items"`
	} `json:"response"`
}

// VkGetWallUploadServer --
type VkGetWallUploadServer struct {
	Response VkGetWallUploadServerResponse `json:"response"`
}

// VkGetWallUploadServerResponse --
type VkGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// VkSaveWallPhoto --
type VkSaveWallPhoto struct {
	Response []VkSaveWallPhotoRes `json:"response"`
}

// VkSaveWallPhotoRes --
type VkSaveWallPhotoRes struct {
	ID      int32 `json:"id"`
	OwnerID int32 `json:"owner_id"`
}
