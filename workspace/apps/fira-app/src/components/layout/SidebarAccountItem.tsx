import {
  Avatar,
  Flex,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';

type SidebarAccountItemProps = {
  label: string;
  avatar: string;
  href: string;
};

function AccountItemButton({ label, avatar }) {
  return (
    <Flex
      alignItems="center"
      py={4}
      px={{ base: 0, md: 4 }}
      justifyContent={{ base: 'flex-start', md: 'center', lg: 'flex-start' }}
    >
      <Avatar size="xs" src={avatar} name={label} />
      <Text fontWeight="bold" ml={4} display={{ base: 'block', md: 'none', lg: 'block' }}>
        {label}
      </Text>
    </Flex>
  );
}

export function SidebarAccountItem({ label, avatar }: SidebarAccountItemProps) {
  useColorModeValue('gray.100', 'gray.700');
  return (
    <Menu isLazy>
      <MenuButton>
        <AccountItemButton label={label} avatar={avatar} />
      </MenuButton>
      <MenuList>
        <MenuItem>Download</MenuItem>
        <MenuItem>Create a Copy</MenuItem>
        <MenuItem>Mark as Draft</MenuItem>
        <MenuItem>Delete</MenuItem>
        <MenuItem>Attend a Workshop</MenuItem>
      </MenuList>
    </Menu>
  );
}
