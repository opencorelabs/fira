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
  useColorMode,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import React from 'react';
import { FiMenu, FiMoon, FiSun } from 'react-icons/fi';

import { WordMark } from '../WordMark';
import { Sidebar } from './Sidebar';

type LayoutProps = {
  children: React.ReactNode;
};

export function Layout({ children }: LayoutProps) {
  const sidebarWidth = useBreakpointValue({ base: '100%', md: '260px' });
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
        py={6}
        display={{ base: isOpen ? 'block' : 'none', md: 'block' }}
        px={2}
      >
        <Box px={4}>
          <WordMark />
        </Box>
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
            icon={<Icon as={FiMenu} />}
            aria-label="Menu"
          />
          <IconButton
            onClick={toggleColorMode}
            variant="ghost"
            icon={<Icon as={colorMode === 'dark' ? FiSun : FiMoon} />}
            aria-label="Toggle Theme"
          />
        </Flex>
        <Box p={6}>{children}</Box>
      </Box>
      <Drawer placement="left" onClose={onClose} isOpen={isOpen}>
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerHeader bg={bg}>
            <WordMark />
          </DrawerHeader>
          <DrawerBody bg={bg}>
            <Sidebar />
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </Flex>
  );
}
