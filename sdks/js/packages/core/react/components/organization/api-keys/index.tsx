import { Flex, Text, EmptyState, Button } from '@raystack/apsara/v1';
import styles from './styles.module.css';
import keyIcon from '~/react/assets/key.svg';
import { Image } from '@raystack/apsara';
import { useFrontier } from '~/react/contexts/FrontierContext';
import { DEFAULT_API_PLATFORM_APP_NAME } from '~/react/utils/constants';

export default function ApiKeys() {
  const { config } = useFrontier();
  const appName = config?.apiPlatform?.appName || DEFAULT_API_PLATFORM_APP_NAME;
  return (
    <Flex direction="column" style={{ width: '100%' }}>
      <Flex className={styles.header}>
        <Text size={6}>API Keys</Text>
      </Flex>
      <Flex justify="center" align="center" className={styles.content}>
        <EmptyState
          icon={
            <Image
              // @ts-ignore
              src={keyIcon}
              alt="keyIcon"
            />
          }
          heading="No service account"
          subHeading={`Create a new account to use the APIs of ${appName}`}
          primaryAction={
            <Button
              data-test-id="frontier-sdk-new-service-account-btn"
              variant="secondary"
            >
              Create new service account
            </Button>
          }
        />
      </Flex>
    </Flex>
  );
}
