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

export function Layout({ children }: LayoutProps) {
  const sidebarWidth = useBreakpointValue({ base: '100%', md: '75px', lg: '240px' });
  const { isOpen, onOpen, onClose } = useDisclosure();
  const bg = useColorModeValue('white', 'gray.800');

  return (
    <Flex height="100vh">
      <Box
        as="nav"
        width={sidebarWidth}
        maxWidth={sidebarWidth}
        bg={bg}
        borderRightWidth="1px"
        borderRightColor={useColorModeValue('gray.200', 'gray.600')}
        display={{ base: 'none', md: 'block' }}
        px={2}
      >
        <Sidebar />
      </Box>
      <Box flex={1} overflow="auto">
        <Header onOpen={onOpen} />
        <Box as="main" p={6}>
          {children}
        </Box>
        <Footer />
      </Box>
      <Drawer placement="left" onClose={onClose} isOpen={isOpen} size="full">
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerBody bg={bg}>
            <Sidebar />
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </Flex>
  );
}
