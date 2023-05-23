import type { NextApiRequest, NextApiResponse } from 'next';

import { withSessionRoute } from 'src/lib/session/session';

export default withSessionRoute(async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  // Used to request a new verification link
  return res.status(405).send('Method Not Allowed');
});
