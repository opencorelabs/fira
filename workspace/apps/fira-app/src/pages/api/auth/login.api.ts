import {
  V1AccountCredentialType,
  V1AccountNamespace,
  V1AccountRegistrationStatus,
} from '@fira/api-sdk';
import type { NextApiRequest, NextApiResponse } from 'next';

import { getApi } from 'src/lib/fira-api';
import { withSessionRoute } from 'src/lib/session/session';

export default withSessionRoute(async function (
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method !== 'POST') {
    return res.status(405).send('Method Not Allowed');
  }

  const credentials = req.body;
  const response = await getApi().firaServiceLoginAccount({
    namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
    credential: {
      credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
      emailCredential: {
        email: credentials?.email,
        password: credentials?.password,
      },
    },
  });

  if (!response.ok) {
    return res.status(response.status).send(response.statusText);
  }

  const me = await getApi().firaServiceGetAccount({
    headers: {
      authorization: `Bearer ${response.data.jwt}`,
    },
  });

  console.info('me', me);

  req.session.user = {
    id: me.data.id,
    token: response.data.jwt,
    verified:
      response.data.status === V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_OK,
    status: response.data.status,
    name: me.data.name,
    avatar: me.data.avatarUrl,
  };
  await req.session.save();

  res.status(response.status).json(response.data);
});
