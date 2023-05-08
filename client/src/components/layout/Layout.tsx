import {
  Box,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  Flex,
  Icon,
  IconButton,
  useBreakpointValue,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import { FiMenu } from 'react-icons/fi';

type LayoutProps = {
  children: React.ReactNode;
};

export function Layout({ children }: LayoutProps) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const sidebar = useBreakpointValue({ base: '100%', md: '260px' });

  return (
    <Flex height="100vh">
      <Box
        as="nav"
        width={sidebar}
        maxWidth={sidebar}
        bg={useColorModeValue('white', 'gray.800')}
        borderRightWidth="1px"
        borderRightColor={useColorModeValue('gray.200', 'gray.600')}
        py="6"
        display={{ base: isOpen ? 'block' : 'none', md: 'block' }}
      >
        {/* Sidebar */}
      </Box>
      <Box flex="1" overflow="auto">
        <Flex
          as="header"
          align="center"
          justify="space-between"
          py="4"
          px="6"
          bg={useColorModeValue('white', 'gray.800')}
          borderBottomWidth="1px"
          borderBottomColor={useColorModeValue('gray.200', 'gray.600')}
        >
          <IconButton
            display={{ base: 'inline-flex', md: 'none' }}
            onClick={onOpen}
            variant="ghost"
            icon={<Icon as={FiMenu} />}
            aria-label="Menu"
          />
        </Flex>
        <Box p={6}>{children}</Box>
      </Box>
      <Drawer placement="left" onClose={onClose} isOpen={isOpen}>
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerHeader>Menu</DrawerHeader>
          <DrawerBody>{/* Sidebar Content */}</DrawerBody>
        </DrawerContent>
      </Drawer>
    </Flex>
  );
}
