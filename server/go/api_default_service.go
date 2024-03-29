/*
 * Translator service between a srb2kart server and json
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"fmt"
	"net/http"

	"go.formulabun.club/functional/array"
	"go.formulabun.club/functional/strings"
	"go.formulabun.club/srb2kart/network"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	Target string
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(target string) DefaultApiServicer {
	return &DefaultApiService{target}
}

// MapsGet - get the installed mods
func (s *DefaultApiService) FilesGet(ctx context.Context) (ImplResponse, error) {
	files, err := network.TellAllFilesNeeded(s.Target)
	if err != nil {
		return Response(500, err), err
	}

	for i, f := range files {
		// oops I forgot null terminator in srb2kart/network
		files[i] = f[:len(f)-1]
	}
	return Response(http.StatusOK, files), nil
}

// PlayerinfoGet - get the player infomation
func (s *DefaultApiService) PlayerinfoGet(ctx context.Context) (ImplResponse, error) {
	_, playerInfo, err := network.AskInfo(s.Target)

	if err != nil {
		return Response(500, err), err
	}
  
  existingNodes := array.Filter(playerInfo, func(player network.PlayerInfo) bool {
    return player.Node != 255
  })
	players := array.Map(existingNodes, func(player network.PlayerInfo) PlayerInfoEntry {
		return PlayerInfoEntry{
			player.Node,
			strings.SafeNullTerminated(player.Name[:]),
			strings.Join(array.Map(player.Address[:], func(v uint8) string { return fmt.Sprint(v) }), "."),
			player.Team,
			player.Skin,
			player.Data,
			player.Score,
			player.TimeInServer,
		}
	})

	return Response(http.StatusOK, PlayerInfo(players)), nil
}

// ServerinfoGet - get the server information
func (s *DefaultApiService) ServerinfoGet(ctx context.Context) (ImplResponse, error) {
	i, _, err := network.AskInfo(s.Target)

	if err != nil {
		return Response(500, err), err
	}

  resp := ServerInfo{
    i.PacketVersion,
    strings.SafeNullTerminated(i.Application[:]),
    i.Version,
    i.Subversion,
    i.NumberOfPlayer,
    i.MaxPlayer,
    i.Gametype,
    i.ModifiedGame > 0,
    i.CheatsEnabled > 0,
    i.KartVars,
    i.FileNeededNum,
    i.Time,
    i.LevelTime,
    strings.SafeNullTerminated(i.ServerName[:]),
    strings.SafeNullTerminated(i.MapName[:]),
    strings.SafeNullTerminated(i.MapTitle[:]),
    i.MapMd5,
    i.ActNum,
    i.IsZone > 0,
    strings.SafeNullTerminated(i.HttpSource[:]),
    []FileNeededInner{},
  }

	return Response(http.StatusOK, resp), nil
}
