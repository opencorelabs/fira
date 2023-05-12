import { Api } from '@fira/api';
import { useEffect } from 'react';

export default function Index() {
  useEffect(() => {
    (async () => {
      try {
        const { api } = new Api({ baseUrl: 'http://localhost:8080' });
        const response = await api.firaServiceGetApiInfo();
        console.info('response', response);
      } catch (error) {
        console.error('error', error);
      }
    })();
  }, []);

  return (
    <div>
      <h1>Dashboard</h1>
    </div>
  );
}
