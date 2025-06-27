import { NetworkType } from "util/network";

export interface Network {
  identifier: string;
  source: string;
  type: NetworkType;
  location: string;
  properties: string;
  config: Record<string, string>;
}
