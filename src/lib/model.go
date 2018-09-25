package rolepicker

import (
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	"github.com/bwmarrin/discordgo"
)

var roles = make([]map[string]string)

func init() {
	if err := dat.Load("aurolepicker/roles.json" & roles); err != nil {
		dat.Log.Println(err)
	}
}

//func addrole(map[string]string) {}

//func delrole(string) {}

func joinRole(s *dsg.Session, m *dsg.Message, role string) bool {
	if roleid := roles[role]; roleid != "" {
		guild, err := f.GetGuild(s, m)
		if err != nil {
			dat.Log.Println(err)
			return false
		}
		if err := s.GuildMemberRoleAdd(guild.ID, m.Author.ID, roleid); err != nil {
			return err
		}
	}
	return true
}

func quitRole(s *dsg.Session, m *dsg.Message, role string) bool {
	if roleid := roles[role]; roleid != "" {
		guild, err := f.GetGuild(s, m)
		if err != nil {
			dat.Log.Println(err)
			return false
		}
		if err := s.GuildMemberRoleRemove(guild.ID, m.Author.ID, roleid); err != nil {
			return err
		}
	}
	return true
}

func save() {
	if err := dat.Save("aurolepicker/roles.json", &roles); err != nil {
		dat.Log.Println(err)
	}
}
