import { DataSourceInstanceSettings } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { LolCloudDataSourceOptions, LolQuery, MatchInfo, MatchParticipant } from 'types';

export class LolDataSource extends DataSourceWithBackend<LolQuery, LolCloudDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<LolCloudDataSourceOptions>) {
    super(instanceSettings);
  }

  fetchMatchIds(type: string, no: number): Promise<string[]> {
    return this.getResource('match/ids', { type, no });
  }

  fetchMatchInfo(matchId: string): Promise<MatchInfo> {
    return this.getResource('match/info', { matchId });
  }

  fetchMatchParticipants(matchId: string): Promise<MatchParticipant[]> {
    return this.getResource('match/participants', { matchId });
  }
}
