import createClient from 'openapi-fetch';

import { paths } from './schema';

export const client = createClient<paths>({
  baseUrl: 'http://localhost:8080',
});

console.log('client', client);
