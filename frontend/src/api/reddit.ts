import { get } from "./apiClient";

type TopPost = {
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

export type TopPostData = {
  children: TopPost[];
};

export async function getTopPosts(name: string) {
  return get<TopPostData[]>(`top/${name}`);
}

export async function searchSubreddits(query: string) {
  return get<string[]>(`search/${query}`);
}
