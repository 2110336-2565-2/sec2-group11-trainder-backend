package controllers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"path/filepath"

	// "mime/multipart"
	"net/http"

	// "os"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

func UploadProfile() gin.HandlerFunc {
	fmt.Println("UploadProfile")
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
			return
		}

		if file.Size <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is empty"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		defer f.Close()

		fileID, err := models.Upload(file.Filename, f)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("File uploaded with ID: %v\n", fileID)

		c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})

		// update profile with avatar URI = fileID
		username, err := tokens.ExtractTokenUsername(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ProfileResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		_, err = models.UpdateAvatarUrl(username, fileID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}

	}
}
func getContentType(filename string) string {
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	return contentType
}

func GetPicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Query("ID")
		username := c.Query("username")
		imageId, err := models.GetAvatarUrl(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		downloadStream, filename, err := models.RetrieveFileFromMongo(imageId)
		if err != nil {
			log.Fatal(err)
		}
		defer downloadStream.Close()

		fmt.Println(getContentType(filename))

		// Set the response headers
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", getContentType(filename))

		// Copy the file data to the response writer
		_, err = io.Copy(c.Writer, downloadStream)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error3": "Error writing file to response"})
			return
		}
		// fmt.Println("c.Writer")
		// // out, err := os.Create(filename)
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }
		// // defer out.Close()

		// // _, err = io.Copy(out, downloadStream)
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }

		log.Println("File send")

	}
}
