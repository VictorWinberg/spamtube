import { get } from "./apiClient";

export type TopPost = {
  data: {
    subreddit: string;
    selftext: string;
    authorFullName: string;
    title: string;
    upvote_ratio: number;
    ups: number;
    isOriginalContent: number;
    score: number;
    thumbnail: string;
    over18: boolean;
    subredditId: string;
    id: string;
    author: string;
    permalink: string;
    url: string;
  };
};

export async function getTopPosts(name: string) {
  return get<TopPost[]>(`top/${name}`);
}

export async function searchSubreddits(query: string) {
  return get<string[]>(`search/${query}`);
}
