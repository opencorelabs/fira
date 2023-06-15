import { PAGE_ROUTES } from 'src/config/routes';

export default function Index() {
  return null;
}

export async function getServerSideProps() {
  return {
    redirect: {
      destination: PAGE_ROUTES.DASHBOARD,
      permanent: false,
    },
  };
}
