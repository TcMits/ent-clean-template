package v1

import (
	"html/template"

	_ "github.com/TcMits/ent-clean-template/docs"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
	"github.com/swaggo/swag"
)

const (
	_swaggerDocPath     = _v1SubPath + "/swagger/doc.json"
	_swaggerSubPath     = "/swagger"
	_swaggerWideSubPath = "/swagger/{any:path}"
	_swaggerDocSubPath  = "/swagger/doc.json"
	_swaggerUIVersion   = "4.5.0"
)

type swaggerUIBundle struct {
	URL         string
	DeepLinking bool
	Version     string
}

func RegisterDocsController(handler iris.Party, l logger.Interface) {
	docHandler := getDocHandler()
	handler.Get(_swaggerDocSubPath, docJSONHandler)
	handler.Get(_swaggerSubPath, docHandler)
	handler.Get(_swaggerWideSubPath, docHandler)
}

func docJSONHandler(ctx iris.Context) {
	ctx.ContentType("application/json")
	doc, err := swag.ReadDoc()
	if err != nil {
		panic(err)
	}
	ctx.Write([]byte(doc))
}

func getDocHandler() iris.Handler {
	// create a template with name
	t := template.New("swagger_index.html")
	index, _ := t.Parse(indexTmpl)
	swagBundle := &swaggerUIBundle{
		URL:         _swaggerDocPath,
		DeepLinking: true,
		Version:     _swaggerUIVersion,
	}

	return func(ctx iris.Context) {
		ctx.ContentType("text/html")
		index.Execute(ctx.ResponseWriter(), swagBundle)
	}
}

const indexTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta
      name="description"
      content="SwaggerUI"
    />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@{{ .Version }}/swagger-ui.css" />
  </head>
  <body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@{{ .Version }}/swagger-ui-bundle.js" crossorigin></script>
  <script src="https://unpkg.com/swagger-ui-dist@{{ .Version }}/swagger-ui-standalone-preset.js" crossorigin></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: "{{.URL}}",
        dom_id: '#swagger-ui',
        validatorUrl: null,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
        deepLinking: {{.DeepLinking}}
      });
    };
  </script>
  </body>
</html>
`
