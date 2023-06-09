/*
 * @Author: flwfdd
 * @Date: 2023-03-20 13:44:32
 * @LastEditTime: 2023-03-20 14:27:33
 * @Description: _(:з」∠)_
 */
package mail

import (
	"BIT101-GO/util/config"
	"net/smtp"
)

func Send(to string, title string, text string) error {
	mail_host := config.Config.Mail.Host
	mail_user := config.Config.Mail.User
	mail_pass := config.Config.Mail.Password
	auth := smtp.PlainAuth("", mail_user, mail_pass, mail_host)

	// 编写发送的消息
	msg := []byte("To: " + to +
		"\r\nFrom: " + mail_user +
		"\r\nSubject: " + title +
		"\r\n\r\n" + text)

	// 调用函数发送邮件
	return smtp.SendMail(mail_host+":25", auth, mail_user, []string{to}, msg)
}
