type LoginRequest = {
  email: string;
  password: string;
};

export async function login(data: LoginRequest, options = {}) {
  const response = await fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
    ...options,
  });
  return response.json();
}

type RegisterRequest = {
  email: string;
  name: string;
  password: string;
};

export async function signup(data: RegisterRequest, options = {}) {
  const response = await fetch('/api/auth/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
    ...options,
  });
  return response.json();
}

export async function logout() {
  return fetch('/api/auth/logout', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });
}
