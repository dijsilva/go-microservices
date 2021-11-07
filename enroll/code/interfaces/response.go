/*
 * @File: models.message.go
 * @Description: Defines Message information will be returned to the clients
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package interfaces

// Message defines the response message
type Response struct {
	Data Message `json:"data"`
}

type Message struct {
	Message string `json:"message"`
}
