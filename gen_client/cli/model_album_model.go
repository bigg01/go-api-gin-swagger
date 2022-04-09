// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"go-api-gin-swagger/client/models"

	"github.com/spf13/cobra"
)

// Schema cli for ModelAlbum

// register flags to command
func registerModelModelAlbumFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModelAlbumArtist(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModelAlbumID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModelAlbumPrice(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModelAlbumTitle(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModelAlbumArtist(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	artistDescription := ``

	var artistFlagName string
	if cmdPrefix == "" {
		artistFlagName = "artist"
	} else {
		artistFlagName = fmt.Sprintf("%v.artist", cmdPrefix)
	}

	var artistFlagDefault string

	_ = cmd.PersistentFlags().String(artistFlagName, artistFlagDefault, artistDescription)

	return nil
}

func registerModelAlbumID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerModelAlbumPrice(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	priceDescription := ``

	var priceFlagName string
	if cmdPrefix == "" {
		priceFlagName = "price"
	} else {
		priceFlagName = fmt.Sprintf("%v.price", cmdPrefix)
	}

	var priceFlagDefault float64

	_ = cmd.PersistentFlags().Float64(priceFlagName, priceFlagDefault, priceDescription)

	return nil
}

func registerModelAlbumTitle(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	titleDescription := ``

	var titleFlagName string
	if cmdPrefix == "" {
		titleFlagName = "title"
	} else {
		titleFlagName = fmt.Sprintf("%v.title", cmdPrefix)
	}

	var titleFlagDefault string

	_ = cmd.PersistentFlags().String(titleFlagName, titleFlagDefault, titleDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModelAlbumFlags(depth int, m *models.ModelAlbum, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, artistAdded := retrieveModelAlbumArtistFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || artistAdded

	err, idAdded := retrieveModelAlbumIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, priceAdded := retrieveModelAlbumPriceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || priceAdded

	err, titleAdded := retrieveModelAlbumTitleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || titleAdded

	return nil, retAdded
}

func retrieveModelAlbumArtistFlags(depth int, m *models.ModelAlbum, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	artistFlagName := fmt.Sprintf("%v.artist", cmdPrefix)
	if cmd.Flags().Changed(artistFlagName) {

		var artistFlagName string
		if cmdPrefix == "" {
			artistFlagName = "artist"
		} else {
			artistFlagName = fmt.Sprintf("%v.artist", cmdPrefix)
		}

		artistFlagValue, err := cmd.Flags().GetString(artistFlagName)
		if err != nil {
			return err, false
		}
		m.Artist = artistFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModelAlbumIDFlags(depth int, m *models.ModelAlbum, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModelAlbumPriceFlags(depth int, m *models.ModelAlbum, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	priceFlagName := fmt.Sprintf("%v.price", cmdPrefix)
	if cmd.Flags().Changed(priceFlagName) {

		var priceFlagName string
		if cmdPrefix == "" {
			priceFlagName = "price"
		} else {
			priceFlagName = fmt.Sprintf("%v.price", cmdPrefix)
		}

		priceFlagValue, err := cmd.Flags().GetFloat64(priceFlagName)
		if err != nil {
			return err, false
		}
		m.Price = priceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModelAlbumTitleFlags(depth int, m *models.ModelAlbum, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	titleFlagName := fmt.Sprintf("%v.title", cmdPrefix)
	if cmd.Flags().Changed(titleFlagName) {

		var titleFlagName string
		if cmdPrefix == "" {
			titleFlagName = "title"
		} else {
			titleFlagName = fmt.Sprintf("%v.title", cmdPrefix)
		}

		titleFlagValue, err := cmd.Flags().GetString(titleFlagName)
		if err != nil {
			return err, false
		}
		m.Title = titleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
