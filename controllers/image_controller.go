package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson/primitive"

	// "mime/multipart"
	"net/http"

	// "os"

	"trainder-api/models"
	"trainder-api/responses"
	"trainder-api/utils/tokens"

	"github.com/gin-gonic/gin"
)

// @Summary Upload profile picture
// @Description Upload profile picture
// @Tags image
// @Accept png
// @Produce json
// @Param image formData file true "file for uploading"
// @Security BearerAuth
// @Success 200 {object} responses.ImageResponse
// @Failure 400 {object} responses.ImageResponse
// @Router /protected/image [POST]
func UploadProfile() gin.HandlerFunc {
	fmt.Println("UploadProfile")
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})

			return
		}

		if file.Size <= 0 {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: "Uploaded file is empty: " + err.Error(),
			})
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file is empty"})
			return
		}

		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		defer f.Close()

		fileID, err := models.Upload(file.Filename, f)
		if err != nil {
			log.Fatal(err)
		}
		// log.Printf("File uploaded with ID: %v\n", fileID)

		// c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})

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
			c.JSON(http.StatusInternalServerError, responses.ProfileResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.ImageResponse{
			Status:  http.StatusOK,
			Message: "Image uploaded successfully",
		})

	}
}
func getContentType(filename string) string {
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	return contentType
}

func isObjectID(value string) bool {
	_, err := primitive.ObjectIDFromHex(value)
	return err == nil
}

// return picture not json

// @Summary retrieve  profile picture
// @Description retrieve profile picture by username return image
// @Tags image
// @Accept json
// @Produce png
// @Security BearerAuth
// @Param username query string true "username of the person you want profile picture"
// @Success 200 {object} responses.ImageResponse
// @Failure 400 {object} responses.ImageResponse
// @Failure 500 {object} responses.ImageResponse
// @Router /protected/image2 [GET]
func GetPicture2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Query("ID")
		username := c.Query("username")
		imageId, err := models.GetAvatarUrl(username)
		if imageId != "" && (!isObjectID(imageId)) {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: "there is no profile picture for this user",
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		downloadStream, filename, err := models.RetrieveFileFromMongo(imageId)
		if err != nil {
			log.Fatal(err)
		}
		defer downloadStream.Close()

		// fmt.Println(getContentType(filename))

		// Set the response headers
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", getContentType(filename))

		// Copy the file data to the response writer
		_, err = io.Copy(c.Writer, downloadStream)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: "Error writing file to response" + err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error3": "Error writing file to response"})
			return
		}

		// log.Println("File send")

	}
}

// @Summary retrieve  profile picture
// @Description retrieve profile picture by username return json
// @Tags image
// @Accept json
// @Produce json
// @Param username query string true "username of the person you want profile picture"
// @Security BearerAuth
// @Success 200 {object} responses.ImageResponse
// @Failure 400 {object} responses.ImageResponse
// @Failure 500 {object} responses.ImageResponse
// @Router /protected/image [GET]
func GetPicture() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Query("ID")
		username := c.Query("username")
		imageId, err := models.GetAvatarUrl(username)
		if imageId != "" && (!isObjectID(imageId)) {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: "there is no profile picture for this user",
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ImageResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		downloadStream, filename, err := models.RetrieveFileFromMongo(imageId)
		if err != nil {
			// log.Fatal(err)
			c.JSON(http.StatusInternalServerError, responses.ImageResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		defer downloadStream.Close()

		// Read the image data into a byte slice
		imgData, err := ioutil.ReadAll(downloadStream)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ImageResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error reading image data:" + err.Error(),
			})
			// c.JSON(http.StatusInternalServerError, gin.H{"error3": "Error reading image data"})
			return
		}

		// Encode the image data as a base64 string
		imgBase64 := base64.StdEncoding.EncodeToString(imgData)

		// Set the response headers and send the JSON response
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses.ImageResponse{
			Status:  http.StatusOK,
			Message: imgBase64,
		})
		// c.JSON(http.StatusOK, responses.ImageResponse{
		// 	Status:  http.StatusOK,
		// 	Message: "imgBase64",
		// })
		// log.Fatal("fff", imgBase64)

	}
}
