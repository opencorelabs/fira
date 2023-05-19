import { Button, Icon, Text } from '@chakra-ui/react';
import NextLink from 'next/link';
import { IconType } from 'react-icons';

type SidebarItemProps = {
  label: string;
  icon: IconType;
  href: string;
};

export function SidebarItem({ label, icon, href }: SidebarItemProps) {
  return (
    <Button
      as={NextLink}
      href={href}
      w="100%"
      variant="ghost"
      alignItems="center"
      px={4}
      py={6}
      justifyContent={{ base: 'flex-start', md: 'center', lg: 'flex-start' }}
    >
      <Icon boxSize={5} as={icon} />
      <Text fontWeight="bold" ml={4} display={{ base: 'block', md: 'none', lg: 'block' }}>
        {label}
      </Text>
    </Button>
  );
}
