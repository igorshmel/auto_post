package ddo

// VKMachine --
type VKMachine struct {
	FileName string
	URL      string
}

// GetWallUploadServer --
type GetWallUploadServer struct {
	Params     map[string]string
	MethodName string
}

// ResSaveWallPhoto --
type ResSaveWallPhoto struct {
	Params     map[string]string
	MethodName string
}

// ReqSaveWallPhoto --
type ReqSaveWallPhoto struct {
	Photo  string
	Hash   string
	Server int
}

// ResPostWallPhoto --
type ResPostWallPhoto struct {
	Params     map[string]string
	MethodName string
}

// ReqPostWallPhoto --
type ReqPostWallPhoto struct {
	AuthURL string
	URL     string
	OwnerID int32
	ID      int32
}

// ResGetUploadServer --
type ResGetUploadServer struct {
	Params     map[string]string
	MethodName string
}

// ReqPhotosSave --
type ReqPhotosSave struct {
	Server    int    `json:"server"`
	AlbumID   int    `json:"aid"`
	Hash      string `json:"hash"`
	PhotoList string `json:"photos_list"`
	URL       string `json:"url"`
}

// ResPhotosSave --
type ResPhotosSave struct {
	Params     map[string]string
	MethodName string
}
