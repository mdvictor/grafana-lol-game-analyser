import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { LolDataSource } from 'datasource';
import { LolCloudDataSourceOptions, LolQuery } from 'types';
import React, { FormEvent } from 'react';
import { MatchSelect } from 'QueryComponents/MatchSelect';
import { MatchTypeSelect } from 'QueryComponents/MatchTypeSelect';
import { InlineFormLabel, InlineSwitch } from '@grafana/ui';
import { PlayerSelect } from 'QueryComponents/PlayerSelect';
import { TimelineDataSelect } from 'QueryComponents/TimelineDataSelect';

type Props = QueryEditorProps<LolDataSource, LolQuery, LolCloudDataSourceOptions>;

export const QueryEditor = (props: Props) => {
  const { datasource, query, onChange, onRunQuery } = props;

  const onMatchSelectChange = (item: SelectableValue<string>) => {
    onChange({
      ...query,
      matchId: item.value!,
    });
  };

  const onMatchTypeChange = (item: SelectableValue<string>) => {
    onChange({
      ...query,
      matchType: item.value!,
    });
  };

  const onSwitchChange = (element: FormEvent<HTMLInputElement>) => {
    onChange({
      ...query,
      useMatchTimerange: !query.useMatchTimerange,
    });

    if (query.player && query.timelineData) {
      onRunQuery();
    }
  };

  const onPlayerChange = (item: SelectableValue<string>) => {
    const championName = item.label!.split(' - ')[0]!;

    onChange({
      ...query,
      player: item.value!,
      championName: championName,
    });

    if (query.timelineData) {
      onRunQuery();
    }
  };

  const onTimelineDataChange = (item: SelectableValue<string>) => {
    onChange({
      ...query,
      timelineData: item.value!,
    });

    if (query.player) {
      onRunQuery();
    }
  };
  return (
    <div>
      <MatchTypeSelect label={'Game type'} onChange={onMatchTypeChange} value={query.matchType} />
      {query.matchType && (
        <MatchSelect
          label={'Match'}
          matchType={query.matchType}
          onChange={onMatchSelectChange}
          value={query.matchId}
          datasource={datasource}
        />
      )}
      {query.matchId && (
        <>
          <PlayerSelect
            label={'Player'}
            onChange={onPlayerChange}
            value={query.player}
            datasource={datasource}
            matchId={query.matchId}
          />
          <TimelineDataSelect label={'Timeline value'} onChange={onTimelineDataChange} value={query.timelineData} />
          <div className="gf-form">
            <InlineFormLabel
              tooltip={
                'The datasource sets the start of the match timeline to Now() - 2h. This switch changes the starting time to the original match time.'
              }
              className="query-keyword"
            >
              Use match time
            </InlineFormLabel>
            <InlineSwitch css={undefined} value={query.useMatchTimerange} onChange={onSwitchChange} />
          </div>
        </>
      )}
    </div>
  );
};
