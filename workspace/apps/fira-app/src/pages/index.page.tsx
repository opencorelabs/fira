import { routes } from 'src/config/routes';

export default function Index() {
  return null;
}

export async function getServerSideProps() {
  return {
    redirect: {
      destination: routes.dashboard,
      permanent: false,
    },
  };
}
