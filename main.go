package natsuki

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func Run(token string) *discordgo.Session {
	client, err := discordgo.New(token)

	HandleError(err)

	HandleError(client.Open())

	return client
}

func HandleError(err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured: %v\n", err)
		os.Exit(1)
		return err
	}
	return err
}
