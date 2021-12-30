import React, { useState, useEffect } from 'react';
import { SelectableValue } from '@grafana/data';
import { LolSelect } from './LolSelect';
import { LolDataSource } from 'datasource';

interface PlayerSelectProps {
  label: string;
  onChange: (item: SelectableValue<string>) => void;
  value: string;
  datasource: LolDataSource;
  matchId: string;
}

const TEAM1 = 100;

export function PlayerSelect(props: PlayerSelectProps) {
  const { value, onChange, label, datasource, matchId } = props;
  const [options, setOptions] = useState<Array<SelectableValue<string>>>([]);

  useEffect(() => {
    datasource.fetchMatchParticipants(matchId).then((participants) => {
      if (participants === null) {
        return;
      }

      const opts = participants.map((participant) => {
        return {
          label:
            participant.championName +
            ' - ' +
            participant.individualPosition +
            ' - KDA: ' +
            participant.kills +
            '/' +
            participant.deaths +
            '/' +
            participant.assists +
            ' - Team ' +
            (participant.teamId === TEAM1 ? '1' : '2'),
          value: participant.puuid,
        };
      });

      setOptions(opts);
    });
  }, [datasource, matchId]);

  return <LolSelect value={value} onChange={onChange} label={label} options={options} width={50} />;
}
