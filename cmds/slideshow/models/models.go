package models

type Operation interface {
	Execute() error
}

type StoredJsonDataHandler interface {
    WriteData() error
    ReadData() (SlideshowDataModel, error)
}

type SlideshowDataModel struct {
    IsRunning bool
    CyclingPictures []string
}
