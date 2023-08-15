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

func (r *Repository) UpdateProfilePic(userID int,
	pic *multipart.FileHeader) (string, *domain.AppError) {
	file, err := pic.Open()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("nombre-del-archivo-en-s3.txt"),
		Body:   file,
		ACL:    aws.String("public-read"),
	}

	_, err = r.s3Client.PutObject(params)
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	photoURL := fmt.Sprintf("https://my-bucket.s3.amazonaws.com/uploads/user%d/%s",
		userID, pic.Filename)

	err = file.Close()
	if err != nil {
		return "", domain.NewAppErrorWithType(domain.RepositoryError)
	}

	return photoURL, nil
}
