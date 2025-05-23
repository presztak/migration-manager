import { Instance } from 'types/instance';
import { APIResponse } from 'types/response';

export const fetchInstances = (): Promise<Instance[]> => {
  return new Promise((resolve, reject) => {
    fetch(`/1.0/instances?recursion=1`)
      .then((response) => response.json())
      .then((data) => resolve(data.metadata))
      .catch(reject);
  });
};

export const fetchInstance = (uuid: string): Promise<Instance> => {
  return new Promise((resolve, reject) => {
    fetch(`/1.0/instances/${uuid}`)
      .then((response) => response.json())
      .then((data) => resolve(data.metadata))
      .catch(reject);
  });
};

export const updateInstanceOverride = (uuid: string, body: string): Promise<APIResponse<null>> => {
  return new Promise((resolve, reject) => {
    fetch(`/1.0/instances/${uuid}/override`, {
      method: "PUT",
      body: body,
    })
      .then((response) => response.json())
      .then((data) => resolve(data))
      .catch(reject);
  });
};

export const deleteInstanceOverride = (uuid: string): Promise<APIResponse<object>> => {
  return new Promise((resolve, reject) => {
    fetch(`/1.0/instances/${uuid}/override`, {method: "DELETE"})
      .then((response) => response.json())
      .then((data) => resolve(data))
      .catch(reject);
  });
};
