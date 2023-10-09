package validator

import (
	"strings"
	"technical_test_go/technical_test_go/internal/consts"
)

const (
	EnumImageSelfieContentType   enum = "validation_image_selfie_content_type"
	EnumImageIdentityContentType enum = "validation_image_identity_content_type"
)

var (
	enumMapping = map[enum]enumValue{
		EnumImageSelfieContentType:   consts.ImageSelfieContentType,
		EnumImageIdentityContentType: consts.ImageIdentityContentType,
	}
)

type enum string

type enumValue []string

func (ev enumValue) String() string {
	return strings.Join(ev, ",")
}
