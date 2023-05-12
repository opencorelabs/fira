import { Api } from '../src';

// This test is not ideal, because it depends on a live instance of the API
describe('The API', () => {
  const api = new Api({ baseUrl: 'http://localhost:8080' });

  it('returns API info', async () => {
    const res = await api.api.firaServiceGetApiInfo();
    expect(res.status).toBe(200);
    expect(res.data.version?.major).toBe(1);
  });
});
