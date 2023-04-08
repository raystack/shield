import { EmptyState, Flex, Table } from "@odpf/apsara";
import { Outlet } from "react-router-dom";
import useSWR from "swr";
import { fetcher, tableStyle } from "~/utils/helper";
import { columns } from "./columns";

export default function UserList() {
  const { data, error } = useSWR("/admin/v1beta1/users", fetcher);
  const { users = [] } = data || { users: [] };
  return (
    <Flex direction="row" css={{ height: "100%", width: "100%" }}>
      <Table
        css={tableStyle}
        columns={columns}
        data={users ?? []}
        noDataChildren={noDataChildren}
      ></Table>
      <Outlet />
    </Flex>
  );
}

export const noDataChildren = (
  <EmptyState>
    <div className="svg-container"></div>
    <h3>0 user created</h3>
    <div className="pera">Try creating a new user.</div>
  </EmptyState>
);
