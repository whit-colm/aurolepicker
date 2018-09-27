package rolepicker

import (
	"fmt"
	"github.com/aurumbot/flags"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
)

func ModRoles(session *dsg.Session, message *dsg.Message) {
	guild, err := f.GetGuild(session, message)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(session, message, err)
		return
	}

	var errors []string
	flgs := flags.Parse(message.Content)
	if len(flgs) <= 1 {
		session.ChannelMessageSend(message.ChannelID, "Please provide a valid flag.")
		return
	}

	for i := range flgs {
		if i == 0 {
			continue
		}
		switch flgs[i].Name {
		case "-a", "--add":
			var rolemap = make(map[string]string)
			for _, role := range strings.Split(flgs[i].Value, ",") {
				r := strings.Split(role, "=")
				if len(r) == 2 {
					rolemap[strings.TrimSpace(r[0])] = strings.TrimSpace(r[1])
				} else {
					errors = append(errors, fmt.Sprintf("Expected a name and ID, seperated by an =. Recieved `%v` instead.", role))
				}
			}
			addRole(guild.ID, rolemap)
		case "-r", "-rm", "--remove":
			delRole(guild.ID, strings.Split(flgs[i].Value, ","))
		default:
			errors = append(errors, fmt.Sprintf("Flag %v (value \"%v\") is not valid", flgs[i].Name, flgs[i].Value))
		}
	}
	save()
	report := "**Done!**\n"
	if len(errors) > 0 {
		report += "Errors were encountered while doing this, they are listed here: "
		for i := range errors {
			report += fmt.Sprintf("%v**,** ", errors[i])
		}
		report += "\nAll other operations were performed successfully."
	}
	session.ChannelMessageSend(message.ChannelID, report)
}
