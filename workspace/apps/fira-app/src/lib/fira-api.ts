import { Api } from '@fira/api-sdk';

export function getApi() {
  const { api } = new Api({ baseUrl: process.env.NEXTAUTH_URL });
  return api;
}
