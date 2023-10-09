package helper

import (
	"bytes"
	"strconv"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/pkg/util"

	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func MultipartFormFile(r *http.Request, field string, maxFileSize int64, extension []string) (*appctx.Response, *presentations.File, error) {
	f, h, err := r.FormFile(field)
	if err == http.ErrMissingFile {
		return appctx.NewResponse().
				WithCode(consts.CodeSuccess).
				WithError([]presentations.MessageErrorValidation{
					{
						Field:    field,
						Messages: []string{fmt.Sprintf("The %s image field required", field)},
					},
				}),
			nil,
			fmt.Errorf("the %s image field required", field)
	}

	if err != nil {
		return appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(`Validation(s) Error`).WithError([]presentations.MessageErrorValidation{{
			Field:    field,
			Messages: []string{fmt.Sprintf("parse form file %s err %v", field, err)},
		},
		}), nil, fmt.Errorf("parse form file %s err %v", field, err)
	}

	var buff bytes.Buffer
	size, err := buff.ReadFrom(f)
	if err != nil {
		return appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(`Validation(s) Error`).WithError([]presentations.MessageErrorValidation{{
			Field:    field,
			Messages: []string{fmt.Sprintf("parse %s error", field)},
		}}), nil, nil
	}

	if maxFileSize != 0 && size > maxFileSize {
		return appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(`Validation(s) Error`).WithError([]presentations.MessageErrorValidation{{
				Field: field,
				Messages: []string{fmt.Sprintf("the %s image file size %s is too large, max allow is %s",
					field,
					strconv.FormatInt(h.Size, 10),
					strconv.FormatInt(maxFileSize, 10)),
				},
			}}),
			nil,
			fmt.Errorf("the %s image file size %s is too large, max allow is %s",
				field,
				strconv.FormatInt(h.Size, 10),
				strconv.FormatInt(maxFileSize, 10))
	}

	ct := ExtractFileExtension(buff.Bytes())
	ext := FileExtension(ct)

	if len(extension) != 0 && !ValidFileExtension(ext, extension) {
		return appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(`Validation(s) Error`).WithError([]presentations.MessageErrorValidation{{
				Field:    field,
				Messages: []string{fmt.Sprintf("the %s image file extension .%s not allowed, request original file content type is %s, only allow: %s", field, ext, ct, strings.Join(extension, ", "))},
			}}),
			nil,
			fmt.Errorf("the %s image file extension  .%s not allowed, request original file content type is %s, only allow: %s", field, ext, ct, strings.Join(extension, ", "))
	}

	fileExt := GetFileNameExtension(h.Filename)
	if len(extension) != 0 && !ValidFileExtension(fileExt, extension) && fileExt != "" {
		return appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(`Validation(s) Error`).WithError([]presentations.MessageErrorValidation{{
			Field:    field,
			Messages: []string{fmt.Sprintf("the %s image file extension .%s not allowed, only know: %s", field, fileExt, strings.Join(extension, ", "))},
		}}), nil, fmt.Errorf(fmt.Sprintf("the %s image file extension .%s not allowed, only know: %s", field, fileExt, strings.Join(extension, ", ")))
	}

	result := &presentations.File{
		Filename:    h.Filename,
		Ext:         fileExt,
		ContentType: ct,
		Size:        size,
		Buffer:      &buff,
	}

	return appctx.NewResponse().WithCode(consts.CodeSuccess), result, nil
}

func ExtractFileExtension(data []byte) string {
	return http.DetectContentType(data)
}

func FileExtension(input string) string {
	if len(input) < 1 {
		return input
	}
	return input[strings.IndexByte(input, '/')+1:]
}

func ValidFileExtension(ext string, extension []string) bool {
	return util.InArray(ext, extension)
}

func GetFileNameExtension(input string) string {
	ex := filepath.Ext(input)
	if ex == "" {
		return ex
	}

	return strings.ToLower(ex[1:len(ex)])
}
