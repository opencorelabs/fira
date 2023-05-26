import { Api } from '@fira/api-sdk';

export function getApi() {
  const { api } = new Api({ baseUrl: process.env.BASE_URL });
  return api;
}
