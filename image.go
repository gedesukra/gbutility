package gbutility

import (
	"image"
	"image/color"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

func ConvPict(input string, output string, width int, height int, logger *CustomLogger) string {
	var result string
	prefixLog := "ConvPict - "
	src, err := imaging.Open(input)
	if err != nil {
		logger.Println(prefixLog+"Failed to open image: %v", err)
	}
	// deal with exif
	efile, err := os.Open(input)
	if err != nil {
		logger.Println(prefixLog+"Failed to open image: %v", err)
	}
	defer efile.Close()
	x, err := exif.Decode(efile)
	if err != nil {
		if x == nil {
			// ignore - image exif data has been already stripped
		}
		logger.Println(prefixLog+"failed reading exif data in [%s]: %s", input, err.Error())
	}
	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			logger.Println(prefixLog+"%s had orientation %s", input, orient.String())
			src = reverseOrientation(src, orient.String(), logger)
		} else {
			logger.Println(prefixLog+"%s had no orientation - implying 1", input)
			src = reverseOrientation(src, "1", logger)
		}

	}

	if src.Bounds().Max.X >= src.Bounds().Max.Y {
		dstImg := imaging.Resize(src, width, 0, imaging.Lanczos)
		err = imaging.Save(dstImg, output)
		if err != nil {
			logger.Println(prefixLog+"Failed to save image: %v", err)
		}
	} else {
		dstImg := imaging.Resize(src, 0, height, imaging.Lanczos)
		err = imaging.Save(dstImg, output)
		if err != nil {
			logger.Println(prefixLog+"Failed to save image: %v", err)
		}
	}

	defer func() {
		if p := recover(); p != nil {
			logger.Println(prefixLog + "Panic, failed to convert image")
			result = ""
		}
	}()

	// RenameFile(input, output)
	DeleteFile(input)

	result = output

	return result
}

func ConvPictProfile(input string, output string, width int, height int, logger *CustomLogger) string {
	var result string
	prefixLog := "ConvPictProfile - "
	src, err := imaging.Open(input)
	if err != nil {
		logger.Println(prefixLog+"Failed to open image: %v", err)
	}
	// deal with exif
	efile, err := os.Open(input)
	if err != nil {
		logger.Println(prefixLog+"Failed to open image: %v", err)
	}
	defer efile.Close()
	x, err := exif.Decode(efile)
	if err != nil {
		if x == nil {
			// ignore - image exif data has been already stripped
		}
		logger.Println(prefixLog+"failed reading exif data in [%s]: %s", input, err.Error())
	}
	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			logger.Println(prefixLog+"%s had orientation %s", input, orient.String())
			src = reverseOrientation(src, orient.String(), logger)
		} else {
			logger.Println(prefixLog+"%s had no orientation - implying 1", input)
			src = reverseOrientation(src, "1", logger)
		}

	}

	if src.Bounds().Max.X >= src.Bounds().Max.Y {
		dstImg := imaging.Resize(src, width, 0, imaging.Lanczos)

		newimg := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
		newimg = imaging.PasteCenter(newimg, dstImg)

		err = imaging.Save(newimg, output)
		if err != nil {
			logger.Println(prefixLog+"Failed to save image: %v", err)
		}
	} else {
		dstImg := imaging.Resize(src, 0, height, imaging.Lanczos)

		newimg := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
		newimg = imaging.PasteCenter(newimg, dstImg)

		err = imaging.Save(newimg, output)
		if err != nil {
			logger.Println(prefixLog+"Failed to save image: %v", err)
		}
	}

	defer func() {
		if p := recover(); p != nil {
			logger.Println(prefixLog + "Panic, failed to convert image")
			result = ""
		}
	}()

	// RenameFile(input, output)
	DeleteFile(input)
	result = output

	return result
}

func ResizeImg(input string, output string, width int, height int) string {
	src, err := imaging.Open(input)
	if err != nil {
		// logger.Println("failed to open image: %v", err)
	}

	if src.Bounds().Max.X >= src.Bounds().Max.Y {
		dstImg := imaging.Resize(src, width, 0, imaging.Lanczos)

		newimg := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
		newimg = imaging.PasteCenter(newimg, dstImg)

		err = imaging.Save(newimg, output)
		if err != nil {
			// logger.Println("failed to save image: %v", err)
		}
	} else {
		dstImg := imaging.Resize(src, 0, height, imaging.Lanczos)

		newimg := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
		newimg = imaging.PasteCenter(newimg, dstImg)

		err = imaging.Save(newimg, output)
		if err != nil {
			// logger.Println("failed to save image: %v", err)
		}
	}

	// RenameFile(input, output)
	DeleteFile(input)

	return output
}

// reverseOrientation amply`s what ever operation is necessary to transform given orientation
// to the orientation 1
func reverseOrientation(img image.Image, o string, logger *CustomLogger) *image.NRGBA {
	prefixLog := "reverseOrientation - "
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	logger.Println(prefixLog+"unknown orientation %s, expect 1-8", o)
	return imaging.Clone(img)
}
