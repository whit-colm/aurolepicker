package rolepicker

import (
	"fmt"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"sort"
	"strings"
)

var roles = make(map[string]map[string]string)

func init() {
	if err := dat.Load("aurolepicker/roles.json", &roles); err != nil {
		dat.Log.Println(err)
	}
}

func joinRole(s *dsg.Session, m *dsg.Message, role string) bool {
	guild, err := f.GetGuild(s, m)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(s, m, err)
		return false
	}
	if roleid := roles[guild.ID][role]; roleid != "" {
		if err := s.GuildMemberRoleAdd(guild.ID, m.Author.ID, roleid); err != nil {
			dat.Log.Println(err)
			return false
		}
	}
	return true
}

func quitRole(s *dsg.Session, m *dsg.Message, role string) bool {
	guild, err := f.GetGuild(s, m)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(s, m, err)
		return false
	}
	if roleid := roles[guild.ID][role]; roleid != "" {
		guild, err := f.GetGuild(s, m)
		if err != nil {
			dat.Log.Println(err)
			return false
		}
		if err := s.GuildMemberRoleRemove(guild.ID, m.Author.ID, roleid); err != nil {
			dat.Log.Println(err)
			return false
		}
	}
	return true
}

func listRoles(guildID string) string {
	var rlist []string
	for k, _ := range roles[guildID] {
		rlist = append(rlist, fmt.Sprintf("- %v\n", k))
	}
	msg := "Available roles:\n"
	sort.Strings(rlist)
	for i := range rlist {
		msg += rlist[i]
	}
	return msg
}

func addRole(guildID string, rolemap map[string]string) {
	if roles[guildID] == nil {
		roles[guildID] = make(map[string]string)
	}
	for k, v := range rolemap {
		roles[guildID][k] = v
	}
}

func delRole(guildID string, rolemap []string) {
	for i := range rolemap {
		delete(roles[guildID], strings.TrimSpace(rolemap[i]))
	}
}

func save() {
	if err := dat.Save("aurolepicker/roles.json", &roles); err != nil {
		dat.Log.Println(err)
	}
}
