import type { NextApiRequest, NextApiResponse } from 'next';

import { client } from 'src/lib/api';
import { withSessionRoute } from 'src/lib/session/session';

export default withSessionRoute(async function (
  req: NextApiRequest,
  res: NextApiResponse
) {
  const response = await client.POST('/api/token/pair', {
    body: req.body,
  });
  console.log('response', response);
  return res.status(405).json({ message: 'Not Implemented' });
  // if (req.method !== 'POST') {
  //   return res.status(405).send('Method Not Allowed');
  // }

  // try {
  // const data = req.body;
  // const response = await getApi().firaServiceLoginAccount({
  //   namespace: V1AccountNamespace.ACCOUNT_NAMESPACE_CONSUMER,
  //   credential: {
  //     credentialType: V1AccountCredentialType.ACCOUNT_CREDENTIAL_TYPE_EMAIL,
  //     emailCredential: {
  //       email: data?.email,
  //       password: data?.password,
  //       verify: true,
  //     },
  //   },
  // });
  // if (response.error) {
  //   return res
  //     .status(response.status)
  //     .json({ error: response.error ?? response.statusText });
  // }
  // let me: HttpResponse<V1Account, RpcStatus> | null = null;
  // if (
  //   response.data.status === V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_OK
  // ) {
  //   me = await getApi().firaServiceGetAccount({
  //     headers: {
  //       authorization: `Bearer ${response.data.jwt}`,
  //     },
  //   });
  // }
  // req.session.user = {
  //   token: response.data.jwt,
  //   verified:
  //     response.data.status ===
  //     V1AccountRegistrationStatus.ACCOUNT_REGISTRATION_STATUS_OK,
  //   status: response.data.status,
  //   ...(me ? { id: me.data.id } : {}),
  //   ...(me ? { name: me.data.name } : {}),
  //   ...(me ? { avatarUrl: me.data.avatarUrl } : {}),
  // };
  // await req.session.save();
  // res.status(response.status).json(response.data);
  // } catch (error) {
  //   res.status(500).json({ error: error.error?.message ?? 'something went wrong' });
  // }
});
