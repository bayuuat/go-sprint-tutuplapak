package domain

type File struct {
	FileID           int    `db:"file_id"`
	FileURI          string `db:"file_uri"`
	FileThumbnailURI string `db:"file_thumbnail_uri"`
}
