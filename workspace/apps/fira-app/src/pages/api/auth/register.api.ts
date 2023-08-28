import type { NextApiRequest, NextApiResponse } from 'next';

import { client } from 'src/lib/api';
import { withSessionRoute } from 'src/lib/session/session';

export default withSessionRoute(async function (
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method !== 'POST') {
    return res.status(405).send('Method Not Allowed');
  }
  try {
    const register = await client.POST('/api/accounts/register', {
      body: req.body,
    });
    if (register.error) {
      return res.status(register.response.status).json({ error: register.error });
    }

    const token = await client.POST('/api/token/pair', {
      body: { email: req.body.email_address, password: req.body.password },
    });

    req.session.user = {
      token: token.data?.access,
      refresh: token.data?.refresh,
      verified: register.data.verified,
      name: register.data.full_name,
      email: register.data.email_address,
    };
    await req.session.save();
    return res.status(register.response.status).json(register.data);
  } catch (error) {
    return res.status(500).json({ error: error?.message ?? 'something went wrong' });
  }
});
