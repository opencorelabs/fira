import { Avatar, Flex, Link, Text, useColorModeValue } from '@chakra-ui/react';
import NextLink from 'next/link';

type SidebarAccountItemProps = {
  label: string;
  avatar: string;
  href: string;
};

export function SidebarAccountItem({ label, avatar, href }: SidebarAccountItemProps) {
  const bg = useColorModeValue('gray.100', 'gray.700');
  return (
    <Link as={NextLink} href={href} w="100%" _hover={{ bg }} borderRadius="md">
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
    </Link>
  );
}
