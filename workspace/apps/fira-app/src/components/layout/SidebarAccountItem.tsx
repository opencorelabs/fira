import {
  Avatar,
  Flex,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
  Text,
} from '@chakra-ui/react';
import Link from 'next/link';
import { signOut } from 'next-auth/react';
import { useCallback } from 'react';

type SidebarAccountItemProps = {
  label: string;
  avatar: string;
};

export function SidebarAccountItem({ label, avatar }: SidebarAccountItemProps) {
  const handleSignout = useCallback(() => {
    signOut();
  }, []);

  return (
    <Menu isLazy>
      <MenuButton>
        <Flex
          alignItems="center"
          py={4}
          px={{ base: 0, md: 4 }}
          justifyContent={{ base: 'flex-start', md: 'center', lg: 'flex-start' }}
        >
          <Avatar size="xs" src={avatar} name={label} />
          <Text
            fontWeight="bold"
            ml={4}
            display={{ base: 'block', md: 'none', lg: 'block' }}
          >
            {label}
          </Text>
        </Flex>
      </MenuButton>
      <MenuList>
        <MenuItem as={Link} href="/settings/account">
          Settings
        </MenuItem>
        <MenuDivider />
        <MenuItem onClick={handleSignout}>Sign out</MenuItem>
      </MenuList>
    </Menu>
  );
}
