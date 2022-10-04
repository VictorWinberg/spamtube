import { get } from "./apiClient";

export async function getTopPosts(name: string) {
  return get<TopPostData[]>(`top/${name}`);
}

export async function searchSubreddits(query: string) {
  return get<string[]>(`search/${query}`);
}
