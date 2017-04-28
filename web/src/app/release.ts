export class Metadata {
  version: string;
  name: string;
  description: string;
}

export class Chart {
  metadata: Metadata;
}

export class Release {
  name: string;
  namespace: string;
  version: number;
  chart: Chart;
}
