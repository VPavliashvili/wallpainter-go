package sharedbehaviour

import (
	"encoding/json"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
)

type dataReader struct {
	json string
}

func (d dataReader) ReadData() (models.SlideshowDataModel, error) {
	data := models.SlideshowDataModel{}
	err := json.Unmarshal([]byte(d.json), &data)

	if err != nil {
		return models.SlideshowDataModel{}, models.JsonDeserializationError{
			WrongJson: d.json,
		}
	}

	return data, nil
}

func GetDataReader(dataAsJson string) models.StoredJsonDataReader {
	return dataReader{json: dataAsJson}
}
