package tools

import (
	"auto_post/app/internal/domains/vk_machine/structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

// URLEncoded --
func URLEncoded(str string) string {
	if r, err := regexp.Compile(","); err == nil {
		str = r.ReplaceAllString(str, "%2C")
		if r, err = regexp.Compile("\n"); err == nil {
			str = r.ReplaceAllString(str, "%0A")
			if r, err = regexp.Compile("\r"); err == nil {
				str = r.ReplaceAllString(str, "%0D")
				if u, err := url.Parse(str); err == nil {
					return u.String()
				}
			} else {
				log.Printf("UrlEncoded error regexp.Compile \r: %v", err)
			}
		} else {
			log.Printf("UrlEncoded error regexp.Compile \n: %v", err)
		}
	} else {
		log.Printf("UrlEncoded error regexp.Compile: %v", err)
	}

	return ""
}

// VideoID --
func VideoID(link string) string {
	// Parse the URL and ensure there are no errors.
	if strings.Contains(link, "attribution_link") {
		return YTP(YTP(link, "u"), "v")
	}
	return YTP(link, "v")
}

// YTP --
func YTP(link string, key string) string {
	if u, err := url.Parse(link); err == nil {
		if fragments, err := url.ParseQuery(u.RawQuery); err == nil {
			if len(fragments[key]) > 0 {
				log.Print(fragments[key][0])
				return fragments[key][0]
			}
			log.Print("not found")
			return ""
		}
		log.Printf("error ParseQuery: %v", err)
	} else {
		log.Printf("error Parse: %v", err)
	}
	return ""
}

// RangeInt --
func RangeInt(min int, max int, n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	var r int
	for r = 0; r <= n-1; r++ {
		arr[r] = rand.Intn(max) + min

	}
	return arr
}

// PhotoWall upload file (on filePath) to given url.
// Return info about uploaded photo.
func PhotoWall(url, filePath string) (*structs.UploadPhotoWallResponse, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("photo", filePath)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			panic("failed close file")
		}
	}(fh)

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic("failed close body")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var uploaded structs.UploadPhotoWallResponse

	err = json.Unmarshal(body, &uploaded)
	if err != nil {
		return nil, err
	}

	return &uploaded, nil
}

/*
func PhotoGroup(url, filePath string) (models.UploadPhotoResponse, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file1", filePath)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}

	fh, err := os.Open(filePath)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}

	var uploaded models.UploadPhotoResponse
	err = json.Unmarshal(body, &uploaded)
	if err != nil {
		return models.UploadPhotoResponse{}, err
	}

	return uploaded, nil
}
*/

// Request provides access to VK API methods.
func Request(s string, method string, params map[string]string, st interface{}) ([]byte, error) {

	var apiURL string
	switch s {
	case "v":
		apiURL = "https://api.vk.com/method/"
	case "y":
		apiURL = "https://www.googleapis.com/youtube/v3/"
	default:
		apiURL = "https://api.vk.com/method/"
	}

	//apiURL  := "https://api.vk.com/method/"

	u, err := url.Parse(apiURL + method)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}

	//query.Set("access_token", vk.AccessToken)
	//query.Set("v", vk.Version)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var handler struct {
		Error    *structs.Error
		Response json.RawMessage
	}
	err = json.Unmarshal(body, &handler)
	err = json.Unmarshal(body, st)

	if handler.Error != nil {
		return nil, handler.Error
	}

	return handler.Response, nil
}

// ResToStruct --
func ResToStruct(b *http.Response, s interface{}) error {
	jss, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return fmt.Errorf("JSON ReadAll failed: %s", err)
	}

	if err := json.Unmarshal(jss, s); err != nil {
		return fmt.Errorf("JSON unmarshaling failed: %s", err)
	}

	return nil
}

// RespToStruct --
func RespToStruct(b []byte, s interface{}) error {

	if err := json.Unmarshal(b, s); err != nil {
		return fmt.Errorf("JSON unmarshaling failed: %s", err)
	}

	return nil
}

// ByteToString --
func ByteToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// ByteToJSON --
func ByteToJSON(b *http.Response) string {
	bs := make([]byte, 1014)
	js := ""
	for true {
		n, err := b.Body.Read(bs)
		js = js + string(bs[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return js
}
