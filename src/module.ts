import { DataSourcePlugin } from '@grafana/data';
import { ConfigEditor } from 'ConfigEditor';
import { LolDataSource } from 'datasource';
import { QueryEditor } from 'QueryEditor';
import { LolCloudDataSourceOptions, LolQuery } from 'types';

export const plugin = new DataSourcePlugin<LolDataSource, LolQuery, LolCloudDataSourceOptions>(LolDataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
