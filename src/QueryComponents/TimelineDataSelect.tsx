import React, { useState, useEffect } from 'react';
import { SelectableValue } from '@grafana/data';
import { LolSelect } from './LolSelect';
import { CHAMPIONS_STATS, DAMAGE_STATS, TimelineData } from 'types';
import { splitCamelCase } from 'utils/utils';

interface TimelineDataProps {
  label: string;
  onChange: (item: SelectableValue<string>) => void;
  value: string;
}

export function TimelineDataSelect(props: TimelineDataProps) {
  const { value, onChange, label } = props;
  const [options, setOptions] = useState<Array<SelectableValue<string>>>([]);

  useEffect(() => {
    let opts = Object.entries(TimelineData)
      .filter(([key, value]) => ![DAMAGE_STATS, CHAMPIONS_STATS].includes(key))
      .map(([key, value]) => {
        return {
          label: splitCamelCase(value as string),
          value: value as string,
        };
      });

    opts = opts.concat(
      Object.entries(TimelineData.DAMAGE_STATS).map(([key, value]) => {
        return {
          label: splitCamelCase(value as string),
          value: value as string,
        };
      })
    );

    opts = opts.concat(
      Object.entries(TimelineData.CHAMPION_STATS).map(([key, value]) => {
        return {
          label: splitCamelCase(value as string),
          value: value as string,
        };
      })
    );

    setOptions(opts);
  }, []);

  return <LolSelect value={value} onChange={onChange} label={label} options={options} width={50} />;
}
