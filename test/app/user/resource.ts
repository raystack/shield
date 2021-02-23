import Lab from '@hapi/lab';
import Sinon from 'sinon';
import * as R from 'ramda';
import Code from 'code';
import { factory } from 'typeorm-seeding';
import * as Config from '../../../config/config';
import { lab } from '../../setup';
import { User } from '../../../model/user';
import { Group } from '../../../model/group';
import * as Resource from '../../../app/user/resource';
import CasbinSingleton from '../../../lib/casbin';

exports.lab = Lab.script();
const Sandbox = Sinon.createSandbox();

lab.afterEach(() => {
  Sandbox.restore();
});

lab.experiment('User::resource', () => {
  lab.experiment('create user', () => {
    let user;

    lab.beforeEach(async () => {
      user = await factory(User)().create();
    });

    lab.afterEach(() => {
      Sandbox.restore();
    });

    lab.test('should update user by email', async () => {
      const response = await Resource.create(user);
      Code.expect(response.displayName).to.equal(user.displayName);
      Code.expect(response.metadata.email).to.equal(user.metadata.email);
    });
  });

  lab.experiment('list users', () => {
    let users, groups, userEntityPolicy;

    lab.beforeEach(async () => {
      // setup data
      const dbUri = Config.get('/postgres').uri;
      const enforcer = await CasbinSingleton.create(dbUri);

      users = await factory(User)().createMany(5);
      groups = await factory(Group)().createMany(2);

      // user group mapping
      await enforcer.addSubjectGroupingJsonPolicy(
        { user: users[0].id },
        { group: groups[0].id }
      );
      await enforcer.addSubjectGroupingJsonPolicy(
        { user: users[1].id },
        { group: groups[0].id }
      );
      await enforcer.addSubjectGroupingJsonPolicy(
        { user: users[3].id },
        { group: groups[1].id }
      );
      await enforcer.addSubjectGroupingJsonPolicy(
        { user: users[2].id },
        { group: groups[1].id }
      );

      // create relavant policies
      await enforcer.addJsonPolicy(
        { group: groups[0].id },
        { entity: 'gojek', privacy: 'public' },
        { action: 'firehose.read' }
      );
      userEntityPolicy = {
        subject: { user: users[2].id },
        resource: { entity: 'gojek' },
        action: { action: 'firehose.read' }
      };
      await enforcer.addJsonPolicy(
        userEntityPolicy.subject,
        userEntityPolicy.resource,
        userEntityPolicy.action
      );
      await enforcer.addJsonPolicy(
        { user: users[2].id },
        { group: groups[1].id },
        { role: 'team.admin' }
      );
    });

    lab.test('should return all users if no filters are specifed', async () => {
      const getListWithFiltersStub = Sandbox.stub(
        Resource,
        'getListWithFilters'
      ).returns(<any>[]);

      const result = await Resource.list();
      Code.expect(result).to.equal(users);
      Sandbox.assert.notCalled(getListWithFiltersStub);
    });

    lab.test('should return users that match the filter', async () => {
      const removeTimestamps = R.omit(['createdAt', 'updatedAt']);

      const filter = {
        entity: 'gojek',
        privacy: 'public',
        action: 'firehose.read'
      };
      const result = (await Resource.list(filter)).map(removeTimestamps);

      const expectedResult = [
        { ...users[0], policies: [] },
        { ...users[1], policies: [] },
        { ...users[2], policies: [userEntityPolicy] }
      ].map(removeTimestamps);

      // ? We need to sort before checking because [1, 2, 3] != [2, 1, 3]
      Code.expect(R.sortBy(R.propOr(null, 'id'), result)).to.equal(
        R.sortBy(R.propOr(null, 'id'), expectedResult)
      );
    });
  });
});
