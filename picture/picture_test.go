package picture

import (
	"github.com/kr/pretty"
	"github.com/zhang555/crawler1/db"
	"github.com/zhang555/crawler1/logger"
	"image"
	"image/jpeg"
	"testing"
)

func TestName(t *testing.T) {
	logger.InitLog()

	db.InitMysql()

	image.RegisterFormat("jpg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)

	//url1 := `https://gdb.voanews.com/1B4B9011-183F-4FB8-8D69-C722705A2EC9_w100_r1.jpg`

	url1 :=`https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTBuTE0vtoheqNfFZ4OUzO6-5EdFoJZqzuS1PRkFPoJhA5Od30nV2X3AOmSjwg&s`

	res, err := GetAndDecodePicture(url1)
	if err != nil {
		return
	}
	//log.Log.Println(res)
	//log.Log.Println(res.Image.Bounds())
	//log.Log.Println(res.Image.Bounds().Min)
	//log.Log.Println(res.Image.Bounds().Min.X)
	//log.Log.Println(res.Image.Bounds().Min.Y)
	//log.Log.Println(res.Image.Bounds().Max)
	//log.Log.Println(res.Image.Bounds().Max.X)
	//log.Log.Println(res.Image.Bounds().Max.Y)

	//
	logger.Log.Println(res.Image.Bounds().Max.X - res.Image.Bounds().Min.X)
	logger.Log.Println(res.Image.Bounds().Max.Y - res.Image.Bounds().Min.Y)

	pretty.Println(res.Image.Bounds())

}

func TestName1(t *testing.T) {
	logger.InitLog()

	db.InitMysql()

	image.RegisterFormat("jpg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)

	//url1 := `https://gdb.voanews.com/1B4B9011-183F-4FB8-8D69-C722705A2EC9_w100_r1.jpg`
	//
	//res, err := GetAndDecodePicture(url1)
	//if err != nil {
	//	return
	//}
	//
	//
	////
	//log.Log.Println(res.Image.Bounds().Max.X - res.Image.Bounds().Min.X)
	//log.Log.Println(res.Image.Bounds().Max.Y - res.Image.Bounds().Min.Y)

	//db.DB.SetLogger()

	db.DB.LogMode(true)

	Picture()
}

//func TestName2(t *testing.T) {
//
//	logger.InitLog()
//
//	db.InitMysql()
//
//	db.DB.LogMode(true)
//
//	//var wikiImages []model.WikiImage
//	var wikiImage model.WikiImage
//	wikiImage.ID = 1
//
//	m := map[string]interface{}{
//		`x`: 1,
//		`y`: 11,
//	}
//
//	err := db.DB.Model(&wikiImage).Updates(m).Error
//
//	//logger.Log.Println(err)
//	log.Println(err)
//
//}
