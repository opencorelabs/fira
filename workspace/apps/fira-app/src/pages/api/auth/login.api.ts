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
    const token = await client.POST('/api/token/pair', {
      body: req.body,
    });

    // TODO: add better error handling here
    if (token.error) {
      // @ts-expect-error type never bs
      return res.status(401).json({ error: token.error.detail });
    }

    const me = await client.GET('/api/accounts/me', {
      headers: {
        authorization: `Bearer ${token.data.access}`,
      },
    });

    // TODO: add better error handling here
    if (!me.data) {
      return res.status(401).json({ error: 'Unauthorized' });
    }

    req.session.user = {
      token: token.data?.access,
      refresh: token.data?.refresh,
      verified: me.data.verified,
      name: me.data.full_name,
      email: me.data.email_address,
    };
    await req.session.save();
    res.status(token.response.status).json(token.data);
  } catch (error) {
    res.status(500).json({ error: error.error?.message ?? 'something went wrong' });
  }
});
