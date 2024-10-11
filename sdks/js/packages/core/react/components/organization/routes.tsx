import { useCallback, useEffect } from 'react';
import {
  Outlet,
  createRoute,
  createRootRouteWithContext,
  useRouteContext,
  RouteComponent
} from '@tanstack/react-router';

import { Flex } from '@raystack/apsara';

import { Toaster } from 'sonner';
import { useFrontier } from '~/react/contexts/FrontierContext';
import Domain from './domain';
import { AddDomain } from './domain/add-domain';
import { VerifyDomain } from './domain/verify-domain';
import GeneralSetting from './general';
import { DeleteOrganization } from './general/delete';
import WorkspaceMembers from './members';
import { InviteMember } from './members/invite';
import UserPreferences from './preferences';

import { default as WorkspaceProjects } from './project';
import { AddProject } from './project/add';
import { DeleteProject } from './project/delete';
import { ProjectPage } from './project/project';

import WorkspaceSecurity from './security';
import { Sidebar } from './sidebar';
import WorkspaceTeams from './teams';
import { AddTeam } from './teams/add';
import { DeleteTeam } from './teams/delete';
import { TeamPage } from './teams/team';
import { UserSetting } from './user';
import { SkeletonTheme } from 'react-loading-skeleton';
import { InviteTeamMembers } from './teams/members/invite';
import { DeleteDomain } from './domain/delete';
import Billing from './billing';
import Tokens from './tokens';
import { ConfirmCycleSwitch } from './billing/cycle-switch';
import Plans from './plans';
import ConfirmPlanChange from './plans/confirm-change';
import MemberRemoveConfirm from './members/MemberRemoveConfirm';

export interface CustomScreen {
  name: string;
  path: string;
  category: 'Organization' | 'User';
  component: RouteComponent;
}

export interface OrganizationProfileProps {
  organizationId: string;
  defaultRoute?: string;
  showBilling?: boolean;
  showTokens?: boolean;
  showPreferences?: boolean;
  hideToast?: boolean;
  customScreens?: CustomScreen[];
}

export interface CustomRoutes {
  Organization: Pick<CustomScreen, 'name' | 'path'>[];
  User: Pick<CustomScreen, 'name' | 'path'>[];
}

type RouterContext = Pick<
  OrganizationProfileProps,
  | 'organizationId'
  | 'showBilling'
  | 'showTokens'
  | 'hideToast'
  | 'showPreferences'
> & { customRoutes: CustomRoutes };

export function getCustomRoutes(customScreens: CustomScreen[] = []) {
  return (
    customScreens?.reduce(
      (acc: CustomRoutes, { name, category, path }) => {
        acc[category].push({ name, path });
        return acc;
      },
      { Organization: [], User: [] }
    ) || { Organization: [], User: [] }
  );
}

const RootRouter = () => {
  const { organizationId, hideToast } = useRouteContext({ from: '__root__' });
  const { client, setActiveOrganization, setIsActiveOrganizationLoading } =
    useFrontier();

  const fetchOrganization = useCallback(async () => {
    try {
      setIsActiveOrganizationLoading(true);
      const resp = await client?.frontierServiceGetOrganization(organizationId);
      const organization = resp?.data.organization;
      setActiveOrganization(organization);
    } catch (err) {
      console.error(err);
    } finally {
      setIsActiveOrganizationLoading(false);
    }
  }, [
    client,
    organizationId,
    setActiveOrganization,
    setIsActiveOrganizationLoading
  ]);

  useEffect(() => {
    if (organizationId) {
      fetchOrganization();
    } else {
      setActiveOrganization(undefined);
    }
  }, [organizationId, fetchOrganization, setActiveOrganization]);

  const visibleToasts = hideToast ? 0 : 1;

  return (
    <SkeletonTheme
      highlightColor="var(--background-base)"
      baseColor="var(--background-base-hover)"
    >
      <Toaster richColors visibleToasts={visibleToasts} />
      <Flex style={{ width: '100%', height: '100%' }}>
        <Sidebar />
        <Outlet />
      </Flex>
    </SkeletonTheme>
  );
};

