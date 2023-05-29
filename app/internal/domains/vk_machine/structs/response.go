package structs

import "fmt"

// Response is the base struct for all responses
// that come back from the Pinterest API.

// Empty --
type Empty struct{}

//type Response struct {
//	Data    interface{} `json:"data"`
//	Message string      `json:"message"`
//	Type    string      `json:"type"`
//	Page    Page        `json:"page"`
//}
////type MyDat struct {
//	Data []Pin `json:"data"`
//	Page Page  `json:"page"`
//}
//type Page struct {
//	Cursor string `json:"cursor"`
//	Next   string `json:"next"`
//}

// SavePhotoWallResponse --
type SavePhotoWallResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// UploadPhotoResponse --
type UploadPhotoResponse struct {
	Server    int    `json:"server"`
	AlbumID   int    `json:"aid"`
	Hash      string `json:"hash"`
	PhotoList string `json:"photos_list"`
}

// Error --
type Error struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

// Error --
func (e *Error) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}
