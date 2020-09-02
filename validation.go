package gbutility

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func ValidateContentVideo(content string) bool {
	res := false
	switch content {
	case "application/vnd.apple.mpegurl":
		res = true
	case "application/x-mpegurl":
		res = true
	case "video/3gpp":
		res = true
	case "video/mp4":
		res = true
	case "video/mpeg":
		res = true
	case "video/ogg":
		res = true
	case "video/quicktime":
		res = true
	case "video/webm":
		res = true
	case "video/x-m4v":
		res = true
	case "video/ms-asf":
		res = true
	case "video/x-ms-wmv":
		res = true
	case "video/x-msvideo":
		res = true
	case "video/x-flv":
		res = true
	default:
		res = false
	}

	return res
}

func ValidateContentAudio(content string) bool {
	res := false
	switch content {
	case "audio/basic":
		res = true
	case "auido/L24":
		res = true
	case "audio/mid":
		res = true
	case "audio/mpeg":
		res = true
	case "audio/mp4":
		res = true
	case "audio/x-aiff":
		res = true
	case "audio/x-mpegurl":
		res = true
	case "audio/vnd.rn-realaudio":
		res = true
	case "audio/ogg":
		res = true
	case "audio/vorbis":
		res = true
	case "audio/vnd.wav":
		res = true
	default:
		res = false
	}

	return res
}

func ValidateContentImg(content string) bool {
	if strings.Contains(content, "image") {
		return true
	} else {
		return false
	}
}
