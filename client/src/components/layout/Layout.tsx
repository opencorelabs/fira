import {
  Box,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerOverlay,
  Flex,
  Icon,
  IconButton,
  useBreakpointValue,
  useColorMode,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import { RiMenuFill, RiMoonLine, RiSunLine } from 'react-icons/ri';

import { Sidebar } from './Sidebar';

type LayoutProps = {
  children: React.ReactNode;
};

export function Layout({ children }: LayoutProps) {
  const sidebarWidth = useBreakpointValue({ base: '100%', md: '75px', lg: '240px' });
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { colorMode, toggleColorMode } = useColorMode();
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
        <Flex
          as="header"
          align="center"
          justify={{ base: 'space-between', md: 'flex-end' }}
          py={4}
          px={6}
          bg={bg}
          borderBottomWidth="1px"
          borderBottomColor={useColorModeValue('gray.200', 'gray.600')}
        >
          <IconButton
            display={{ base: 'inline-flex', md: 'none' }}
            onClick={onOpen}
            variant="ghost"
            icon={<Icon as={RiMenuFill} />}
            aria-label="Menu"
          />
          <IconButton
            display="inline-flex"
            onClick={toggleColorMode}
            variant="ghost"
            icon={<Icon as={colorMode === 'dark' ? RiSunLine : RiMoonLine} />}
            aria-label="Toggle Theme"
          />
        </Flex>
        <Box p={6}>{children}</Box>
      </Box>
      <Drawer placement="left" onClose={onClose} isOpen={isOpen}>
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
