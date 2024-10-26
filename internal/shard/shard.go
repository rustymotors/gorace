package shard

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/rustymotors/gorace/internal/helpers"
)

type ShardEntry struct {
	name                string
	description         string
	shardId             int
	loginServerIp       string
	loginServerPort     int
	lobbyServerIp       string
	lobbyServerPort     int
	mcotsServerIp       string
	statusId            int
	statusReason        string
	serverGroupName     string
	population          int
	maxPersonasPerUser  int
	dianosticServerHost string
	dianosticServerPort int
}

func (s *ShardEntry) NewShardEntry(
	name string,
	description string,
	id int,
	loginServerIp string,
	loginServerPort int,
	lobbyServerIp string,
	lobbyServerPort int,
	mcotsServerIp string,
	statusId int,
	statusReason string,
	sreverGroupName string,
	population int,
	maxPersonasPerUser int,
	dianosticServerHost string,
	dianosticServerPort int,
) *ShardEntry {
	return &ShardEntry{
		name:                name,
		description:         description,
		shardId:             id,
		loginServerIp:       loginServerIp,
		loginServerPort:     loginServerPort,
		lobbyServerIp:       lobbyServerIp,
		lobbyServerPort:     lobbyServerPort,
		mcotsServerIp:       mcotsServerIp,
		statusId:            statusId,
		statusReason:        statusReason,
		serverGroupName:     sreverGroupName,
		population:          population,
		maxPersonasPerUser:  maxPersonasPerUser,
		dianosticServerHost: dianosticServerHost,
		dianosticServerPort: dianosticServerPort,
	}
}

func (s *ShardEntry) formatShardEntry() string {
	return "[" + s.name + "]\n" +
		"Description=" + s.description + "\n" + 
		"ShardId=" + strconv.Itoa(s.shardId) + "\n" + 
		"LoginServerIP=" + s.loginServerIp + "\n" + 
		"LoginServerPort=" + strconv.Itoa(s.loginServerPort) + "\n" + 
		"LobbyServerIP=" + s.lobbyServerIp + "\n" + 
		"LobbyServerPort=" + strconv.Itoa(s.lobbyServerPort) + "\n" + 
		"MCOTSServerIP=" + s.mcotsServerIp + "\n" + 
		"StatusId=" + strconv.Itoa(s.statusId) + "\n" + 
		"Status_Reason=" + s.statusReason + "\n" + 
		"ServerGroup_Name=" + s.serverGroupName + "\n" + 
		"Population=" + strconv.Itoa(s.population) + "\n" + 
		"MaxPersonasPerUser=" + strconv.Itoa(s.maxPersonasPerUser) + "\n" + 
		"DianosticServerHost=" + s.dianosticServerHost + "\n" + 
		"DianosticServerPort=" + strconv.Itoa(s.dianosticServerPort)
}

func HandleShardList(w http.ResponseWriter, r *http.Request) {
	shardIp := "10.10.5.20"

	shard1 := ShardEntry{}
	shard1.name = "Shard 1"
	shard1.description = "Shard 1 Description"
	shard1.shardId = 1
	shard1.loginServerIp = shardIp
	shard1.loginServerPort = 8226
	shard1.lobbyServerIp = shardIp
	shard1.lobbyServerPort = 8228
	shard1.mcotsServerIp = shardIp
	shard1.statusId = 0
	shard1.statusReason = ""
	shard1.serverGroupName = "Group1"
	shard1.population = 100
	shard1.maxPersonasPerUser = 5
	shard1.dianosticServerHost = shardIp
	shard1.dianosticServerPort = 3000

	fmt.Println("Sending shard list")
	helpers.WriteResponse(w, shard1.formatShardEntry())
}
