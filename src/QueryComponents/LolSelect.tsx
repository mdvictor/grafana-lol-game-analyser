import React, { useMemo } from 'react';
import { SelectableValue } from '@grafana/data';
import { InlineFormLabel, Select } from '@grafana/ui';

interface LolSelectProps {
  label: string;
  onChange: (item: SelectableValue<string>) => void;
  value: string;
  options: Array<SelectableValue<string>>;
  width?: number;
}

export function LolSelect(props: LolSelectProps) {
  const { value, onChange, label, options, width = 20 } = props;

  const selected = useMemo(() => {
    return options.find((option) => option.value === value);
  }, [value, options]);

  return (
    <div className="gf-form">
      <InlineFormLabel className="query-keyword">{label}</InlineFormLabel>
      <Select menuPlacement="bottom" options={options} value={selected} onChange={onChange} width={width} />
    </div>
  );
}
