package base

import "github.com/spf13/pflag"

type EmailOptions struct {
	Host       string `json:"host" mapstructure:"host"`
	Port       string `json:"port" mapstructure:"port"`
	Username   string `json:"username" mapstructure:"username"`
	Password   string `json:"password" mapstructure:"password"`
	SenderName string `json:"sender-name" mapstructure:"sender-name"`
	AdminEmail string `json:"admin-email" mapstructure:"admin-email"`
}

func NewEmailOptions() *EmailOptions {
	return &EmailOptions{
		Host:       "",
		Port:       "",
		Username:   "",
		Password:   "",
		SenderName: "",
		AdminEmail: "",
	}
}

func (s *EmailOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Host, "qycms.email.host", s.Host, "发邮件服务器地址.")
	fs.StringVar(&s.Port, "qycms.email.host", s.Port, "发邮件服务器端口号.")
	fs.StringVar(&s.Username, "qycms.email.host", s.Username, "邮箱用户名.")

	fs.StringVar(&s.Password, "qycms.email.host", s.Password, "邮箱客户端密码.")
	fs.StringVar(&s.SenderName, "qycms.email.sender-name", s.SenderName, "发送人昵称.")
	fs.StringVar(&s.AdminEmail, "qycms.email.admin-email", s.AdminEmail, "发送人邮箱.")
}

func (s *EmailOptions) Validate() []error {
	errs := []error{}

	return errs
}
