import {
  Avatar,
  Box,
  Button,
  Divider,
  Flex,
  PlacementWithLogical,
  Portal,
  Switch,
  Text,
  useBreakpointValue,
  useColorMode,
  useDisclosure,
  VStack,
} from '@chakra-ui/react';
import { Popover, PopoverContent, PopoverTrigger } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useCallback } from 'react';
import { RiLogoutBoxRLine, RiMoonLine, RiSettings2Line } from 'react-icons/ri';

import { logout } from 'src/lib/auth';

export function AccountMenu() {
  const { colorMode, toggleColorMode } = useColorMode();
  const { isOpen, onClose, onToggle } = useDisclosure();
  const router = useRouter();

  const placement: PlacementWithLogical | undefined = useBreakpointValue({
    base: 'auto',
    md: 'right-end',
  });

  const handleLogout = useCallback(async () => {
    await logout();
    router.push('/');
  }, [router]);

  const handleLink = (href: string) => () => {
    router.push(href);
    onClose();
  };

  const name = 'John Doe';

  return (
    <Popover isLazy placement={placement} closeOnBlur onClose={onClose} isOpen={isOpen}>
      <PopoverTrigger>
        <Button
          onClick={onToggle}
          variant="ghost"
          px={4}
          py={6}
          justifyContent={{ base: 'flex-start', md: 'center', lg: 'flex-start' }}
        >
          <Avatar size="xs" src={'avatar'} name={name} />
          <Text
            fontWeight="bold"
            ml={4}
            display={{ base: 'block', md: 'none', lg: 'block' }}
          >
            {name}
          </Text>
        </Button>
      </PopoverTrigger>
      <PopoverContent minW={{ base: '100vw', md: '100%' }}>
        <VStack align="flex-start" p={2}>
          <Flex
            as={Button}
            align="center"
            onClick={handleLink('/settings/account')}
            leftIcon={<RiSettings2Line />}
            justifyContent="flex-start"
            w="100%"
            variant="ghost"
          >
            Settings
          </Flex>
          <Flex
            as={Button}
            align="center"
            leftIcon={<RiMoonLine />}
            justifyContent="flex-start"
            w="100%"
            variant="ghost"
          >
            <Flex justify="space-between" w="100%">
              <Box>Dark Mode</Box>
              <Switch isChecked={colorMode === 'dark'} onChange={toggleColorMode} />
            </Flex>
          </Flex>
        </VStack>
        <Divider />
        <VStack align="flex-start" p={2}>
          <Flex
            as={Button}
            align="center"
            leftIcon={<RiLogoutBoxRLine />}
            justifyContent="flex-start"
            w="100%"
            variant="ghost"
            onClick={handleLogout}
          >
            Sign out
          </Flex>
        </VStack>
      </PopoverContent>
      {isOpen && (
        <Portal>
          <Box
            position={'fixed'}
            top={0}
            left={0}
            width={'100%'}
            height={'100%'}
            backgroundColor={'rgba(0, 0, 0, 0.5)'}
            zIndex={0}
            onClick={onClose}
          />
        </Portal>
      )}
    </Popover>
  );
}
