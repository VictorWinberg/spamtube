const BASE_URL = "/api";

export interface HTTPResponseInterface<T> extends Response {
  parsedBody?: T;
}

export async function get<T>(
  path: string,
  config: RequestInit = { method: "get" },
  baseUrl?: string
): Promise<HTTPResponseInterface<T>> {
  return client<T>(path, config, baseUrl);
}

export async function post<T>(
  path: string,
  body?: string,
  config: RequestInit = { method: "post", body },
  baseUrl?: string
): Promise<HTTPResponseInterface<T>> {
  return client<T>(path, config, baseUrl);
}

export async function put<T>(
  path: string,
  body?: string,
  config: RequestInit = { method: "put", body },
  baseUrl?: string
): Promise<HTTPResponseInterface<T>> {
  return client<T>(path, config, baseUrl);
}

export async function remove<T>(
  path: string,
  body?: string,
  config: RequestInit = { method: "delete", body },
  baseUrl?: string
): Promise<HTTPResponseInterface<T>> {
  return await client<T>(path, config, baseUrl);
}

async function client<T>(
  path: string,
  config: RequestInit = { method: "get" },
  baseUrl: string = BASE_URL
): Promise<HTTPResponseInterface<T>> {
  config.headers = new Headers({
    "Content-Type": "application/json",
  });
  const request = new Request(`${baseUrl}/${path}`, config);
  const response: HTTPResponseInterface<T> = await fetch(request);

  let body;

  try {
    body = await response.json();
  } catch {
    throw new Error("Could not parse response");
  }

  if (response.ok) {
    response.parsedBody = body ?? response.parsedBody;
    return response;
  }

  throw new Error(response.statusText);
}
