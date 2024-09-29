package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"
)

func ConvertImages(Base64 string) (images image.Image, Type string, err error) {
	if strings.Contains(Base64, "image/jpeg") || strings.Contains(Base64, "image/jpg") {
		var base64images = strings.Split(Base64, ",")
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64images[1]))

		images, err = jpeg.Decode(reader)
		if err != nil {
			return images, Type, err
		}
		return images, ".jpg", err
	} else if strings.Contains(Base64, "image/png") {
		var base64images = strings.Split(Base64, ",")
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64images[1]))
		images, err = png.Decode(reader)
		if err != nil {
			return images, "", err
		}
		return images, ".png", err
	} else {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(Base64))

		images, _, err = image.Decode(reader)
		if err != nil {
			return images, "", err
		}
		return images, ".png", err
	}

}

func SaveImage(img image.Image, types, DataPath string) (err error) {

	outputFile, err := os.Create(DataPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if types == ".png" {
		err = png.Encode(outputFile, img)
		if err != nil {
			return err
		}
	} else if types == ".jpg" {
		var opt jpeg.Options
		opt.Quality = 80

		err = jpeg.Encode(outputFile, img, &opt)
		if err != nil {
			return err
		}
	}

	return nil
}

func RandomNumber(n int) int {
	var res []int

	for i := 0; i < n; i++ {
		angka := rand.Intn(10)
		res = append(res, angka)
	}

	return sliceToInt(res)
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func ShortDur(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d Jam %02d Menit", h, m)
}
