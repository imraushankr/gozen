// package email

// import (
// 	"fmt"
// 	"gozen/src/configs"
// 	"net/smtp"
// 	"strings"
// )

// type Email struct {
// 	From    string
// 	To      []string
// 	Subject string
// 	Body    string
// }

// type Service struct {
// 	emailConfig configs.EmailConfig
// }

// func NewService(cfg *configs.EmailConfig) *Service {
// 	return &Service{
// 		emailConfig: *cfg,
// 	}
// }

// func (s *Service) Send(email Email) error {
// 	auth := smtp.PlainAuth("", s.emailConfig.SMTP.Username, s.emailConfig.SMTP.Password, s.emailConfig.SMTP.Host)

// 	msg := fmt.Sprintf(
// 		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
// 		email.From,
// 		strings.Join(email.To, ","),
// 		email.Subject,
// 		email.Body,
// 	)

// 	return smtp.SendMail(
// 		fmt.Sprintf("%s:%d", s.emailConfig.SMTP.Host, s.emailConfig.SMTP.Port),
// 		auth,
// 		email.From,
// 		email.To,
// 		[]byte(msg),
// 	)
// }

package email

import (
	"bytes"
	"fmt"
	"gozen/src/configs"
	"net/mail"
	"net/smtp"
	"strings"
	"text/template"
)

type Email struct {
	From      string
	To        []string
	Subject   string
	Body      string
	HTMLBody  string // Added HTML support
	Headers   map[string]string // Additional headers
}

type Service struct {
	config     configs.SMTPConfig
	auth       smtp.Auth
	tlsEnabled bool
}

func NewService(cfg *configs.EmailConfig) *Service {
	return &Service{
		config:     cfg.SMTP,
		auth:       smtp.PlainAuth("", cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host),
		tlsEnabled: cfg.SMTP.UseTLS,
	}
}

func (s *Service) Send(email Email) error {
	// Validate email addresses
	if _, err := mail.ParseAddress(email.From); err != nil {
		return fmt.Errorf("invalid from address: %w", err)
	}

	for _, to := range email.To {
		if _, err := mail.ParseAddress(to); err != nil {
			return fmt.Errorf("invalid recipient address %s: %w", to, err)
		}
	}

	// Build message
	msg, err := s.buildMessage(email)
	if err != nil {
		return fmt.Errorf("failed to build email message: %w", err)
	}

	// Send email
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	if err := smtp.SendMail(addr, s.auth, email.From, email.To, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (s *Service) buildMessage(email Email) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	// Standard headers
	buf.WriteString(fmt.Sprintf("From: %s\r\n", email.From))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ",")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", email.Subject))

	// Additional headers
	for k, v := range email.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	// Determine content type
	switch {
	case email.HTMLBody != "" && email.Body != "":
		buf.WriteString("Content-Type: multipart/alternative; boundary=boundary123\r\n\r\n")
		buf.WriteString("--boundary123\r\n")
		buf.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
		buf.WriteString(email.Body)
		buf.WriteString("\r\n--boundary123\r\n")
		buf.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
		buf.WriteString(email.HTMLBody)
		buf.WriteString("\r\n--boundary123--\r\n")
	case email.HTMLBody != "":
		buf.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
		buf.WriteString(email.HTMLBody)
	default:
		buf.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
		buf.WriteString(email.Body)
	}

	return buf.Bytes(), nil
}

func (s *Service) SendTemplate(email Email, templatePath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	email.Body = body.String()
	return s.Send(email)
}