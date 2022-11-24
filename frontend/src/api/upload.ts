import { post } from "./apiClient";

export async function startUploadFlow(
  title: string,
  description: string,
  image: string,
  voice: string,
  service: string
) {
  const body = {
    ref: "master",
    inputs: {
      title,
      description,
      image,
      voice,
      service,
    },
  };
  return post<string>(`generate`, JSON.stringify(body));
}
