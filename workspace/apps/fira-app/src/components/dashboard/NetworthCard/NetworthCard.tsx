import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Divider,
  Flex,
  Heading,
  StatGroup,
  StatLabel,
} from '@chakra-ui/react';
import { Stat, StatArrow, StatHelpText, StatNumber } from '@chakra-ui/react';
import { useCallback } from 'react';
import { FiPlus } from 'react-icons/fi';

import { AddAccountModal } from 'src/components/AddAccountModal/AddAccountModal';
import { useModal } from 'src/context/ModalContext';

export function NetworthCard() {
  const { open, close } = useModal();

  const handleAddAccount = useCallback(() => {
    open(<AddAccountModal onClose={close} />);
  }, [close, open]);

  // Each dashboard component is responsible for its own data fetching
  // if (error) return <div>failed to load</div>;

  //   Stats should be based on the following:
  //   - Total value of all accounts
  //   - Total value of all assets
  //   - Total value of all liabilities
  //   - Total value of all investments
  //   - Total value of all cash
  //   - Total value of all credit cards
  //   - Total value of all loans

  return (
    <Card>
      <CardHeader>
        <Flex justify="space-between" align="center">
          <Heading size="md">Networth</Heading>
          <Button leftIcon={<FiPlus />} onClick={handleAddAccount}>
            Add Account
          </Button>
        </Flex>
      </CardHeader>
      <Divider color="gray.200" />
      <CardBody>
        <Stat variant="networth" size="xl">
          <StatLabel>Total</StatLabel>
          <StatNumber>$750,000</StatNumber>
          <StatHelpText color="green.500">
            <StatArrow type="increase" />
            $12,321 (5.23%) last month
          </StatHelpText>
        </Stat>
        <StatGroup>
          <Stat size="sm">
            <StatLabel>Cash</StatLabel>
            <StatNumber>$250,000</StatNumber>
            <StatHelpText fontSize="sm" color="green.500">
              <StatArrow type="increase" />
              $12,321 (5.23%) last month
            </StatHelpText>
          </Stat>
          <Stat size="sm">
            <StatLabel>Investments</StatLabel>
            <StatNumber>$250,000</StatNumber>
            <StatHelpText fontSize="sm" color="red.500">
              <StatArrow type="decrease" />
              $(12,321) (-5.23%) last month
            </StatHelpText>
          </Stat>
          <Stat size="sm">
            <StatLabel>Crypto</StatLabel>
            <StatNumber>$250,000</StatNumber>
            <StatHelpText fontSize="sm" color="red.500">
              <StatArrow type="decrease" />
              $(12,321) (-5.23%) last month
            </StatHelpText>
          </Stat>
        </StatGroup>
      </CardBody>
    </Card>
  );
}
