import {
  Flex,
  Icon,
  IconButton,
  useColorMode,
  useColorModeValue,
} from '@chakra-ui/react';
import React from 'react';
import { RiMenuFill, RiMoonLine, RiSunLine } from 'react-icons/ri';

type HeaderProps = {
  onOpen: () => void;
};

export function Header({ onOpen }: HeaderProps) {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <Flex
      as="header"
      align="center"
      justify={{ base: 'space-between', md: 'flex-end' }}
      py={4}
      px={6}
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
  );
}
