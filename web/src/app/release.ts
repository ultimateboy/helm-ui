
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
}

export class Release {
  name: string;
  namespace: string;
  version: number;
  chart: Chart;
  loading: boolean;
}
