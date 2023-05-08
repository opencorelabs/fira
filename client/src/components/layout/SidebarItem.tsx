import { Flex, Icon, Link, Text, useColorModeValue } from '@chakra-ui/react';
import NextLink from 'next/link';
import { IconType } from 'react-icons';

type SidebarItemProps = {
  label: string;
  icon: IconType;
  href: string;
};

export function SidebarItem({ label, icon, href }: SidebarItemProps) {
  const bg = useColorModeValue('gray.100', 'gray.700');
  return (
    <Link as={NextLink} href={href} w="100%" _hover={{ bg }} borderRadius="md">
      <Flex align="center" p={4}>
        <Icon as={icon} />
        <Text ml={2}>{label}</Text>
      </Flex>
    </Link>
  );
}
