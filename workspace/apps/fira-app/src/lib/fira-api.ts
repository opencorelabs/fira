import { Api } from '@fira/api-sdk';

export function getApi() {
  const { api } = new Api({ baseUrl: process.env.NEXT_PUBLIC_BASE_URL });
  return api;
}
