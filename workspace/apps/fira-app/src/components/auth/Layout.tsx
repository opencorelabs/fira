import { Box } from '@chakra-ui/react';

type AuthLayoutProps = {
  children: React.ReactNode;
};

export function AuthLayout({ children }: AuthLayoutProps) {
  return <Box>{children}</Box>;
}
