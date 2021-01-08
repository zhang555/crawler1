package picture

import (
	"github.com/zhang555/crawler1/db"
	"github.com/zhang555/crawler1/logger"
	"github.com/zhang555/crawler1/model"
	"image"
	"net/http"
)

func Picture() {

	pageSize := 10

	for i := 0; i < 10000; i++ {

		var wikiImages []model.WikiImage
		db.DB.Where(`x = 0 and y = 0`).
			Offset(i * pageSize).Limit(pageSize).Find(&wikiImages)

		if len(wikiImages) == 0 {
			return
		}

		for _, wikiImage := range wikiImages {

			res, err := GetAndDecodePicture(wikiImage.ImageUrl)
			if err != nil {
				continue
			}

			logger.Log.Println(res)
			logger.Log.Println(wikiImage)

			err = db.DB.Model(&wikiImage).Updates(map[string]interface{}{
				`x`: res.Image.Bounds().Max.X - res.Image.Bounds().Min.X,
				`y`: res.Image.Bounds().Max.Y - res.Image.Bounds().Min.Y,
			}).Error
			if err != nil {
				logger.Log.Println(err)
				continue
			} else {
				logger.Log.Println(`success`)

			}
		}
	}

}

type GetAndDecodePictureRes struct {
	Image image.Image
}

func GetAndDecodePicture(url1 string) (*GetAndDecodePictureRes, error) {

	resp, err := http.Get(url1)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//判断图片的类型
	image1, _, err := image.Decode(resp.Body)
	if err != nil {
		logger.Log.Error("GetAndDecodePicture - image.Decode(readers[0]) ",
			err,
		)
		return nil, err
	}

	//log.Log.Println(image1)
	return &GetAndDecodePictureRes{Image: image1}, nil

}
