// Package consts
package consts

import (
	"fmt"
	"github.com/inhies/go-bytesize"
)

const (
	//ServiceTypeHTTP marks the usecase service type for HTTP operation
	ServiceTypeHTTP = "http"

	//ServiceTypeMQ marks the usecasee service type for MQ Operation
	ServiceTypeConsumer = "consumer"
)

type memory int

func (m memory) String() string {
	bytesize.Format = "%.0f "
	return fmt.Sprintf("%s", bytesize.New(float64(m)))
}
