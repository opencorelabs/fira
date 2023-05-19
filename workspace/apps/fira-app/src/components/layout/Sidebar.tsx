import { Box, Flex, Stack } from '@chakra-ui/react';
import React from 'react';
import { RiBankLine, RiPieChart2Line } from 'react-icons/ri';

import { WordMark } from '../WordMark';
import { SidebarAccountItem } from './SidebarAccountItem';
import { SidebarItem } from './SidebarItem';

export function Sidebar() {
  return (
    <Flex direction="column" justify="space-between" h="100%" pb={2}>
      <Box>
        <Flex
          alignItems="center"
          px={{ base: 0, lg: 4 }}
          h="75px"
          justifyContent={{ base: 'flex-start', md: 'center', lg: 'flex-start' }}
        >
          <WordMark size={{ base: 'md', lg: 'xl' }} />
        </Flex>
        <Stack spacing={1}>
          <SidebarItem label="Net Worth" icon={RiPieChart2Line} href="/dashboard" />
          <SidebarItem label="Accounts" icon={RiBankLine} href="/accounts" />
        </Stack>
      </Box>
      <SidebarAccountItem avatar="" label="Harry Hexhash" />
    </Flex>
  );
}
