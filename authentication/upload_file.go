package authentication

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/thgamejam/pkg/uuid"
)

// UploadFileClaims 单个上传文件的token
type UploadFileClaims struct {
	Bucket    string    `json:"bucket"` // 需要上传文件所在的oss桶
	Name      string    `json:"name"`   // oss中设置的文件名
	UUID      uuid.UUID `json:"uuid"`   // 链接的唯一id 实现幂等性
	ExpiresAt uint32    `json:"exp"`    // 链接到期时间的时间戳
	CRC       string    `json:"crc"`    // 上传文件的crc-32-hash值
	SHA1      string    `json:"sha1"`   // 上传文件的sha1-hash值
}

var infoToMD5 = func(v interface{}, secretKey *string) (string, error) {
	byteJson, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	byteJson = append(byteJson, *secretKey...)
	byteHash := md5.Sum(byteJson)
	hash := hex.EncodeToString(byteHash[:])
	return hash, nil
}

const UploadFileURL = "/web/v1/file/%s/%s?uuid=%s&exp=%d&hash=%s&crc=%s&sha1=%s"

// CreateUploadURL 创建单个文件上传使用的URL
func CreateUploadURL(claims *UploadFileClaims, secretKey *string) (string, error) {
	calculatedHash, err := infoToMD5(claims, secretKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(UploadFileURL,
		claims.Bucket,
		claims.Name,
		claims.UUID.String(),
		claims.ExpiresAt,
		calculatedHash,
		claims.CRC,
		claims.SHA1,
	), nil
}

// ValidateUploadInfo 验证单个文件上传信息
func ValidateUploadInfo(claims *UploadFileClaims, secretKey *string, hash *string) (bool, error) {
	calculatedHash, err := infoToMD5(claims, secretKey)
	if err != nil {
		return false, err
	}

	if calculatedHash != *hash {
		return false, nil
	}

	timestamp := uint32(time.Now().Unix())
	if timestamp > claims.ExpiresAt {
		return false, nil
	}

	return true, nil
}

// UploadSliceFileClaims 多文件分片上传文件的token
type UploadSliceFileClaims struct {
	UploadFileClaims
	SliceID uint32 `json:"slice_id"` // 文件分配的id
	Sum     uint32 `json:"sum"`      // 分片的数量总和
}

const UploadSliceFileURL = "/web/v1/files/%s/%s/%d?sum=%d&uuid=%s&exp=%d&hash=%s&crc=%s&sha1=%s"

// CreateSliceUploadURL 创建大文件分片文件上传使用的URL
func CreateSliceUploadURL(claims *UploadSliceFileClaims, secretKey *string) (string, error) {
	calculatedHash, err := infoToMD5(claims, secretKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(UploadSliceFileURL,
		claims.Bucket,
		claims.Name,
		claims.SliceID,
		claims.Sum,
		claims.UUID.String(),
		claims.ExpiresAt,
		calculatedHash,
		claims.CRC,
		claims.SHA1,
	), nil
}

// ValidateSliceUploadInfo 验证分片文件上传信息
func ValidateSliceUploadInfo(claims *UploadSliceFileClaims, secretKey *string, hash *string) (bool, error) {
	calculatedHash, err := infoToMD5(claims, secretKey)
	if err != nil {
		return false, err
	}

	if calculatedHash != *hash {
		return false, nil
	}

	timestamp := uint32(time.Now().Unix())
	if timestamp > claims.ExpiresAt {
		return false, nil
	}

	return true, nil
}
