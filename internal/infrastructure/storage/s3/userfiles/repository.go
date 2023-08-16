package userfiles

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"mime/multipart"
)

type Repository struct {
	s3Client *s3.S3
}

const (
	bucketName = "cafelatte-profile_pics"
	ACL        = "public-read"
)

func (r *Repository) UpdateProfilePic(userID int,
	pic *multipart.FileHeader, picName string) (string, *domain.AppError) {
	file, err := pic.Open()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

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
		bucketName, userID, pic.Filename)

	err = file.Close()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	return photoURL, nil
}

func NewUserFilesRepository(s3Client *s3.S3) *Repository {
	return &Repository{s3Client}
}
