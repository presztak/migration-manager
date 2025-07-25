export interface InstanceDeviceInfo {
  type: string;
  label: string;
  summary: string;
}

export interface InstancePropertiesDisk {
  name: string;
  capacity: number;
  shared: boolean;
  supported: boolean;
}

export interface InstancePropertiesNIC {
  id: string;
  hardware_address: string;
  network: string;
}

export interface InstancePropertiesSnapshot {
  name: string;
}

export interface InstanceProperties {
  uuid: string;
  name: string;
  description: string;
  cpus: number;
  memory: number;
  location: string;
  os: string;
  os_version: string;
  secure_boot: boolean;
  legacy_boot: boolean;
  tpm: boolean;
  background_import: boolean;
  architecture: string;
  nics: InstancePropertiesNIC[];
  disks: InstancePropertiesDisk[];
  snapshots: InstanceSnapshotInfo[];
}

export interface InstancePropertiesConfigurable {
  description: string;
  cpus: number;
  memory: number;
  config: Record<string, string>;
}

export interface InstanceOverride {
  last_update: string;
  comment: string;
  disable_migration: boolean;
  properties: InstancePropertiesConfigurable;
}

export interface Instance {
  last_update_from_source: string;
  source: string;
  source_type: SourceType;
  properties: InstanceProperties;
  overrides: InstanceOverride;
}

export interface InstanceOverrideFormValues {
  comment: string;
  disable_migration: string;
  cpus: number;
  memory: string;
  config: Record<string, string>;
}
