import { FC, useState } from 'react';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import { useParams } from 'react-router';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { useFormik } from 'formik';
import {
  createInstanceOverride,
  deleteInstanceOverride,
  updateInstanceOverride,
  fetchInstance
} from 'api/instances';
import { useNotification } from 'context/notification';
import { APIResponse } from 'types/response';
import {
  bytesToHumanReadable,
  hasOverride,
  humanReadableToBytes
} from 'util/instance';

const InstanceOverrides: FC = () => {
  const { notify } = useNotification();
  const queryClient = useQueryClient();
  const { uuid } = useParams<{uuid:string}>();
  const [showOverrideDeleteModal, setShowOverrideDeleteModal] = useState(false);

  const {
    data: instance,
    error,
    isLoading,
  } = useQuery({
    queryKey: ['instances', uuid],
    queryFn: () => {
      return fetchInstance(uuid ?? "");
    }
    });

  const overrideExists = hasOverride(instance);

  let formikInitialValues = {
      comment: '',
      disable_migration: 'false',
      number_cpus: 0,
      memory_in_bytes: '',
  };

  if (instance && overrideExists) {
    const overrides = instance.overrides;
    formikInitialValues = {
      comment: overrides.comment,
      disable_migration: overrides.disable_migration.toString(),
      number_cpus: overrides.number_cpus,
      memory_in_bytes: bytesToHumanReadable(overrides.memory_in_bytes),
    };
  }

  const handleSuccessResponse = (response: APIResponse<null>) => {
    if (response.error_code == 0) {
      void queryClient.invalidateQueries({queryKey: ['instances', uuid]});
      notify.success(`Override for the instance with ${uuid} saved.`);
      return;
    }
    notify.error(`Failed to save override for ${uuid}. ${response.error}`);
  }

  const handleErrorResponse = (e: Error) => {
    notify.error(`Failed to save override for ${uuid}. ${e}`);
  }

  const validateForm = (values: any) => {
    const errors: any = {};

    if (values.memory_in_bytes) {
      try {
        humanReadableToBytes(values.memory_in_bytes);
      } catch(e: any) {
        errors.memory_in_bytes = e.toString();
      }
    }

    return errors;
  };

  const formik = useFormik({
    initialValues: formikInitialValues,
    validate: validateForm,
    enableReinitialize: true,
    onSubmit: (values) => {
      let memoryInBytes = 0;

      if (values.memory_in_bytes) {
        try {
          memoryInBytes = humanReadableToBytes(values.memory_in_bytes);
        } catch(e) {
          notify.error(`Failed to save override for ${uuid}. ${e}`);
          return;
        }
      }

      const modifiedValues = {
        ...values,
        uuid: uuid,
        memory_in_bytes: memoryInBytes,
        disable_migration: values.disable_migration == 'true',
      };
      if (!overrideExists) {
        createInstanceOverride(uuid ?? '', JSON.stringify(modifiedValues, null, 2))
          .then((response) => {
            handleSuccessResponse(response);
          })
          .catch((e) => {
            handleErrorResponse(e);
          });
      } else {
        updateInstanceOverride(uuid ?? '', JSON.stringify(modifiedValues, null, 2))
          .then((response) => {
            handleSuccessResponse(response);
          })
          .catch((e) => {
            handleErrorResponse(e);
          });
      }
     },
   });

  const handleDelete = () => {
    deleteInstanceOverride(uuid ?? "")
      .then((response) => {
        handleOverrideModalClose();
        if (response.error_code == 0) {
          void queryClient.invalidateQueries({queryKey: ['instances', uuid]});
          notify.success(`Override for the instance with ${uuid} deleted.`);
          return;
        }
        notify.error(`Failed to save override for ${uuid}. ${response.error}`);
    })
      .catch((e) => {
        handleOverrideModalClose();
        notify.error(`Failed to delete override for ${uuid}. ${e}`);
    });
  };


  const handleOverrideModalClose = () => setShowOverrideDeleteModal(false);
  const handleOverrideModalShow = () => setShowOverrideDeleteModal(true);

  if (isLoading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error when fetching data.</div>
  }

  return (
    <div className="form-container">
      <Form noValidate>
        <Form.Group className="mb-3" controlId="comment">
          <Form.Label>Comment</Form.Label>
          <Form.Control
            type="text"
            name="comment"
            value={formik.values.comment}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            isInvalid={!!formik.errors.comment && formik.touched.comment}/>
          <Form.Control.Feedback type="invalid">
            {formik.errors.comment}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group className="mb-3" controlId="disable_migration">
          <Form.Label>Disable migration</Form.Label>
          <Form.Select
            name="disable_migration"
            value={formik.values.disable_migration}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            isInvalid={!!formik.errors.disable_migration && formik.touched.disable_migration}>
              <option value="false">no</option>
              <option value="true">yes</option>
          </Form.Select>
          <Form.Control.Feedback type="invalid">
            {formik.errors.disable_migration}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group className="mb-3" controlId="number_cpus">
          <Form.Label>Num VCPUS</Form.Label>
          <Form.Control
            type="number"
            name="number_cpus"
            value={formik.values.number_cpus}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            isInvalid={!!formik.errors.number_cpus && formik.touched.number_cpus}/>
          <Form.Control.Feedback type="invalid">
            {formik.errors.number_cpus}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group className="mb-3" controlId="memory_in_bytes">
          <Form.Label>Memory</Form.Label>
          <Form.Control
            type="text"
            name="memory_in_bytes"
            value={formik.values.memory_in_bytes}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            isInvalid={!!formik.errors.memory_in_bytes && formik.touched.memory_in_bytes}/>
          <Form.Control.Feedback type="invalid">
            {formik.errors.memory_in_bytes}
          </Form.Control.Feedback>
        </Form.Group>
        <Button className="float-end" variant="success" onClick={() => formik.handleSubmit()}>
          Save
        </Button>
        { overrideExists && (<Button className="float-end me-2" variant="danger" onClick={() => handleOverrideModalShow()}>
          Delete
        </Button>
        )}
      </Form>

      <Modal show={showOverrideDeleteModal} onHide={handleOverrideModalClose}>
        <Modal.Header closeButton>
          <Modal.Title>Delete instance override?</Modal.Title>
        </Modal.Header>
        <Modal.Body>Are you sure you want to delete the override for {uuid}? This action cannot be undone.</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleOverrideModalClose}>
            Close
          </Button>
          <Button variant="danger" onClick={handleDelete}>
            Delete
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
}

export default InstanceOverrides;
