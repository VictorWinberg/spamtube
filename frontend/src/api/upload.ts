import { post } from "./apiClient";

export async function startUploadFlow(
  title: string,
  description: string,
  image: string,
  voice: string
) {
  const body = {
    ref: "master",
    inputs: {
      title,
      description,
      image,
      voice,
    },
  };
  return post<string>(`generate`, JSON.stringify(body));
}
