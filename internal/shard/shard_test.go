package shard

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormatShardEntry(t *testing.T) {
	tests := []struct {
		name     string
		shard    ShardEntry
		expected string
	}{
		{
			name: "Test Shard 1",
			shard: ShardEntry{
				name:                "Shard 1",
				description:         "Shard 1 Description",
				shardId:             1,
				loginServerIp:       "10.10.5.20",
				loginServerPort:     8226,
				lobbyServerIp:       "10.10.5.20",
				lobbyServerPort:     8228,
				mcotsServerIp:       "10.10.5.20",
				statusId:            0,
				statusReason:        "",
				serverGroupName:     "Group1",
				population:          100,
				maxPersonasPerUser:  5,
				dianosticServerHost: "10.10.5.20",
				dianosticServerPort: 3000,
			},
			expected: `[Shard 1]
Description=Shard 1 Description
ShardId=1
LoginServerIP=10.10.5.20
LoginServerPort=8226
LobbyServerIP=10.10.5.20
LobbyServerPort=8228
MCOTSServerIP=10.10.5.20
StatusId=0
Status_Reason=
ServerGroup_Name=Group1
Population=100
MaxPersonasPerUser=5
DianosticServerHost=10.10.5.20
DianosticServerPort=3000`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shard.formatShardEntry()
			if result != tt.expected {
				t.Errorf("formatShardEntry() = %v, want %v", result, tt.expected)
			}
		})
	}
}


func TestHandleShardList(t *testing.T) {
	req, err := http.NewRequest("GET", "/shardlist", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleShardList)

	handler.ServeHTTP(rr, req)

	expected := `[Shard 1]
Description=Shard 1 Description
ShardId=1
LoginServerIP=10.10.5.20
LoginServerPort=8226
LobbyServerIP=10.10.5.20
LobbyServerPort=8228
MCOTSServerIP=10.10.5.20
StatusId=0
Status_Reason=
ServerGroup_Name=Group1
Population=100
MaxPersonasPerUser=5
DianosticServerHost=10.10.5.20
DianosticServerPort=3000`

	if rr.Body.String() != expected {
		t.Errorf("HandleShardList() = %v, want %v", rr.Body.String(), expected)
	}
}

