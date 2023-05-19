import {
  Button,
  Flex,
  GridItem,
  Icon,
  Input,
  InputGroup,
  InputLeftElement,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  SimpleGrid,
} from '@chakra-ui/react';
import Image from 'next/image';
import { FiSearch } from 'react-icons/fi';

const accounts = [
  {
    id: '1',
    name: 'Bank of America',
    type: 'bank',
    logo: './logos/bankofamerica.svg',
  },
  {
    id: '2',
    name: 'Chase',
    type: 'bank',
    logo: './logos/chase.svg',
  },
  {
    id: '3',
    name: 'Wells Fargo',
    type: 'bank',
    logo: './logos/wellsfargo.svg',
  },
  {
    id: '4',
    name: 'Vanguard',
    type: 'brokerage',
    logo: './logos/vanguard.svg',
  },
  {
    id: '5',
    name: 'Coinbase',
    type: 'crypto',
    logo: './logos/coinbase.svg',
  },
  {
    id: '6',
    name: 'Robinhood',
    type: 'brokerage',
    logo: './logos/robinhood.svg',
  },
];

export function AddAccountModal({ onClose }) {
  return (
    <Modal isOpen onClose={onClose} size="2xl">
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Add Account</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          {/* TODO: convert to autocomplete and fetch available FI's */}
          <InputGroup>
            <InputLeftElement pointerEvents="none">
              <Icon as={FiSearch} color="gray.300" />
            </InputLeftElement>
            <Input placeholder="Search for banks, brokerages, financial accounts..." />
          </InputGroup>
          {/* TODO: Start with default list and update based on autocomplete results */}
          <SimpleGrid
            mt={4}
            templateColumns="repeat(3, 1fr)"
            gap={4}
            alignItems="center"
            justifyItems="center"
          >
            {accounts.map((account) => (
              <GridItem key={account.id} w="100%" h={100} as={Button} variant="outline">
                <Flex align="center" justify="center" h="100%">
                  <Image alt={account.name} src={account.logo} width={100} height={100} />
                </Flex>
              </GridItem>
            ))}
          </SimpleGrid>
          <Button mt={4} w="100%" variant="outline">
            Add a manual account or holding
          </Button>
        </ModalBody>
        <ModalFooter />
      </ModalContent>
    </Modal>
  );
}
