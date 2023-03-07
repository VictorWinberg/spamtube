import { get } from "./apiClient";

export type YoutubeVideoData = {
  Group: {
    Community: {
      StarRating: {
        Average: string;
        Count: string;
      };
      Statistics: {
        Views: string;
      };
    };
    Description: string;
    Thumbnail: {
      URL: string;
    };
  };
  Link: {
    Href: string;
  };
  Published: string;
  Title: string;
  Updated: string;
  VideoId: string;
};

export async function getUploadedYoutubeVideos() {
  return get<YoutubeVideoData[]>(`videos`);
}
