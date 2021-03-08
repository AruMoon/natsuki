/*
Copyright 2021 Aru Moon

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package core

import (
	"fmt"
	"os"

	. "github.com/bwmarrin/discordgo"
)

func Run(token string) *Session {
	client, err := New(token)

	HandleError(err)

	HandleError(client.Open())

	return client
}

func HandleError(err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured: %v\n", err)
		os.Exit(1) // Comment this line to ignore errors
		return err
	}
	return err
}
