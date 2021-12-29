import { DataSourcePluginOptionsEditorProps, SelectableValue } from '@grafana/data';
import { Input, Select, LegacyForms } from '@grafana/ui';
import React, { ChangeEvent } from 'react';
import { LolCloudDataSourceOptions, LolCloudSecureJsonData, RoutingKeys } from 'types';

const { SecretFormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<LolCloudDataSourceOptions, LolCloudSecureJsonData> {}

export const ConfigEditor = (props: Props) => {
  const platformValues = Object.values(RoutingKeys).map((val) => {
    return { label: val, value: val };
  });

  const onAPITokenChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = props;
    onOptionsChange({
      ...options,
      secureJsonData: {
        apiToken: event.target.value,
      },
    });
  };

  const onResetAPIToken = () => {
    const { onOptionsChange, options } = props;
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        apiToken: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        apiToken: '',
      },
    });
  };

  const onPlatformChange = (selected: SelectableValue<string>) => {
    const { onOptionsChange, options } = props;
    onOptionsChange({
      ...options,
      jsonData: {
        ...options.jsonData,
        platform: selected.value!,
      },
    });
  };

  const onSummonerNameChange = (event: React.FormEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = props;
    onOptionsChange({
      ...options,
      jsonData: {
        ...options.jsonData,
        summonerName: event.currentTarget.value,
      },
    });
  };

  const { secureJsonFields, secureJsonData = {}, jsonData } = props.options;

  return (
    <>
      <div className="gf-form-group">
        <div className="gf-form">
          <SecretFormField
            isConfigured={secureJsonFields && secureJsonFields.apiToken}
            value={secureJsonData.apiToken || ''}
            label="API Token"
            placeholder="Your Riot API Token"
            labelWidth={6}
            inputWidth={20}
            onReset={onResetAPIToken}
            onChange={onAPITokenChange}
          />
        </div>
        <div className="gf-form">
          <label className="gf-form-label width-9">Riot Summoner name</label>
          <div className="gf-form-select-wrapper max-width-17">
            <Input
              css={undefined}
              onChange={onSummonerNameChange}
              value={jsonData.summonerName}
              placeholder="Your Summoner name"
            />
          </div>
        </div>
        <div className="gf-form">
          <label className="gf-form-label width-6">Platform</label>
          <div className="gf-form-select-wrapper max-width-20">
            <Select options={platformValues} value={jsonData.platform} onChange={onPlatformChange} />
          </div>
        </div>
      </div>
    </>
  );
};
