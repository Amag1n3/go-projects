package imageutils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// downloadImage downloads an image from a URL and saves it to a local file.
func DownloadImage(imageURL string) (string, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download image: %s", resp.Status)
	}

	tmpFile, err := os.CreateTemp("", "recipe-image-*.png")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

// displayImage uses chafa to display the image.
func DisplayImage(imagePath string) error {
	cmd := exec.Command("chafa", imagePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s", err)
	}
	fmt.Println(string(output))
	return nil
}

// cleanUp removes the temporary image file.
func CleanUp(imagePath string) error {
	return os.Remove(imagePath)
}
