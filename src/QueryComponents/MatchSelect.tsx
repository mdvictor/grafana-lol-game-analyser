import React, { useState, useEffect } from 'react';
import { SelectableValue } from '@grafana/data';
import { LolDataSource } from 'datasource';
import { LolSelect } from './LolSelect';

interface MatchSelectProps {
  label: string;
  onChange: (item: SelectableValue<string>) => void;
  value: string;
  matchType: string;
  datasource: LolDataSource;
}

const NO_OF_MATCHES = 10;

export function MatchSelect(props: MatchSelectProps) {
  const { datasource, value, onChange, label, matchType } = props;
  const [options, setOptions] = useState<Array<SelectableValue<string>>>([]);

  useEffect(() => {
    datasource.fetchMatchIds(matchType, NO_OF_MATCHES).then((matches) => {
      if (matches === null) {
        return;
      }

      const promises = matches.map((matchId) => datasource.fetchMatchSelfInfo(matchId));

      Promise.all(promises).then((MatchParticipantsInfo) => {
        setOptions(
          MatchParticipantsInfo.map((matchParticipantInfo: any) => {
            return {
              label:
                matchParticipantInfo.championName +
                ' - ' +
                matchParticipantInfo.individualPosition +
                ' - KDA: ' +
                matchParticipantInfo.kills +
                '/' +
                matchParticipantInfo.deaths +
                '/' +
                matchParticipantInfo.assists,
              value: matchParticipantInfo.matchId,
            };
          })
        );
      });
    });
  }, [datasource, matchType]);

  return <LolSelect label={label} value={value} onChange={onChange} options={options} width={50} />;
}
