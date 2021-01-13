package middlewares

import "fmt"

func flowNaming(methodHTTP, path string) string {
	fullPath := fmt.Sprintf("%s %s", methodHTTP, path)

	switch fullPath {
	case "GET /prueba-log":
		return "flujo de prueba"
	default:
		return "no flow found"
	}
}
