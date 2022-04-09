// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"go-api-gin-swagger/client/client/album"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAlbumGetAlbAlbumsCmd returns a cmd to handle operation getAlbAlbums
func makeOperationAlbumGetAlbAlbumsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetAlbAlbums",
		Short: `return a list of Albums`,
		RunE:  runOperationAlbumGetAlbAlbums,
	}

	if err := registerOperationAlbumGetAlbAlbumsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAlbumGetAlbAlbums uses cmd flags to call endpoint api
func runOperationAlbumGetAlbAlbums(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := album.NewGetAlbAlbumsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAlbumGetAlbAlbumsResult(appCli.Album.GetAlbAlbums(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAlbumGetAlbAlbumsParamFlags registers all flags needed to fill params
func registerOperationAlbumGetAlbAlbumsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationAlbumGetAlbAlbumsResult parses request result and return the string content
func parseOperationAlbumGetAlbAlbumsResult(resp0 *album.GetAlbAlbumsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*album.GetAlbAlbumsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
