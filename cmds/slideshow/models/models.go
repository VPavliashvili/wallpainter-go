package models

type Operation interface {
	Execute() error
}

type StoredJsonDataReader interface {
    ReadData() (SlideshowDataModel, error)
}

type SlideshowDataModel struct {
    IsRunning bool
    SlideshowPictures []string
}

type JsonReaderFactory interface{
    GetReader() (StoredJsonDataReader, error)
}
