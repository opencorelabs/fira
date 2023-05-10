import {
  Avatar,
  Box,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerOverlay,
  Flex,
  Icon,
  IconButton,
  Link,
  Stack,
  Text,
  useBreakpointValue,
  useColorMode,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import NextLink from 'next/link';
import React from 'react';
import { RiMenuFill, RiMoonLine, RiSunLine } from 'react-icons/ri';
import { RiBankLine, RiPieChart2Line } from 'react-icons/ri';

import { WordMark } from '../WordMark';
import { SidebarItem } from './SidebarItem';

type LayoutProps = {
  children: React.ReactNode;
};

export function Layout({ children }: LayoutProps) {
  const sidebarWidth = useBreakpointValue({ base: '100%', md: '75px', lg: '240px' });
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { colorMode, toggleColorMode } = useColorMode();
  const bg = useColorModeValue('white', 'gray.800');
  const menubg = useColorModeValue('gray.100', 'gray.700');

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
        <Flex direction="column" justify="space-between" h="100%" pb={2}>
          <Box>
            <Flex
              alignItems="center"
              px={4}
              h="75px"
              justifyContent={{ base: 'center', lg: 'flex-start' }}
            >
              <WordMark size={{ base: 'md', lg: 'xl' }} />
            </Flex>
            <Stack spacing={1}>
              <SidebarItem label="Net Worth" icon={RiPieChart2Line} href="/" />
              <SidebarItem label="Accounts" icon={RiBankLine} href="/accounts" />
            </Stack>
          </Box>
          <Link as={NextLink} href="" w="100%" _hover={{ bg: menubg }} borderRadius="md">
            <Flex
              alignItems="center"
              py={4}
              px={{ base: 0, md: 4 }}
              justifyContent={{ base: 'center', lg: 'flex-start' }}
            >
              <Avatar boxSize={5} src="" />
              <Text
                fontWeight="bold"
                ml={4}
                display={{ base: 'block', md: 'none', lg: 'block' }}
              >
                Harry Hexhash
              </Text>
            </Flex>
          </Link>
        </Flex>
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
            <Flex direction="column" justify="space-between" h="100%" pb={2}>
              <Box>
                <Flex
                  alignItems="center"
                  px={4}
                  h="75px"
                  justifyContent={{ base: 'center', lg: 'flex-start' }}
                >
                  <WordMark size={{ base: 'md', lg: 'xl' }} />
                </Flex>
                <Stack spacing={1}>
                  <SidebarItem label="Net Worth" icon={RiPieChart2Line} href="/" />
                  <SidebarItem label="Accounts" icon={RiBankLine} href="/accounts" />
                </Stack>
              </Box>
              <Link
                as={NextLink}
                href=""
                w="100%"
                _hover={{ bg: menubg }}
                borderRadius="md"
              >
                <Flex
                  alignItems="center"
                  py={4}
                  px={{ base: 0, md: 4 }}
                  justifyContent={{ base: 'center', lg: 'flex-start' }}
                >
                  <Avatar boxSize={5} src="" />
                  <Text
                    fontWeight="bold"
                    ml={4}
                    display={{ base: 'block', md: 'none', lg: 'block' }}
                  >
                    Harry Hexhash
                  </Text>
                </Flex>
              </Link>
            </Flex>
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </Flex>
  );
}
