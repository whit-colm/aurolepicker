package rolepicker

import (
	"fmt"
	"github.com/aurumbot/flags"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"sort"
	"strings"
)

/*
func ModRoles(session *dsg.Session, message *dsg.Message) {
	flgs := flags.Parse(message.Content)
	if len(flgs) <= 2 {
		session.ChannelMessageSend(m.ChannelID, "Please provide a valid flag.")
		return
	}

	for i := range flgs {
		if i == 0 {
			continue
		}
		switch flgs[i].Name {
		case "-a", "--add":
			var rolemap = make(map[string]string)
			for role := range strings.Split(flgs[i].Value, ",") {
				r := strings.Split(role, "=")
				if len(r) == 2 {
					rolemap[strings.TrimSpace(r[0])] = strings.TrimSpace(r[1])
				} else {
					// Tell actor they did it wrong
				}
			}
			addRole(rolemap)
		case "-r", "-rm", "--remove":
			var rolemap = make(map[string]string)
			for role := range strings.Split(flgs[i].Value, ",") {
				delRole(strings.TrimSpace(role))
			}
			//default:
			//etcIssue = append(etcIssue, fmt.Sprintf("Flag %v (value \"%v\") is not valid", flgs[i].Name, flgs[i].Value))
		}
	}
	save()
	report := "**Done!**\n"
	s.ChannelMessageSend(message.ChannelID, report)
}*/
