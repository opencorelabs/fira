import { AuthLayout } from './AuthLayout';

type AuthLayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: AuthLayoutProps) {
  return <AuthLayout>{children}</AuthLayout>;
}
