package v3

import (
	"os"

	"github.com/swaggo/swag"
)

const _jsonFile = "docs/v3/openapi.json"

func init() {
	spec, err := os.ReadFile(_jsonFile)
	if err != nil {
		spec = []byte{}
	}
	swag.Register(swag.Name, &swag.Spec{SwaggerTemplate: string(spec)})
}
