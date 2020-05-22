package crocgodyl

import "encoding/json"

// --------------------------------------------------------------
// Client API

// Client Server API

// ClientServers is the default all servers view for the client API.
// GET this from the '/api/client' endpoint
type ClientServers struct {
	Object       string         `json:"object"`
	ClientServer []ClientServer `json:"data"`
	Meta         struct {
		Pagination struct {
			Total       int           `json:"total"`
			Count       int           `json:"count"`
			PerPage     int           `json:"per_page"`
			CurrentPage int           `json:"current_page"`
			TotalPages  int           `json:"total_pages"`
			Links       []interface{} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

// ClientServer is the server object view returning single server information.
// GET this from the '/api/client/servers/<server_ID>' endpoint
type ClientServer struct {
	Object     string `json:"object"`
	Attributes struct {
		ServerOwner bool   `json:"server_owner"`
		Identifier  string `json:"identifier"`
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Limits      struct {
			Memory int `json:"memory"`
			Swap   int `json:"swap"`
			Disk   int `json:"disk"`
			Io     int `json:"io"`
			CPU    int `json:"cpu"`
		} `json:"limits"`
		FeatureLimits struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations"`
		} `json:"feature_limits"`
	} `json:"attributes"`
}

// ClientServerUtilization is the server statistics reported by the daemon.
// GET this from the '/api/client/servers/<server_ID>/utilization' endpoint
type ClientServerUtilization struct {
	Object     string `json:"object"`
	Attributes struct {
		State  string `json:"state"`
		Memory struct {
			Current int `json:"current"`
			Limit   int `json:"limit"`
		} `json:"memory"`
		CPU struct {
			Current float64   `json:"current"`
			Cores   []float64 `json:"cores"`
			Limit   int       `json:"limit"`
		} `json:"cpu"`
		Disk struct {
			Current int `json:"current"`
			Limit   int `json:"limit"`
		} `json:"disk"`
	} `json:"attributes"`
}

// ClientServerConsoleCommand is the struct for sending a command for the server console
// GET this from the '/api/client/servers/<server_ID>/command' endpoint
type ClientServerConsoleCommand struct {
	Command string `json:"command"`
}

// ClientServerPowerAction is the struct for sending a power command for the server
// GET this from the '/api/client/servers/<server_ID>/power' endpoint
type ClientServerPowerAction struct {
	Signal string `json:"signal"`
}

//ExecuteCommand executes a command
//It requires a serverID as an int, a command as a string and a config
func (config *CrocConfig) ExecuteCommand(serverID string, command string) error {
	esbytes, err := json.Marshal(&ClientServerConsoleCommand{Command: command})
	if err != nil {
		return err
	}
	_, err = config.queryClientAPI("servers/"+serverID+"/command", "post", esbytes)
	if err != nil {
		return err
	}

	return nil
}
