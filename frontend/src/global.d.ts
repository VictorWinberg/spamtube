declare global {
  declare module "*.svg" {
    const filePath: string;
    export default filePath;
  }

  declare module "*.png" {
    const value: string;
    export = value;
  }

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
}

export {};
