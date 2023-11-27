package groups

import (
	"strconv"
	"strings"
)

type Group struct {
	Priority                  int      `json:"priority"`
	IsProxy                   bool     `json:"isProxy"`
	HiddenAtProxyGroups       []string `json:"hiddenAtProxyGroups"`
	Name                      string   `json:"name"`
	TemplateName              string   `json:"templateName"`
	MaxMemory                 int      `json:"maxMemory"`
	MaxCpuUsage               int      `json:"maxCpuUsage"`
	MaxDiskSpace              int      `json:"maxDiskSpace"`
	MaxIOUsage                int      `json:"maxIOUsage"`
	MaxProcessCount           int      `json:"maxProcessCount"`
	MaxPlayers                int      `json:"maxPlayers"`
	MinimumOnlineServiceCount int      `json:"minimumOnlineServiceCount"`
	MaximumOnlineServiceCount int      `json:"maximumOnlineServiceCount"`
	Maintenance               bool     `json:"maintenance"`
	Static                    bool     `json:"static"`
	PercentToStartNewService  int      `json:"percentToStartNewService"`
	StartPriority             int      `json:"startPriority"`
	PermissionToAccess        string   `json:"permissions"`
	StateUpdating             bool     `json:"stateUpdating"`
	ServiceVersion            string   `json:"serviceVersion"`
}

func (g *Group) CreateStartupCommand() []string {
	command := "java -Xms128M -Xmx" + strconv.Itoa(g.MaxMemory) + "M -Dterminal.jline=false -Dterminal.ansi=true --add-modules=jdk.incubator.vector -XX:+UseG1GC -XX:+ParallelRefProcEnabled -XX:MaxGCPauseMillis=200 -XX:+UnlockExperimentalVMOptions -XX:+DisableExplicitGC -XX:+AlwaysPreTouch -XX:G1NewSizePercent=40 -XX:G1MaxNewSizePercent=50 -XX:G1HeapRegionSize=16M -XX:G1ReservePercent=15 -XX:G1HeapWastePercent=5 -XX:G1MixedGCCountTarget=4 -XX:InitiatingHeapOccupancyPercent=20 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:SurvivorRatio=32 -XX:+PerfDisableSharedMem -XX:MaxTenuringThreshold=1 -Dusing.aikars.flags=https://mcflags.emc.gs -Daikars.new.flags=true -jar server.jar"
	return strings.Split(command, " ")
}
