import fs from 'node:fs';
import openapi from 'openapi-typescript';

(async () => {
  try {
    const output = await openapi('http://localhost:8080/api/openapi.json');
    await fs.promises.writeFile('./src/lib/api/schema.ts', output);
  } catch (error) {
    console.error('error', error);
  }
})();
