import { Button, DataTable, EmptyState, Flex } from '@raystack/apsara';
import { useNavigate, useParams } from '@tanstack/react-router';
import { useMemo } from 'react';
import { usePermissions } from '~/react/hooks/usePermissions';
import { V1Beta1User } from '~/src';
import { Role } from '~/src/types';
import { PERMISSIONS, shouldShowComponent } from '~/utils';
import { getColumns } from './member.columns';

export type MembersProps = {
  members?: V1Beta1User[];
  memberRoles?: Record<string, Role[]>;
  isLoading?: boolean;
};

export const Members = ({
  members,
  memberRoles,
  isLoading: isMemberLoading
}: MembersProps) => {
  const navigate = useNavigate({ from: '/projects/$projectId' });
  const { projectId } = useParams({ from: '/projects/$projectId' });

  const resource = `app/project:${projectId}`;
  const listOfPermissionsToCheck = useMemo(
    () => [
      {
        permission: PERMISSIONS.UpdatePermission,
        resource
      }
    ],
    [resource]
  );

  const { permissions, isFetching: isPermissionsFetching } = usePermissions(
    listOfPermissionsToCheck,
    !!projectId
  );

  const { canUpdateProject } = useMemo(() => {
    return {
      canUpdateProject: shouldShowComponent(
        permissions,
        `${PERMISSIONS.UpdatePermission}::${resource}`
      )
    };
  }, [permissions, resource]);

  const tableStyle = members?.length
    ? { width: '100%' }
    : { width: '100%', height: '100%' };

  const isLoading = isMemberLoading || isPermissionsFetching;

  const columns = useMemo(
    () => getColumns(memberRoles, isLoading),
    [memberRoles, isLoading]
  );

  const updatedUsers = useMemo(() => {
    return isLoading
      ? ([{ id: 1 }, { id: 2 }, { id: 3 }] as any)
      : members?.length
      ? members
      : [];
  }, [members, isLoading]);

  return (
    <Flex direction="column" style={{ paddingTop: '32px' }}>
      <DataTable
        data={updatedUsers}
        columns={columns}
        emptyState={noDataChildren}
        parentStyle={{ height: 'calc(100vh - 212px)' }}
        style={tableStyle}
      >
        <DataTable.Toolbar style={{ padding: 0, border: 0 }}>
          <Flex justify="between" gap="small">
            <Flex style={{ maxWidth: '360px', width: '100%' }}>
              <DataTable.GloabalSearch
                placeholder="Search by name or email"
                size="medium"
              />
            </Flex>
            {canUpdateProject ? (
              <Button
                variant="primary"
                style={{ width: 'fit-content' }}
                onClick={() =>
                  navigate({
                    to: '/projects/$projectId/invite',
                    params: { projectId: projectId }
                  })
                }
              >
                Add Team
              </Button>
            ) : null}
          </Flex>
        </DataTable.Toolbar>
      </DataTable>
    </Flex>
  );
};

const noDataChildren = (
  <EmptyState>
    <div className="svg-container"></div>
    <h3>0 members in your team</h3>
    <div className="pera">Try adding new members.</div>
  </EmptyState>
);
