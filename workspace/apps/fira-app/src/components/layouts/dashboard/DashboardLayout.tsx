import {
  Box,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerOverlay,
  Flex,
  useBreakpointValue,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';

import { Footer } from './Footer';
import { Header } from './Header';
import { Sidebar } from './Sidebar';

type LayoutProps = {
  children: React.ReactNode;
};

export function DashboardLayout({ children }: LayoutProps) {
  const sidebarWidth = useBreakpointValue({ base: '100%', md: '75px', lg: '240px' });
  const { isOpen, onOpen, onClose } = useDisclosure();
  const sidebarBg = useColorModeValue('white', 'gray.800');
  const mainBg = useColorModeValue('gray.100', 'gray.800');

  return (
    <Flex height="100vh">
      <Box
        as="nav"
        width={sidebarWidth}
        maxWidth={sidebarWidth}
        bg={sidebarBg}
        borderRightWidth="1px"
        borderRightColor={useColorModeValue('gray.200', 'gray.600')}
        display={{ base: 'none', md: 'block' }}
        px={2}
      >
        <Sidebar />
      </Box>
      <Flex flex={1} overflow="auto" bg={mainBg} minH="100vh" direction="column">
        <Header onOpen={onOpen} />
        <Box as="main" p={6} flex={1}>
          {children}
        </Box>
        <Footer />
      </Flex>
      <Drawer placement="left" onClose={onClose} isOpen={isOpen} size="full">
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerBody bg={sidebarBg}>
            <Sidebar />
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </Flex>
  );
}
