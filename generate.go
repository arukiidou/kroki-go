package kroki

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// FromString takes a string and returns the image generated by Kroki
func (c *Client) FromString(input string, graphFormat GraphFormat, imageFormat ImageFormat) (string, error) {
	payload, err := CreatePayload(input)
	if err != nil {
		return "", err
	}
	return c.GetRequest(payload, graphFormat, imageFormat)
}

// FromFile takes a file path and returns the image generated by Kroki
func (c *Client) FromFile(path string, graphFormat GraphFormat, imageFormat ImageFormat) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, "fail to read file '%s'", path)
	}
	payload, err := CreatePayload(string(content))
	if err != nil {
		return "", err
	}
	return c.GetRequest(payload, graphFormat, imageFormat)
}

// WriteToFile takes a file path and a string
// write the string into the file
func (c *Client) WriteToFile(path string, result string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "fail to create file '%s'", path)
	}
	defer file.Close()
	_, err = file.Write([]byte(result))
	if err != nil {
		return errors.Wrapf(err, "fail to write to file '%s'", path)
	}
	err = file.Sync()
	if err != nil {
		return errors.Wrapf(err, "fail to sync file '%s'", path)
	}
	return nil
}