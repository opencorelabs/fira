import { useEffect, useState } from 'react';

import { api } from 'src/lib/fira-api';

export default function Index() {
  const [info, setInfo] = useState(null);
  useEffect(() => {
    (async () => {
      try {
        const response = await api.firaServiceGetApiInfo();
        setInfo(response);
      } catch (error) {
        console.error('error', error);
      }
    })();
  }, []);

  return (
    <div>
      <h1>Dashboard</h1>
      <pre>{JSON.stringify(info, null, 2)}</pre>
    </div>
  );
}
