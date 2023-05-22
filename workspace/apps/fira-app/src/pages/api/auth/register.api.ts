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

  try {
    const data = req.body;
    const response = await getApi().firaServiceCreateAccount({
      namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
      credential: {
        credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
        emailCredential: {
          email: data.email,
          name: data.name,
          password: data.password,
          verificationBaseUrl: process.env.NEXT_PUBLIC_VERIFICATION_BASE_URL,
        },
      },
    });

    if (!response.ok) {
      return res.status(response.status).send(response.statusText);
    }
    req.session.user = {
      token: response.data.jwt,
      status: response.data.status,
      verified:
        response.data.status ===
        V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_OK,
    };
    await req.session.save();
    res.status(response.status).json(response.data);
  } catch (error) {
    res.status(500).json({ error: error.error?.message ?? 'something went wrong' });
  }
});