const rootRoute = createRootRouteWithContext<RouterContext>()({
  component: RootRouter
});
const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/',
  component: GeneralSetting
});

const deleteOrgRoute = createRoute({
  getParentRoute: () => indexRoute,
  path: '/delete',
  component: DeleteOrganization
});

const securityRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/security',
  component: WorkspaceSecurity
});

const membersRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/members',
  component: WorkspaceMembers
});

const inviteMemberRoute = createRoute({
  getParentRoute: () => membersRoute,
  path: '/modal',
  component: InviteMember
});

const removeMemberRoute = createRoute({
  getParentRoute: () => membersRoute,
  path: '/remove-member/$memberId/$invited',
  component: MemberRemoveConfirm
});

const teamsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/teams',
  component: WorkspaceTeams
});

const addTeamRoute = createRoute({
  getParentRoute: () => teamsRoute,
  path: '/modal',
  component: AddTeam
});

const domainsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/domains',
  component: Domain
});

const verifyDomainRoute = createRoute({
  getParentRoute: () => domainsRoute,
  path: '/$domainId/verify',
  component: VerifyDomain
});

const deleteDomainRoute = createRoute({
  getParentRoute: () => domainsRoute,
  path: '/$domainId/delete',
  component: DeleteDomain
});

const addDomainRoute = createRoute({
  getParentRoute: () => domainsRoute,
  path: '/modal',
  component: AddDomain
});

const teamRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/teams/$teamId',
  component: TeamPage
});

const inviteTeamMembersRoute = createRoute({
  getParentRoute: () => teamRoute,
  path: '/invite',
  component: InviteTeamMembers
});

const deleteTeamRoute = createRoute({
  getParentRoute: () => teamRoute,
  path: '/delete',
  component: DeleteTeam
});

const projectsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/projects',
  component: WorkspaceProjects
});

const addProjectRoute = createRoute({
  getParentRoute: () => projectsRoute,
  path: '/modal',
  component: AddProject
});

const projectPageRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/projects/$projectId',
  component: ProjectPage
});

const deleteProjectRoute = createRoute({
  getParentRoute: () => projectPageRoute,
  path: '/delete',
  component: DeleteProject
});

const profileRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/profile',
  component: UserSetting
});

const preferencesRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/preferences',
  component: UserPreferences
});

const billingRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/billing',
  component: Billing
});

const switchBillingCycleModalRoute = createRoute({
  getParentRoute: () => billingRoute,
  path: '/cycle-switch/$planId',
  component: ConfirmCycleSwitch
});

const plansRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/plans',
  component: Plans
});

const planDowngradeRoute = createRoute({
  getParentRoute: () => plansRoute,
  path: '/confirm-change/$planId',
  component: ConfirmPlanChange
});

const tokensRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '/tokens',
  component: Tokens
});

interface getRootTreeOptions {
  customScreens?: CustomScreen[];
}

export function getRootTree({ customScreens = [] }: getRootTreeOptions) {
  return rootRoute.addChildren([
    indexRoute.addChildren([deleteOrgRoute]),
    securityRoute,
    membersRoute.addChildren([inviteMemberRoute, removeMemberRoute]),
    teamsRoute.addChildren([addTeamRoute]),
    domainsRoute.addChildren([
      addDomainRoute,
      verifyDomainRoute,
      deleteDomainRoute
    ]),
    teamRoute.addChildren([deleteTeamRoute, inviteTeamMembersRoute]),
    projectsRoute.addChildren([addProjectRoute]),
    projectPageRoute.addChildren([deleteProjectRoute]),
    profileRoute,
    preferencesRoute,
    billingRoute.addChildren([switchBillingCycleModalRoute]),
    plansRoute.addChildren([planDowngradeRoute]),
    tokensRoute,
    ...customScreens.map(cc =>
      createRoute({
        path: cc.path,
        component: cc.component,
        getParentRoute: () => rootRoute
      })
    )
  ]);
}