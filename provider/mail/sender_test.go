package mail

import (
	"errors"
	"iTask/config"

	"github.com/spf13/viper"
)

func LoadConfigTest() (*config.Config, error) {
	v := viper.New()

	v.AddConfigPath("../../config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// check is not found file config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c config.Config // Unmarshal data config have get in file config then get into c
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// func TestSendEmailWithGmail(t *testing.T) {
// 	cfg, err := LoadConfigTest()
// 	require.NoError(t, err)
// 	if err != nil {
// 		log.Fatalln("Get config error", err)
// 		return
// 	}

// 	sender := NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)
// 	subject := "Test send email"
// 	content := `
// 		<h1>Test send email</h1>
// 		<p>Test send email</p>
// 		`
// 	to := []string{"phatbfbf@gmail.com"}
// 	attachFileq := []string{"../../tmp/test_send_mail.md"}
// 	err = sender.SendEmail(subject, content, to, nil, nil, attachFileq)
// 	require.NoError(t, err)
// }
