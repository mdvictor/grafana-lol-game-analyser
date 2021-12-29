import React, { useState, useEffect } from 'react';
import { SelectableValue } from '@grafana/data';
import { MatchTypes } from 'types';
import { LolSelect } from './LolSelect';

interface MatchSelectProps {
  label: string;
  onChange: (item: SelectableValue<string>) => void;
  value: string;
}

export function MatchTypeSelect(props: MatchSelectProps) {
  const { value, onChange, label } = props;
  const [options, setOptions] = useState<Array<SelectableValue<string>>>([]);

  useEffect(() => {
    const opts = Object.entries(MatchTypes).map(([key, value]) => {
      return {
        label: key,
        value: value,
      };
    });

    setOptions(opts);
  }, []);

  return <LolSelect value={value} onChange={onChange} label={label} options={options} width={15} />;
}
