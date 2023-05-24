import type { NextApiRequest, NextApiResponse } from 'next';

import { withSessionRoute } from 'src/lib/session/session';

export default withSessionRoute(async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method !== 'POST') {
    return res.status(405).send('Method Not Allowed');
  }
  req.session.destroy();
  res.status(200).json({ result: 'ok' });
});
