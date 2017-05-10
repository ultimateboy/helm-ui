
export class CTemplate {
  name: string;
  data: string;
}
export class Metadata {
  version: string;
  name: string;
  description: string;
}

export class Chart {
  metadata: Metadata;
  templates: CTemplate[];
  values: Values;
}

export class Values {
  raw: string;
}
export class Config {
  raw: string;
}

export class HelmTime {
  nanos: number;
  second: number;
}
export class Status {
  code: number;
  notes: string;
  resources: string;
}
export class Info {
  Description: string;
  first_deployed: HelmTime;
  last_deployed: HelmTime;
  status: Status;
}
export class Release {
  name: string;
  namespace: string;
  version: number;
  chart: Chart;
  config: Config;
  info: Info;
  loading: boolean;
}
