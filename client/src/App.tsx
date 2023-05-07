import { AppProviders } from 'src/AppProviders';

export function App({ Component, pageProps }) {
  return (
    <AppProviders>
      <Component {...pageProps} />
    </AppProviders>
  );
}
