
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

export class Release {
  name: string;
  namespace: string;
  version: number;
  chart: Chart;
  config: Config;
  loading: boolean;
}
