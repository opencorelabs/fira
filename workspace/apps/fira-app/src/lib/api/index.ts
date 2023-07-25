import createClient from 'openapi-fetch';

export const client = createClient({
  baseUrl: 'http://localhost:8080/api',
});
