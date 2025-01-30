package dto

type File struct {
}

type FileReq struct {
}

type FileData struct {
	FileID           string `json:"file_id"`
	FileURI          string `json:"file_uri"`
	FileThumbnailURI string `json:"file_thumbnail_uri"`
}

type UpdateFileReq struct {
}
