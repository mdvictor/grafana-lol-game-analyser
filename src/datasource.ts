import { DataSourceInstanceSettings } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { LolCloudDataSourceOptions, LolQuery, MatchParticipantInfo } from 'types';

export class LolDataSource extends DataSourceWithBackend<LolQuery, LolCloudDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<LolCloudDataSourceOptions>) {
    super(instanceSettings);
  }

  fetchMatchIds(type: string, no: number): Promise<string[]> {
    return this.getResource('match/ids', { type, no });
  }

  fetchMatchSelfInfo(matchId: string): Promise<MatchParticipantInfo> {
    return this.getResource('match/self-info', { matchId });
  }

  fetchMatchParticipants(matchId: string): Promise<MatchParticipantInfo[]> {
    return this.getResource('match/participants', { matchId });
  }
}
