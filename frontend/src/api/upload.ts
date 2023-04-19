import { post } from "./apiClient";

export async function startUploadFlow(params: {
  title: string;
  description: string;
  image: string;
  voice?: string;
  style?: string;
  background?: string;
  textContent: string;
}) {
  const body = {
    ref: "master",
    inputs: {
      title: params.title,
      description: params.description,
      image: params.image,
      text_content: params.textContent,
      custom_voice: params.voice,
      custom_style: params.style,
      custom_background: params.background,
    },
  };
  return post<string>(`generate`, JSON.stringify(body));
}
