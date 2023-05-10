import { Grid, Icon, Link, Text, useColorModeValue } from '@chakra-ui/react';
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
      <Grid
        alignItems="center"
        py={4}
        px={{ base: 0, md: 4 }}
        justifyContent={{ base: 'center', lg: 'flex-start' }}
        gridTemplateColumns="15% 85%"
      >
        <Icon boxSize={5} as={icon} />
        <Text
          fontWeight="bold"
          ml={2}
          display={{ base: 'block', md: 'none', lg: 'block' }}
        >
          {label}
        </Text>
      </Grid>
    </Link>
  );
}
