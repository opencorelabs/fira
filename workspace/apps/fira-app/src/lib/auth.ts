import { api } from 'src/config/routes';

type LoginRequest = {
  email: string;
  password: string;
};

export async function login(data: LoginRequest, options = {}) {
  const response = await fetch(api.login, {
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
  const response = await fetch(api.register, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
    ...options,
  });
  return response.json();
}

export async function logout() {
  const response = await fetch(api.logout, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
  });
  return response.json();
}
