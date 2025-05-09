import { FC, useEffect, useState } from "react";
import { Button, Form, Table } from "react-bootstrap";
import { BsPlus, BsTrash } from "react-icons/bs";
import { MigrationWindow } from "types/batch";
import { formatDate } from 'util/date';

interface Props {
  value: MigrationWindow[];
  onChange: (value: MigrationWindow[]) => void;
}

const MigrationWindowWidget: FC<Props> = ({ value, onChange }) => {
  const [entries, setEntries] = useState<MigrationWindow[]>(value || []);
  const [start, setStart] = useState("");
  const [end, setEnd] = useState("");
  const [lockout, setLockout] = useState("");

  const handleAdd = () => {
    const newValues = [...entries, {"start": start, "end": end, "lockout": lockout}];
    setEntries(newValues);
    onChange(newValues);
    setStart("");
    setEnd("");
  };

  useEffect(() => {
    setEntries(value || {});
  }, [value]);

  const handleDelete = (index: number) => {
    const updated = entries.filter((_, idx) => idx != index);
    setEntries(updated);
    onChange(updated);
  };

  const handleEdit = (index: number, field: string, value: string) => {
    const newValue = entries[index];
    if (field === "start") {
      newValue.start = value;
    } else if (field == "end") {
      newValue.end = value;
    } else {
      newValue.lockout = value;
    }

    const newValues = entries.map((item, idx) =>
      idx === index ? newValue : item
    );
    setEntries(newValues);
    onChange(newValues);
  };

  return (
    <div>
      <Table borderless>
        <tbody>
          {entries.map((item, index) => (
            <tr key={index}>
              <td style={{ display: 'flex', gap: '8px' }}>
                <Form.Control
                  type="text"
                  size="sm"
                  value={item.start}
                  onChange={(e) => handleEdit(index, "start", e.target.value)}
                />
                <Form.Control
                  type="text"
                  size="sm"
                  value={item.end}
                  onChange={(e) => handleEdit(index, "end", e.target.value)}
                />
                <Form.Control
                  type="text"
                  size="sm"
                  value={item.lockout}
                  onChange={(e) => handleEdit(index, "lockout", e.target.value)}
                />
              </td>
              <td>
                <Button title="Delete" size="sm" variant="outline-secondary" className="bg-white border no-hover" onClick={() => handleDelete(index)}>
                  <BsTrash />
                </Button>
              </td>
            </tr>
          ))}
          <tr>
            <td style={{ display: 'flex', gap: '8px' }}>
              <Form.Control
                type="text"
                size="sm"
                placeholder="Start"
                value={start}
                onChange={(e) => setStart(e.target.value)}
              />
              <Form.Control
                type="text"
                size="sm"
                placeholder="End"
                value={end}
                onChange={(e) => setEnd(e.target.value)}
              />
              <Form.Control
                type="text"
                size="sm"
                placeholder="Lockout"
                value={lockout}
                onChange={(e) => setLockout(e.target.value)}
              />
            </td>
            <td>
              <Button title="Add" size="sm" variant="outline-secondary" className="bg-white border no-hover" onClick={handleAdd}>
                <BsPlus />
              </Button>
            </td>
          </tr>
          <tr>
            <td className="text-muted small">YYYY-MM-DD HH:MM:SS / YYYY-MM-DD HH:MM:SS UTC (e.g., {formatDate(new Date().toISOString())})</td>
          </tr>
        </tbody>
      </Table>
    </div>
  );
};

export default MigrationWindowWidget;
