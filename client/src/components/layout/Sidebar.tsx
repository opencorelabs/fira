import { VStack } from '@chakra-ui/react';
import { CiBank } from 'react-icons/ci';
import { FiPieChart } from 'react-icons/fi';

import { SidebarItem } from './SidebarItem';

export function Sidebar() {
  return (
    <VStack gap={1}>
      <SidebarItem label="Net Worth" icon={FiPieChart} href="/" />
      <SidebarItem label="Accounts" icon={CiBank} href="/accounts" />
    </VStack>
  );
}
