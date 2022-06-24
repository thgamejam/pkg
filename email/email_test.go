package email

import (
	"github.com/stretchr/testify/assert"
	"github.com/thgamejam/pkg/conf"
	"testing"
)

func TestNewEmailService(t *testing.T) {
	s, err := NewEmailService(&conf.Email{
		Host:     "smtpdm.aliyun.com",
		Port:     80,
		Username: "mail@mailpush.test.thjam.cc",
		Password: "password",
	})
	if err != nil {
		return
	}
	err = s.SendEmail("null122", "测试", "测试测试测试测试", "test@test.com")
	assert.NoError(t, err)
}
