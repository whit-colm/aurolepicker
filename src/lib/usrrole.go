package rolepicker

import (
	"fmt"
	"github.com/aurumbot/flags"
	//	"github.com/aurumbot/lib/dat"
	//	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
)

func UsrRoles(session *dsg.Session, message *dsg.Message) {
	var joinSucc []string
	var joinFail []string
	var quitSucc []string
	var quitFail []string
	var etcIssue []string
	var listed = false

	flgs := flags.Parse(message.Content)
	if len(flgs) <= 2 {
		session.ChannelMessageSend(message.ChannelID, "Please provide a valid flag.")
		return
	}

	for i := range flgs {
		if i == 0 {
			continue
		}
		switch flgs[i].Name {
		case "-j", "--join", "-a", "--add":
			for _, role := range strings.Split(flgs[i].Value, ",") {
				ok := joinRole(session, message, strings.TrimSpace(role))
				if ok {
					joinSucc = append(joinSucc, role)
				} else {
					joinFail = append(joinFail, role)
				}
			}
		case "-q", "--quit", "--leave":
			for _, role := range strings.Split(flgs[i].Value, ",") {
				ok := quitRole(session, message, strings.TrimSpace(role))
				if ok {
					quitSucc = append(quitSucc, role)
				} else {
					quitFail = append(quitFail, role)
				}
			}
		case "-l", "-ls", "--list":
			if !listed {
				session.ChannelMessageSend(message.ChannelID, listRoles())
			}
			listed = true
		default:
			etcIssue = append(etcIssue, fmt.Sprintf("Flag %v (value \"%v\") is not valid", flgs[i].Name, flgs[i].Value))
		}
	}

	report := "**Done**\n"

	session.ChannelMessageSend(message.ChannelID, report)
}
