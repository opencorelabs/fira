import {
  V1AccountCredentialType,
  V1AccountNamespace,
  V1AccountRegistrationStatus,
} from '@fira/api-sdk';
import type { NextApiRequest, NextApiResponse } from 'next';

import { getApi } from 'src/lib/fira-api';
import { withSessionRoute } from 'src/lib/session';

export default withSessionRoute(async function (
  req: NextApiRequest,
  res: NextApiResponse
) {
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

  req.session.user = {
    id: 230,
    token: response.data.jwt,
    verified:
      response.data.status === V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_OK,
  };
  await req.session.save();

  if (
    response.data.status ===
    V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL
  ) {
    res.status(response.status).json({ ...response.data, verifyEmail: true });
    return;
  }

  res.status(response.status).json(response.data);
});
