package project

import (
	"bytes"
	"encoding/json"

	adminSdk "github.com/heyframe/go-heyframe-admin-api-sdk"

	"github.com/heyframe/heyframe-cli/logging"
	"github.com/heyframe/heyframe-cli/platform"
)

type SystemConfigSync struct{}

func (SystemConfigSync) Push(ctx adminSdk.ApiContext, client *adminSdk.Client, config *platform.Config, operation *ConfigSyncOperation) error {
	if config.Sync == nil {
		return nil
	}

	c := adminSdk.Criteria{}
	c.Includes = map[string][]string{"sales_channel": {"id", "name"}}
	channelResponse, resp, err := client.Repository.Channel.SearchAll(ctx, c)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logging.FromContext(ctx.Context).Errorf("SystemConfigSync/Push: %v", err)
		}
	}()

	for _, config := range config.Sync.Config {
		if config.Channel != nil && len(*config.Channel) != 32 {
			foundId := false

			for _, scRow := range channelResponse.Data {
				scRow := scRow
				if *config.Channel == scRow.Name {
					config.Channel = &scRow.Id

					foundId = true
				}
			}

			if !foundId {
				logging.FromContext(ctx.Context).Errorf("Cannot find sales channel id for %s", *config.Channel)
				continue
			}
		}

		currentConfig, err := readSystemConfig(ctx, client, config.Channel)
		if err != nil {
			return err
		}

		for newK, newV := range config.Settings {
			_, ok := operation.SystemSettings[config.Channel]

			if !ok {
				operation.SystemSettings[config.Channel] = map[string]interface{}{}
			}

			foundKey := false

			for _, existingConfig := range currentConfig.Data {
				if existingConfig.ConfigurationKey == newK {
					foundKey = true

					encodedSource, _ := json.Marshal(existingConfig.ConfigurationValue)
					encodedTarget, _ := json.Marshal(newV)

					if !bytes.Equal(encodedSource, encodedTarget) {
						operation.SystemSettings[config.Channel][newK] = newV
					}

					break
				}
			}

			if !foundKey {
				operation.SystemSettings[config.Channel][newK] = newV
			}
		}
	}

	return nil
}

func (SystemConfigSync) Pull(ctx adminSdk.ApiContext, client *adminSdk.Client, config *platform.Config) error {
	config.Sync.Config = make([]platform.ConfigSyncConfig, 0)

	c := adminSdk.Criteria{}
	c.Includes = map[string][]string{"sales_channel": {"id", "name"}}
	channelResponse, resp, err := client.Repository.Channel.SearchAll(ctx, c)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logging.FromContext(ctx.Context).Errorf("SystemConfigSync/Pull: %v", err)
		}
	}()

	channelList := make([]adminSdk.Channel, 0)
	channelList = append(channelList, adminSdk.Channel{Id: ""})
	channelList = append(channelList, channelResponse.Data...)

	for _, sc := range channelList {
		var sysConfigs *adminSdk.EntityCollection[adminSdk.SystemConfig]
		var err error

		cfg := platform.ConfigSyncConfig{
			Settings: map[string]interface{}{},
		}

		if sc.Id == "" {
			sysConfigs, err = readSystemConfig(ctx, client, nil)
		} else {
			scName := sc.Name
			cfg.Channel = &scName

			sysConfigs, err = readSystemConfig(ctx, client, &sc.Id)
		}

		if err != nil {
			return err
		}

		for _, record := range sysConfigs.Data {
			// app system shopId
			if record.ConfigurationKey == "core.app.shopId" {
				continue
			}

			cfg.Settings[record.ConfigurationKey] = record.ConfigurationValue
		}

		config.Sync.Config = append(config.Sync.Config, cfg)
	}

	return nil
}
