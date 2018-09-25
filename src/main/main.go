package main

import (
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	"github.com/bwmarrin/discordgo"
	"github.com/whitman-colm/aurolepicker"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["csrole"] = &f.Command{
		Name: "Get a novelty role",
		Help: `This allows the actor to add or remove novelty roles to themself for 
things such as name colours, or gain access to semi-private channels.
Flags:
**<-j|--join> <role1, role2]** : join roles, multiple roles are supported seperated by commas.
**<-q|--quit> <role1, role2]** : quit roles, multiple roles are supported seperated by commas.
**<-l|--list>** : list avaialable roles.
Usage:
` + f.Config.Prefix + `csrole -j orange, memes -q grey`,
		Perms:   -1,
		Version: "2.0.0",
		action:  rolepicker.UsrRoles,
	}
	/*
		Commands["csrolemod"] = &f.Command{
			Name: "Add novelty roles to the role list.",
			Help: `This adds novelty roles to the list in *csrole*.
	Flags:
	**<-a|--add> <role1=roleID, role2=roleID]** : add role(s) to the list. Must be formatted by name=ID. Multiple entries are comma separeted.
	**<-r|--remove> <role1, role2]** : remove a role from the list. Multiple entries are comma separated.
	Usage:
	` + f.Config.Prefix + `csrolemod -a notification squad=470036470779084801`,
			Perms:   PermissionManageRoles,
			Version: "1.0.0",
			action:  rolepicker.ModRoles,
		}*/
}
