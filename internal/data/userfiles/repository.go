package userfiles

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

const ACL = "public-read"

type Repository struct {
	s3Client *s3.S3
}

func New(s3Client *s3.S3) *Repository {
	return &Repository{s3Client}
}

func (r *Repository) UpdateProfilePicById(userId int, pic *multipart.FileHeader,
	picName string) (string, *domain.AppError) {
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

	_, err = r.s3Client.PutObject(params)
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	photoURL := fmt.Sprintf("https://%s.s3.amazonaws.com/uploads/user%d/%s",
		bucketName,
		userId,
		pic.Filename)

	err = file.Close()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	return photoURL, nil
}
