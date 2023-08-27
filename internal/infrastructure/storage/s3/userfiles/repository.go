package userfiles

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"mime/multipart"
	"os"
)

type Repository struct {
	s3Client *s3.S3
}

const ACL = "public-read"

func (repository *Repository) UpdateProfilePicById(userID int,
	pic *multipart.FileHeader, picName string) (string, *domain.AppError) {
	file, err := pic.Open()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	bucketName := os.Getenv("FILES_S3BUCKET")

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(picName),
		Body:   file,
		ACL:    aws.String(ACL),
	}

	_, err = repository.s3Client.PutObject(params)
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	photoURL := fmt.Sprintf("https://%s.s3.amazonaws.com/uploads/user%d/%s",
		bucketName, userID, pic.Filename)

	err = file.Close()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	return photoURL, nil
}

func New(s3Client *s3.S3) *Repository {
	return &Repository{s3Client}
}
