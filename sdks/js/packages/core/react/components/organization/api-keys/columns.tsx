import { TrashIcon } from '@radix-ui/react-icons';
import { ApsaraColumnDef } from '@raystack/apsara';
import { Button, Flex, Text } from '@raystack/apsara/v1';
import dayjs from 'dayjs';
import { V1Beta1ServiceUser } from '~/api-client';

export const getColumns = ({
  dateFormat
}: {
  dateFormat: string;
}): ApsaraColumnDef<V1Beta1ServiceUser>[] => {
  return [
    {
      header: 'Name',
      accessorKey: 'title',
      cell: ({ row, getValue }) => {
        return (
          <Flex direction="column">
            <Text>{getValue()}</Text>
          </Flex>
        );
      }
    },
    {
      header: 'Created on',
      accessorKey: 'created_at',
      cell: ({ row, getValue }) => {
        const value = getValue();
        return (
          <Flex direction="column">
            <Text>{dayjs(value).format(dateFormat)}</Text>
          </Flex>
        );
      }
    },
    {
      header: '',
      accessorKey: 'id',
      enableSorting: false,
      meta: {
        style: {
          padding: 0
        }
      },
      cell: ({ row, getValue }) => {
        return (
          <Button
            variant="text"
            size="small"
            data-test-id="frontier-sdk-delete-service-account-btn"
          >
            <TrashIcon />
          </Button>
        );
      }
    }
  ];
};