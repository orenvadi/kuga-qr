package app

import (
	"net/http"

	"github.com/orenvadi/kuga-lms/openapi/pkg/api"
)

func serveOpenAPI(w http.ResponseWriter, r *http.Request) {
	spec, err := api.GetSwagger()
	if err != nil {
		http.Error(w, "Failed to load OpenAPI spec", http.StatusInternalServerError)
		return
	}

	data, err := spec.MarshalJSON()
	if err != nil {
		http.Error(w, "Failed to serialize OpenAPI spec", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/yaml")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	// Use an official CDN
	const swaggerPage = `
<!DOCTYPE html>
<html>
  <head>
    <title>Swagger UI - Kuga LMS</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js"></script>
    <script>
      SwaggerUIBundle({
        url: '/openapi.json',
        dom_id: '#swagger-ui',
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIBundle.SwaggerUIStandalonePreset
        ],
        layout: "BaseLayout"
      });
    </script>
  </body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(swaggerPage))
}
