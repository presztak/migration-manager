import { useQuery } from '@tanstack/react-query'
import { fetchInstances } from 'api/instances'
import DataTable from 'components/DataTable.tsx'

const Instance = () => {
  const {
    data: instances = [],
    error,
    isLoading,
  } = useQuery({ queryKey: ['instances'], queryFn: fetchInstances })

  const headers = ["UUID", "Source", "Inventory path", "OS version", "CPU", "Memory", "Migration status"];
  const rows = instances.map((item) => {
    return [
      {
        content: item.uuid
      },
      {
        content: item.source_id
      },
      {
        content: item.inventory_path
      },
      {
        content: item.os_version
      },
      {
        content: item.cpu.number_cpus
      },
      {
        content: item.memory.memory_in_bytes
      },
      {
        content: item.migration_status_string
      }];
  });

  if (isLoading) {
    return (
      <div>Loading instances...</div>
    );
  }

  if (error) {
    return (
      <div>Error while loading instances</div>
    );
  }

  return <DataTable headers={headers} rows={rows} />;
};

export default Instance;
