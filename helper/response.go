package helper

// Helper to generate error response
func ErrorResponseBuilder(status, msg string) map[string]interface{} {
	res := map[string]interface{}{
		"status":  status,
		"message": msg,
	}
	return res

}

// Helper to generate respons
func ResponseBuilder(msg string, data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"data":    data,
	}
	return res
}
