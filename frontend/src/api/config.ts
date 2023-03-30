import { get, put, remove } from "./apiClient";

export type ConfigEntry = {
  id: string;
  name: string;
  createdAt: string;
  cron: string;
};

export type ConfigEntryRequestBody = {
  id?: string;
  name?: string;
  cron?: string;
};

export async function upsertConfig(body: ConfigEntryRequestBody) {
  return put<ConfigEntry>("subreddits", JSON.stringify(body));
}

export async function deleteConfigEntry(id: string) {
  return remove<string>(`subreddits/${id}`);
}

export async function getConfigurations() {
  return get<ConfigEntry[]>("subreddits");
}
