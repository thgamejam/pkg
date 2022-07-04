package authentication

import (
	"github.com/thgamejam/pkg/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	secret       = "61bc667fe31f47f7a312ee177be915dd"
	uploadClaims = UploadFileClaims{
		Bucket:    "bucket",
		Name:      "name.jpg",
		UUID:      uuid.New(),
		ExpiresAt: uint32(time.Now().Unix()) + 3600,
		CRC:       "70930f27",
		SHA1:      "7c4a8d09ca3762af61e59520943dc26494f8941b",
	}
	calculatedHash, _ = infoToMD5(&uploadClaims, &secret)
)

func TestCreateUploadURL(t *testing.T) {
	// 创建url
	url, err := CreateUploadURL(&uploadClaims, &secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, url)
	t.Logf("log url:=%v\n", url)
}

func TestValidateUploadInfo(t *testing.T) {
	// 校验hash
	success, err := ValidateUploadInfo(&uploadClaims, &secret, &calculatedHash)
	assert.NoError(t, err)
	assert.True(t, success) // 判断hash校验是否正确

	// 创建过期的claims
	overdueClaims := uploadClaims
	overdueClaims.ExpiresAt = uint32(time.Now().Unix())
	overdueClaimsHash, _ := infoToMD5(&overdueClaims, &secret)

	// 延时
	time.Sleep(time.Second)

	success, err = ValidateUploadInfo(&overdueClaims, &secret, &overdueClaimsHash)
	assert.NoError(t, err)
	assert.False(t, success)
}

var (
	uploadSliceClaims = UploadSliceFileClaims{
		UploadFileClaims: uploadClaims,
		SliceID:          1,
		Sum:              10,
	}
	calculatedSliceHash, _ = infoToMD5(&uploadSliceClaims, &secret)
)

func TestCreateSliceUploadURL(t *testing.T) {
	// 创建url
	url, err := CreateSliceUploadURL(&uploadSliceClaims, &secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, url)
	t.Logf("log url:=%v\n", url)
}

func TestValidateSliceUploadInfo(t *testing.T) {
	// 校验hash
	success, err := ValidateSliceUploadInfo(&uploadSliceClaims, &secret, &calculatedSliceHash)
	assert.NoError(t, err)
	assert.True(t, success) // 判断hash校验是否正确

	// 创建过期的claims
	overdueClaims := uploadSliceClaims
	overdueClaims.ExpiresAt = uint32(time.Now().Unix())
	overdueClaimsHash, _ := infoToMD5(&overdueClaims, &secret)

	// 延时
	time.Sleep(time.Second)

	success, err = ValidateSliceUploadInfo(&overdueClaims, &secret, &overdueClaimsHash)
	assert.NoError(t, err)
	assert.False(t, success)
}
