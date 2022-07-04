package emails

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	mail "github.com/xhit/go-simple-mail/v2"
	"time"
)

var (
	AllStates = []string{"Delivered", "Canceled", "On-The-Way"}

	// email properties
	EmailHTMLBody = ``
)


// Creates Default SMTP Client...
func createSMTPClient() (*mail.SMTPClient, error){
	// creates SMTP Client for managing emails.
	port, error :=  strconv.Atoi(os.Getenv("SMTP_PORT"))
	if error != nil {
		return nil, error
	}

	client := mail.NewSMTPClient()
	client.Encryption = mail.EncryptionSTARTTLS
	client.Username = os.Getenv("SMTP_EMAIL")
	client.Password = os.Getenv("SMTP_PASSWORD")
	client.Port = port
	client.Host = os.Getenv("SMTP_HOST")
	client.ConnectTimeout = 10 * time.Second
	client.SendTimeout = 10 * time.Second
	smtpClient, error := client.Connect()
	if error == nil && smtpClient != nil {
		return smtpClient, nil
	}
	return nil, errors.New("Failed to create SMTP Client.")
}



// func RenderEmailOrderCheckoutImage(OrderContent map[string]string) (string, error){
// 	basePattern := ``
// 	return "", nil 
// }


// Sends Email Notification using mail golang SDK 
func sendEmailNotification(message string, customerEmail string) (bool, error){
	// some logic of sending email...
	client, error := createSMTPClient()
	if error != nil{
		return false, nil
	}
	EmailMessage := mail.NewMSG()
	EmailMessage.AddTo(customerEmail).SetSubject(message)
	EmailMessage.SetBody(mail.TextHTML, EmailHTMLBody)
	sended_error := EmailMessage.Send(client)
	if sended_error != nil {
		return false, sended_error
	}
	return true, nil
}



// Base Method for managing Email Notifications.
func NotifyEmailOrder(customerEmail string, message string) (bool, error){

	response, error_ := sendEmailNotification(message, customerEmail)
	if response != false && error_ == nil {
		return true, nil
	}
	return false, errors.New("Failure to send Email.")
}



// Method Is used for sending Email Notification to the customer Email, that the order has been rejected.
// Prepares the message and calls `NotifyOrder` method that sends email.
func NotifyEmailOrderRejected(customerEmail string, orderId string) (bool, error){
	message := fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
    "has been rejected. Check you profile for more details.",
	customerEmail, orderId)
	sended, error := NotifyEmailOrder(customerEmail, message)
	if sended != true || error != nil {return false, errors.New(
	"Failed To Send Reject Email Notification.")} else {return true, nil}
}


// Method Is used for sending Email Notification to the customer Email, that the order has been Accepted.
func NotifyEmailOrderAccepted(customerEmail string, orderId string) (bool, error){
	message := fmt.Sprintf("Hello, %s, Unfortunately we need to notify you, that your Order `%s`" +
    "has been Accepted. Check you profile for more details.",
	customerEmail, orderId)
	sended, error := NotifyEmailOrder(customerEmail, message)
	if sended != true || error != nil {return false, errors.New(
	"Failed To Send Accept Email Notification.")} else {return true, nil}
}



