package bot

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func generateContentType(params map[string]string, files map[string]string, custom_byte *bytes.Buffer) (string, error) {
	writer := multipart.NewWriter(custom_byte)

	for key, value := range params {
		writer.WriteField(key, value)
	}

	if len(files) > 0 {
		for key, value := range files {
			writeMultipart(value, writer, key)
		}
	}

	err := writer.Close()

	return writer.FormDataContentType(), err
}

func writeMultipart(path string, writer *multipart.Writer, key string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file content: %w", err)
	}
	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file content: %w", err)
	}
	file.Close()

	part, err := writer.CreateFormFile(key, fi.Name())
	if err != nil {
		return fmt.Errorf("failed to create multipart form: %w", err)
	}

	part.Write(fileContent)

	return err
}
