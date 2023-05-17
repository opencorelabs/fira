/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface GetApiInfoResponseVersion {
  /** @format int32 */
  major?: number;
  /** @format int32 */
  minor?: number;
  /** @format int32 */
  patch?: number;
}

export interface ProtobufAny {
  "@type"?: string;
  [key: string]: any;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

export interface V1Account {
  /** Accounts can be created within one of these namespaces. */
  namespace?: V1AccountNamespace;
  status?: V1AccountRegistrationStatus;
  credentialType?: V1AccountCredentialType;
  id?: string;
  /** The account holder's name. */
  name?: string;
  /** The account holder's email. */
  email?: string;
  /** The account holder's avatar URL. */
  avatarUrl?: string;
  /**
   * The account creation timestamp.
   * @format date-time
   */
  createdAt?: string;
  /**
   * The account last updated timestamp.
   * @format date-time
   */
  updatedAt?: string;
}

/** @default "ACCOUNT_CREDENTIAL_TYPE_UNSPECIFIED" */
export enum V1AccountCredentialType {
  ACCOUNT_CREDENTIAL_TYPE_UNSPECIFIED = "ACCOUNT_CREDENTIAL_TYPE_UNSPECIFIED",
  ACCOUNT_CREDENTIAL_TYPE_EMAIL = "ACCOUNT_CREDENTIAL_TYPE_EMAIL",
  ACCOUNT_CREDENTIAL_TYPE_OAUTH_GITHUB = "ACCOUNT_CREDENTIAL_TYPE_OAUTH_GITHUB",
}

/**
 * Accounts can be created within one of these namespaces.
 * @default "ACCOUNT_NAMESPACE_UNSPECIFIED"
 */
export enum V1AccountNamespace {
  ACCOUNT_NAMESPACE_UNSPECIFIED = "ACCOUNT_NAMESPACE_UNSPECIFIED",
  ACCOUNT_NAMESPACE_DEVELOPER = "ACCOUNT_NAMESPACE_DEVELOPER",
  ACCOUNT_NAMESPACE_CONSUMER = "ACCOUNT_NAMESPACE_CONSUMER",
}

/** @default "ACCOUNT_REGISTRATION_STATUS_UNSPECIFIED" */
export enum V1AccountRegistrationStatus {
  ACCOUNT_REGISTRATION_STATUS_UNSPECIFIED = "ACCOUNT_REGISTRATION_STATUS_UNSPECIFIED",
  ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL = "ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL",
  ACCOUNT_REGISTRATION_STATUS_OK = "ACCOUNT_REGISTRATION_STATUS_OK",
  ACCOUNT_REGISTRATION_STATUS_ERROR = "ACCOUNT_REGISTRATION_STATUS_ERROR",
}

export interface V1BeginPasswordResetRequest {
  /** Namespace to log the user into - must match the original namespace used to create the account. */
  namespace?: V1AccountNamespace;
  /**
   * This will be used for building the URL in the password reset email.
   * For example, with a base URL of "https://myapp.net" will the link in the reset email will a "reset_token"
   * URL argument appended to it: "http://myapp.net?reset_token=zxb38rjs9nh"
   */
  resetBaseUrl?: string;
  /** Email address of the account to reset. */
  email?: string;
}

/** Placeholder - there is no information to return right now. If the email exists, a reset email will be sent. */
export type V1BeginPasswordResetResponse = object;

export interface V1CompletePasswordResetRequest {
  /** The reset_token provided in the URL. */
  token?: string;
  /** New password for the account. */
  password?: string;
}

export interface V1CompletePasswordResetResponse {
  /** Status of the password reset. */
  status?: V1AccountRegistrationStatus;
  /** If status is ACCOUNT_REGISTRATION_STATUS_ERROR this will be a human-readable message. */
  errorMessage?: string;
  /** If the status is ACCOUNT_REGISTRATION_STATUS_OK, this will be a new login JWT. */
  jwt?: string;
}

export type V1Connection = object;

export interface V1CreateAccountRequest {
  /**
   * Which namespace to create the account in
   * Accounts can be created within one of these namespaces.
   */
  namespace?: V1AccountNamespace;
  /** Login credentials for the new account */
  credential?: V1LoginCredential;
}

export interface V1CreateAccountResponse {
  /**
   * Status of the account registration. Most accounts will immediately go into the
   * ACCOUNT_REGISTRATION_STATUS_VERIFY_EMAIL status.
   */
  status?: V1AccountRegistrationStatus;
  /** If the status was ACCOUNT_REGISTRATION_STATUS_ERROR, this will be the human-readable error message. */
  errorMessage?: string;
  /** If the status is ACCOUNT_REGISTRATION_STATUS_OK, this will be the login JWT. */
  jwt?: string;
}

export interface V1CreateLinkSessionRequest {
  /** This is the URL that will be redirected to once the user session is comp */
  redirectUri?: string;
  /** Use connection_id when reconnecting a lapsed connection. */
  connectionId?: string;
}

export interface V1CreateLinkSessionResponse {
  /**
   * A link session represents an attempt by a user to connect to their financial institution, select accounts,
   * and provide consent to data sharing.
   */
  linkSession?: V1LinkSession;
}

export interface V1CredentialTypeEmail {
  /** The email address to use for the account. */
  email?: string;
  /** The password to use for the account. */
  password?: string;
  /**
   * This will be used for building the URL in the verification email.
   * For example, with a base URL of "https://myapp.net" will the link in the verification email will a "verification_token"
   * URL argument appended to it: "http://myapp.net?verification_token=zxb38rjs9nh"
   * Use only for registration or re-verification
   */
  verificationBaseUrl?: string;
  /**
   * Common name for the new account i.e. "Jane Smith" - provide if the credential type is email.
   * Use only for registration.
   */
  name?: string;
  /** If true, re-send the email verification. */
  verify?: boolean;
}

export interface V1CredentialTypeGithub {
  /** Public client ID for the GitHub OAuth app - must match the client secret stored on the server. */
  clientId?: string;
  /** The redirect URI configured for the GitHub OAuth app. */
  redirectUri?: string;
  /** The OAuth code returned by GitHub after the user has authorized the app. */
  code?: string;
}

export interface V1GetAccountResponse {
  account?: V1Account;
}

export interface V1GetApiInfoResponse {
  version?: GetApiInfoResponseVersion;
}

export interface V1GetLinkSessionResponse {
  /**
   * A link session represents an attempt by a user to connect to their financial institution, select accounts,
   * and provide consent to data sharing.
   */
  linkSession?: V1LinkSession;
}

export type V1LinkError = object;

/**
 * A link session represents an attempt by a user to connect to their financial institution, select accounts,
 * and provide consent to data sharing.
 */
export interface V1LinkSession {
  /** The link ID. Use in URL to fetch status. */
  linkId?: string;
  /** The configured redirect URI */
  redirectUri?: string;
  /** The configured connection to reconnect. */
  connectionId?: string;
  /** Status of the current link. A timed out link will show up as failed. */
  status?: V1LinkSessionStatus;
  /** If the link failed, this will be the LinkError object */
  error?: V1LinkError;
  /** If the link was successful, this will be the resulting Connection object. */
  connection?: V1Connection;
}

/** @default "LINK_SESSION_STATUS_UNSPECIFIED" */
export enum V1LinkSessionStatus {
  LINK_SESSION_STATUS_UNSPECIFIED = "LINK_SESSION_STATUS_UNSPECIFIED",
  LINK_SESSION_STATUS_CREATED = "LINK_SESSION_STATUS_CREATED",
  LINK_SESSION_STATUS_IN_PROGRESS = "LINK_SESSION_STATUS_IN_PROGRESS",
  LINK_SESSION_STATUS_COMPLETE = "LINK_SESSION_STATUS_COMPLETE",
  LINK_SESSION_STATUS_FAILED = "LINK_SESSION_STATUS_FAILED",
}

export interface V1LoginAccountRequest {
  /** Namespace to log the user into - must match the original namespace used to create the account. */
  namespace?: V1AccountNamespace;
  /** Login credentials for the account */
  credential?: V1LoginCredential;
}

export interface V1LoginAccountResponse {
  /** Status of the account */
  status?: V1AccountRegistrationStatus;
  /** If status is ACCOUNT_REGISTRATION_STATUS_ERROR this will be a human-readable message. */
  accountErrorMessage?: string;
  /** Whether the provided credentials were valid. If not, the JWT will not be provided. */
  credentialsValid?: boolean;
  /** If the credentials were valid, this will be a new login JWT. */
  jwt?: string;
}

export interface V1LoginCredential {
  credentialType?: V1AccountCredentialType;
  emailCredential?: V1CredentialTypeEmail;
  githubCredential?: V1CredentialTypeGithub;
}

/** @default "VERIFICATION_TYPE_UNSPECIFIED" */
export enum V1VerificationType {
  VERIFICATION_TYPE_UNSPECIFIED = "VERIFICATION_TYPE_UNSPECIFIED",
  VERIFICATION_TYPE_EMAIL = "VERIFICATION_TYPE_EMAIL",
}

export interface V1VerifyAccountRequest {
  /** The verification_token provided in the URL. */
  token?: string;
  /** The namespace to verify the account in. */
  namespace?: V1AccountNamespace;
  /** The verification type, will also be provided in the URL. */
  type?: V1VerificationType;
}

export interface V1VerifyAccountResponse {
  /**
   * Status of the account registration. If the token was valid and not expired, this should be
   * ACCOUNT_REGISTRATION_STATUS_OK. Otherwise, the user may have to re-verify. To re-verify, use the log in endpoint
   * with email and password, setting the verify flag to true.
   */
  status?: V1AccountRegistrationStatus;
  /** If the status is ACCOUNT_REGISTRATION_STATUS_ERROR this will be a human-readable message. */
  errorMessage?: string;
  /** If the token was valid, this will be a new login JWT. */
  jwt?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.Text]: (input: any) => (input !== null && typeof input !== "string" ? JSON.stringify(input) : input),
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title Fira API
 * @version 1.0.0
 * @license GNU Affero General Public License v3.0 (https://github.com/opencorelabs/fira/blob/main/LICENSE)
 * @contact Open Core Labs <pnwx@opencorelabs.org> (https://github.com/opencorelabs/fira)
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  api = {
    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceCompletePasswordReset
     * @summary Complete the password reset process. Provide the token from the password reset email and a new password.
     * @request POST:/api/v1/accounts/complete_password_reset
     * @response `200` `V1CompletePasswordResetResponse` A successful response.
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceCompletePasswordReset: (body: V1CompletePasswordResetRequest, params: RequestParams = {}) =>
      this.request<V1CompletePasswordResetResponse, RpcStatus>({
        path: `/api/v1/accounts/complete_password_reset`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceLoginAccount
     * @summary Exchange email/password credentials for a session JWT
     * @request POST:/api/v1/accounts/login
     * @response `200` `V1LoginAccountResponse` A successful response.
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceLoginAccount: (body: V1LoginAccountRequest, params: RequestParams = {}) =>
      this.request<V1LoginAccountResponse, RpcStatus>({
        path: `/api/v1/accounts/login`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceGetAccount
     * @summary Retrieve the account information for the currently logged in account holder.
     * @request GET:/api/v1/accounts/me
     * @response `200` `V1Account`
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceGetAccount: (params: RequestParams = {}) =>
      this.request<V1Account, RpcStatus>({
        path: `/api/v1/accounts/me`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
 * No description
 *
 * @tags FiraService
 * @name FiraServiceCreateAccount
 * @summary Register a new account with the server. This will trigger an out-of-band verification for the user. As of now,
this will take the form of email verification.
 * @request POST:/api/v1/accounts/register
 * @response `200` `V1CreateAccountResponse` A successful response.
 * @response `default` `RpcStatus` An unexpected error response.
 */
    firaServiceCreateAccount: (body: V1CreateAccountRequest, params: RequestParams = {}) =>
      this.request<V1CreateAccountResponse, RpcStatus>({
        path: `/api/v1/accounts/register`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceBeginPasswordReset
     * @summary Start the password reset process. If the email is found, a password reset email will be sent to the account holder.
     * @request POST:/api/v1/accounts/request_password_reset
     * @response `200` `V1BeginPasswordResetResponse` A successful response.
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceBeginPasswordReset: (body: V1BeginPasswordResetRequest, params: RequestParams = {}) =>
      this.request<V1BeginPasswordResetResponse, RpcStatus>({
        path: `/api/v1/accounts/request_password_reset`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceVerifyAccount
     * @summary Verify an unverified account using a token exchanged out-of-band with the account owner.
     * @request POST:/api/v1/accounts/verify
     * @response `200` `V1VerifyAccountResponse` A successful response.
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceVerifyAccount: (body: V1VerifyAccountRequest, params: RequestParams = {}) =>
      this.request<V1VerifyAccountResponse, RpcStatus>({
        path: `/api/v1/accounts/verify`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceGetApiInfo
     * @request GET:/api/v1/info
     * @response `200` `V1GetApiInfoResponse` A successful response.
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceGetApiInfo: (params: RequestParams = {}) =>
      this.request<V1GetApiInfoResponse, RpcStatus>({
        path: `/api/v1/info`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
 * No description
 *
 * @tags FiraService
 * @name FiraServiceCreateLinkSession
 * @summary Create a new link session. This will return a URL to redirect the user to where they will be able to select a
financial institution to connect to and log in.
 * @request POST:/api/v1/link
 * @response `200` `V1LinkSession`
 * @response `default` `RpcStatus` An unexpected error response.
 */
    firaServiceCreateLinkSession: (body: V1CreateLinkSessionRequest, params: RequestParams = {}) =>
      this.request<V1LinkSession, RpcStatus>({
        path: `/api/v1/link`,
        method: "POST",
        body: body,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags FiraService
     * @name FiraServiceGetLinkSession
     * @summary Retrieve the status of a link session.
     * @request GET:/api/v1/link/{linkId}
     * @response `200` `V1LinkSession`
     * @response `default` `RpcStatus` An unexpected error response.
     */
    firaServiceGetLinkSession: (linkId: string, params: RequestParams = {}) =>
      this.request<V1LinkSession, RpcStatus>({
        path: `/api/v1/link/${linkId}`,
        method: "GET",
        format: "json",
        ...params,
      }),
  };
}
